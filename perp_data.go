package d8x_futures

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func CreatePerpetualManagerInstance(config Config) *IPerpetualManager {
	rpc, err := ethclient.Dial(config.NodeURL)
	if err != nil {
		log.Fatalf("Failed to connect to RPC %v", err)
	}
	defer rpc.Close()
	proxy, err := NewIPerpetualManager(config.ProxyAddr, rpc)
	if err != nil {
		log.Fatalf("Failed to instantiate Proxy contract: %v", err)
	}
	return proxy
}

func queryPoolStaticInfo(proxyContract *IPerpetualManager, config Config) {
	rpc, err := ethclient.Dial(config.NodeURL)
	if err != nil {
		log.Fatalf("Failed to connect to RPC %v", err)
	}
	defer rpc.Close()

	const idxFrom = 1
	const idxTo = 20
	//([][]*big.Int, []common.Address, []common.Address, common.Address, error)
	nestedPerpetualIds, poolShareTokenAddr, poolMarginTokenAddr, oracleFactory, error := proxyContract.GetPoolStaticInfo(nil, idxFrom, idxTo)
	if err != nil {
		panic(error)
	}
	fmt.Println(nestedPerpetualIds)
	fmt.Println(poolShareTokenAddr)
	fmt.Println(poolMarginTokenAddr)
	fmt.Println(oracleFactory)
}
