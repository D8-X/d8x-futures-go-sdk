package d8x_futures

import (
	"fmt"
	"math/big"
	"slices"
	"sync"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/config"
	"github.com/D8-X/d8x-futures-go-sdk/pkg/contracts"
	"github.com/D8-X/d8x-futures-go-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// handles on-chain oracle queries

type ChainOracles struct {
	Config     []utils.PriceFeedOnChainConfig
	LastResult map[string]*OracleObs
	ChainId    int
}

type OracleObs struct {
	Px          float64
	Ts          int64 //when did our query happen?
	UpdatedAtTs int64 //from oracle
	muRW        *sync.Mutex
	muChain     *sync.Mutex
}

func NewChainOracles(chainId int) (*ChainOracles, error) {
	var oracles ChainOracles
	var err error
	oracles.Config, err = config.GetPriceFeedOnChain()
	if err != nil {
		return nil, err
	}
	oracles.ChainId = chainId
	oracles.LastResult = make(map[string]*OracleObs)
	var wg sync.WaitGroup
	for _, c := range oracles.Config {
		if !slices.Contains(c.RelevantChains, chainId) {
			// only add oracles to the slice if we
			// are on a "relevantChain"
			continue
		}
		oracles.LastResult[c.Name] = &OracleObs{
			Px:          0,
			Ts:          0,
			UpdatedAtTs: 0,
			muRW:        &sync.Mutex{},
			muChain:     &sync.Mutex{},
		}
		wg.Add(1)
		go oracles.updatePrice(c.Name, false, &wg)
	}
	// wait for all on-chain prices to finish
	wg.Wait()
	return &oracles, nil
}

// GetPrice returns the last queried price and the updatedAt-timestamp from on-chain. It also queries a new price
// if max-feed age is exceeded but returns the last queried.
func (ch *ChainOracles) GetPrice(name string, doLog bool) (float64, int64, error) {
	j, err := ch.findConfigIdx(name)
	if err != nil {
		return 0, 0, err
	}
	obs, exists := ch.LastResult[name]
	if !exists {
		return 0, 0, fmt.Errorf("no price feed %s", name)
	}
	nowTs := time.Now().Unix()
	obs.muRW.Lock()
	ts := obs.Ts
	px := obs.Px
	updated := obs.UpdatedAtTs
	obs.muRW.Unlock()
	if nowTs-ts > ch.Config[j].MaxFeedAgeSec {
		go ch.updatePrice(name, doLog, nil)
	}
	return px, updated, nil
}

// updatePrice updates the price from on-chain unless there is another routine
// that is already updating
func (ch *ChainOracles) updatePrice(name string, doLog bool, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	obs := ch.LastResult[name]
	if !obs.muChain.TryLock() {
		// If the mutex is already locked, return
		return
	}
	if doLog {
		fmt.Printf("updating price for %s\n", name)
	}
	defer obs.muChain.Unlock()
	px, ts, err := ch.FetchPrice(name, doLog)
	if err != nil {
		fmt.Println("unable to fetch on-chain price in go-routine:" + err.Error())
		return
	}
	obs.muRW.Lock()
	defer obs.muRW.Unlock()
	obs.Px = px
	obs.UpdatedAtTs = ts
	obs.Ts = time.Now().Unix()
}

// FetchPrice gets the price for the given symbol from on-chain
func (ch *ChainOracles) FetchPrice(name string, doLog bool) (float64, int64, error) {
	j, err := ch.findConfigIdx(name)
	if err != nil {
		return 0, 0, err
	}
	return ch.fetchPrice(j, doLog)
}

func (ch *ChainOracles) findConfigIdx(name string) (int, error) {
	j := -1
	for k, n := range ch.Config {
		if n.Name == name {
			j = k
			break
		}
	}
	if j == -1 {
		return 0, fmt.Errorf("FetchPrice: no chain oracle found for name %s", name)
	}
	return j, nil
}

// fetchPrice gets the price for the symbol with config idx j from on-chain
func (ch *ChainOracles) fetchPrice(j int, doLog bool) (float64, int64, error) {
	var err error
	var rpc *ethclient.Client
	rpcIdx := -1
	maxTrial := 10
	for trial := 0; trial < maxTrial; trial++ {
		rpcIdx = (rpcIdx + 1) % len(ch.Config[j].RPCs)
		rpc, err = ethclient.Dial(ch.Config[j].RPCs[rpcIdx])
		if err != nil {
			if doLog {
				fmt.Printf("FetchPrice: could not initiate rpc %s: %s. Retrying", ch.Config[j].RPCs[rpcIdx], err.Error())
			}
			continue
		}
		if ch.Config[j].Type == "redstone" {
			px, ts, err := fetchRedstonePx(&ch.Config[j], rpc)
			if err != nil {
				if doLog {
					fmt.Printf("FetchPrice: could not get price rpc %s: %s. Retrying", ch.Config[j].RPCs[rpcIdx], err.Error())
				}
				time.Sleep(100 * time.Millisecond)
				continue
			}
			return px, ts, nil
		} else if ch.Config[j].Type == "angle" {
			// angle
			px, err := STUSDToUSDC(rpc)
			if err != nil {
				if doLog {
					fmt.Printf("FetchPrice: could not get angle price with rpc %s: %s. Retrying", ch.Config[j].RPCs[rpcIdx], err.Error())
				}
				time.Sleep(100 * time.Millisecond)
				continue
			}
			ts := time.Now().Unix()
			return px, ts, nil
		} else {
			return 0, 0, fmt.Errorf("unknown type %s", ch.Config[j].Type)
		}
	}
	return 0, 0, fmt.Errorf("no working call for token %s", ch.Config[j].Name)
}

func fetchRedstonePx(config *utils.PriceFeedOnChainConfig, rpc *ethclient.Client) (float64, int64, error) {
	oracleCtrct, err := contracts.NewRedStoneOracle(common.HexToAddress(config.PxFeedAddress), rpc)
	if err != nil {
		return 0, 0, fmt.Errorf("could not initiate contract: %s", err.Error())
	}
	var data struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	}
	data, err = oracleCtrct.LatestRoundData(nil)
	if err != nil {
		return 0, 0, fmt.Errorf("could not get price: %s", err.Error())
	}
	ts := data.UpdatedAt.Int64()
	return utils.DecNToFloat(data.Answer, uint8(config.Decimals)), ts, nil
}
