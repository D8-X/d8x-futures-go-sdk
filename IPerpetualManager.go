// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package d8x_futures

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AMMPerpLogicAMMVariables is an auto generated low-level Go binding around an user-defined struct.
type AMMPerpLogicAMMVariables struct {
	FLockedValue1             *big.Int
	FPoolM1                   *big.Int
	FPoolM2                   *big.Int
	FPoolM3                   *big.Int
	FAMMK2                    *big.Int
	FCurrentTraderExposureEMA *big.Int
}

// AMMPerpLogicMarketVariables is an auto generated low-level Go binding around an user-defined struct.
type AMMPerpLogicMarketVariables struct {
	FIndexPriceS2 *big.Int
	FIndexPriceS3 *big.Int
	FSigma2       *big.Int
	FSigma3       *big.Int
	FRho23        *big.Int
}

// IPerpetualGetterPerpetualStaticInfo is an auto generated low-level Go binding around an user-defined struct.
type IPerpetualGetterPerpetualStaticInfo struct {
	Id                     *big.Int
	LimitOrderBookAddr     common.Address
	FInitialMarginRate     int32
	FMaintenanceMarginRate int32
	PerpetualState         uint8
	CollCurrencyType       uint8
	S2BaseCCY              [4]byte
	S2QuoteCCY             [4]byte
	S3BaseCCY              [4]byte
	S3QuoteCCY             [4]byte
	FLotSizeBC             *big.Int
	FReferralRebateCC      *big.Int
	PriceIds               [][32]byte
	IsPyth                 []bool
}

// IPerpetualOrderOrder is an auto generated low-level Go binding around an user-defined struct.
type IPerpetualOrderOrder struct {
	LeverageTDR        uint16
	BrokerFeeTbps      uint16
	IPerpetualId       *big.Int
	TraderAddr         common.Address
	ExecutionTimestamp uint32
	BrokerAddr         common.Address
	SubmittedTimestamp uint32
	Flags              uint32
	IDeadline          uint32
	ReferrerAddr       common.Address
	FAmount            *big.Int
	FLimitPrice        *big.Int
	FTriggerPrice      *big.Int
	BrokerSignature    []byte
}

// IPerpetualTreasuryWithdrawRequest is an auto generated low-level Go binding around an user-defined struct.
type IPerpetualTreasuryWithdrawRequest struct {
	Lp                common.Address
	ShareTokens       *big.Int
	WithdrawTimestamp uint64
}

// PerpStorageLiquidityPoolData is an auto generated low-level Go binding around an user-defined struct.
type PerpStorageLiquidityPoolData struct {
	IsRunning                        bool
	IPerpetualCount                  uint8
	Id                               uint8
	ITargetPoolSizeUpdateTime        uint16
	MarginTokenAddress               common.Address
	FFundAllocationNormalizationCC   *big.Int
	FDefaultFundCashCC               *big.Int
	PrevAnchor                       uint64
	FRedemptionRate                  int32
	ShareTokenAddress                common.Address
	FPnLparticipantsCashCC           *big.Int
	FAMMFundCashCC                   *big.Int
	FTargetAMMFundSize               *big.Int
	FTargetDFSize                    *big.Int
	FMaxTransferPerConvergencePeriod *big.Int
	FBrokerCollateralLotSize         *big.Int
	PrevTokenAmount                  *big.Int
	NextTokenAmount                  *big.Int
	TotalSupplyShareToken            *big.Int
}

// PerpStorageMarginAccount is an auto generated low-level Go binding around an user-defined struct.
type PerpStorageMarginAccount struct {
	FLockedInValueQC             *big.Int
	FCashCC                      *big.Int
	FPositionBC                  *big.Int
	FUnitAccumulatedFundingStart *big.Int
	ILastOpenTimestamp           uint64
	FeeTbps                      uint16
	BrokerFeeTbps                uint16
	PositionId                   [16]byte
}

// PerpStoragePerpetualData is an auto generated low-level Go binding around an user-defined struct.
type PerpStoragePerpetualData struct {
	PoolId                              uint8
	Id                                  *big.Int
	FInitialMarginRate                  int32
	FMaintenanceMarginRate              int32
	FSigma2                             int32
	ILastFundingTime                    uint32
	ILastSettlementPriceUpdateTimestamp uint32
	ILastPriceJumpTimestamp             uint32
	ILastDefaultFundTransfer            uint32
	State                               uint8
	ECollateralCurrency                 uint8
	MinimalSpreadTbps                   uint16
	S2BaseCCY                           [4]byte
	IncentiveSpreadTbps                 uint16
	S2QuoteCCY                          [4]byte
	JumpSpreadTbps                      uint16
	S3BaseCCY                           [4]byte
	LiquidationPenaltyRateTbps          uint16
	S3QuoteCCY                          [4]byte
	FSigma3                             int32
	CurrentMarkPremiumRate              PerpStoragePriceTimeData
	FDFLambda                           [2]*big.Int
	FCurrentAMMExposureEMA              [2]*big.Int
	FAMMTargetDD                        [2]*big.Int
	FStressReturnS2                     [2]*big.Int
	FStressReturnS3                     [2]*big.Int
	PremiumRatesEMA                     *big.Int
	FTargetDFSize                       *big.Int
	FFundAllocationWeightCC             *big.Int
	FOpenInterest                       *big.Int
	FTargetAMMFundSize                  *big.Int
	FCurrentTraderExposureEMA           *big.Int
	FCurrentFundingRate                 *big.Int
	FUnitAccumulatedFunding             *big.Int
	FLotSizeBC                          *big.Int
	FkStar                              *big.Int
	FAMMFundCashCC                      *big.Int
	FAMMMinSizeCC                       *big.Int
	FMinimalTraderExposureEMA           *big.Int
	FMinimalAMMExposureEMA              *big.Int
	FTotalMarginBalance                 *big.Int
	FSettlementS3PriceData              *big.Int
	FReferralRebateCC                   *big.Int
	FSettlementS2PriceData              *big.Int
	ILastTargetPoolSizeTime             uint32
	FMarkPriceEMALambda                 int32
	FFundingRateClamp                   int32
	FMaximalTradeSizeBumpUp             int32
	FRho23                              int32
	FDFCoverNRate                       int32
}

// PerpStoragePriceTimeData is an auto generated low-level Go binding around an user-defined struct.
type PerpStoragePriceTimeData struct {
	FPrice *big.Int
	Time   uint64
}

// IPerpetualManagerMetaData contains all meta data concerning the IPerpetualManager contract.
var IPerpetualManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numLots\",\"type\":\"uint32\"}],\"name\":\"BrokerLotsTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fVolume\",\"type\":\"int128\"}],\"name\":\"BrokerVolumeTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"Clear\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"protocolFeeCC\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"participationFundFeeCC\",\"type\":\"int128\"}],\"name\":\"DistributeFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"positionId\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"amountLiquidatedBC\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"liquidationPrice\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"newPositionSizeBC\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fFeeCC\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fPnlCC\",\"type\":\"int128\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"poolId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shareAmount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"marginTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"shareTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"iTargetPoolSizeUpdateTime\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fMaxTransferPerConvergencePeriod\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fBrokerCollateralLotSize\",\"type\":\"int128\"}],\"name\":\"LiquidityPoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseOn\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"}],\"name\":\"LiquidityProvisionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"poolId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shareAmount\",\"type\":\"uint256\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"poolId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shareAmount\",\"type\":\"uint256\"}],\"name\":\"LiquidityWithdrawalInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int128[7]\",\"name\":\"baseParams\",\"type\":\"int128[7]\"},{\"indexed\":false,\"internalType\":\"int128[5]\",\"name\":\"underlyingRiskParams\",\"type\":\"int128[5]\"},{\"indexed\":false,\"internalType\":\"int128[13]\",\"name\":\"defaultFundRiskParams\",\"type\":\"int128[13]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eCollateralCurrency\",\"type\":\"uint256\"}],\"name\":\"PerpetualCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"PerpetualLimitOrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_liqPoolID\",\"type\":\"uint8\"}],\"name\":\"RunLiquidityPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"delay\",\"type\":\"uint8\"}],\"name\":\"SetBlockDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32[]\",\"name\":\"designations\",\"type\":\"uint32[]\"},{\"indexed\":false,\"internalType\":\"uint16[]\",\"name\":\"fees\",\"type\":\"uint16[]\"}],\"name\":\"SetBrokerDesignations\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tiers\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint16[]\",\"name\":\"feesTbps\",\"type\":\"uint16[]\"}],\"name\":\"SetBrokerTiers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tiers\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint16[]\",\"name\":\"feesTbps\",\"type\":\"uint16[]\"}],\"name\":\"SetBrokerVolumeTiers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"}],\"name\":\"SetClearedState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fSettlementMarkPremiumRate\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fSettlementS2Price\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fSettlementS3Price\",\"type\":\"int128\"}],\"name\":\"SetEmergencyState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"}],\"name\":\"SetNormalState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bytes4[2]\",\"name\":\"baseQuoteS2\",\"type\":\"bytes4[2]\"},{\"indexed\":false,\"internalType\":\"bytes4[2]\",\"name\":\"baseQuoteS3\",\"type\":\"bytes4[2]\"}],\"name\":\"SetOracles\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"value\",\"type\":\"int128\"}],\"name\":\"SetParameter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"value1\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"value2\",\"type\":\"int128\"}],\"name\":\"SetParameterPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int128[7]\",\"name\":\"baseParams\",\"type\":\"int128[7]\"}],\"name\":\"SetPerpetualBaseParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int128[5]\",\"name\":\"underlyingRiskParams\",\"type\":\"int128[5]\"},{\"indexed\":false,\"internalType\":\"int128[13]\",\"name\":\"defaultFundRiskParams\",\"type\":\"int128[13]\"}],\"name\":\"SetPerpetualRiskParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"value\",\"type\":\"int128\"}],\"name\":\"SetPoolParameter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tiers\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint16[]\",\"name\":\"feesTbps\",\"type\":\"uint16[]\"}],\"name\":\"SetTraderTiers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tiers\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint16[]\",\"name\":\"feesTbps\",\"type\":\"uint16[]\"}],\"name\":\"SetTraderVolumeTiers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"SetUtilityToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"}],\"name\":\"Settle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"}],\"name\":\"SettleState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"amount\",\"type\":\"int128\"}],\"name\":\"TokensDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"amount\",\"type\":\"int128\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"positionId\",\"type\":\"bytes16\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"leverageTDR\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executionTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"submittedTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iDeadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fAmount\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fLimitPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTriggerPrice\",\"type\":\"int128\"},{\"internalType\":\"bytes\",\"name\":\"brokerSignature\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structIPerpetualOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"newPositionSizeBC\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"price\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fFeeCC\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fPnlCC\",\"type\":\"int128\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOBFactory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOBFactory\",\"type\":\"address\"}],\"name\":\"TransferAddressTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fEarnings\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"newDefaultFundSize\",\"type\":\"int128\"}],\"name\":\"TransferEarningsToTreasury\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"broker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"feeCC\",\"type\":\"int128\"}],\"name\":\"TransferFeeToBroker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"referralRebate\",\"type\":\"int128\"}],\"name\":\"TransferFeeToReferrer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fNewAMMFundCash\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fNewLiqPoolTotalAMMFundsCash\",\"type\":\"int128\"}],\"name\":\"UpdateAMMFundCash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"liquidityPoolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fAMMFundCashCCInPerpetual\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fTargetAMMFundSizeInPerpetual\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fAMMFundCashCCInPool\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fTargetAMMFundSizeInPool\",\"type\":\"int128\"}],\"name\":\"UpdateAMMFundTargetSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"iLots\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"iNewBrokerLots\",\"type\":\"uint32\"}],\"name\":\"UpdateBrokerAddedCash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fDeltaAmountCC\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fNewFundCash\",\"type\":\"int128\"}],\"name\":\"UpdateDefaultFundCash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"liquidityPoolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fDefaultFundCashCC\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fTargetDFSize\",\"type\":\"int128\"}],\"name\":\"UpdateDefaultFundTargetSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fNormalizedPerpetualWeight\",\"type\":\"int128\"}],\"name\":\"UpdateFundAllocationWeight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fFundingRate\",\"type\":\"int128\"}],\"name\":\"UpdateFundingRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"positionId\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fPositionBC\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fCashCC\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fLockedInValueQC\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fFundingPaymentCC\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fOpenInterestBC\",\"type\":\"int128\"}],\"name\":\"UpdateMarginAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fMidPricePremium\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fMarkPricePremium\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fSpotIndexPrice\",\"type\":\"int128\"}],\"name\":\"UpdateMarkPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fDeltaAmountCC\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fNewFundCash\",\"type\":\"int128\"}],\"name\":\"UpdateParticipationFundCash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fCurrentTraderExposureEMA\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fCurrentAMMExposureEMAShort\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"fCurrentAMMExposureEMALong\",\"type\":\"int128\"}],\"name\":\"UpdateReprTradeSizes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"perpetualId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"unitAccumulativeFunding\",\"type\":\"int128\"}],\"name\":\"UpdateUnitAccumulatedFunding\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"}],\"name\":\"activatePerpetual\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128\",\"name\":\"_fTokenAmount\",\"type\":\"int128\"}],\"name\":\"addAMMLiquidityToPerpetual\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_iPoolIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128\",\"name\":\"_fSettlementS2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fSettlementS3\",\"type\":\"int128\"}],\"name\":\"adjustSettlementPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_iLots\",\"type\":\"uint32\"}],\"name\":\"brokerDepositToDefaultFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int128\",\"name\":\"fLockedValue1\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fPoolM1\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fPoolM2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fPoolM3\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fAMM_K2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fCurrentTraderExposureEMA\",\"type\":\"int128\"}],\"internalType\":\"structAMMPerpLogic.AMMVariables\",\"name\":\"_ammVars\",\"type\":\"tuple\"},{\"internalType\":\"int128\",\"name\":\"_fTradeAmount\",\"type\":\"int128\"}],\"name\":\"calculateBoundedSlippage\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128[2]\",\"name\":\"_fK2AMM\",\"type\":\"int128[2]\"},{\"internalType\":\"int128\",\"name\":\"_fk2Trader\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fCoverN\",\"type\":\"int128\"},{\"internalType\":\"int128[2]\",\"name\":\"fStressRet2\",\"type\":\"int128[2]\"},{\"internalType\":\"int128[2]\",\"name\":\"fStressRet3\",\"type\":\"int128[2]\"},{\"internalType\":\"int128[2]\",\"name\":\"fIndexPrices\",\"type\":\"int128[2]\"},{\"internalType\":\"enumAMMPerpLogic.CollateralCurrency\",\"name\":\"_eCCY\",\"type\":\"uint8\"}],\"name\":\"calculateDefaultFundSize\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int128\",\"name\":\"fLockedValue1\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fPoolM1\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fPoolM2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fPoolM3\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fAMM_K2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fCurrentTraderExposureEMA\",\"type\":\"int128\"}],\"internalType\":\"structAMMPerpLogic.AMMVariables\",\"name\":\"_ammVars\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int128\",\"name\":\"fIndexPriceS2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fIndexPriceS3\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSigma2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSigma3\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fRho23\",\"type\":\"int128\"}],\"internalType\":\"structAMMPerpLogic.MarketVariables\",\"name\":\"_mktVars\",\"type\":\"tuple\"},{\"internalType\":\"int128\",\"name\":\"_fTradeAmount\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fBidAskSpread\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fIncentiveSpread\",\"type\":\"int128\"}],\"name\":\"calculatePerpetualPrice\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int128\",\"name\":\"fLockedValue1\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fPoolM1\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fPoolM2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fPoolM3\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fAMM_K2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fCurrentTraderExposureEMA\",\"type\":\"int128\"}],\"internalType\":\"structAMMPerpLogic.AMMVariables\",\"name\":\"_ammVars\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int128\",\"name\":\"fIndexPriceS2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fIndexPriceS3\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSigma2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSigma3\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fRho23\",\"type\":\"int128\"}],\"internalType\":\"structAMMPerpLogic.MarketVariables\",\"name\":\"_mktVars\",\"type\":\"tuple\"},{\"internalType\":\"int128\",\"name\":\"_fTradeAmount\",\"type\":\"int128\"},{\"internalType\":\"bool\",\"name\":\"_withCDF\",\"type\":\"bool\"}],\"name\":\"calculateRiskNeutralPD\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"leverageTDR\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executionTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"submittedTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iDeadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fAmount\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fLimitPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTriggerPrice\",\"type\":\"int128\"},{\"internalType\":\"bytes\",\"name\":\"brokerSignature\",\"type\":\"bytes\"}],\"internalType\":\"structIPerpetualOrder.Order\",\"name\":\"_order\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"_feeTbps\",\"type\":\"uint16\"}],\"name\":\"chargePostingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"}],\"name\":\"countActivePerpAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marginTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_iTargetPoolSizeUpdateTime\",\"type\":\"uint16\"},{\"internalType\":\"int128\",\"name\":\"_fMaxTransferPerConvergencePeriod\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fBrokerCollateralLotSize\",\"type\":\"int128\"}],\"name\":\"createLiquidityPool\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_iPoolId\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[2]\",\"name\":\"_baseQuoteS2\",\"type\":\"bytes4[2]\"},{\"internalType\":\"bytes4[2]\",\"name\":\"_baseQuoteS3\",\"type\":\"bytes4[2]\"},{\"internalType\":\"int128[7]\",\"name\":\"_baseParams\",\"type\":\"int128[7]\"},{\"internalType\":\"int128[5]\",\"name\":\"_underlyingRiskParams\",\"type\":\"int128[5]\"},{\"internalType\":\"int128[13]\",\"name\":\"_defaultFundRiskParams\",\"type\":\"int128[13]\"},{\"internalType\":\"uint256\",\"name\":\"_eCollateralCurrency\",\"type\":\"uint256\"}],\"name\":\"createPerpetual\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_iPoolIdx\",\"type\":\"uint8\"},{\"internalType\":\"int128\",\"name\":\"_fAmount\",\"type\":\"int128\"}],\"name\":\"decreasePoolCash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128\",\"name\":\"_fAmount\",\"type\":\"int128\"},{\"internalType\":\"bytes[]\",\"name\":\"_updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_publishTimes\",\"type\":\"uint64[]\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128\",\"name\":\"_fDepositRequired\",\"type\":\"int128\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"leverageTDR\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executionTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"submittedTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iDeadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fAmount\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fLimitPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTriggerPrice\",\"type\":\"int128\"},{\"internalType\":\"bytes\",\"name\":\"brokerSignature\",\"type\":\"bytes\"}],\"internalType\":\"structIPerpetualOrder.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"depositMarginForOpeningTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"int128\",\"name\":\"_fAmount\",\"type\":\"int128\"}],\"name\":\"depositToDefaultFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"leverageTDR\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executionTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"submittedTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iDeadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fAmount\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fLimitPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTriggerPrice\",\"type\":\"int128\"},{\"internalType\":\"bytes\",\"name\":\"brokerSignature\",\"type\":\"bytes\"}],\"internalType\":\"structIPerpetualOrder.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"determineExchangeFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"leverageTDR\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executionTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"submittedTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iDeadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fAmount\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fLimitPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTriggerPrice\",\"type\":\"int128\"},{\"internalType\":\"bytes\",\"name\":\"brokerSignature\",\"type\":\"bytes\"}],\"internalType\":\"structIPerpetualOrder.Order\",\"name\":\"_order\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_hasOpened\",\"type\":\"bool\"}],\"name\":\"distributeFees\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_traderAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"_fDeltaPositionBC\",\"type\":\"int128\"}],\"name\":\"distributeFeesLiquidation\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"_digest\",\"type\":\"bytes32\"}],\"name\":\"executeCancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_lpAddr\",\"type\":\"address\"}],\"name\":\"executeLiquidityWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_traderAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"_fTraderPos\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fTradeAmount\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fPrice\",\"type\":\"int128\"},{\"internalType\":\"bool\",\"name\":\"_isClose\",\"type\":\"bool\"}],\"name\":\"executeTrade\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAMMPerpLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128[2]\",\"name\":\"_fIndexPrice\",\"type\":\"int128[2]\"}],\"name\":\"getAMMState\",\"outputs\":[{\"internalType\":\"int128[15]\",\"name\":\"\",\"type\":\"int128[15]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"}],\"name\":\"getActivePerpAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getActivePerpAccountsByChunks\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_brokerAddr\",\"type\":\"address\"}],\"name\":\"getBrokerDesignation\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_brokerAddr\",\"type\":\"address\"}],\"name\":\"getBrokerInducedFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getCollateralTokenAmountForPricing\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_brokerAddr\",\"type\":\"address\"}],\"name\":\"getCurrentBrokerVolume\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_traderAddr\",\"type\":\"address\"}],\"name\":\"getCurrentTraderVolume\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"_fPosition0\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fBalance0\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fTradeAmount\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fTargetLeverage\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fS2Mark\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fS3\",\"type\":\"int128\"}],\"name\":\"getDepositAmountForLvgPosition\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_brokerDesignation\",\"type\":\"uint32\"}],\"name\":\"getFeeForBrokerDesignation\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"}],\"name\":\"getFeeForBrokerStake\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_brokerAddr\",\"type\":\"address\"}],\"name\":\"getFeeForBrokerVolume\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"}],\"name\":\"getFeeForTraderStake\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_traderAddr\",\"type\":\"address\"}],\"name\":\"getFeeForTraderVolume\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"}],\"name\":\"getLastPerpetualBaseToUSDConversion\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128[2]\",\"name\":\"_fIndexPrice\",\"type\":\"int128[2]\"}],\"name\":\"getLiquidatableAccounts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"unsafeAccounts\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getLiquidityPool\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isRunning\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"iPerpetualCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"iTargetPoolSizeUpdateTime\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"marginTokenAddress\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fFundAllocationNormalizationCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fDefaultFundCashCC\",\"type\":\"int128\"},{\"internalType\":\"uint64\",\"name\":\"prevAnchor\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"fRedemptionRate\",\"type\":\"int32\"},{\"internalType\":\"address\",\"name\":\"shareTokenAddress\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fPnLparticipantsCashCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fAMMFundCashCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTargetAMMFundSize\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTargetDFSize\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fMaxTransferPerConvergencePeriod\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fBrokerCollateralLotSize\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"prevTokenAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"nextTokenAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalSupplyShareToken\",\"type\":\"uint128\"}],\"internalType\":\"structPerpStorage.LiquidityPoolData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolIdFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_poolIdTo\",\"type\":\"uint8\"}],\"name\":\"getLiquidityPools\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isRunning\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"iPerpetualCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"iTargetPoolSizeUpdateTime\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"marginTokenAddress\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fFundAllocationNormalizationCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fDefaultFundCashCC\",\"type\":\"int128\"},{\"internalType\":\"uint64\",\"name\":\"prevAnchor\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"fRedemptionRate\",\"type\":\"int32\"},{\"internalType\":\"address\",\"name\":\"shareTokenAddress\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fPnLparticipantsCashCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fAMMFundCashCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTargetAMMFundSize\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTargetDFSize\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fMaxTransferPerConvergencePeriod\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fBrokerCollateralLotSize\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"prevTokenAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"nextTokenAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalSupplyShareToken\",\"type\":\"uint128\"}],\"internalType\":\"structPerpStorage.LiquidityPoolData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_traderAddress\",\"type\":\"address\"}],\"name\":\"getMarginAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"int128\",\"name\":\"fLockedInValueQC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fCashCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fPositionBC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fUnitAccumulatedFundingStart\",\"type\":\"int128\"},{\"internalType\":\"uint64\",\"name\":\"iLastOpenTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"feeTbps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"bytes16\",\"name\":\"positionId\",\"type\":\"bytes16\"}],\"internalType\":\"structPerpStorage.MarginAccount\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128\",\"name\":\"_fCurrentTraderPos\",\"type\":\"int128\"},{\"internalType\":\"bool\",\"name\":\"_isBuy\",\"type\":\"bool\"}],\"name\":\"getMaxSignedOpenTradeSizeForPos\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128[2]\",\"name\":\"_fIndexPrice\",\"type\":\"int128[2]\"}],\"name\":\"getNextLiquidatableTrader\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracleFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4[2]\",\"name\":\"_baseQuote\",\"type\":\"bytes4[2]\"}],\"name\":\"getOraclePrice\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"fPrice\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"}],\"name\":\"getOrderBookAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrderBookFactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"}],\"name\":\"getPerpetual\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"internalType\":\"int32\",\"name\":\"fInitialMarginRate\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"fMaintenanceMarginRate\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"fSigma2\",\"type\":\"int32\"},{\"internalType\":\"uint32\",\"name\":\"iLastFundingTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iLastSettlementPriceUpdateTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iLastPriceJumpTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iLastDefaultFundTransfer\",\"type\":\"uint32\"},{\"internalType\":\"enumPerpStorage.PerpetualState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumAMMPerpLogic.CollateralCurrency\",\"name\":\"eCollateralCurrency\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"minimalSpreadTbps\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"S2BaseCCY\",\"type\":\"bytes4\"},{\"internalType\":\"uint16\",\"name\":\"incentiveSpreadTbps\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"S2QuoteCCY\",\"type\":\"bytes4\"},{\"internalType\":\"uint16\",\"name\":\"jumpSpreadTbps\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"S3BaseCCY\",\"type\":\"bytes4\"},{\"internalType\":\"uint16\",\"name\":\"liquidationPenaltyRateTbps\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"S3QuoteCCY\",\"type\":\"bytes4\"},{\"internalType\":\"int32\",\"name\":\"fSigma3\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"int128\",\"name\":\"fPrice\",\"type\":\"int128\"},{\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"}],\"internalType\":\"structPerpStorage.PriceTimeData\",\"name\":\"currentMarkPremiumRate\",\"type\":\"tuple\"},{\"internalType\":\"int128[2]\",\"name\":\"fDFLambda\",\"type\":\"int128[2]\"},{\"internalType\":\"int128[2]\",\"name\":\"fCurrentAMMExposureEMA\",\"type\":\"int128[2]\"},{\"internalType\":\"int128[2]\",\"name\":\"fAMMTargetDD\",\"type\":\"int128[2]\"},{\"internalType\":\"int128[2]\",\"name\":\"fStressReturnS2\",\"type\":\"int128[2]\"},{\"internalType\":\"int128[2]\",\"name\":\"fStressReturnS3\",\"type\":\"int128[2]\"},{\"internalType\":\"int128\",\"name\":\"premiumRatesEMA\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTargetDFSize\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fFundAllocationWeightCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fOpenInterest\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTargetAMMFundSize\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fCurrentTraderExposureEMA\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fCurrentFundingRate\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fUnitAccumulatedFunding\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fLotSizeBC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fkStar\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fAMMFundCashCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fAMMMinSizeCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fMinimalTraderExposureEMA\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fMinimalAMMExposureEMA\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTotalMarginBalance\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSettlementS3PriceData\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fReferralRebateCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSettlementS2PriceData\",\"type\":\"int128\"},{\"internalType\":\"uint32\",\"name\":\"iLastTargetPoolSizeTime\",\"type\":\"uint32\"},{\"internalType\":\"int32\",\"name\":\"fMarkPriceEMALambda\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"fFundingRateClamp\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"fMaximalTradeSizeBumpUp\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"fRho23\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"fDFCoverNRate\",\"type\":\"int32\"}],\"internalType\":\"structPerpStorage.PerpetualData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getPerpetualCountInPool\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_perpetualIndex\",\"type\":\"uint8\"}],\"name\":\"getPerpetualId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24[]\",\"name\":\"perpetualIds\",\"type\":\"uint24[]\"}],\"name\":\"getPerpetualStaticInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"limitOrderBookAddr\",\"type\":\"address\"},{\"internalType\":\"int32\",\"name\":\"fInitialMarginRate\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"fMaintenanceMarginRate\",\"type\":\"int32\"},{\"internalType\":\"uint8\",\"name\":\"perpetualState\",\"type\":\"uint8\"},{\"internalType\":\"enumAMMPerpLogic.CollateralCurrency\",\"name\":\"collCurrencyType\",\"type\":\"uint8\"},{\"internalType\":\"bytes4\",\"name\":\"S2BaseCCY\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"S2QuoteCCY\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"S3BaseCCY\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"S3QuoteCCY\",\"type\":\"bytes4\"},{\"internalType\":\"int128\",\"name\":\"fLotSizeBC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fReferralRebateCC\",\"type\":\"int128\"},{\"internalType\":\"bytes32[]\",\"name\":\"priceIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"isPyth\",\"type\":\"bool[]\"}],\"internalType\":\"structIPerpetualGetter.PerpetualStaticInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24[]\",\"name\":\"perpetualIds\",\"type\":\"uint24[]\"}],\"name\":\"getPerpetuals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"internalType\":\"int32\",\"name\":\"fInitialMarginRate\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"fMaintenanceMarginRate\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"fSigma2\",\"type\":\"int32\"},{\"internalType\":\"uint32\",\"name\":\"iLastFundingTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iLastSettlementPriceUpdateTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iLastPriceJumpTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iLastDefaultFundTransfer\",\"type\":\"uint32\"},{\"internalType\":\"enumPerpStorage.PerpetualState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumAMMPerpLogic.CollateralCurrency\",\"name\":\"eCollateralCurrency\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"minimalSpreadTbps\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"S2BaseCCY\",\"type\":\"bytes4\"},{\"internalType\":\"uint16\",\"name\":\"incentiveSpreadTbps\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"S2QuoteCCY\",\"type\":\"bytes4\"},{\"internalType\":\"uint16\",\"name\":\"jumpSpreadTbps\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"S3BaseCCY\",\"type\":\"bytes4\"},{\"internalType\":\"uint16\",\"name\":\"liquidationPenaltyRateTbps\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"S3QuoteCCY\",\"type\":\"bytes4\"},{\"internalType\":\"int32\",\"name\":\"fSigma3\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"int128\",\"name\":\"fPrice\",\"type\":\"int128\"},{\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"}],\"internalType\":\"structPerpStorage.PriceTimeData\",\"name\":\"currentMarkPremiumRate\",\"type\":\"tuple\"},{\"internalType\":\"int128[2]\",\"name\":\"fDFLambda\",\"type\":\"int128[2]\"},{\"internalType\":\"int128[2]\",\"name\":\"fCurrentAMMExposureEMA\",\"type\":\"int128[2]\"},{\"internalType\":\"int128[2]\",\"name\":\"fAMMTargetDD\",\"type\":\"int128[2]\"},{\"internalType\":\"int128[2]\",\"name\":\"fStressReturnS2\",\"type\":\"int128[2]\"},{\"internalType\":\"int128[2]\",\"name\":\"fStressReturnS3\",\"type\":\"int128[2]\"},{\"internalType\":\"int128\",\"name\":\"premiumRatesEMA\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTargetDFSize\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fFundAllocationWeightCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fOpenInterest\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTargetAMMFundSize\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fCurrentTraderExposureEMA\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fCurrentFundingRate\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fUnitAccumulatedFunding\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fLotSizeBC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fkStar\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fAMMFundCashCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fAMMMinSizeCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fMinimalTraderExposureEMA\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fMinimalAMMExposureEMA\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTotalMarginBalance\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSettlementS3PriceData\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fReferralRebateCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSettlementS2PriceData\",\"type\":\"int128\"},{\"internalType\":\"uint32\",\"name\":\"iLastTargetPoolSizeTime\",\"type\":\"uint32\"},{\"internalType\":\"int32\",\"name\":\"fMarkPriceEMALambda\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"fFundingRateClamp\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"fMaximalTradeSizeBumpUp\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"fRho23\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"fDFCoverNRate\",\"type\":\"int32\"}],\"internalType\":\"structPerpStorage.PerpetualData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"}],\"name\":\"getPoolIdByPerpetualId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolFromIdx\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_poolToIdx\",\"type\":\"uint8\"}],\"name\":\"getPoolStaticInfo\",\"outputs\":[{\"internalType\":\"uint24[][]\",\"name\":\"\",\"type\":\"uint24[][]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_oracleFactoryAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"}],\"name\":\"getPriceInfo\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getShareTokenFactory\",\"outputs\":[{\"internalType\":\"contractIShareTokenFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getShareTokenPriceD18\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"_fK2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fL1\",\"type\":\"int128\"},{\"components\":[{\"internalType\":\"int128\",\"name\":\"fIndexPriceS2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fIndexPriceS3\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSigma2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSigma3\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fRho23\",\"type\":\"int128\"}],\"internalType\":\"structAMMPerpLogic.MarketVariables\",\"name\":\"_mktVars\",\"type\":\"tuple\"},{\"internalType\":\"int128\",\"name\":\"_fTargetDD\",\"type\":\"int128\"}],\"name\":\"getTargetCollateralM1\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"_fK2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fL1\",\"type\":\"int128\"},{\"components\":[{\"internalType\":\"int128\",\"name\":\"fIndexPriceS2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fIndexPriceS3\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSigma2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSigma3\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fRho23\",\"type\":\"int128\"}],\"internalType\":\"structAMMPerpLogic.MarketVariables\",\"name\":\"_mktVars\",\"type\":\"tuple\"},{\"internalType\":\"int128\",\"name\":\"_fTargetDD\",\"type\":\"int128\"}],\"name\":\"getTargetCollateralM2\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"_fK2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fL1\",\"type\":\"int128\"},{\"components\":[{\"internalType\":\"int128\",\"name\":\"fIndexPriceS2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fIndexPriceS3\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSigma2\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fSigma3\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fRho23\",\"type\":\"int128\"}],\"internalType\":\"structAMMPerpLogic.MarketVariables\",\"name\":\"_mktVars\",\"type\":\"tuple\"},{\"internalType\":\"int128\",\"name\":\"_fTargetDD\",\"type\":\"int128\"}],\"name\":\"getTargetCollateralM3\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_shareAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenAmountToReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_traderAddress\",\"type\":\"address\"},{\"internalType\":\"int128[2]\",\"name\":\"_fIndexPrice\",\"type\":\"int128[2]\"}],\"name\":\"getTraderState\",\"outputs\":[{\"internalType\":\"int128[11]\",\"name\":\"\",\"type\":\"int128[11]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_fromIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numRequests\",\"type\":\"uint256\"}],\"name\":\"getWithdrawRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shareTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"withdrawTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIPerpetualTreasury.WithdrawRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numBlockSinceLastOpen\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"_fLambda\",\"type\":\"int128\"}],\"name\":\"holdingPeriodPenalty\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_iPoolIdx\",\"type\":\"uint8\"},{\"internalType\":\"int128\",\"name\":\"_fAmount\",\"type\":\"int128\"}],\"name\":\"increasePoolCash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_traderAddress\",\"type\":\"address\"}],\"name\":\"isActiveAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_baseCurrency\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"_quoteCurrency\",\"type\":\"bytes4\"}],\"name\":\"isMarketClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"isOrderCanceled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"isOrderExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"}],\"name\":\"isPerpMarketClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isClosed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_traderAddress\",\"type\":\"address\"}],\"name\":\"isTraderMaintenanceMarginSafe\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualIndex\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_liquidatorAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_traderAddr\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"_updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_publishTimes\",\"type\":\"uint64[]\"}],\"name\":\"liquidateByAMM\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"liquidatedAmount\",\"type\":\"int128\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_pauseOn\",\"type\":\"bool\"}],\"name\":\"pauseLiquidityProvision\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"leverageTDR\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executionTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"submittedTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iDeadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fAmount\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fLimitPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTriggerPrice\",\"type\":\"int128\"},{\"internalType\":\"bytes\",\"name\":\"brokerSignature\",\"type\":\"bytes\"}],\"internalType\":\"structIPerpetualOrder.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"preTrade\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_traderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_brokerAddr\",\"type\":\"address\"}],\"name\":\"queryExchangeFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24[]\",\"name\":\"perpetualIds\",\"type\":\"uint24[]\"},{\"internalType\":\"int128[]\",\"name\":\"idxPriceDataPairs\",\"type\":\"int128[]\"}],\"name\":\"queryMidPrices\",\"outputs\":[{\"internalType\":\"int128[]\",\"name\":\"\",\"type\":\"int128[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128\",\"name\":\"_fTradeAmountBC\",\"type\":\"int128\"},{\"internalType\":\"int128[2]\",\"name\":\"_fIndexPrice\",\"type\":\"int128[2]\"}],\"name\":\"queryPerpetualPrice\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"leverageTDR\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executionTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"submittedTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iDeadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fAmount\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fLimitPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTriggerPrice\",\"type\":\"int128\"},{\"internalType\":\"bytes\",\"name\":\"brokerSignature\",\"type\":\"bytes\"}],\"internalType\":\"structIPerpetualOrder.Order\",\"name\":\"_order\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"_feeTbps\",\"type\":\"uint16\"}],\"name\":\"rebatePostingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_traderAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"_fAmountToWithdraw\",\"type\":\"int128\"}],\"name\":\"reduceMarginCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"_fDeltaPos\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fTreasuryFeeRate\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fPnLPartRate\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fReferralRebate\",\"type\":\"int128\"},{\"internalType\":\"address\",\"name\":\"_referrerAddr\",\"type\":\"address\"}],\"name\":\"relativeFeeToCCAmount\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_liqPoolID\",\"type\":\"uint8\"}],\"name\":\"runLiquidityPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_AMMPerpLogic\",\"type\":\"address\"}],\"name\":\"setAMMPerpLogic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_delay\",\"type\":\"uint8\"}],\"name\":\"setBlockDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"_feesTbps\",\"type\":\"uint16[]\"}],\"name\":\"setBrokerTiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"_feesTbps\",\"type\":\"uint16[]\"}],\"name\":\"setBrokerVolumeTiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"}],\"name\":\"setEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"_designations\",\"type\":\"uint32[]\"},{\"internalType\":\"uint16[]\",\"name\":\"_fees\",\"type\":\"uint16[]\"}],\"name\":\"setFeesForDesignation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"}],\"name\":\"setInitialFundAllocationWeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_brokerAddr\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_feeTbps\",\"type\":\"uint16\"}],\"name\":\"setInitialVolumeForFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"}],\"name\":\"setNormalState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleFactory\",\"type\":\"address\"}],\"name\":\"setOracleFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_oracleAddr\",\"type\":\"address\"}],\"name\":\"setOracleFactoryForPerpetual\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_orderBookFactory\",\"type\":\"address\"}],\"name\":\"setOrderBookFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128[7]\",\"name\":\"_baseParams\",\"type\":\"int128[7]\"}],\"name\":\"setPerpetualBaseParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"bytes4[2]\",\"name\":\"_baseQuoteS2\",\"type\":\"bytes4[2]\"},{\"internalType\":\"bytes4[2]\",\"name\":\"_baseQuoteS3\",\"type\":\"bytes4[2]\"}],\"name\":\"setPerpetualOracles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"string\",\"name\":\"_varName\",\"type\":\"string\"},{\"internalType\":\"int128\",\"name\":\"_value\",\"type\":\"int128\"}],\"name\":\"setPerpetualParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"int128\",\"name\":\"_value1\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_value2\",\"type\":\"int128\"}],\"name\":\"setPerpetualParamPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_shareTokenFactory\",\"type\":\"address\"}],\"name\":\"setPerpetualPoolFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128[5]\",\"name\":\"_underlyingRiskParams\",\"type\":\"int128[5]\"},{\"internalType\":\"int128[13]\",\"name\":\"_defaultFundRiskParams\",\"type\":\"int128[13]\"}],\"name\":\"setPerpetualRiskParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"int128\",\"name\":\"_value\",\"type\":\"int128\"}],\"name\":\"setPoolParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"_feesTbps\",\"type\":\"uint16[]\"}],\"name\":\"setTraderTiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"_feesTbps\",\"type\":\"uint16[]\"}],\"name\":\"setTraderVolumeTiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"setUtilityTokenAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualID\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_traderAddr\",\"type\":\"address\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_id\",\"type\":\"uint8\"}],\"name\":\"settleNextTraderInPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"}],\"name\":\"splitProtocolFee\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"}],\"name\":\"togglePerpEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"leverageTDR\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"brokerFeeTbps\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"traderAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"executionTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"brokerAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"submittedTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"iDeadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fAmount\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fLimitPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTriggerPrice\",\"type\":\"int128\"},{\"internalType\":\"bytes\",\"name\":\"brokerSignature\",\"type\":\"bytes\"}],\"internalType\":\"structIPerpetualOrder.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"tradeViaOrderBook\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_transferToAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_lots\",\"type\":\"uint32\"}],\"name\":\"transferBrokerLots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_transferToAddr\",\"type\":\"address\"}],\"name\":\"transferBrokerOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"int128\",\"name\":\"_fAmount\",\"type\":\"int128\"}],\"name\":\"transferEarningsToTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferValueToTreasury\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128\",\"name\":\"fTargetFundSize\",\"type\":\"int128\"}],\"name\":\"updateAMMTargetFundSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"}],\"name\":\"updateDefaultFundTargetSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_iPoolIndex\",\"type\":\"uint8\"}],\"name\":\"updateDefaultFundTargetSizeRandom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId0\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId1\",\"type\":\"uint24\"}],\"name\":\"updateFundingAndPricesAfter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId0\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId1\",\"type\":\"uint24\"}],\"name\":\"updateFundingAndPricesBefore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_perpetualId\",\"type\":\"uint24\"},{\"internalType\":\"bytes[]\",\"name\":\"_updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_publishTimes\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxAcceptableFeedAge\",\"type\":\"uint256\"}],\"name\":\"updatePriceFeeds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_traderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_brokerAddr\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"_tradeAmountBC\",\"type\":\"int128\"}],\"name\":\"updateVolumeEMAOnNewTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"int128\",\"name\":\"_fMarkPrice\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_fTriggerPrice\",\"type\":\"int128\"}],\"name\":\"validateStopPrice\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_jumpTbps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_MinimalSpreadTbps\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_numBlockSinceJump\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"_fLambda\",\"type\":\"int128\"}],\"name\":\"volatilitySpread\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"int128\",\"name\":\"_fAmount\",\"type\":\"int128\"},{\"internalType\":\"bytes[]\",\"name\":\"_updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_publishTimes\",\"type\":\"uint64[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"bytes[]\",\"name\":\"_updateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_publishTimes\",\"type\":\"uint64[]\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_iPerpetualId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_traderAddr\",\"type\":\"address\"}],\"name\":\"withdrawDepositFromMarginAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"int128\",\"name\":\"_fAmount\",\"type\":\"int128\"}],\"name\":\"withdrawFromDefaultFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_iPoolIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_shareAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPerpetualManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IPerpetualManagerMetaData.ABI instead.
var IPerpetualManagerABI = IPerpetualManagerMetaData.ABI

// IPerpetualManager is an auto generated Go binding around an Ethereum contract.
type IPerpetualManager struct {
	IPerpetualManagerCaller     // Read-only binding to the contract
	IPerpetualManagerTransactor // Write-only binding to the contract
	IPerpetualManagerFilterer   // Log filterer for contract events
}

// IPerpetualManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPerpetualManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPerpetualManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPerpetualManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPerpetualManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPerpetualManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPerpetualManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPerpetualManagerSession struct {
	Contract     *IPerpetualManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IPerpetualManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPerpetualManagerCallerSession struct {
	Contract *IPerpetualManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IPerpetualManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPerpetualManagerTransactorSession struct {
	Contract     *IPerpetualManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IPerpetualManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPerpetualManagerRaw struct {
	Contract *IPerpetualManager // Generic contract binding to access the raw methods on
}

// IPerpetualManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPerpetualManagerCallerRaw struct {
	Contract *IPerpetualManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IPerpetualManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPerpetualManagerTransactorRaw struct {
	Contract *IPerpetualManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPerpetualManager creates a new instance of IPerpetualManager, bound to a specific deployed contract.
func NewIPerpetualManager(address common.Address, backend bind.ContractBackend) (*IPerpetualManager, error) {
	contract, err := bindIPerpetualManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManager{IPerpetualManagerCaller: IPerpetualManagerCaller{contract: contract}, IPerpetualManagerTransactor: IPerpetualManagerTransactor{contract: contract}, IPerpetualManagerFilterer: IPerpetualManagerFilterer{contract: contract}}, nil
}

// NewIPerpetualManagerCaller creates a new read-only instance of IPerpetualManager, bound to a specific deployed contract.
func NewIPerpetualManagerCaller(address common.Address, caller bind.ContractCaller) (*IPerpetualManagerCaller, error) {
	contract, err := bindIPerpetualManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerCaller{contract: contract}, nil
}

// NewIPerpetualManagerTransactor creates a new write-only instance of IPerpetualManager, bound to a specific deployed contract.
func NewIPerpetualManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IPerpetualManagerTransactor, error) {
	contract, err := bindIPerpetualManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerTransactor{contract: contract}, nil
}

// NewIPerpetualManagerFilterer creates a new log filterer instance of IPerpetualManager, bound to a specific deployed contract.
func NewIPerpetualManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IPerpetualManagerFilterer, error) {
	contract, err := bindIPerpetualManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerFilterer{contract: contract}, nil
}

// bindIPerpetualManager binds a generic wrapper to an already deployed contract.
func bindIPerpetualManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPerpetualManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPerpetualManager *IPerpetualManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPerpetualManager.Contract.IPerpetualManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPerpetualManager *IPerpetualManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.IPerpetualManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPerpetualManager *IPerpetualManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.IPerpetualManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPerpetualManager *IPerpetualManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPerpetualManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPerpetualManager *IPerpetualManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPerpetualManager *IPerpetualManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.contract.Transact(opts, method, params...)
}

// CalculateBoundedSlippage is a free data retrieval call binding the contract method 0xd3a243ff.
//
// Solidity: function calculateBoundedSlippage((int128,int128,int128,int128,int128,int128) _ammVars, int128 _fTradeAmount) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) CalculateBoundedSlippage(opts *bind.CallOpts, _ammVars AMMPerpLogicAMMVariables, _fTradeAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "calculateBoundedSlippage", _ammVars, _fTradeAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateBoundedSlippage is a free data retrieval call binding the contract method 0xd3a243ff.
//
// Solidity: function calculateBoundedSlippage((int128,int128,int128,int128,int128,int128) _ammVars, int128 _fTradeAmount) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) CalculateBoundedSlippage(_ammVars AMMPerpLogicAMMVariables, _fTradeAmount *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.CalculateBoundedSlippage(&_IPerpetualManager.CallOpts, _ammVars, _fTradeAmount)
}

// CalculateBoundedSlippage is a free data retrieval call binding the contract method 0xd3a243ff.
//
// Solidity: function calculateBoundedSlippage((int128,int128,int128,int128,int128,int128) _ammVars, int128 _fTradeAmount) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) CalculateBoundedSlippage(_ammVars AMMPerpLogicAMMVariables, _fTradeAmount *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.CalculateBoundedSlippage(&_IPerpetualManager.CallOpts, _ammVars, _fTradeAmount)
}

// CalculateDefaultFundSize is a free data retrieval call binding the contract method 0x68021b7f.
//
// Solidity: function calculateDefaultFundSize(int128[2] _fK2AMM, int128 _fk2Trader, int128 _fCoverN, int128[2] fStressRet2, int128[2] fStressRet3, int128[2] fIndexPrices, uint8 _eCCY) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) CalculateDefaultFundSize(opts *bind.CallOpts, _fK2AMM [2]*big.Int, _fk2Trader *big.Int, _fCoverN *big.Int, fStressRet2 [2]*big.Int, fStressRet3 [2]*big.Int, fIndexPrices [2]*big.Int, _eCCY uint8) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "calculateDefaultFundSize", _fK2AMM, _fk2Trader, _fCoverN, fStressRet2, fStressRet3, fIndexPrices, _eCCY)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateDefaultFundSize is a free data retrieval call binding the contract method 0x68021b7f.
//
// Solidity: function calculateDefaultFundSize(int128[2] _fK2AMM, int128 _fk2Trader, int128 _fCoverN, int128[2] fStressRet2, int128[2] fStressRet3, int128[2] fIndexPrices, uint8 _eCCY) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) CalculateDefaultFundSize(_fK2AMM [2]*big.Int, _fk2Trader *big.Int, _fCoverN *big.Int, fStressRet2 [2]*big.Int, fStressRet3 [2]*big.Int, fIndexPrices [2]*big.Int, _eCCY uint8) (*big.Int, error) {
	return _IPerpetualManager.Contract.CalculateDefaultFundSize(&_IPerpetualManager.CallOpts, _fK2AMM, _fk2Trader, _fCoverN, fStressRet2, fStressRet3, fIndexPrices, _eCCY)
}

// CalculateDefaultFundSize is a free data retrieval call binding the contract method 0x68021b7f.
//
// Solidity: function calculateDefaultFundSize(int128[2] _fK2AMM, int128 _fk2Trader, int128 _fCoverN, int128[2] fStressRet2, int128[2] fStressRet3, int128[2] fIndexPrices, uint8 _eCCY) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) CalculateDefaultFundSize(_fK2AMM [2]*big.Int, _fk2Trader *big.Int, _fCoverN *big.Int, fStressRet2 [2]*big.Int, fStressRet3 [2]*big.Int, fIndexPrices [2]*big.Int, _eCCY uint8) (*big.Int, error) {
	return _IPerpetualManager.Contract.CalculateDefaultFundSize(&_IPerpetualManager.CallOpts, _fK2AMM, _fk2Trader, _fCoverN, fStressRet2, fStressRet3, fIndexPrices, _eCCY)
}

// CalculatePerpetualPrice is a free data retrieval call binding the contract method 0x7655d385.
//
// Solidity: function calculatePerpetualPrice((int128,int128,int128,int128,int128,int128) _ammVars, (int128,int128,int128,int128,int128) _mktVars, int128 _fTradeAmount, int128 _fBidAskSpread, int128 _fIncentiveSpread) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) CalculatePerpetualPrice(opts *bind.CallOpts, _ammVars AMMPerpLogicAMMVariables, _mktVars AMMPerpLogicMarketVariables, _fTradeAmount *big.Int, _fBidAskSpread *big.Int, _fIncentiveSpread *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "calculatePerpetualPrice", _ammVars, _mktVars, _fTradeAmount, _fBidAskSpread, _fIncentiveSpread)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePerpetualPrice is a free data retrieval call binding the contract method 0x7655d385.
//
// Solidity: function calculatePerpetualPrice((int128,int128,int128,int128,int128,int128) _ammVars, (int128,int128,int128,int128,int128) _mktVars, int128 _fTradeAmount, int128 _fBidAskSpread, int128 _fIncentiveSpread) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) CalculatePerpetualPrice(_ammVars AMMPerpLogicAMMVariables, _mktVars AMMPerpLogicMarketVariables, _fTradeAmount *big.Int, _fBidAskSpread *big.Int, _fIncentiveSpread *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.CalculatePerpetualPrice(&_IPerpetualManager.CallOpts, _ammVars, _mktVars, _fTradeAmount, _fBidAskSpread, _fIncentiveSpread)
}

// CalculatePerpetualPrice is a free data retrieval call binding the contract method 0x7655d385.
//
// Solidity: function calculatePerpetualPrice((int128,int128,int128,int128,int128,int128) _ammVars, (int128,int128,int128,int128,int128) _mktVars, int128 _fTradeAmount, int128 _fBidAskSpread, int128 _fIncentiveSpread) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) CalculatePerpetualPrice(_ammVars AMMPerpLogicAMMVariables, _mktVars AMMPerpLogicMarketVariables, _fTradeAmount *big.Int, _fBidAskSpread *big.Int, _fIncentiveSpread *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.CalculatePerpetualPrice(&_IPerpetualManager.CallOpts, _ammVars, _mktVars, _fTradeAmount, _fBidAskSpread, _fIncentiveSpread)
}

// CalculateRiskNeutralPD is a free data retrieval call binding the contract method 0xa4fb7972.
//
// Solidity: function calculateRiskNeutralPD((int128,int128,int128,int128,int128,int128) _ammVars, (int128,int128,int128,int128,int128) _mktVars, int128 _fTradeAmount, bool _withCDF) view returns(int128, int128)
func (_IPerpetualManager *IPerpetualManagerCaller) CalculateRiskNeutralPD(opts *bind.CallOpts, _ammVars AMMPerpLogicAMMVariables, _mktVars AMMPerpLogicMarketVariables, _fTradeAmount *big.Int, _withCDF bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "calculateRiskNeutralPD", _ammVars, _mktVars, _fTradeAmount, _withCDF)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// CalculateRiskNeutralPD is a free data retrieval call binding the contract method 0xa4fb7972.
//
// Solidity: function calculateRiskNeutralPD((int128,int128,int128,int128,int128,int128) _ammVars, (int128,int128,int128,int128,int128) _mktVars, int128 _fTradeAmount, bool _withCDF) view returns(int128, int128)
func (_IPerpetualManager *IPerpetualManagerSession) CalculateRiskNeutralPD(_ammVars AMMPerpLogicAMMVariables, _mktVars AMMPerpLogicMarketVariables, _fTradeAmount *big.Int, _withCDF bool) (*big.Int, *big.Int, error) {
	return _IPerpetualManager.Contract.CalculateRiskNeutralPD(&_IPerpetualManager.CallOpts, _ammVars, _mktVars, _fTradeAmount, _withCDF)
}

// CalculateRiskNeutralPD is a free data retrieval call binding the contract method 0xa4fb7972.
//
// Solidity: function calculateRiskNeutralPD((int128,int128,int128,int128,int128,int128) _ammVars, (int128,int128,int128,int128,int128) _mktVars, int128 _fTradeAmount, bool _withCDF) view returns(int128, int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) CalculateRiskNeutralPD(_ammVars AMMPerpLogicAMMVariables, _mktVars AMMPerpLogicMarketVariables, _fTradeAmount *big.Int, _withCDF bool) (*big.Int, *big.Int, error) {
	return _IPerpetualManager.Contract.CalculateRiskNeutralPD(&_IPerpetualManager.CallOpts, _ammVars, _mktVars, _fTradeAmount, _withCDF)
}

// CountActivePerpAccounts is a free data retrieval call binding the contract method 0xda326ae6.
//
// Solidity: function countActivePerpAccounts(uint24 _perpetualId) view returns(uint256)
func (_IPerpetualManager *IPerpetualManagerCaller) CountActivePerpAccounts(opts *bind.CallOpts, _perpetualId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "countActivePerpAccounts", _perpetualId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountActivePerpAccounts is a free data retrieval call binding the contract method 0xda326ae6.
//
// Solidity: function countActivePerpAccounts(uint24 _perpetualId) view returns(uint256)
func (_IPerpetualManager *IPerpetualManagerSession) CountActivePerpAccounts(_perpetualId *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.CountActivePerpAccounts(&_IPerpetualManager.CallOpts, _perpetualId)
}

// CountActivePerpAccounts is a free data retrieval call binding the contract method 0xda326ae6.
//
// Solidity: function countActivePerpAccounts(uint24 _perpetualId) view returns(uint256)
func (_IPerpetualManager *IPerpetualManagerCallerSession) CountActivePerpAccounts(_perpetualId *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.CountActivePerpAccounts(&_IPerpetualManager.CallOpts, _perpetualId)
}

// DetermineExchangeFee is a free data retrieval call binding the contract method 0xf4ec8aa9.
//
// Solidity: function determineExchangeFee((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCaller) DetermineExchangeFee(opts *bind.CallOpts, _order IPerpetualOrderOrder) (uint16, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "determineExchangeFee", _order)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DetermineExchangeFee is a free data retrieval call binding the contract method 0xf4ec8aa9.
//
// Solidity: function determineExchangeFee((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerSession) DetermineExchangeFee(_order IPerpetualOrderOrder) (uint16, error) {
	return _IPerpetualManager.Contract.DetermineExchangeFee(&_IPerpetualManager.CallOpts, _order)
}

// DetermineExchangeFee is a free data retrieval call binding the contract method 0xf4ec8aa9.
//
// Solidity: function determineExchangeFee((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCallerSession) DetermineExchangeFee(_order IPerpetualOrderOrder) (uint16, error) {
	return _IPerpetualManager.Contract.DetermineExchangeFee(&_IPerpetualManager.CallOpts, _order)
}

// GetAMMPerpLogic is a free data retrieval call binding the contract method 0xf7a68dcd.
//
// Solidity: function getAMMPerpLogic() view returns(address)
func (_IPerpetualManager *IPerpetualManagerCaller) GetAMMPerpLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getAMMPerpLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAMMPerpLogic is a free data retrieval call binding the contract method 0xf7a68dcd.
//
// Solidity: function getAMMPerpLogic() view returns(address)
func (_IPerpetualManager *IPerpetualManagerSession) GetAMMPerpLogic() (common.Address, error) {
	return _IPerpetualManager.Contract.GetAMMPerpLogic(&_IPerpetualManager.CallOpts)
}

// GetAMMPerpLogic is a free data retrieval call binding the contract method 0xf7a68dcd.
//
// Solidity: function getAMMPerpLogic() view returns(address)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetAMMPerpLogic() (common.Address, error) {
	return _IPerpetualManager.Contract.GetAMMPerpLogic(&_IPerpetualManager.CallOpts)
}

// GetAMMState is a free data retrieval call binding the contract method 0x1249b910.
//
// Solidity: function getAMMState(uint24 _perpetualId, int128[2] _fIndexPrice) view returns(int128[15])
func (_IPerpetualManager *IPerpetualManagerCaller) GetAMMState(opts *bind.CallOpts, _perpetualId *big.Int, _fIndexPrice [2]*big.Int) ([15]*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getAMMState", _perpetualId, _fIndexPrice)

	if err != nil {
		return *new([15]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([15]*big.Int)).(*[15]*big.Int)

	return out0, err

}

// GetAMMState is a free data retrieval call binding the contract method 0x1249b910.
//
// Solidity: function getAMMState(uint24 _perpetualId, int128[2] _fIndexPrice) view returns(int128[15])
func (_IPerpetualManager *IPerpetualManagerSession) GetAMMState(_perpetualId *big.Int, _fIndexPrice [2]*big.Int) ([15]*big.Int, error) {
	return _IPerpetualManager.Contract.GetAMMState(&_IPerpetualManager.CallOpts, _perpetualId, _fIndexPrice)
}

// GetAMMState is a free data retrieval call binding the contract method 0x1249b910.
//
// Solidity: function getAMMState(uint24 _perpetualId, int128[2] _fIndexPrice) view returns(int128[15])
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetAMMState(_perpetualId *big.Int, _fIndexPrice [2]*big.Int) ([15]*big.Int, error) {
	return _IPerpetualManager.Contract.GetAMMState(&_IPerpetualManager.CallOpts, _perpetualId, _fIndexPrice)
}

// GetActivePerpAccounts is a free data retrieval call binding the contract method 0x5b23c303.
//
// Solidity: function getActivePerpAccounts(uint24 _perpetualId) view returns(address[])
func (_IPerpetualManager *IPerpetualManagerCaller) GetActivePerpAccounts(opts *bind.CallOpts, _perpetualId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getActivePerpAccounts", _perpetualId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActivePerpAccounts is a free data retrieval call binding the contract method 0x5b23c303.
//
// Solidity: function getActivePerpAccounts(uint24 _perpetualId) view returns(address[])
func (_IPerpetualManager *IPerpetualManagerSession) GetActivePerpAccounts(_perpetualId *big.Int) ([]common.Address, error) {
	return _IPerpetualManager.Contract.GetActivePerpAccounts(&_IPerpetualManager.CallOpts, _perpetualId)
}

// GetActivePerpAccounts is a free data retrieval call binding the contract method 0x5b23c303.
//
// Solidity: function getActivePerpAccounts(uint24 _perpetualId) view returns(address[])
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetActivePerpAccounts(_perpetualId *big.Int) ([]common.Address, error) {
	return _IPerpetualManager.Contract.GetActivePerpAccounts(&_IPerpetualManager.CallOpts, _perpetualId)
}

// GetActivePerpAccountsByChunks is a free data retrieval call binding the contract method 0xb276de70.
//
// Solidity: function getActivePerpAccountsByChunks(uint24 _perpetualId, uint256 _from, uint256 _to) view returns(address[])
func (_IPerpetualManager *IPerpetualManagerCaller) GetActivePerpAccountsByChunks(opts *bind.CallOpts, _perpetualId *big.Int, _from *big.Int, _to *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getActivePerpAccountsByChunks", _perpetualId, _from, _to)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActivePerpAccountsByChunks is a free data retrieval call binding the contract method 0xb276de70.
//
// Solidity: function getActivePerpAccountsByChunks(uint24 _perpetualId, uint256 _from, uint256 _to) view returns(address[])
func (_IPerpetualManager *IPerpetualManagerSession) GetActivePerpAccountsByChunks(_perpetualId *big.Int, _from *big.Int, _to *big.Int) ([]common.Address, error) {
	return _IPerpetualManager.Contract.GetActivePerpAccountsByChunks(&_IPerpetualManager.CallOpts, _perpetualId, _from, _to)
}

// GetActivePerpAccountsByChunks is a free data retrieval call binding the contract method 0xb276de70.
//
// Solidity: function getActivePerpAccountsByChunks(uint24 _perpetualId, uint256 _from, uint256 _to) view returns(address[])
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetActivePerpAccountsByChunks(_perpetualId *big.Int, _from *big.Int, _to *big.Int) ([]common.Address, error) {
	return _IPerpetualManager.Contract.GetActivePerpAccountsByChunks(&_IPerpetualManager.CallOpts, _perpetualId, _from, _to)
}

// GetBrokerDesignation is a free data retrieval call binding the contract method 0x18b097bd.
//
// Solidity: function getBrokerDesignation(uint8 _poolId, address _brokerAddr) view returns(uint32)
func (_IPerpetualManager *IPerpetualManagerCaller) GetBrokerDesignation(opts *bind.CallOpts, _poolId uint8, _brokerAddr common.Address) (uint32, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getBrokerDesignation", _poolId, _brokerAddr)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetBrokerDesignation is a free data retrieval call binding the contract method 0x18b097bd.
//
// Solidity: function getBrokerDesignation(uint8 _poolId, address _brokerAddr) view returns(uint32)
func (_IPerpetualManager *IPerpetualManagerSession) GetBrokerDesignation(_poolId uint8, _brokerAddr common.Address) (uint32, error) {
	return _IPerpetualManager.Contract.GetBrokerDesignation(&_IPerpetualManager.CallOpts, _poolId, _brokerAddr)
}

// GetBrokerDesignation is a free data retrieval call binding the contract method 0x18b097bd.
//
// Solidity: function getBrokerDesignation(uint8 _poolId, address _brokerAddr) view returns(uint32)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetBrokerDesignation(_poolId uint8, _brokerAddr common.Address) (uint32, error) {
	return _IPerpetualManager.Contract.GetBrokerDesignation(&_IPerpetualManager.CallOpts, _poolId, _brokerAddr)
}

// GetBrokerInducedFee is a free data retrieval call binding the contract method 0xab633c4a.
//
// Solidity: function getBrokerInducedFee(uint8 _poolId, address _brokerAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCaller) GetBrokerInducedFee(opts *bind.CallOpts, _poolId uint8, _brokerAddr common.Address) (uint16, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getBrokerInducedFee", _poolId, _brokerAddr)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetBrokerInducedFee is a free data retrieval call binding the contract method 0xab633c4a.
//
// Solidity: function getBrokerInducedFee(uint8 _poolId, address _brokerAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerSession) GetBrokerInducedFee(_poolId uint8, _brokerAddr common.Address) (uint16, error) {
	return _IPerpetualManager.Contract.GetBrokerInducedFee(&_IPerpetualManager.CallOpts, _poolId, _brokerAddr)
}

// GetBrokerInducedFee is a free data retrieval call binding the contract method 0xab633c4a.
//
// Solidity: function getBrokerInducedFee(uint8 _poolId, address _brokerAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetBrokerInducedFee(_poolId uint8, _brokerAddr common.Address) (uint16, error) {
	return _IPerpetualManager.Contract.GetBrokerInducedFee(&_IPerpetualManager.CallOpts, _poolId, _brokerAddr)
}

// GetCollateralTokenAmountForPricing is a free data retrieval call binding the contract method 0x71fff802.
//
// Solidity: function getCollateralTokenAmountForPricing(uint8 _poolId) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) GetCollateralTokenAmountForPricing(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getCollateralTokenAmountForPricing", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralTokenAmountForPricing is a free data retrieval call binding the contract method 0x71fff802.
//
// Solidity: function getCollateralTokenAmountForPricing(uint8 _poolId) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) GetCollateralTokenAmountForPricing(_poolId uint8) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetCollateralTokenAmountForPricing(&_IPerpetualManager.CallOpts, _poolId)
}

// GetCollateralTokenAmountForPricing is a free data retrieval call binding the contract method 0x71fff802.
//
// Solidity: function getCollateralTokenAmountForPricing(uint8 _poolId) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetCollateralTokenAmountForPricing(_poolId uint8) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetCollateralTokenAmountForPricing(&_IPerpetualManager.CallOpts, _poolId)
}

// GetCurrentBrokerVolume is a free data retrieval call binding the contract method 0xd6e4394b.
//
// Solidity: function getCurrentBrokerVolume(uint8 _poolId, address _brokerAddr) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) GetCurrentBrokerVolume(opts *bind.CallOpts, _poolId uint8, _brokerAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getCurrentBrokerVolume", _poolId, _brokerAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBrokerVolume is a free data retrieval call binding the contract method 0xd6e4394b.
//
// Solidity: function getCurrentBrokerVolume(uint8 _poolId, address _brokerAddr) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) GetCurrentBrokerVolume(_poolId uint8, _brokerAddr common.Address) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetCurrentBrokerVolume(&_IPerpetualManager.CallOpts, _poolId, _brokerAddr)
}

// GetCurrentBrokerVolume is a free data retrieval call binding the contract method 0xd6e4394b.
//
// Solidity: function getCurrentBrokerVolume(uint8 _poolId, address _brokerAddr) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetCurrentBrokerVolume(_poolId uint8, _brokerAddr common.Address) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetCurrentBrokerVolume(&_IPerpetualManager.CallOpts, _poolId, _brokerAddr)
}

// GetCurrentTraderVolume is a free data retrieval call binding the contract method 0xa46281ec.
//
// Solidity: function getCurrentTraderVolume(uint8 _poolId, address _traderAddr) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) GetCurrentTraderVolume(opts *bind.CallOpts, _poolId uint8, _traderAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getCurrentTraderVolume", _poolId, _traderAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTraderVolume is a free data retrieval call binding the contract method 0xa46281ec.
//
// Solidity: function getCurrentTraderVolume(uint8 _poolId, address _traderAddr) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) GetCurrentTraderVolume(_poolId uint8, _traderAddr common.Address) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetCurrentTraderVolume(&_IPerpetualManager.CallOpts, _poolId, _traderAddr)
}

// GetCurrentTraderVolume is a free data retrieval call binding the contract method 0xa46281ec.
//
// Solidity: function getCurrentTraderVolume(uint8 _poolId, address _traderAddr) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetCurrentTraderVolume(_poolId uint8, _traderAddr common.Address) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetCurrentTraderVolume(&_IPerpetualManager.CallOpts, _poolId, _traderAddr)
}

// GetDepositAmountForLvgPosition is a free data retrieval call binding the contract method 0x20f1b336.
//
// Solidity: function getDepositAmountForLvgPosition(int128 _fPosition0, int128 _fBalance0, int128 _fTradeAmount, int128 _fTargetLeverage, int128 _fPrice, int128 _fS2Mark, int128 _fS3) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) GetDepositAmountForLvgPosition(opts *bind.CallOpts, _fPosition0 *big.Int, _fBalance0 *big.Int, _fTradeAmount *big.Int, _fTargetLeverage *big.Int, _fPrice *big.Int, _fS2Mark *big.Int, _fS3 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getDepositAmountForLvgPosition", _fPosition0, _fBalance0, _fTradeAmount, _fTargetLeverage, _fPrice, _fS2Mark, _fS3)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositAmountForLvgPosition is a free data retrieval call binding the contract method 0x20f1b336.
//
// Solidity: function getDepositAmountForLvgPosition(int128 _fPosition0, int128 _fBalance0, int128 _fTradeAmount, int128 _fTargetLeverage, int128 _fPrice, int128 _fS2Mark, int128 _fS3) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) GetDepositAmountForLvgPosition(_fPosition0 *big.Int, _fBalance0 *big.Int, _fTradeAmount *big.Int, _fTargetLeverage *big.Int, _fPrice *big.Int, _fS2Mark *big.Int, _fS3 *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetDepositAmountForLvgPosition(&_IPerpetualManager.CallOpts, _fPosition0, _fBalance0, _fTradeAmount, _fTargetLeverage, _fPrice, _fS2Mark, _fS3)
}

// GetDepositAmountForLvgPosition is a free data retrieval call binding the contract method 0x20f1b336.
//
// Solidity: function getDepositAmountForLvgPosition(int128 _fPosition0, int128 _fBalance0, int128 _fTradeAmount, int128 _fTargetLeverage, int128 _fPrice, int128 _fS2Mark, int128 _fS3) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetDepositAmountForLvgPosition(_fPosition0 *big.Int, _fBalance0 *big.Int, _fTradeAmount *big.Int, _fTargetLeverage *big.Int, _fPrice *big.Int, _fS2Mark *big.Int, _fS3 *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetDepositAmountForLvgPosition(&_IPerpetualManager.CallOpts, _fPosition0, _fBalance0, _fTradeAmount, _fTargetLeverage, _fPrice, _fS2Mark, _fS3)
}

// GetFeeForBrokerDesignation is a free data retrieval call binding the contract method 0x530232fe.
//
// Solidity: function getFeeForBrokerDesignation(uint32 _brokerDesignation) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCaller) GetFeeForBrokerDesignation(opts *bind.CallOpts, _brokerDesignation uint32) (uint16, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getFeeForBrokerDesignation", _brokerDesignation)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFeeForBrokerDesignation is a free data retrieval call binding the contract method 0x530232fe.
//
// Solidity: function getFeeForBrokerDesignation(uint32 _brokerDesignation) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerSession) GetFeeForBrokerDesignation(_brokerDesignation uint32) (uint16, error) {
	return _IPerpetualManager.Contract.GetFeeForBrokerDesignation(&_IPerpetualManager.CallOpts, _brokerDesignation)
}

// GetFeeForBrokerDesignation is a free data retrieval call binding the contract method 0x530232fe.
//
// Solidity: function getFeeForBrokerDesignation(uint32 _brokerDesignation) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetFeeForBrokerDesignation(_brokerDesignation uint32) (uint16, error) {
	return _IPerpetualManager.Contract.GetFeeForBrokerDesignation(&_IPerpetualManager.CallOpts, _brokerDesignation)
}

// GetFeeForBrokerStake is a free data retrieval call binding the contract method 0x25d758b3.
//
// Solidity: function getFeeForBrokerStake(address brokerAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCaller) GetFeeForBrokerStake(opts *bind.CallOpts, brokerAddr common.Address) (uint16, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getFeeForBrokerStake", brokerAddr)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFeeForBrokerStake is a free data retrieval call binding the contract method 0x25d758b3.
//
// Solidity: function getFeeForBrokerStake(address brokerAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerSession) GetFeeForBrokerStake(brokerAddr common.Address) (uint16, error) {
	return _IPerpetualManager.Contract.GetFeeForBrokerStake(&_IPerpetualManager.CallOpts, brokerAddr)
}

// GetFeeForBrokerStake is a free data retrieval call binding the contract method 0x25d758b3.
//
// Solidity: function getFeeForBrokerStake(address brokerAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetFeeForBrokerStake(brokerAddr common.Address) (uint16, error) {
	return _IPerpetualManager.Contract.GetFeeForBrokerStake(&_IPerpetualManager.CallOpts, brokerAddr)
}

// GetFeeForBrokerVolume is a free data retrieval call binding the contract method 0xb6efc782.
//
// Solidity: function getFeeForBrokerVolume(uint8 _poolId, address _brokerAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCaller) GetFeeForBrokerVolume(opts *bind.CallOpts, _poolId uint8, _brokerAddr common.Address) (uint16, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getFeeForBrokerVolume", _poolId, _brokerAddr)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFeeForBrokerVolume is a free data retrieval call binding the contract method 0xb6efc782.
//
// Solidity: function getFeeForBrokerVolume(uint8 _poolId, address _brokerAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerSession) GetFeeForBrokerVolume(_poolId uint8, _brokerAddr common.Address) (uint16, error) {
	return _IPerpetualManager.Contract.GetFeeForBrokerVolume(&_IPerpetualManager.CallOpts, _poolId, _brokerAddr)
}

// GetFeeForBrokerVolume is a free data retrieval call binding the contract method 0xb6efc782.
//
// Solidity: function getFeeForBrokerVolume(uint8 _poolId, address _brokerAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetFeeForBrokerVolume(_poolId uint8, _brokerAddr common.Address) (uint16, error) {
	return _IPerpetualManager.Contract.GetFeeForBrokerVolume(&_IPerpetualManager.CallOpts, _poolId, _brokerAddr)
}

// GetFeeForTraderStake is a free data retrieval call binding the contract method 0xbf7b6bb5.
//
// Solidity: function getFeeForTraderStake(address traderAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCaller) GetFeeForTraderStake(opts *bind.CallOpts, traderAddr common.Address) (uint16, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getFeeForTraderStake", traderAddr)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFeeForTraderStake is a free data retrieval call binding the contract method 0xbf7b6bb5.
//
// Solidity: function getFeeForTraderStake(address traderAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerSession) GetFeeForTraderStake(traderAddr common.Address) (uint16, error) {
	return _IPerpetualManager.Contract.GetFeeForTraderStake(&_IPerpetualManager.CallOpts, traderAddr)
}

// GetFeeForTraderStake is a free data retrieval call binding the contract method 0xbf7b6bb5.
//
// Solidity: function getFeeForTraderStake(address traderAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetFeeForTraderStake(traderAddr common.Address) (uint16, error) {
	return _IPerpetualManager.Contract.GetFeeForTraderStake(&_IPerpetualManager.CallOpts, traderAddr)
}

// GetFeeForTraderVolume is a free data retrieval call binding the contract method 0x849a5fd6.
//
// Solidity: function getFeeForTraderVolume(uint8 _poolId, address _traderAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCaller) GetFeeForTraderVolume(opts *bind.CallOpts, _poolId uint8, _traderAddr common.Address) (uint16, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getFeeForTraderVolume", _poolId, _traderAddr)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFeeForTraderVolume is a free data retrieval call binding the contract method 0x849a5fd6.
//
// Solidity: function getFeeForTraderVolume(uint8 _poolId, address _traderAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerSession) GetFeeForTraderVolume(_poolId uint8, _traderAddr common.Address) (uint16, error) {
	return _IPerpetualManager.Contract.GetFeeForTraderVolume(&_IPerpetualManager.CallOpts, _poolId, _traderAddr)
}

// GetFeeForTraderVolume is a free data retrieval call binding the contract method 0x849a5fd6.
//
// Solidity: function getFeeForTraderVolume(uint8 _poolId, address _traderAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetFeeForTraderVolume(_poolId uint8, _traderAddr common.Address) (uint16, error) {
	return _IPerpetualManager.Contract.GetFeeForTraderVolume(&_IPerpetualManager.CallOpts, _poolId, _traderAddr)
}

// GetLastPerpetualBaseToUSDConversion is a free data retrieval call binding the contract method 0xc0822eb5.
//
// Solidity: function getLastPerpetualBaseToUSDConversion(uint24 _iPerpetualId) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) GetLastPerpetualBaseToUSDConversion(opts *bind.CallOpts, _iPerpetualId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getLastPerpetualBaseToUSDConversion", _iPerpetualId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastPerpetualBaseToUSDConversion is a free data retrieval call binding the contract method 0xc0822eb5.
//
// Solidity: function getLastPerpetualBaseToUSDConversion(uint24 _iPerpetualId) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) GetLastPerpetualBaseToUSDConversion(_iPerpetualId *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetLastPerpetualBaseToUSDConversion(&_IPerpetualManager.CallOpts, _iPerpetualId)
}

// GetLastPerpetualBaseToUSDConversion is a free data retrieval call binding the contract method 0xc0822eb5.
//
// Solidity: function getLastPerpetualBaseToUSDConversion(uint24 _iPerpetualId) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetLastPerpetualBaseToUSDConversion(_iPerpetualId *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetLastPerpetualBaseToUSDConversion(&_IPerpetualManager.CallOpts, _iPerpetualId)
}

// GetLiquidatableAccounts is a free data retrieval call binding the contract method 0xe70ac514.
//
// Solidity: function getLiquidatableAccounts(uint24 _perpetualId, int128[2] _fIndexPrice) view returns(address[] unsafeAccounts)
func (_IPerpetualManager *IPerpetualManagerCaller) GetLiquidatableAccounts(opts *bind.CallOpts, _perpetualId *big.Int, _fIndexPrice [2]*big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getLiquidatableAccounts", _perpetualId, _fIndexPrice)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLiquidatableAccounts is a free data retrieval call binding the contract method 0xe70ac514.
//
// Solidity: function getLiquidatableAccounts(uint24 _perpetualId, int128[2] _fIndexPrice) view returns(address[] unsafeAccounts)
func (_IPerpetualManager *IPerpetualManagerSession) GetLiquidatableAccounts(_perpetualId *big.Int, _fIndexPrice [2]*big.Int) ([]common.Address, error) {
	return _IPerpetualManager.Contract.GetLiquidatableAccounts(&_IPerpetualManager.CallOpts, _perpetualId, _fIndexPrice)
}

// GetLiquidatableAccounts is a free data retrieval call binding the contract method 0xe70ac514.
//
// Solidity: function getLiquidatableAccounts(uint24 _perpetualId, int128[2] _fIndexPrice) view returns(address[] unsafeAccounts)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetLiquidatableAccounts(_perpetualId *big.Int, _fIndexPrice [2]*big.Int) ([]common.Address, error) {
	return _IPerpetualManager.Contract.GetLiquidatableAccounts(&_IPerpetualManager.CallOpts, _perpetualId, _fIndexPrice)
}

// GetLiquidityPool is a free data retrieval call binding the contract method 0x30af7c72.
//
// Solidity: function getLiquidityPool(uint8 _poolId) view returns((bool,uint8,uint8,uint16,address,int128,int128,uint64,int32,address,int128,int128,int128,int128,int128,int128,uint128,uint128,uint128))
func (_IPerpetualManager *IPerpetualManagerCaller) GetLiquidityPool(opts *bind.CallOpts, _poolId uint8) (PerpStorageLiquidityPoolData, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getLiquidityPool", _poolId)

	if err != nil {
		return *new(PerpStorageLiquidityPoolData), err
	}

	out0 := *abi.ConvertType(out[0], new(PerpStorageLiquidityPoolData)).(*PerpStorageLiquidityPoolData)

	return out0, err

}

// GetLiquidityPool is a free data retrieval call binding the contract method 0x30af7c72.
//
// Solidity: function getLiquidityPool(uint8 _poolId) view returns((bool,uint8,uint8,uint16,address,int128,int128,uint64,int32,address,int128,int128,int128,int128,int128,int128,uint128,uint128,uint128))
func (_IPerpetualManager *IPerpetualManagerSession) GetLiquidityPool(_poolId uint8) (PerpStorageLiquidityPoolData, error) {
	return _IPerpetualManager.Contract.GetLiquidityPool(&_IPerpetualManager.CallOpts, _poolId)
}

// GetLiquidityPool is a free data retrieval call binding the contract method 0x30af7c72.
//
// Solidity: function getLiquidityPool(uint8 _poolId) view returns((bool,uint8,uint8,uint16,address,int128,int128,uint64,int32,address,int128,int128,int128,int128,int128,int128,uint128,uint128,uint128))
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetLiquidityPool(_poolId uint8) (PerpStorageLiquidityPoolData, error) {
	return _IPerpetualManager.Contract.GetLiquidityPool(&_IPerpetualManager.CallOpts, _poolId)
}

// GetLiquidityPools is a free data retrieval call binding the contract method 0xe6d16785.
//
// Solidity: function getLiquidityPools(uint8 _poolIdFrom, uint8 _poolIdTo) view returns((bool,uint8,uint8,uint16,address,int128,int128,uint64,int32,address,int128,int128,int128,int128,int128,int128,uint128,uint128,uint128)[])
func (_IPerpetualManager *IPerpetualManagerCaller) GetLiquidityPools(opts *bind.CallOpts, _poolIdFrom uint8, _poolIdTo uint8) ([]PerpStorageLiquidityPoolData, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getLiquidityPools", _poolIdFrom, _poolIdTo)

	if err != nil {
		return *new([]PerpStorageLiquidityPoolData), err
	}

	out0 := *abi.ConvertType(out[0], new([]PerpStorageLiquidityPoolData)).(*[]PerpStorageLiquidityPoolData)

	return out0, err

}

// GetLiquidityPools is a free data retrieval call binding the contract method 0xe6d16785.
//
// Solidity: function getLiquidityPools(uint8 _poolIdFrom, uint8 _poolIdTo) view returns((bool,uint8,uint8,uint16,address,int128,int128,uint64,int32,address,int128,int128,int128,int128,int128,int128,uint128,uint128,uint128)[])
func (_IPerpetualManager *IPerpetualManagerSession) GetLiquidityPools(_poolIdFrom uint8, _poolIdTo uint8) ([]PerpStorageLiquidityPoolData, error) {
	return _IPerpetualManager.Contract.GetLiquidityPools(&_IPerpetualManager.CallOpts, _poolIdFrom, _poolIdTo)
}

// GetLiquidityPools is a free data retrieval call binding the contract method 0xe6d16785.
//
// Solidity: function getLiquidityPools(uint8 _poolIdFrom, uint8 _poolIdTo) view returns((bool,uint8,uint8,uint16,address,int128,int128,uint64,int32,address,int128,int128,int128,int128,int128,int128,uint128,uint128,uint128)[])
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetLiquidityPools(_poolIdFrom uint8, _poolIdTo uint8) ([]PerpStorageLiquidityPoolData, error) {
	return _IPerpetualManager.Contract.GetLiquidityPools(&_IPerpetualManager.CallOpts, _poolIdFrom, _poolIdTo)
}

// GetMarginAccount is a free data retrieval call binding the contract method 0x7aa744d6.
//
// Solidity: function getMarginAccount(uint24 _perpetualId, address _traderAddress) view returns((int128,int128,int128,int128,uint64,uint16,uint16,bytes16))
func (_IPerpetualManager *IPerpetualManagerCaller) GetMarginAccount(opts *bind.CallOpts, _perpetualId *big.Int, _traderAddress common.Address) (PerpStorageMarginAccount, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getMarginAccount", _perpetualId, _traderAddress)

	if err != nil {
		return *new(PerpStorageMarginAccount), err
	}

	out0 := *abi.ConvertType(out[0], new(PerpStorageMarginAccount)).(*PerpStorageMarginAccount)

	return out0, err

}

// GetMarginAccount is a free data retrieval call binding the contract method 0x7aa744d6.
//
// Solidity: function getMarginAccount(uint24 _perpetualId, address _traderAddress) view returns((int128,int128,int128,int128,uint64,uint16,uint16,bytes16))
func (_IPerpetualManager *IPerpetualManagerSession) GetMarginAccount(_perpetualId *big.Int, _traderAddress common.Address) (PerpStorageMarginAccount, error) {
	return _IPerpetualManager.Contract.GetMarginAccount(&_IPerpetualManager.CallOpts, _perpetualId, _traderAddress)
}

// GetMarginAccount is a free data retrieval call binding the contract method 0x7aa744d6.
//
// Solidity: function getMarginAccount(uint24 _perpetualId, address _traderAddress) view returns((int128,int128,int128,int128,uint64,uint16,uint16,bytes16))
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetMarginAccount(_perpetualId *big.Int, _traderAddress common.Address) (PerpStorageMarginAccount, error) {
	return _IPerpetualManager.Contract.GetMarginAccount(&_IPerpetualManager.CallOpts, _perpetualId, _traderAddress)
}

// GetMaxSignedOpenTradeSizeForPos is a free data retrieval call binding the contract method 0xfbc3ac3c.
//
// Solidity: function getMaxSignedOpenTradeSizeForPos(uint24 _perpetualId, int128 _fCurrentTraderPos, bool _isBuy) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) GetMaxSignedOpenTradeSizeForPos(opts *bind.CallOpts, _perpetualId *big.Int, _fCurrentTraderPos *big.Int, _isBuy bool) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getMaxSignedOpenTradeSizeForPos", _perpetualId, _fCurrentTraderPos, _isBuy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxSignedOpenTradeSizeForPos is a free data retrieval call binding the contract method 0xfbc3ac3c.
//
// Solidity: function getMaxSignedOpenTradeSizeForPos(uint24 _perpetualId, int128 _fCurrentTraderPos, bool _isBuy) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) GetMaxSignedOpenTradeSizeForPos(_perpetualId *big.Int, _fCurrentTraderPos *big.Int, _isBuy bool) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetMaxSignedOpenTradeSizeForPos(&_IPerpetualManager.CallOpts, _perpetualId, _fCurrentTraderPos, _isBuy)
}

// GetMaxSignedOpenTradeSizeForPos is a free data retrieval call binding the contract method 0xfbc3ac3c.
//
// Solidity: function getMaxSignedOpenTradeSizeForPos(uint24 _perpetualId, int128 _fCurrentTraderPos, bool _isBuy) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetMaxSignedOpenTradeSizeForPos(_perpetualId *big.Int, _fCurrentTraderPos *big.Int, _isBuy bool) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetMaxSignedOpenTradeSizeForPos(&_IPerpetualManager.CallOpts, _perpetualId, _fCurrentTraderPos, _isBuy)
}

// GetNextLiquidatableTrader is a free data retrieval call binding the contract method 0xe3a68c47.
//
// Solidity: function getNextLiquidatableTrader(uint24 _perpetualId, int128[2] _fIndexPrice) view returns(address traderAddr)
func (_IPerpetualManager *IPerpetualManagerCaller) GetNextLiquidatableTrader(opts *bind.CallOpts, _perpetualId *big.Int, _fIndexPrice [2]*big.Int) (common.Address, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getNextLiquidatableTrader", _perpetualId, _fIndexPrice)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNextLiquidatableTrader is a free data retrieval call binding the contract method 0xe3a68c47.
//
// Solidity: function getNextLiquidatableTrader(uint24 _perpetualId, int128[2] _fIndexPrice) view returns(address traderAddr)
func (_IPerpetualManager *IPerpetualManagerSession) GetNextLiquidatableTrader(_perpetualId *big.Int, _fIndexPrice [2]*big.Int) (common.Address, error) {
	return _IPerpetualManager.Contract.GetNextLiquidatableTrader(&_IPerpetualManager.CallOpts, _perpetualId, _fIndexPrice)
}

// GetNextLiquidatableTrader is a free data retrieval call binding the contract method 0xe3a68c47.
//
// Solidity: function getNextLiquidatableTrader(uint24 _perpetualId, int128[2] _fIndexPrice) view returns(address traderAddr)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetNextLiquidatableTrader(_perpetualId *big.Int, _fIndexPrice [2]*big.Int) (common.Address, error) {
	return _IPerpetualManager.Contract.GetNextLiquidatableTrader(&_IPerpetualManager.CallOpts, _perpetualId, _fIndexPrice)
}

// GetOracleFactory is a free data retrieval call binding the contract method 0x47619264.
//
// Solidity: function getOracleFactory() view returns(address)
func (_IPerpetualManager *IPerpetualManagerCaller) GetOracleFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getOracleFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracleFactory is a free data retrieval call binding the contract method 0x47619264.
//
// Solidity: function getOracleFactory() view returns(address)
func (_IPerpetualManager *IPerpetualManagerSession) GetOracleFactory() (common.Address, error) {
	return _IPerpetualManager.Contract.GetOracleFactory(&_IPerpetualManager.CallOpts)
}

// GetOracleFactory is a free data retrieval call binding the contract method 0x47619264.
//
// Solidity: function getOracleFactory() view returns(address)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetOracleFactory() (common.Address, error) {
	return _IPerpetualManager.Contract.GetOracleFactory(&_IPerpetualManager.CallOpts)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x002ebb4a.
//
// Solidity: function getOraclePrice(bytes4[2] _baseQuote) view returns(int128 fPrice)
func (_IPerpetualManager *IPerpetualManagerCaller) GetOraclePrice(opts *bind.CallOpts, _baseQuote [2][4]byte) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getOraclePrice", _baseQuote)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOraclePrice is a free data retrieval call binding the contract method 0x002ebb4a.
//
// Solidity: function getOraclePrice(bytes4[2] _baseQuote) view returns(int128 fPrice)
func (_IPerpetualManager *IPerpetualManagerSession) GetOraclePrice(_baseQuote [2][4]byte) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetOraclePrice(&_IPerpetualManager.CallOpts, _baseQuote)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x002ebb4a.
//
// Solidity: function getOraclePrice(bytes4[2] _baseQuote) view returns(int128 fPrice)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetOraclePrice(_baseQuote [2][4]byte) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetOraclePrice(&_IPerpetualManager.CallOpts, _baseQuote)
}

// GetOrderBookAddress is a free data retrieval call binding the contract method 0x1e780ba4.
//
// Solidity: function getOrderBookAddress(uint24 _perpetualId) view returns(address)
func (_IPerpetualManager *IPerpetualManagerCaller) GetOrderBookAddress(opts *bind.CallOpts, _perpetualId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getOrderBookAddress", _perpetualId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOrderBookAddress is a free data retrieval call binding the contract method 0x1e780ba4.
//
// Solidity: function getOrderBookAddress(uint24 _perpetualId) view returns(address)
func (_IPerpetualManager *IPerpetualManagerSession) GetOrderBookAddress(_perpetualId *big.Int) (common.Address, error) {
	return _IPerpetualManager.Contract.GetOrderBookAddress(&_IPerpetualManager.CallOpts, _perpetualId)
}

// GetOrderBookAddress is a free data retrieval call binding the contract method 0x1e780ba4.
//
// Solidity: function getOrderBookAddress(uint24 _perpetualId) view returns(address)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetOrderBookAddress(_perpetualId *big.Int) (common.Address, error) {
	return _IPerpetualManager.Contract.GetOrderBookAddress(&_IPerpetualManager.CallOpts, _perpetualId)
}

// GetOrderBookFactoryAddress is a free data retrieval call binding the contract method 0x40239dc5.
//
// Solidity: function getOrderBookFactoryAddress() view returns(address)
func (_IPerpetualManager *IPerpetualManagerCaller) GetOrderBookFactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getOrderBookFactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOrderBookFactoryAddress is a free data retrieval call binding the contract method 0x40239dc5.
//
// Solidity: function getOrderBookFactoryAddress() view returns(address)
func (_IPerpetualManager *IPerpetualManagerSession) GetOrderBookFactoryAddress() (common.Address, error) {
	return _IPerpetualManager.Contract.GetOrderBookFactoryAddress(&_IPerpetualManager.CallOpts)
}

// GetOrderBookFactoryAddress is a free data retrieval call binding the contract method 0x40239dc5.
//
// Solidity: function getOrderBookFactoryAddress() view returns(address)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetOrderBookFactoryAddress() (common.Address, error) {
	return _IPerpetualManager.Contract.GetOrderBookFactoryAddress(&_IPerpetualManager.CallOpts)
}

// GetPerpetual is a free data retrieval call binding the contract method 0xc7fd35bf.
//
// Solidity: function getPerpetual(uint24 _perpetualId) view returns((uint8,uint24,int32,int32,int32,uint32,uint32,uint32,uint32,uint8,uint8,uint16,bytes4,uint16,bytes4,uint16,bytes4,uint16,bytes4,int32,(int128,uint64),int128[2],int128[2],int128[2],int128[2],int128[2],int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,uint32,int32,int32,int32,int32,int32))
func (_IPerpetualManager *IPerpetualManagerCaller) GetPerpetual(opts *bind.CallOpts, _perpetualId *big.Int) (PerpStoragePerpetualData, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getPerpetual", _perpetualId)

	if err != nil {
		return *new(PerpStoragePerpetualData), err
	}

	out0 := *abi.ConvertType(out[0], new(PerpStoragePerpetualData)).(*PerpStoragePerpetualData)

	return out0, err

}

// GetPerpetual is a free data retrieval call binding the contract method 0xc7fd35bf.
//
// Solidity: function getPerpetual(uint24 _perpetualId) view returns((uint8,uint24,int32,int32,int32,uint32,uint32,uint32,uint32,uint8,uint8,uint16,bytes4,uint16,bytes4,uint16,bytes4,uint16,bytes4,int32,(int128,uint64),int128[2],int128[2],int128[2],int128[2],int128[2],int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,uint32,int32,int32,int32,int32,int32))
func (_IPerpetualManager *IPerpetualManagerSession) GetPerpetual(_perpetualId *big.Int) (PerpStoragePerpetualData, error) {
	return _IPerpetualManager.Contract.GetPerpetual(&_IPerpetualManager.CallOpts, _perpetualId)
}

// GetPerpetual is a free data retrieval call binding the contract method 0xc7fd35bf.
//
// Solidity: function getPerpetual(uint24 _perpetualId) view returns((uint8,uint24,int32,int32,int32,uint32,uint32,uint32,uint32,uint8,uint8,uint16,bytes4,uint16,bytes4,uint16,bytes4,uint16,bytes4,int32,(int128,uint64),int128[2],int128[2],int128[2],int128[2],int128[2],int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,uint32,int32,int32,int32,int32,int32))
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetPerpetual(_perpetualId *big.Int) (PerpStoragePerpetualData, error) {
	return _IPerpetualManager.Contract.GetPerpetual(&_IPerpetualManager.CallOpts, _perpetualId)
}

// GetPerpetualCountInPool is a free data retrieval call binding the contract method 0xe5f26270.
//
// Solidity: function getPerpetualCountInPool(uint8 _poolId) view returns(uint8)
func (_IPerpetualManager *IPerpetualManagerCaller) GetPerpetualCountInPool(opts *bind.CallOpts, _poolId uint8) (uint8, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getPerpetualCountInPool", _poolId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPerpetualCountInPool is a free data retrieval call binding the contract method 0xe5f26270.
//
// Solidity: function getPerpetualCountInPool(uint8 _poolId) view returns(uint8)
func (_IPerpetualManager *IPerpetualManagerSession) GetPerpetualCountInPool(_poolId uint8) (uint8, error) {
	return _IPerpetualManager.Contract.GetPerpetualCountInPool(&_IPerpetualManager.CallOpts, _poolId)
}

// GetPerpetualCountInPool is a free data retrieval call binding the contract method 0xe5f26270.
//
// Solidity: function getPerpetualCountInPool(uint8 _poolId) view returns(uint8)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetPerpetualCountInPool(_poolId uint8) (uint8, error) {
	return _IPerpetualManager.Contract.GetPerpetualCountInPool(&_IPerpetualManager.CallOpts, _poolId)
}

// GetPerpetualId is a free data retrieval call binding the contract method 0x77a5216a.
//
// Solidity: function getPerpetualId(uint8 _poolId, uint8 _perpetualIndex) view returns(uint24)
func (_IPerpetualManager *IPerpetualManagerCaller) GetPerpetualId(opts *bind.CallOpts, _poolId uint8, _perpetualIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getPerpetualId", _poolId, _perpetualIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPerpetualId is a free data retrieval call binding the contract method 0x77a5216a.
//
// Solidity: function getPerpetualId(uint8 _poolId, uint8 _perpetualIndex) view returns(uint24)
func (_IPerpetualManager *IPerpetualManagerSession) GetPerpetualId(_poolId uint8, _perpetualIndex uint8) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetPerpetualId(&_IPerpetualManager.CallOpts, _poolId, _perpetualIndex)
}

// GetPerpetualId is a free data retrieval call binding the contract method 0x77a5216a.
//
// Solidity: function getPerpetualId(uint8 _poolId, uint8 _perpetualIndex) view returns(uint24)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetPerpetualId(_poolId uint8, _perpetualIndex uint8) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetPerpetualId(&_IPerpetualManager.CallOpts, _poolId, _perpetualIndex)
}

// GetPerpetualStaticInfo is a free data retrieval call binding the contract method 0x68805619.
//
// Solidity: function getPerpetualStaticInfo(uint24[] perpetualIds) view returns((uint24,address,int32,int32,uint8,uint8,bytes4,bytes4,bytes4,bytes4,int128,int128,bytes32[],bool[])[])
func (_IPerpetualManager *IPerpetualManagerCaller) GetPerpetualStaticInfo(opts *bind.CallOpts, perpetualIds []*big.Int) ([]IPerpetualGetterPerpetualStaticInfo, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getPerpetualStaticInfo", perpetualIds)

	if err != nil {
		return *new([]IPerpetualGetterPerpetualStaticInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPerpetualGetterPerpetualStaticInfo)).(*[]IPerpetualGetterPerpetualStaticInfo)

	return out0, err

}

// GetPerpetualStaticInfo is a free data retrieval call binding the contract method 0x68805619.
//
// Solidity: function getPerpetualStaticInfo(uint24[] perpetualIds) view returns((uint24,address,int32,int32,uint8,uint8,bytes4,bytes4,bytes4,bytes4,int128,int128,bytes32[],bool[])[])
func (_IPerpetualManager *IPerpetualManagerSession) GetPerpetualStaticInfo(perpetualIds []*big.Int) ([]IPerpetualGetterPerpetualStaticInfo, error) {
	return _IPerpetualManager.Contract.GetPerpetualStaticInfo(&_IPerpetualManager.CallOpts, perpetualIds)
}

// GetPerpetualStaticInfo is a free data retrieval call binding the contract method 0x68805619.
//
// Solidity: function getPerpetualStaticInfo(uint24[] perpetualIds) view returns((uint24,address,int32,int32,uint8,uint8,bytes4,bytes4,bytes4,bytes4,int128,int128,bytes32[],bool[])[])
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetPerpetualStaticInfo(perpetualIds []*big.Int) ([]IPerpetualGetterPerpetualStaticInfo, error) {
	return _IPerpetualManager.Contract.GetPerpetualStaticInfo(&_IPerpetualManager.CallOpts, perpetualIds)
}

// GetPerpetuals is a free data retrieval call binding the contract method 0x36b6b9dd.
//
// Solidity: function getPerpetuals(uint24[] perpetualIds) view returns((uint8,uint24,int32,int32,int32,uint32,uint32,uint32,uint32,uint8,uint8,uint16,bytes4,uint16,bytes4,uint16,bytes4,uint16,bytes4,int32,(int128,uint64),int128[2],int128[2],int128[2],int128[2],int128[2],int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,uint32,int32,int32,int32,int32,int32)[])
func (_IPerpetualManager *IPerpetualManagerCaller) GetPerpetuals(opts *bind.CallOpts, perpetualIds []*big.Int) ([]PerpStoragePerpetualData, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getPerpetuals", perpetualIds)

	if err != nil {
		return *new([]PerpStoragePerpetualData), err
	}

	out0 := *abi.ConvertType(out[0], new([]PerpStoragePerpetualData)).(*[]PerpStoragePerpetualData)

	return out0, err

}

// GetPerpetuals is a free data retrieval call binding the contract method 0x36b6b9dd.
//
// Solidity: function getPerpetuals(uint24[] perpetualIds) view returns((uint8,uint24,int32,int32,int32,uint32,uint32,uint32,uint32,uint8,uint8,uint16,bytes4,uint16,bytes4,uint16,bytes4,uint16,bytes4,int32,(int128,uint64),int128[2],int128[2],int128[2],int128[2],int128[2],int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,uint32,int32,int32,int32,int32,int32)[])
func (_IPerpetualManager *IPerpetualManagerSession) GetPerpetuals(perpetualIds []*big.Int) ([]PerpStoragePerpetualData, error) {
	return _IPerpetualManager.Contract.GetPerpetuals(&_IPerpetualManager.CallOpts, perpetualIds)
}

// GetPerpetuals is a free data retrieval call binding the contract method 0x36b6b9dd.
//
// Solidity: function getPerpetuals(uint24[] perpetualIds) view returns((uint8,uint24,int32,int32,int32,uint32,uint32,uint32,uint32,uint8,uint8,uint16,bytes4,uint16,bytes4,uint16,bytes4,uint16,bytes4,int32,(int128,uint64),int128[2],int128[2],int128[2],int128[2],int128[2],int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,int128,uint32,int32,int32,int32,int32,int32)[])
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetPerpetuals(perpetualIds []*big.Int) ([]PerpStoragePerpetualData, error) {
	return _IPerpetualManager.Contract.GetPerpetuals(&_IPerpetualManager.CallOpts, perpetualIds)
}

// GetPoolCount is a free data retrieval call binding the contract method 0x8eec5d70.
//
// Solidity: function getPoolCount() view returns(uint8)
func (_IPerpetualManager *IPerpetualManagerCaller) GetPoolCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getPoolCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPoolCount is a free data retrieval call binding the contract method 0x8eec5d70.
//
// Solidity: function getPoolCount() view returns(uint8)
func (_IPerpetualManager *IPerpetualManagerSession) GetPoolCount() (uint8, error) {
	return _IPerpetualManager.Contract.GetPoolCount(&_IPerpetualManager.CallOpts)
}

// GetPoolCount is a free data retrieval call binding the contract method 0x8eec5d70.
//
// Solidity: function getPoolCount() view returns(uint8)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetPoolCount() (uint8, error) {
	return _IPerpetualManager.Contract.GetPoolCount(&_IPerpetualManager.CallOpts)
}

// GetPoolIdByPerpetualId is a free data retrieval call binding the contract method 0xef3ac40d.
//
// Solidity: function getPoolIdByPerpetualId(uint24 _perpetualId) view returns(uint8)
func (_IPerpetualManager *IPerpetualManagerCaller) GetPoolIdByPerpetualId(opts *bind.CallOpts, _perpetualId *big.Int) (uint8, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getPoolIdByPerpetualId", _perpetualId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPoolIdByPerpetualId is a free data retrieval call binding the contract method 0xef3ac40d.
//
// Solidity: function getPoolIdByPerpetualId(uint24 _perpetualId) view returns(uint8)
func (_IPerpetualManager *IPerpetualManagerSession) GetPoolIdByPerpetualId(_perpetualId *big.Int) (uint8, error) {
	return _IPerpetualManager.Contract.GetPoolIdByPerpetualId(&_IPerpetualManager.CallOpts, _perpetualId)
}

// GetPoolIdByPerpetualId is a free data retrieval call binding the contract method 0xef3ac40d.
//
// Solidity: function getPoolIdByPerpetualId(uint24 _perpetualId) view returns(uint8)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetPoolIdByPerpetualId(_perpetualId *big.Int) (uint8, error) {
	return _IPerpetualManager.Contract.GetPoolIdByPerpetualId(&_IPerpetualManager.CallOpts, _perpetualId)
}

// GetPoolStaticInfo is a free data retrieval call binding the contract method 0x88402688.
//
// Solidity: function getPoolStaticInfo(uint8 _poolFromIdx, uint8 _poolToIdx) view returns(uint24[][], address[], address[], address _oracleFactoryAddress)
func (_IPerpetualManager *IPerpetualManagerCaller) GetPoolStaticInfo(opts *bind.CallOpts, _poolFromIdx uint8, _poolToIdx uint8) ([][]*big.Int, []common.Address, []common.Address, common.Address, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getPoolStaticInfo", _poolFromIdx, _poolToIdx)

	if err != nil {
		return *new([][]*big.Int), *new([]common.Address), *new([]common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, err

}

// GetPoolStaticInfo is a free data retrieval call binding the contract method 0x88402688.
//
// Solidity: function getPoolStaticInfo(uint8 _poolFromIdx, uint8 _poolToIdx) view returns(uint24[][], address[], address[], address _oracleFactoryAddress)
func (_IPerpetualManager *IPerpetualManagerSession) GetPoolStaticInfo(_poolFromIdx uint8, _poolToIdx uint8) ([][]*big.Int, []common.Address, []common.Address, common.Address, error) {
	return _IPerpetualManager.Contract.GetPoolStaticInfo(&_IPerpetualManager.CallOpts, _poolFromIdx, _poolToIdx)
}

// GetPoolStaticInfo is a free data retrieval call binding the contract method 0x88402688.
//
// Solidity: function getPoolStaticInfo(uint8 _poolFromIdx, uint8 _poolToIdx) view returns(uint24[][], address[], address[], address _oracleFactoryAddress)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetPoolStaticInfo(_poolFromIdx uint8, _poolToIdx uint8) ([][]*big.Int, []common.Address, []common.Address, common.Address, error) {
	return _IPerpetualManager.Contract.GetPoolStaticInfo(&_IPerpetualManager.CallOpts, _poolFromIdx, _poolToIdx)
}

// GetPriceInfo is a free data retrieval call binding the contract method 0x8dd46a3b.
//
// Solidity: function getPriceInfo(uint24 _perpetualId) view returns(bytes32[], bool[])
func (_IPerpetualManager *IPerpetualManagerCaller) GetPriceInfo(opts *bind.CallOpts, _perpetualId *big.Int) ([][32]byte, []bool, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getPriceInfo", _perpetualId)

	if err != nil {
		return *new([][32]byte), *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	out1 := *abi.ConvertType(out[1], new([]bool)).(*[]bool)

	return out0, out1, err

}

// GetPriceInfo is a free data retrieval call binding the contract method 0x8dd46a3b.
//
// Solidity: function getPriceInfo(uint24 _perpetualId) view returns(bytes32[], bool[])
func (_IPerpetualManager *IPerpetualManagerSession) GetPriceInfo(_perpetualId *big.Int) ([][32]byte, []bool, error) {
	return _IPerpetualManager.Contract.GetPriceInfo(&_IPerpetualManager.CallOpts, _perpetualId)
}

// GetPriceInfo is a free data retrieval call binding the contract method 0x8dd46a3b.
//
// Solidity: function getPriceInfo(uint24 _perpetualId) view returns(bytes32[], bool[])
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetPriceInfo(_perpetualId *big.Int) ([][32]byte, []bool, error) {
	return _IPerpetualManager.Contract.GetPriceInfo(&_IPerpetualManager.CallOpts, _perpetualId)
}

// GetShareTokenFactory is a free data retrieval call binding the contract method 0x1fb94c4f.
//
// Solidity: function getShareTokenFactory() view returns(address)
func (_IPerpetualManager *IPerpetualManagerCaller) GetShareTokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getShareTokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetShareTokenFactory is a free data retrieval call binding the contract method 0x1fb94c4f.
//
// Solidity: function getShareTokenFactory() view returns(address)
func (_IPerpetualManager *IPerpetualManagerSession) GetShareTokenFactory() (common.Address, error) {
	return _IPerpetualManager.Contract.GetShareTokenFactory(&_IPerpetualManager.CallOpts)
}

// GetShareTokenFactory is a free data retrieval call binding the contract method 0x1fb94c4f.
//
// Solidity: function getShareTokenFactory() view returns(address)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetShareTokenFactory() (common.Address, error) {
	return _IPerpetualManager.Contract.GetShareTokenFactory(&_IPerpetualManager.CallOpts)
}

// GetShareTokenPriceD18 is a free data retrieval call binding the contract method 0x2faee618.
//
// Solidity: function getShareTokenPriceD18(uint8 _poolId) view returns(uint256 price)
func (_IPerpetualManager *IPerpetualManagerCaller) GetShareTokenPriceD18(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getShareTokenPriceD18", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetShareTokenPriceD18 is a free data retrieval call binding the contract method 0x2faee618.
//
// Solidity: function getShareTokenPriceD18(uint8 _poolId) view returns(uint256 price)
func (_IPerpetualManager *IPerpetualManagerSession) GetShareTokenPriceD18(_poolId uint8) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetShareTokenPriceD18(&_IPerpetualManager.CallOpts, _poolId)
}

// GetShareTokenPriceD18 is a free data retrieval call binding the contract method 0x2faee618.
//
// Solidity: function getShareTokenPriceD18(uint8 _poolId) view returns(uint256 price)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetShareTokenPriceD18(_poolId uint8) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetShareTokenPriceD18(&_IPerpetualManager.CallOpts, _poolId)
}

// GetTargetCollateralM1 is a free data retrieval call binding the contract method 0xed2db96f.
//
// Solidity: function getTargetCollateralM1(int128 _fK2, int128 _fL1, (int128,int128,int128,int128,int128) _mktVars, int128 _fTargetDD) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) GetTargetCollateralM1(opts *bind.CallOpts, _fK2 *big.Int, _fL1 *big.Int, _mktVars AMMPerpLogicMarketVariables, _fTargetDD *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getTargetCollateralM1", _fK2, _fL1, _mktVars, _fTargetDD)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetCollateralM1 is a free data retrieval call binding the contract method 0xed2db96f.
//
// Solidity: function getTargetCollateralM1(int128 _fK2, int128 _fL1, (int128,int128,int128,int128,int128) _mktVars, int128 _fTargetDD) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) GetTargetCollateralM1(_fK2 *big.Int, _fL1 *big.Int, _mktVars AMMPerpLogicMarketVariables, _fTargetDD *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetTargetCollateralM1(&_IPerpetualManager.CallOpts, _fK2, _fL1, _mktVars, _fTargetDD)
}

// GetTargetCollateralM1 is a free data retrieval call binding the contract method 0xed2db96f.
//
// Solidity: function getTargetCollateralM1(int128 _fK2, int128 _fL1, (int128,int128,int128,int128,int128) _mktVars, int128 _fTargetDD) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetTargetCollateralM1(_fK2 *big.Int, _fL1 *big.Int, _mktVars AMMPerpLogicMarketVariables, _fTargetDD *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetTargetCollateralM1(&_IPerpetualManager.CallOpts, _fK2, _fL1, _mktVars, _fTargetDD)
}

// GetTargetCollateralM2 is a free data retrieval call binding the contract method 0xf2fa682a.
//
// Solidity: function getTargetCollateralM2(int128 _fK2, int128 _fL1, (int128,int128,int128,int128,int128) _mktVars, int128 _fTargetDD) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) GetTargetCollateralM2(opts *bind.CallOpts, _fK2 *big.Int, _fL1 *big.Int, _mktVars AMMPerpLogicMarketVariables, _fTargetDD *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getTargetCollateralM2", _fK2, _fL1, _mktVars, _fTargetDD)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetCollateralM2 is a free data retrieval call binding the contract method 0xf2fa682a.
//
// Solidity: function getTargetCollateralM2(int128 _fK2, int128 _fL1, (int128,int128,int128,int128,int128) _mktVars, int128 _fTargetDD) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) GetTargetCollateralM2(_fK2 *big.Int, _fL1 *big.Int, _mktVars AMMPerpLogicMarketVariables, _fTargetDD *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetTargetCollateralM2(&_IPerpetualManager.CallOpts, _fK2, _fL1, _mktVars, _fTargetDD)
}

// GetTargetCollateralM2 is a free data retrieval call binding the contract method 0xf2fa682a.
//
// Solidity: function getTargetCollateralM2(int128 _fK2, int128 _fL1, (int128,int128,int128,int128,int128) _mktVars, int128 _fTargetDD) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetTargetCollateralM2(_fK2 *big.Int, _fL1 *big.Int, _mktVars AMMPerpLogicMarketVariables, _fTargetDD *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetTargetCollateralM2(&_IPerpetualManager.CallOpts, _fK2, _fL1, _mktVars, _fTargetDD)
}

// GetTargetCollateralM3 is a free data retrieval call binding the contract method 0x2fe01862.
//
// Solidity: function getTargetCollateralM3(int128 _fK2, int128 _fL1, (int128,int128,int128,int128,int128) _mktVars, int128 _fTargetDD) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) GetTargetCollateralM3(opts *bind.CallOpts, _fK2 *big.Int, _fL1 *big.Int, _mktVars AMMPerpLogicMarketVariables, _fTargetDD *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getTargetCollateralM3", _fK2, _fL1, _mktVars, _fTargetDD)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetCollateralM3 is a free data retrieval call binding the contract method 0x2fe01862.
//
// Solidity: function getTargetCollateralM3(int128 _fK2, int128 _fL1, (int128,int128,int128,int128,int128) _mktVars, int128 _fTargetDD) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) GetTargetCollateralM3(_fK2 *big.Int, _fL1 *big.Int, _mktVars AMMPerpLogicMarketVariables, _fTargetDD *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetTargetCollateralM3(&_IPerpetualManager.CallOpts, _fK2, _fL1, _mktVars, _fTargetDD)
}

// GetTargetCollateralM3 is a free data retrieval call binding the contract method 0x2fe01862.
//
// Solidity: function getTargetCollateralM3(int128 _fK2, int128 _fL1, (int128,int128,int128,int128,int128) _mktVars, int128 _fTargetDD) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetTargetCollateralM3(_fK2 *big.Int, _fL1 *big.Int, _mktVars AMMPerpLogicMarketVariables, _fTargetDD *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetTargetCollateralM3(&_IPerpetualManager.CallOpts, _fK2, _fL1, _mktVars, _fTargetDD)
}

// GetTokenAmountToReturn is a free data retrieval call binding the contract method 0x642d7a94.
//
// Solidity: function getTokenAmountToReturn(uint8 _poolId, uint256 _shareAmount) view returns(uint256)
func (_IPerpetualManager *IPerpetualManagerCaller) GetTokenAmountToReturn(opts *bind.CallOpts, _poolId uint8, _shareAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getTokenAmountToReturn", _poolId, _shareAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenAmountToReturn is a free data retrieval call binding the contract method 0x642d7a94.
//
// Solidity: function getTokenAmountToReturn(uint8 _poolId, uint256 _shareAmount) view returns(uint256)
func (_IPerpetualManager *IPerpetualManagerSession) GetTokenAmountToReturn(_poolId uint8, _shareAmount *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetTokenAmountToReturn(&_IPerpetualManager.CallOpts, _poolId, _shareAmount)
}

// GetTokenAmountToReturn is a free data retrieval call binding the contract method 0x642d7a94.
//
// Solidity: function getTokenAmountToReturn(uint8 _poolId, uint256 _shareAmount) view returns(uint256)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetTokenAmountToReturn(_poolId uint8, _shareAmount *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.GetTokenAmountToReturn(&_IPerpetualManager.CallOpts, _poolId, _shareAmount)
}

// GetTraderState is a free data retrieval call binding the contract method 0x7a9fa6c2.
//
// Solidity: function getTraderState(uint24 _perpetualId, address _traderAddress, int128[2] _fIndexPrice) view returns(int128[11])
func (_IPerpetualManager *IPerpetualManagerCaller) GetTraderState(opts *bind.CallOpts, _perpetualId *big.Int, _traderAddress common.Address, _fIndexPrice [2]*big.Int) ([11]*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getTraderState", _perpetualId, _traderAddress, _fIndexPrice)

	if err != nil {
		return *new([11]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([11]*big.Int)).(*[11]*big.Int)

	return out0, err

}

// GetTraderState is a free data retrieval call binding the contract method 0x7a9fa6c2.
//
// Solidity: function getTraderState(uint24 _perpetualId, address _traderAddress, int128[2] _fIndexPrice) view returns(int128[11])
func (_IPerpetualManager *IPerpetualManagerSession) GetTraderState(_perpetualId *big.Int, _traderAddress common.Address, _fIndexPrice [2]*big.Int) ([11]*big.Int, error) {
	return _IPerpetualManager.Contract.GetTraderState(&_IPerpetualManager.CallOpts, _perpetualId, _traderAddress, _fIndexPrice)
}

// GetTraderState is a free data retrieval call binding the contract method 0x7a9fa6c2.
//
// Solidity: function getTraderState(uint24 _perpetualId, address _traderAddress, int128[2] _fIndexPrice) view returns(int128[11])
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetTraderState(_perpetualId *big.Int, _traderAddress common.Address, _fIndexPrice [2]*big.Int) ([11]*big.Int, error) {
	return _IPerpetualManager.Contract.GetTraderState(&_IPerpetualManager.CallOpts, _perpetualId, _traderAddress, _fIndexPrice)
}

// GetTreasuryAddress is a free data retrieval call binding the contract method 0xe0024604.
//
// Solidity: function getTreasuryAddress() view returns(address)
func (_IPerpetualManager *IPerpetualManagerCaller) GetTreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getTreasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasuryAddress is a free data retrieval call binding the contract method 0xe0024604.
//
// Solidity: function getTreasuryAddress() view returns(address)
func (_IPerpetualManager *IPerpetualManagerSession) GetTreasuryAddress() (common.Address, error) {
	return _IPerpetualManager.Contract.GetTreasuryAddress(&_IPerpetualManager.CallOpts)
}

// GetTreasuryAddress is a free data retrieval call binding the contract method 0xe0024604.
//
// Solidity: function getTreasuryAddress() view returns(address)
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetTreasuryAddress() (common.Address, error) {
	return _IPerpetualManager.Contract.GetTreasuryAddress(&_IPerpetualManager.CallOpts)
}

// GetWithdrawRequests is a free data retrieval call binding the contract method 0x1a8d52e1.
//
// Solidity: function getWithdrawRequests(uint8 poolId, uint256 _fromIdx, uint256 numRequests) view returns((address,uint256,uint64)[])
func (_IPerpetualManager *IPerpetualManagerCaller) GetWithdrawRequests(opts *bind.CallOpts, poolId uint8, _fromIdx *big.Int, numRequests *big.Int) ([]IPerpetualTreasuryWithdrawRequest, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "getWithdrawRequests", poolId, _fromIdx, numRequests)

	if err != nil {
		return *new([]IPerpetualTreasuryWithdrawRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPerpetualTreasuryWithdrawRequest)).(*[]IPerpetualTreasuryWithdrawRequest)

	return out0, err

}

// GetWithdrawRequests is a free data retrieval call binding the contract method 0x1a8d52e1.
//
// Solidity: function getWithdrawRequests(uint8 poolId, uint256 _fromIdx, uint256 numRequests) view returns((address,uint256,uint64)[])
func (_IPerpetualManager *IPerpetualManagerSession) GetWithdrawRequests(poolId uint8, _fromIdx *big.Int, numRequests *big.Int) ([]IPerpetualTreasuryWithdrawRequest, error) {
	return _IPerpetualManager.Contract.GetWithdrawRequests(&_IPerpetualManager.CallOpts, poolId, _fromIdx, numRequests)
}

// GetWithdrawRequests is a free data retrieval call binding the contract method 0x1a8d52e1.
//
// Solidity: function getWithdrawRequests(uint8 poolId, uint256 _fromIdx, uint256 numRequests) view returns((address,uint256,uint64)[])
func (_IPerpetualManager *IPerpetualManagerCallerSession) GetWithdrawRequests(poolId uint8, _fromIdx *big.Int, numRequests *big.Int) ([]IPerpetualTreasuryWithdrawRequest, error) {
	return _IPerpetualManager.Contract.GetWithdrawRequests(&_IPerpetualManager.CallOpts, poolId, _fromIdx, numRequests)
}

// HoldingPeriodPenalty is a free data retrieval call binding the contract method 0x24e03d33.
//
// Solidity: function holdingPeriodPenalty(uint256 _numBlockSinceLastOpen, int128 _fLambda) pure returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCaller) HoldingPeriodPenalty(opts *bind.CallOpts, _numBlockSinceLastOpen *big.Int, _fLambda *big.Int) (uint16, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "holdingPeriodPenalty", _numBlockSinceLastOpen, _fLambda)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// HoldingPeriodPenalty is a free data retrieval call binding the contract method 0x24e03d33.
//
// Solidity: function holdingPeriodPenalty(uint256 _numBlockSinceLastOpen, int128 _fLambda) pure returns(uint16)
func (_IPerpetualManager *IPerpetualManagerSession) HoldingPeriodPenalty(_numBlockSinceLastOpen *big.Int, _fLambda *big.Int) (uint16, error) {
	return _IPerpetualManager.Contract.HoldingPeriodPenalty(&_IPerpetualManager.CallOpts, _numBlockSinceLastOpen, _fLambda)
}

// HoldingPeriodPenalty is a free data retrieval call binding the contract method 0x24e03d33.
//
// Solidity: function holdingPeriodPenalty(uint256 _numBlockSinceLastOpen, int128 _fLambda) pure returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCallerSession) HoldingPeriodPenalty(_numBlockSinceLastOpen *big.Int, _fLambda *big.Int) (uint16, error) {
	return _IPerpetualManager.Contract.HoldingPeriodPenalty(&_IPerpetualManager.CallOpts, _numBlockSinceLastOpen, _fLambda)
}

// IsActiveAccount is a free data retrieval call binding the contract method 0x1efb7dc0.
//
// Solidity: function isActiveAccount(uint24 _perpetualId, address _traderAddress) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerCaller) IsActiveAccount(opts *bind.CallOpts, _perpetualId *big.Int, _traderAddress common.Address) (bool, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "isActiveAccount", _perpetualId, _traderAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveAccount is a free data retrieval call binding the contract method 0x1efb7dc0.
//
// Solidity: function isActiveAccount(uint24 _perpetualId, address _traderAddress) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerSession) IsActiveAccount(_perpetualId *big.Int, _traderAddress common.Address) (bool, error) {
	return _IPerpetualManager.Contract.IsActiveAccount(&_IPerpetualManager.CallOpts, _perpetualId, _traderAddress)
}

// IsActiveAccount is a free data retrieval call binding the contract method 0x1efb7dc0.
//
// Solidity: function isActiveAccount(uint24 _perpetualId, address _traderAddress) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerCallerSession) IsActiveAccount(_perpetualId *big.Int, _traderAddress common.Address) (bool, error) {
	return _IPerpetualManager.Contract.IsActiveAccount(&_IPerpetualManager.CallOpts, _perpetualId, _traderAddress)
}

// IsMarketClosed is a free data retrieval call binding the contract method 0x62c6ebd7.
//
// Solidity: function isMarketClosed(bytes4 _baseCurrency, bytes4 _quoteCurrency) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerCaller) IsMarketClosed(opts *bind.CallOpts, _baseCurrency [4]byte, _quoteCurrency [4]byte) (bool, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "isMarketClosed", _baseCurrency, _quoteCurrency)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarketClosed is a free data retrieval call binding the contract method 0x62c6ebd7.
//
// Solidity: function isMarketClosed(bytes4 _baseCurrency, bytes4 _quoteCurrency) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerSession) IsMarketClosed(_baseCurrency [4]byte, _quoteCurrency [4]byte) (bool, error) {
	return _IPerpetualManager.Contract.IsMarketClosed(&_IPerpetualManager.CallOpts, _baseCurrency, _quoteCurrency)
}

// IsMarketClosed is a free data retrieval call binding the contract method 0x62c6ebd7.
//
// Solidity: function isMarketClosed(bytes4 _baseCurrency, bytes4 _quoteCurrency) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerCallerSession) IsMarketClosed(_baseCurrency [4]byte, _quoteCurrency [4]byte) (bool, error) {
	return _IPerpetualManager.Contract.IsMarketClosed(&_IPerpetualManager.CallOpts, _baseCurrency, _quoteCurrency)
}

// IsOrderCanceled is a free data retrieval call binding the contract method 0xb12dff78.
//
// Solidity: function isOrderCanceled(bytes32 digest) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerCaller) IsOrderCanceled(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "isOrderCanceled", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrderCanceled is a free data retrieval call binding the contract method 0xb12dff78.
//
// Solidity: function isOrderCanceled(bytes32 digest) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerSession) IsOrderCanceled(digest [32]byte) (bool, error) {
	return _IPerpetualManager.Contract.IsOrderCanceled(&_IPerpetualManager.CallOpts, digest)
}

// IsOrderCanceled is a free data retrieval call binding the contract method 0xb12dff78.
//
// Solidity: function isOrderCanceled(bytes32 digest) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerCallerSession) IsOrderCanceled(digest [32]byte) (bool, error) {
	return _IPerpetualManager.Contract.IsOrderCanceled(&_IPerpetualManager.CallOpts, digest)
}

// IsOrderExecuted is a free data retrieval call binding the contract method 0x4bbec667.
//
// Solidity: function isOrderExecuted(bytes32 digest) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerCaller) IsOrderExecuted(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "isOrderExecuted", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrderExecuted is a free data retrieval call binding the contract method 0x4bbec667.
//
// Solidity: function isOrderExecuted(bytes32 digest) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerSession) IsOrderExecuted(digest [32]byte) (bool, error) {
	return _IPerpetualManager.Contract.IsOrderExecuted(&_IPerpetualManager.CallOpts, digest)
}

// IsOrderExecuted is a free data retrieval call binding the contract method 0x4bbec667.
//
// Solidity: function isOrderExecuted(bytes32 digest) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerCallerSession) IsOrderExecuted(digest [32]byte) (bool, error) {
	return _IPerpetualManager.Contract.IsOrderExecuted(&_IPerpetualManager.CallOpts, digest)
}

// IsPerpMarketClosed is a free data retrieval call binding the contract method 0x7cb8ff18.
//
// Solidity: function isPerpMarketClosed(uint24 _perpetualId) view returns(bool isClosed)
func (_IPerpetualManager *IPerpetualManagerCaller) IsPerpMarketClosed(opts *bind.CallOpts, _perpetualId *big.Int) (bool, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "isPerpMarketClosed", _perpetualId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPerpMarketClosed is a free data retrieval call binding the contract method 0x7cb8ff18.
//
// Solidity: function isPerpMarketClosed(uint24 _perpetualId) view returns(bool isClosed)
func (_IPerpetualManager *IPerpetualManagerSession) IsPerpMarketClosed(_perpetualId *big.Int) (bool, error) {
	return _IPerpetualManager.Contract.IsPerpMarketClosed(&_IPerpetualManager.CallOpts, _perpetualId)
}

// IsPerpMarketClosed is a free data retrieval call binding the contract method 0x7cb8ff18.
//
// Solidity: function isPerpMarketClosed(uint24 _perpetualId) view returns(bool isClosed)
func (_IPerpetualManager *IPerpetualManagerCallerSession) IsPerpMarketClosed(_perpetualId *big.Int) (bool, error) {
	return _IPerpetualManager.Contract.IsPerpMarketClosed(&_IPerpetualManager.CallOpts, _perpetualId)
}

// IsTraderMaintenanceMarginSafe is a free data retrieval call binding the contract method 0xbfdbeb76.
//
// Solidity: function isTraderMaintenanceMarginSafe(uint24 _perpetualId, address _traderAddress) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerCaller) IsTraderMaintenanceMarginSafe(opts *bind.CallOpts, _perpetualId *big.Int, _traderAddress common.Address) (bool, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "isTraderMaintenanceMarginSafe", _perpetualId, _traderAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTraderMaintenanceMarginSafe is a free data retrieval call binding the contract method 0xbfdbeb76.
//
// Solidity: function isTraderMaintenanceMarginSafe(uint24 _perpetualId, address _traderAddress) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerSession) IsTraderMaintenanceMarginSafe(_perpetualId *big.Int, _traderAddress common.Address) (bool, error) {
	return _IPerpetualManager.Contract.IsTraderMaintenanceMarginSafe(&_IPerpetualManager.CallOpts, _perpetualId, _traderAddress)
}

// IsTraderMaintenanceMarginSafe is a free data retrieval call binding the contract method 0xbfdbeb76.
//
// Solidity: function isTraderMaintenanceMarginSafe(uint24 _perpetualId, address _traderAddress) view returns(bool)
func (_IPerpetualManager *IPerpetualManagerCallerSession) IsTraderMaintenanceMarginSafe(_perpetualId *big.Int, _traderAddress common.Address) (bool, error) {
	return _IPerpetualManager.Contract.IsTraderMaintenanceMarginSafe(&_IPerpetualManager.CallOpts, _perpetualId, _traderAddress)
}

// QueryExchangeFee is a free data retrieval call binding the contract method 0x62028985.
//
// Solidity: function queryExchangeFee(uint8 _poolId, address _traderAddr, address _brokerAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCaller) QueryExchangeFee(opts *bind.CallOpts, _poolId uint8, _traderAddr common.Address, _brokerAddr common.Address) (uint16, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "queryExchangeFee", _poolId, _traderAddr, _brokerAddr)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// QueryExchangeFee is a free data retrieval call binding the contract method 0x62028985.
//
// Solidity: function queryExchangeFee(uint8 _poolId, address _traderAddr, address _brokerAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerSession) QueryExchangeFee(_poolId uint8, _traderAddr common.Address, _brokerAddr common.Address) (uint16, error) {
	return _IPerpetualManager.Contract.QueryExchangeFee(&_IPerpetualManager.CallOpts, _poolId, _traderAddr, _brokerAddr)
}

// QueryExchangeFee is a free data retrieval call binding the contract method 0x62028985.
//
// Solidity: function queryExchangeFee(uint8 _poolId, address _traderAddr, address _brokerAddr) view returns(uint16)
func (_IPerpetualManager *IPerpetualManagerCallerSession) QueryExchangeFee(_poolId uint8, _traderAddr common.Address, _brokerAddr common.Address) (uint16, error) {
	return _IPerpetualManager.Contract.QueryExchangeFee(&_IPerpetualManager.CallOpts, _poolId, _traderAddr, _brokerAddr)
}

// QueryMidPrices is a free data retrieval call binding the contract method 0xa0b6e127.
//
// Solidity: function queryMidPrices(uint24[] perpetualIds, int128[] idxPriceDataPairs) view returns(int128[])
func (_IPerpetualManager *IPerpetualManagerCaller) QueryMidPrices(opts *bind.CallOpts, perpetualIds []*big.Int, idxPriceDataPairs []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "queryMidPrices", perpetualIds, idxPriceDataPairs)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// QueryMidPrices is a free data retrieval call binding the contract method 0xa0b6e127.
//
// Solidity: function queryMidPrices(uint24[] perpetualIds, int128[] idxPriceDataPairs) view returns(int128[])
func (_IPerpetualManager *IPerpetualManagerSession) QueryMidPrices(perpetualIds []*big.Int, idxPriceDataPairs []*big.Int) ([]*big.Int, error) {
	return _IPerpetualManager.Contract.QueryMidPrices(&_IPerpetualManager.CallOpts, perpetualIds, idxPriceDataPairs)
}

// QueryMidPrices is a free data retrieval call binding the contract method 0xa0b6e127.
//
// Solidity: function queryMidPrices(uint24[] perpetualIds, int128[] idxPriceDataPairs) view returns(int128[])
func (_IPerpetualManager *IPerpetualManagerCallerSession) QueryMidPrices(perpetualIds []*big.Int, idxPriceDataPairs []*big.Int) ([]*big.Int, error) {
	return _IPerpetualManager.Contract.QueryMidPrices(&_IPerpetualManager.CallOpts, perpetualIds, idxPriceDataPairs)
}

// QueryPerpetualPrice is a free data retrieval call binding the contract method 0x0a0a3aa7.
//
// Solidity: function queryPerpetualPrice(uint24 _iPerpetualId, int128 _fTradeAmountBC, int128[2] _fIndexPrice) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) QueryPerpetualPrice(opts *bind.CallOpts, _iPerpetualId *big.Int, _fTradeAmountBC *big.Int, _fIndexPrice [2]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "queryPerpetualPrice", _iPerpetualId, _fTradeAmountBC, _fIndexPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryPerpetualPrice is a free data retrieval call binding the contract method 0x0a0a3aa7.
//
// Solidity: function queryPerpetualPrice(uint24 _iPerpetualId, int128 _fTradeAmountBC, int128[2] _fIndexPrice) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) QueryPerpetualPrice(_iPerpetualId *big.Int, _fTradeAmountBC *big.Int, _fIndexPrice [2]*big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.QueryPerpetualPrice(&_IPerpetualManager.CallOpts, _iPerpetualId, _fTradeAmountBC, _fIndexPrice)
}

// QueryPerpetualPrice is a free data retrieval call binding the contract method 0x0a0a3aa7.
//
// Solidity: function queryPerpetualPrice(uint24 _iPerpetualId, int128 _fTradeAmountBC, int128[2] _fIndexPrice) view returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) QueryPerpetualPrice(_iPerpetualId *big.Int, _fTradeAmountBC *big.Int, _fIndexPrice [2]*big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.QueryPerpetualPrice(&_IPerpetualManager.CallOpts, _iPerpetualId, _fTradeAmountBC, _fIndexPrice)
}

// RelativeFeeToCCAmount is a free data retrieval call binding the contract method 0xb7048d59.
//
// Solidity: function relativeFeeToCCAmount(int128 _fDeltaPos, int128 _fTreasuryFeeRate, int128 _fPnLPartRate, int128 _fReferralRebate, address _referrerAddr) pure returns(int128, int128, int128)
func (_IPerpetualManager *IPerpetualManagerCaller) RelativeFeeToCCAmount(opts *bind.CallOpts, _fDeltaPos *big.Int, _fTreasuryFeeRate *big.Int, _fPnLPartRate *big.Int, _fReferralRebate *big.Int, _referrerAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "relativeFeeToCCAmount", _fDeltaPos, _fTreasuryFeeRate, _fPnLPartRate, _fReferralRebate, _referrerAddr)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// RelativeFeeToCCAmount is a free data retrieval call binding the contract method 0xb7048d59.
//
// Solidity: function relativeFeeToCCAmount(int128 _fDeltaPos, int128 _fTreasuryFeeRate, int128 _fPnLPartRate, int128 _fReferralRebate, address _referrerAddr) pure returns(int128, int128, int128)
func (_IPerpetualManager *IPerpetualManagerSession) RelativeFeeToCCAmount(_fDeltaPos *big.Int, _fTreasuryFeeRate *big.Int, _fPnLPartRate *big.Int, _fReferralRebate *big.Int, _referrerAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _IPerpetualManager.Contract.RelativeFeeToCCAmount(&_IPerpetualManager.CallOpts, _fDeltaPos, _fTreasuryFeeRate, _fPnLPartRate, _fReferralRebate, _referrerAddr)
}

// RelativeFeeToCCAmount is a free data retrieval call binding the contract method 0xb7048d59.
//
// Solidity: function relativeFeeToCCAmount(int128 _fDeltaPos, int128 _fTreasuryFeeRate, int128 _fPnLPartRate, int128 _fReferralRebate, address _referrerAddr) pure returns(int128, int128, int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) RelativeFeeToCCAmount(_fDeltaPos *big.Int, _fTreasuryFeeRate *big.Int, _fPnLPartRate *big.Int, _fReferralRebate *big.Int, _referrerAddr common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _IPerpetualManager.Contract.RelativeFeeToCCAmount(&_IPerpetualManager.CallOpts, _fDeltaPos, _fTreasuryFeeRate, _fPnLPartRate, _fReferralRebate, _referrerAddr)
}

// SplitProtocolFee is a free data retrieval call binding the contract method 0xc4519f3b.
//
// Solidity: function splitProtocolFee(uint16 fee) pure returns(int128, int128)
func (_IPerpetualManager *IPerpetualManagerCaller) SplitProtocolFee(opts *bind.CallOpts, fee uint16) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "splitProtocolFee", fee)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// SplitProtocolFee is a free data retrieval call binding the contract method 0xc4519f3b.
//
// Solidity: function splitProtocolFee(uint16 fee) pure returns(int128, int128)
func (_IPerpetualManager *IPerpetualManagerSession) SplitProtocolFee(fee uint16) (*big.Int, *big.Int, error) {
	return _IPerpetualManager.Contract.SplitProtocolFee(&_IPerpetualManager.CallOpts, fee)
}

// SplitProtocolFee is a free data retrieval call binding the contract method 0xc4519f3b.
//
// Solidity: function splitProtocolFee(uint16 fee) pure returns(int128, int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) SplitProtocolFee(fee uint16) (*big.Int, *big.Int, error) {
	return _IPerpetualManager.Contract.SplitProtocolFee(&_IPerpetualManager.CallOpts, fee)
}

// ValidateStopPrice is a free data retrieval call binding the contract method 0x81982b1b.
//
// Solidity: function validateStopPrice(bool _isLong, int128 _fMarkPrice, int128 _fTriggerPrice) pure returns()
func (_IPerpetualManager *IPerpetualManagerCaller) ValidateStopPrice(opts *bind.CallOpts, _isLong bool, _fMarkPrice *big.Int, _fTriggerPrice *big.Int) error {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "validateStopPrice", _isLong, _fMarkPrice, _fTriggerPrice)

	if err != nil {
		return err
	}

	return err

}

// ValidateStopPrice is a free data retrieval call binding the contract method 0x81982b1b.
//
// Solidity: function validateStopPrice(bool _isLong, int128 _fMarkPrice, int128 _fTriggerPrice) pure returns()
func (_IPerpetualManager *IPerpetualManagerSession) ValidateStopPrice(_isLong bool, _fMarkPrice *big.Int, _fTriggerPrice *big.Int) error {
	return _IPerpetualManager.Contract.ValidateStopPrice(&_IPerpetualManager.CallOpts, _isLong, _fMarkPrice, _fTriggerPrice)
}

// ValidateStopPrice is a free data retrieval call binding the contract method 0x81982b1b.
//
// Solidity: function validateStopPrice(bool _isLong, int128 _fMarkPrice, int128 _fTriggerPrice) pure returns()
func (_IPerpetualManager *IPerpetualManagerCallerSession) ValidateStopPrice(_isLong bool, _fMarkPrice *big.Int, _fTriggerPrice *big.Int) error {
	return _IPerpetualManager.Contract.ValidateStopPrice(&_IPerpetualManager.CallOpts, _isLong, _fMarkPrice, _fTriggerPrice)
}

// VolatilitySpread is a free data retrieval call binding the contract method 0x0730b613.
//
// Solidity: function volatilitySpread(uint16 _jumpTbps, uint16 _MinimalSpreadTbps, uint256 _numBlockSinceJump, int128 _fLambda) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerCaller) VolatilitySpread(opts *bind.CallOpts, _jumpTbps uint16, _MinimalSpreadTbps uint16, _numBlockSinceJump *big.Int, _fLambda *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IPerpetualManager.contract.Call(opts, &out, "volatilitySpread", _jumpTbps, _MinimalSpreadTbps, _numBlockSinceJump, _fLambda)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VolatilitySpread is a free data retrieval call binding the contract method 0x0730b613.
//
// Solidity: function volatilitySpread(uint16 _jumpTbps, uint16 _MinimalSpreadTbps, uint256 _numBlockSinceJump, int128 _fLambda) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) VolatilitySpread(_jumpTbps uint16, _MinimalSpreadTbps uint16, _numBlockSinceJump *big.Int, _fLambda *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.VolatilitySpread(&_IPerpetualManager.CallOpts, _jumpTbps, _MinimalSpreadTbps, _numBlockSinceJump, _fLambda)
}

// VolatilitySpread is a free data retrieval call binding the contract method 0x0730b613.
//
// Solidity: function volatilitySpread(uint16 _jumpTbps, uint16 _MinimalSpreadTbps, uint256 _numBlockSinceJump, int128 _fLambda) pure returns(int128)
func (_IPerpetualManager *IPerpetualManagerCallerSession) VolatilitySpread(_jumpTbps uint16, _MinimalSpreadTbps uint16, _numBlockSinceJump *big.Int, _fLambda *big.Int) (*big.Int, error) {
	return _IPerpetualManager.Contract.VolatilitySpread(&_IPerpetualManager.CallOpts, _jumpTbps, _MinimalSpreadTbps, _numBlockSinceJump, _fLambda)
}

// ActivatePerpetual is a paid mutator transaction binding the contract method 0xa65dbc99.
//
// Solidity: function activatePerpetual(uint24 _perpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) ActivatePerpetual(opts *bind.TransactOpts, _perpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "activatePerpetual", _perpetualId)
}

// ActivatePerpetual is a paid mutator transaction binding the contract method 0xa65dbc99.
//
// Solidity: function activatePerpetual(uint24 _perpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerSession) ActivatePerpetual(_perpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.ActivatePerpetual(&_IPerpetualManager.TransactOpts, _perpetualId)
}

// ActivatePerpetual is a paid mutator transaction binding the contract method 0xa65dbc99.
//
// Solidity: function activatePerpetual(uint24 _perpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) ActivatePerpetual(_perpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.ActivatePerpetual(&_IPerpetualManager.TransactOpts, _perpetualId)
}

// AddAMMLiquidityToPerpetual is a paid mutator transaction binding the contract method 0x50403b45.
//
// Solidity: function addAMMLiquidityToPerpetual(uint24 _iPerpetualId, int128 _fTokenAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) AddAMMLiquidityToPerpetual(opts *bind.TransactOpts, _iPerpetualId *big.Int, _fTokenAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "addAMMLiquidityToPerpetual", _iPerpetualId, _fTokenAmount)
}

// AddAMMLiquidityToPerpetual is a paid mutator transaction binding the contract method 0x50403b45.
//
// Solidity: function addAMMLiquidityToPerpetual(uint24 _iPerpetualId, int128 _fTokenAmount) returns()
func (_IPerpetualManager *IPerpetualManagerSession) AddAMMLiquidityToPerpetual(_iPerpetualId *big.Int, _fTokenAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.AddAMMLiquidityToPerpetual(&_IPerpetualManager.TransactOpts, _iPerpetualId, _fTokenAmount)
}

// AddAMMLiquidityToPerpetual is a paid mutator transaction binding the contract method 0x50403b45.
//
// Solidity: function addAMMLiquidityToPerpetual(uint24 _iPerpetualId, int128 _fTokenAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) AddAMMLiquidityToPerpetual(_iPerpetualId *big.Int, _fTokenAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.AddAMMLiquidityToPerpetual(&_IPerpetualManager.TransactOpts, _iPerpetualId, _fTokenAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x8cd81c00.
//
// Solidity: function addLiquidity(uint8 _iPoolIndex, uint256 _tokenAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) AddLiquidity(opts *bind.TransactOpts, _iPoolIndex uint8, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "addLiquidity", _iPoolIndex, _tokenAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x8cd81c00.
//
// Solidity: function addLiquidity(uint8 _iPoolIndex, uint256 _tokenAmount) returns()
func (_IPerpetualManager *IPerpetualManagerSession) AddLiquidity(_iPoolIndex uint8, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.AddLiquidity(&_IPerpetualManager.TransactOpts, _iPoolIndex, _tokenAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x8cd81c00.
//
// Solidity: function addLiquidity(uint8 _iPoolIndex, uint256 _tokenAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) AddLiquidity(_iPoolIndex uint8, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.AddLiquidity(&_IPerpetualManager.TransactOpts, _iPoolIndex, _tokenAmount)
}

// AdjustSettlementPrice is a paid mutator transaction binding the contract method 0xce985429.
//
// Solidity: function adjustSettlementPrice(uint24 _perpetualId, int128 _fSettlementS2, int128 _fSettlementS3) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) AdjustSettlementPrice(opts *bind.TransactOpts, _perpetualId *big.Int, _fSettlementS2 *big.Int, _fSettlementS3 *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "adjustSettlementPrice", _perpetualId, _fSettlementS2, _fSettlementS3)
}

// AdjustSettlementPrice is a paid mutator transaction binding the contract method 0xce985429.
//
// Solidity: function adjustSettlementPrice(uint24 _perpetualId, int128 _fSettlementS2, int128 _fSettlementS3) returns()
func (_IPerpetualManager *IPerpetualManagerSession) AdjustSettlementPrice(_perpetualId *big.Int, _fSettlementS2 *big.Int, _fSettlementS3 *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.AdjustSettlementPrice(&_IPerpetualManager.TransactOpts, _perpetualId, _fSettlementS2, _fSettlementS3)
}

// AdjustSettlementPrice is a paid mutator transaction binding the contract method 0xce985429.
//
// Solidity: function adjustSettlementPrice(uint24 _perpetualId, int128 _fSettlementS2, int128 _fSettlementS3) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) AdjustSettlementPrice(_perpetualId *big.Int, _fSettlementS2 *big.Int, _fSettlementS3 *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.AdjustSettlementPrice(&_IPerpetualManager.TransactOpts, _perpetualId, _fSettlementS2, _fSettlementS3)
}

// BrokerDepositToDefaultFund is a paid mutator transaction binding the contract method 0x91a2fd59.
//
// Solidity: function brokerDepositToDefaultFund(uint8 _poolId, uint32 _iLots) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) BrokerDepositToDefaultFund(opts *bind.TransactOpts, _poolId uint8, _iLots uint32) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "brokerDepositToDefaultFund", _poolId, _iLots)
}

// BrokerDepositToDefaultFund is a paid mutator transaction binding the contract method 0x91a2fd59.
//
// Solidity: function brokerDepositToDefaultFund(uint8 _poolId, uint32 _iLots) returns()
func (_IPerpetualManager *IPerpetualManagerSession) BrokerDepositToDefaultFund(_poolId uint8, _iLots uint32) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.BrokerDepositToDefaultFund(&_IPerpetualManager.TransactOpts, _poolId, _iLots)
}

// BrokerDepositToDefaultFund is a paid mutator transaction binding the contract method 0x91a2fd59.
//
// Solidity: function brokerDepositToDefaultFund(uint8 _poolId, uint32 _iLots) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) BrokerDepositToDefaultFund(_poolId uint8, _iLots uint32) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.BrokerDepositToDefaultFund(&_IPerpetualManager.TransactOpts, _poolId, _iLots)
}

// ChargePostingFee is a paid mutator transaction binding the contract method 0x5ad59a1c.
//
// Solidity: function chargePostingFee((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order, uint16 _feeTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) ChargePostingFee(opts *bind.TransactOpts, _order IPerpetualOrderOrder, _feeTbps uint16) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "chargePostingFee", _order, _feeTbps)
}

// ChargePostingFee is a paid mutator transaction binding the contract method 0x5ad59a1c.
//
// Solidity: function chargePostingFee((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order, uint16 _feeTbps) returns()
func (_IPerpetualManager *IPerpetualManagerSession) ChargePostingFee(_order IPerpetualOrderOrder, _feeTbps uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.ChargePostingFee(&_IPerpetualManager.TransactOpts, _order, _feeTbps)
}

// ChargePostingFee is a paid mutator transaction binding the contract method 0x5ad59a1c.
//
// Solidity: function chargePostingFee((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order, uint16 _feeTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) ChargePostingFee(_order IPerpetualOrderOrder, _feeTbps uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.ChargePostingFee(&_IPerpetualManager.TransactOpts, _order, _feeTbps)
}

// CreateLiquidityPool is a paid mutator transaction binding the contract method 0xf20bcadb.
//
// Solidity: function createLiquidityPool(address _marginTokenAddress, uint16 _iTargetPoolSizeUpdateTime, int128 _fMaxTransferPerConvergencePeriod, int128 _fBrokerCollateralLotSize) returns(uint8)
func (_IPerpetualManager *IPerpetualManagerTransactor) CreateLiquidityPool(opts *bind.TransactOpts, _marginTokenAddress common.Address, _iTargetPoolSizeUpdateTime uint16, _fMaxTransferPerConvergencePeriod *big.Int, _fBrokerCollateralLotSize *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "createLiquidityPool", _marginTokenAddress, _iTargetPoolSizeUpdateTime, _fMaxTransferPerConvergencePeriod, _fBrokerCollateralLotSize)
}

// CreateLiquidityPool is a paid mutator transaction binding the contract method 0xf20bcadb.
//
// Solidity: function createLiquidityPool(address _marginTokenAddress, uint16 _iTargetPoolSizeUpdateTime, int128 _fMaxTransferPerConvergencePeriod, int128 _fBrokerCollateralLotSize) returns(uint8)
func (_IPerpetualManager *IPerpetualManagerSession) CreateLiquidityPool(_marginTokenAddress common.Address, _iTargetPoolSizeUpdateTime uint16, _fMaxTransferPerConvergencePeriod *big.Int, _fBrokerCollateralLotSize *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.CreateLiquidityPool(&_IPerpetualManager.TransactOpts, _marginTokenAddress, _iTargetPoolSizeUpdateTime, _fMaxTransferPerConvergencePeriod, _fBrokerCollateralLotSize)
}

// CreateLiquidityPool is a paid mutator transaction binding the contract method 0xf20bcadb.
//
// Solidity: function createLiquidityPool(address _marginTokenAddress, uint16 _iTargetPoolSizeUpdateTime, int128 _fMaxTransferPerConvergencePeriod, int128 _fBrokerCollateralLotSize) returns(uint8)
func (_IPerpetualManager *IPerpetualManagerTransactorSession) CreateLiquidityPool(_marginTokenAddress common.Address, _iTargetPoolSizeUpdateTime uint16, _fMaxTransferPerConvergencePeriod *big.Int, _fBrokerCollateralLotSize *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.CreateLiquidityPool(&_IPerpetualManager.TransactOpts, _marginTokenAddress, _iTargetPoolSizeUpdateTime, _fMaxTransferPerConvergencePeriod, _fBrokerCollateralLotSize)
}

// CreatePerpetual is a paid mutator transaction binding the contract method 0x14ef624e.
//
// Solidity: function createPerpetual(uint8 _iPoolId, bytes4[2] _baseQuoteS2, bytes4[2] _baseQuoteS3, int128[7] _baseParams, int128[5] _underlyingRiskParams, int128[13] _defaultFundRiskParams, uint256 _eCollateralCurrency) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) CreatePerpetual(opts *bind.TransactOpts, _iPoolId uint8, _baseQuoteS2 [2][4]byte, _baseQuoteS3 [2][4]byte, _baseParams [7]*big.Int, _underlyingRiskParams [5]*big.Int, _defaultFundRiskParams [13]*big.Int, _eCollateralCurrency *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "createPerpetual", _iPoolId, _baseQuoteS2, _baseQuoteS3, _baseParams, _underlyingRiskParams, _defaultFundRiskParams, _eCollateralCurrency)
}

// CreatePerpetual is a paid mutator transaction binding the contract method 0x14ef624e.
//
// Solidity: function createPerpetual(uint8 _iPoolId, bytes4[2] _baseQuoteS2, bytes4[2] _baseQuoteS3, int128[7] _baseParams, int128[5] _underlyingRiskParams, int128[13] _defaultFundRiskParams, uint256 _eCollateralCurrency) returns()
func (_IPerpetualManager *IPerpetualManagerSession) CreatePerpetual(_iPoolId uint8, _baseQuoteS2 [2][4]byte, _baseQuoteS3 [2][4]byte, _baseParams [7]*big.Int, _underlyingRiskParams [5]*big.Int, _defaultFundRiskParams [13]*big.Int, _eCollateralCurrency *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.CreatePerpetual(&_IPerpetualManager.TransactOpts, _iPoolId, _baseQuoteS2, _baseQuoteS3, _baseParams, _underlyingRiskParams, _defaultFundRiskParams, _eCollateralCurrency)
}

// CreatePerpetual is a paid mutator transaction binding the contract method 0x14ef624e.
//
// Solidity: function createPerpetual(uint8 _iPoolId, bytes4[2] _baseQuoteS2, bytes4[2] _baseQuoteS3, int128[7] _baseParams, int128[5] _underlyingRiskParams, int128[13] _defaultFundRiskParams, uint256 _eCollateralCurrency) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) CreatePerpetual(_iPoolId uint8, _baseQuoteS2 [2][4]byte, _baseQuoteS3 [2][4]byte, _baseParams [7]*big.Int, _underlyingRiskParams [5]*big.Int, _defaultFundRiskParams [13]*big.Int, _eCollateralCurrency *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.CreatePerpetual(&_IPerpetualManager.TransactOpts, _iPoolId, _baseQuoteS2, _baseQuoteS3, _baseParams, _underlyingRiskParams, _defaultFundRiskParams, _eCollateralCurrency)
}

// DecreasePoolCash is a paid mutator transaction binding the contract method 0x733e8816.
//
// Solidity: function decreasePoolCash(uint8 _iPoolIdx, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) DecreasePoolCash(opts *bind.TransactOpts, _iPoolIdx uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "decreasePoolCash", _iPoolIdx, _fAmount)
}

// DecreasePoolCash is a paid mutator transaction binding the contract method 0x733e8816.
//
// Solidity: function decreasePoolCash(uint8 _iPoolIdx, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerSession) DecreasePoolCash(_iPoolIdx uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.DecreasePoolCash(&_IPerpetualManager.TransactOpts, _iPoolIdx, _fAmount)
}

// DecreasePoolCash is a paid mutator transaction binding the contract method 0x733e8816.
//
// Solidity: function decreasePoolCash(uint8 _iPoolIdx, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) DecreasePoolCash(_iPoolIdx uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.DecreasePoolCash(&_IPerpetualManager.TransactOpts, _iPoolIdx, _fAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1285bf92.
//
// Solidity: function deposit(uint24 _iPerpetualId, int128 _fAmount, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) Deposit(opts *bind.TransactOpts, _iPerpetualId *big.Int, _fAmount *big.Int, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "deposit", _iPerpetualId, _fAmount, _updateData, _publishTimes)
}

// Deposit is a paid mutator transaction binding the contract method 0x1285bf92.
//
// Solidity: function deposit(uint24 _iPerpetualId, int128 _fAmount, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_IPerpetualManager *IPerpetualManagerSession) Deposit(_iPerpetualId *big.Int, _fAmount *big.Int, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.Deposit(&_IPerpetualManager.TransactOpts, _iPerpetualId, _fAmount, _updateData, _publishTimes)
}

// Deposit is a paid mutator transaction binding the contract method 0x1285bf92.
//
// Solidity: function deposit(uint24 _iPerpetualId, int128 _fAmount, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) Deposit(_iPerpetualId *big.Int, _fAmount *big.Int, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.Deposit(&_IPerpetualManager.TransactOpts, _iPerpetualId, _fAmount, _updateData, _publishTimes)
}

// DepositMarginForOpeningTrade is a paid mutator transaction binding the contract method 0xce5c4563.
//
// Solidity: function depositMarginForOpeningTrade(uint24 _iPerpetualId, int128 _fDepositRequired, (uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order) returns(bool)
func (_IPerpetualManager *IPerpetualManagerTransactor) DepositMarginForOpeningTrade(opts *bind.TransactOpts, _iPerpetualId *big.Int, _fDepositRequired *big.Int, _order IPerpetualOrderOrder) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "depositMarginForOpeningTrade", _iPerpetualId, _fDepositRequired, _order)
}

// DepositMarginForOpeningTrade is a paid mutator transaction binding the contract method 0xce5c4563.
//
// Solidity: function depositMarginForOpeningTrade(uint24 _iPerpetualId, int128 _fDepositRequired, (uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order) returns(bool)
func (_IPerpetualManager *IPerpetualManagerSession) DepositMarginForOpeningTrade(_iPerpetualId *big.Int, _fDepositRequired *big.Int, _order IPerpetualOrderOrder) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.DepositMarginForOpeningTrade(&_IPerpetualManager.TransactOpts, _iPerpetualId, _fDepositRequired, _order)
}

// DepositMarginForOpeningTrade is a paid mutator transaction binding the contract method 0xce5c4563.
//
// Solidity: function depositMarginForOpeningTrade(uint24 _iPerpetualId, int128 _fDepositRequired, (uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order) returns(bool)
func (_IPerpetualManager *IPerpetualManagerTransactorSession) DepositMarginForOpeningTrade(_iPerpetualId *big.Int, _fDepositRequired *big.Int, _order IPerpetualOrderOrder) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.DepositMarginForOpeningTrade(&_IPerpetualManager.TransactOpts, _iPerpetualId, _fDepositRequired, _order)
}

// DepositToDefaultFund is a paid mutator transaction binding the contract method 0xc63f7935.
//
// Solidity: function depositToDefaultFund(uint8 _poolId, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) DepositToDefaultFund(opts *bind.TransactOpts, _poolId uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "depositToDefaultFund", _poolId, _fAmount)
}

// DepositToDefaultFund is a paid mutator transaction binding the contract method 0xc63f7935.
//
// Solidity: function depositToDefaultFund(uint8 _poolId, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerSession) DepositToDefaultFund(_poolId uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.DepositToDefaultFund(&_IPerpetualManager.TransactOpts, _poolId, _fAmount)
}

// DepositToDefaultFund is a paid mutator transaction binding the contract method 0xc63f7935.
//
// Solidity: function depositToDefaultFund(uint8 _poolId, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) DepositToDefaultFund(_poolId uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.DepositToDefaultFund(&_IPerpetualManager.TransactOpts, _poolId, _fAmount)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x21ba98a8.
//
// Solidity: function distributeFees((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order, bool _hasOpened) returns(int128)
func (_IPerpetualManager *IPerpetualManagerTransactor) DistributeFees(opts *bind.TransactOpts, _order IPerpetualOrderOrder, _hasOpened bool) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "distributeFees", _order, _hasOpened)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x21ba98a8.
//
// Solidity: function distributeFees((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order, bool _hasOpened) returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) DistributeFees(_order IPerpetualOrderOrder, _hasOpened bool) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.DistributeFees(&_IPerpetualManager.TransactOpts, _order, _hasOpened)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x21ba98a8.
//
// Solidity: function distributeFees((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order, bool _hasOpened) returns(int128)
func (_IPerpetualManager *IPerpetualManagerTransactorSession) DistributeFees(_order IPerpetualOrderOrder, _hasOpened bool) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.DistributeFees(&_IPerpetualManager.TransactOpts, _order, _hasOpened)
}

// DistributeFeesLiquidation is a paid mutator transaction binding the contract method 0x9baac3c4.
//
// Solidity: function distributeFeesLiquidation(uint24 _iPerpetualId, address _traderAddr, int128 _fDeltaPositionBC) returns(int128)
func (_IPerpetualManager *IPerpetualManagerTransactor) DistributeFeesLiquidation(opts *bind.TransactOpts, _iPerpetualId *big.Int, _traderAddr common.Address, _fDeltaPositionBC *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "distributeFeesLiquidation", _iPerpetualId, _traderAddr, _fDeltaPositionBC)
}

// DistributeFeesLiquidation is a paid mutator transaction binding the contract method 0x9baac3c4.
//
// Solidity: function distributeFeesLiquidation(uint24 _iPerpetualId, address _traderAddr, int128 _fDeltaPositionBC) returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) DistributeFeesLiquidation(_iPerpetualId *big.Int, _traderAddr common.Address, _fDeltaPositionBC *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.DistributeFeesLiquidation(&_IPerpetualManager.TransactOpts, _iPerpetualId, _traderAddr, _fDeltaPositionBC)
}

// DistributeFeesLiquidation is a paid mutator transaction binding the contract method 0x9baac3c4.
//
// Solidity: function distributeFeesLiquidation(uint24 _iPerpetualId, address _traderAddr, int128 _fDeltaPositionBC) returns(int128)
func (_IPerpetualManager *IPerpetualManagerTransactorSession) DistributeFeesLiquidation(_iPerpetualId *big.Int, _traderAddr common.Address, _fDeltaPositionBC *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.DistributeFeesLiquidation(&_IPerpetualManager.TransactOpts, _iPerpetualId, _traderAddr, _fDeltaPositionBC)
}

// ExecuteCancelOrder is a paid mutator transaction binding the contract method 0x7200c7c5.
//
// Solidity: function executeCancelOrder(uint24 _perpetualId, bytes32 _digest) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) ExecuteCancelOrder(opts *bind.TransactOpts, _perpetualId *big.Int, _digest [32]byte) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "executeCancelOrder", _perpetualId, _digest)
}

// ExecuteCancelOrder is a paid mutator transaction binding the contract method 0x7200c7c5.
//
// Solidity: function executeCancelOrder(uint24 _perpetualId, bytes32 _digest) returns()
func (_IPerpetualManager *IPerpetualManagerSession) ExecuteCancelOrder(_perpetualId *big.Int, _digest [32]byte) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.ExecuteCancelOrder(&_IPerpetualManager.TransactOpts, _perpetualId, _digest)
}

// ExecuteCancelOrder is a paid mutator transaction binding the contract method 0x7200c7c5.
//
// Solidity: function executeCancelOrder(uint24 _perpetualId, bytes32 _digest) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) ExecuteCancelOrder(_perpetualId *big.Int, _digest [32]byte) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.ExecuteCancelOrder(&_IPerpetualManager.TransactOpts, _perpetualId, _digest)
}

// ExecuteLiquidityWithdrawal is a paid mutator transaction binding the contract method 0xf0064f37.
//
// Solidity: function executeLiquidityWithdrawal(uint8 _poolId, address _lpAddr) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) ExecuteLiquidityWithdrawal(opts *bind.TransactOpts, _poolId uint8, _lpAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "executeLiquidityWithdrawal", _poolId, _lpAddr)
}

// ExecuteLiquidityWithdrawal is a paid mutator transaction binding the contract method 0xf0064f37.
//
// Solidity: function executeLiquidityWithdrawal(uint8 _poolId, address _lpAddr) returns()
func (_IPerpetualManager *IPerpetualManagerSession) ExecuteLiquidityWithdrawal(_poolId uint8, _lpAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.ExecuteLiquidityWithdrawal(&_IPerpetualManager.TransactOpts, _poolId, _lpAddr)
}

// ExecuteLiquidityWithdrawal is a paid mutator transaction binding the contract method 0xf0064f37.
//
// Solidity: function executeLiquidityWithdrawal(uint8 _poolId, address _lpAddr) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) ExecuteLiquidityWithdrawal(_poolId uint8, _lpAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.ExecuteLiquidityWithdrawal(&_IPerpetualManager.TransactOpts, _poolId, _lpAddr)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0xc9636384.
//
// Solidity: function executeTrade(uint24 _iPerpetualId, address _traderAddr, int128 _fTraderPos, int128 _fTradeAmount, int128 _fPrice, bool _isClose) returns(int128)
func (_IPerpetualManager *IPerpetualManagerTransactor) ExecuteTrade(opts *bind.TransactOpts, _iPerpetualId *big.Int, _traderAddr common.Address, _fTraderPos *big.Int, _fTradeAmount *big.Int, _fPrice *big.Int, _isClose bool) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "executeTrade", _iPerpetualId, _traderAddr, _fTraderPos, _fTradeAmount, _fPrice, _isClose)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0xc9636384.
//
// Solidity: function executeTrade(uint24 _iPerpetualId, address _traderAddr, int128 _fTraderPos, int128 _fTradeAmount, int128 _fPrice, bool _isClose) returns(int128)
func (_IPerpetualManager *IPerpetualManagerSession) ExecuteTrade(_iPerpetualId *big.Int, _traderAddr common.Address, _fTraderPos *big.Int, _fTradeAmount *big.Int, _fPrice *big.Int, _isClose bool) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.ExecuteTrade(&_IPerpetualManager.TransactOpts, _iPerpetualId, _traderAddr, _fTraderPos, _fTradeAmount, _fPrice, _isClose)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0xc9636384.
//
// Solidity: function executeTrade(uint24 _iPerpetualId, address _traderAddr, int128 _fTraderPos, int128 _fTradeAmount, int128 _fPrice, bool _isClose) returns(int128)
func (_IPerpetualManager *IPerpetualManagerTransactorSession) ExecuteTrade(_iPerpetualId *big.Int, _traderAddr common.Address, _fTraderPos *big.Int, _fTradeAmount *big.Int, _fPrice *big.Int, _isClose bool) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.ExecuteTrade(&_IPerpetualManager.TransactOpts, _iPerpetualId, _traderAddr, _fTraderPos, _fTradeAmount, _fPrice, _isClose)
}

// IncreasePoolCash is a paid mutator transaction binding the contract method 0x00db5b22.
//
// Solidity: function increasePoolCash(uint8 _iPoolIdx, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) IncreasePoolCash(opts *bind.TransactOpts, _iPoolIdx uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "increasePoolCash", _iPoolIdx, _fAmount)
}

// IncreasePoolCash is a paid mutator transaction binding the contract method 0x00db5b22.
//
// Solidity: function increasePoolCash(uint8 _iPoolIdx, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerSession) IncreasePoolCash(_iPoolIdx uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.IncreasePoolCash(&_IPerpetualManager.TransactOpts, _iPoolIdx, _fAmount)
}

// IncreasePoolCash is a paid mutator transaction binding the contract method 0x00db5b22.
//
// Solidity: function increasePoolCash(uint8 _iPoolIdx, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) IncreasePoolCash(_iPoolIdx uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.IncreasePoolCash(&_IPerpetualManager.TransactOpts, _iPoolIdx, _fAmount)
}

// LiquidateByAMM is a paid mutator transaction binding the contract method 0xb8d56eb3.
//
// Solidity: function liquidateByAMM(uint24 _perpetualIndex, address _liquidatorAddr, address _traderAddr, bytes[] _updateData, uint64[] _publishTimes) payable returns(int128 liquidatedAmount)
func (_IPerpetualManager *IPerpetualManagerTransactor) LiquidateByAMM(opts *bind.TransactOpts, _perpetualIndex *big.Int, _liquidatorAddr common.Address, _traderAddr common.Address, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "liquidateByAMM", _perpetualIndex, _liquidatorAddr, _traderAddr, _updateData, _publishTimes)
}

// LiquidateByAMM is a paid mutator transaction binding the contract method 0xb8d56eb3.
//
// Solidity: function liquidateByAMM(uint24 _perpetualIndex, address _liquidatorAddr, address _traderAddr, bytes[] _updateData, uint64[] _publishTimes) payable returns(int128 liquidatedAmount)
func (_IPerpetualManager *IPerpetualManagerSession) LiquidateByAMM(_perpetualIndex *big.Int, _liquidatorAddr common.Address, _traderAddr common.Address, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.LiquidateByAMM(&_IPerpetualManager.TransactOpts, _perpetualIndex, _liquidatorAddr, _traderAddr, _updateData, _publishTimes)
}

// LiquidateByAMM is a paid mutator transaction binding the contract method 0xb8d56eb3.
//
// Solidity: function liquidateByAMM(uint24 _perpetualIndex, address _liquidatorAddr, address _traderAddr, bytes[] _updateData, uint64[] _publishTimes) payable returns(int128 liquidatedAmount)
func (_IPerpetualManager *IPerpetualManagerTransactorSession) LiquidateByAMM(_perpetualIndex *big.Int, _liquidatorAddr common.Address, _traderAddr common.Address, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.LiquidateByAMM(&_IPerpetualManager.TransactOpts, _perpetualIndex, _liquidatorAddr, _traderAddr, _updateData, _publishTimes)
}

// PauseLiquidityProvision is a paid mutator transaction binding the contract method 0x59a39709.
//
// Solidity: function pauseLiquidityProvision(uint8 _poolId, bool _pauseOn) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) PauseLiquidityProvision(opts *bind.TransactOpts, _poolId uint8, _pauseOn bool) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "pauseLiquidityProvision", _poolId, _pauseOn)
}

// PauseLiquidityProvision is a paid mutator transaction binding the contract method 0x59a39709.
//
// Solidity: function pauseLiquidityProvision(uint8 _poolId, bool _pauseOn) returns()
func (_IPerpetualManager *IPerpetualManagerSession) PauseLiquidityProvision(_poolId uint8, _pauseOn bool) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.PauseLiquidityProvision(&_IPerpetualManager.TransactOpts, _poolId, _pauseOn)
}

// PauseLiquidityProvision is a paid mutator transaction binding the contract method 0x59a39709.
//
// Solidity: function pauseLiquidityProvision(uint8 _poolId, bool _pauseOn) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) PauseLiquidityProvision(_poolId uint8, _pauseOn bool) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.PauseLiquidityProvision(&_IPerpetualManager.TransactOpts, _poolId, _pauseOn)
}

// PreTrade is a paid mutator transaction binding the contract method 0xa600b9bd.
//
// Solidity: function preTrade(uint24 _iPerpetualId, (uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order) returns(int128, int128)
func (_IPerpetualManager *IPerpetualManagerTransactor) PreTrade(opts *bind.TransactOpts, _iPerpetualId *big.Int, _order IPerpetualOrderOrder) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "preTrade", _iPerpetualId, _order)
}

// PreTrade is a paid mutator transaction binding the contract method 0xa600b9bd.
//
// Solidity: function preTrade(uint24 _iPerpetualId, (uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order) returns(int128, int128)
func (_IPerpetualManager *IPerpetualManagerSession) PreTrade(_iPerpetualId *big.Int, _order IPerpetualOrderOrder) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.PreTrade(&_IPerpetualManager.TransactOpts, _iPerpetualId, _order)
}

// PreTrade is a paid mutator transaction binding the contract method 0xa600b9bd.
//
// Solidity: function preTrade(uint24 _iPerpetualId, (uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order) returns(int128, int128)
func (_IPerpetualManager *IPerpetualManagerTransactorSession) PreTrade(_iPerpetualId *big.Int, _order IPerpetualOrderOrder) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.PreTrade(&_IPerpetualManager.TransactOpts, _iPerpetualId, _order)
}

// Rebalance is a paid mutator transaction binding the contract method 0xa25d6a28.
//
// Solidity: function rebalance(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) Rebalance(opts *bind.TransactOpts, _iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "rebalance", _iPerpetualId)
}

// Rebalance is a paid mutator transaction binding the contract method 0xa25d6a28.
//
// Solidity: function rebalance(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerSession) Rebalance(_iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.Rebalance(&_IPerpetualManager.TransactOpts, _iPerpetualId)
}

// Rebalance is a paid mutator transaction binding the contract method 0xa25d6a28.
//
// Solidity: function rebalance(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) Rebalance(_iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.Rebalance(&_IPerpetualManager.TransactOpts, _iPerpetualId)
}

// RebatePostingFee is a paid mutator transaction binding the contract method 0xaa0a9079.
//
// Solidity: function rebatePostingFee((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order, uint16 _feeTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) RebatePostingFee(opts *bind.TransactOpts, _order IPerpetualOrderOrder, _feeTbps uint16) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "rebatePostingFee", _order, _feeTbps)
}

// RebatePostingFee is a paid mutator transaction binding the contract method 0xaa0a9079.
//
// Solidity: function rebatePostingFee((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order, uint16 _feeTbps) returns()
func (_IPerpetualManager *IPerpetualManagerSession) RebatePostingFee(_order IPerpetualOrderOrder, _feeTbps uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.RebatePostingFee(&_IPerpetualManager.TransactOpts, _order, _feeTbps)
}

// RebatePostingFee is a paid mutator transaction binding the contract method 0xaa0a9079.
//
// Solidity: function rebatePostingFee((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order, uint16 _feeTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) RebatePostingFee(_order IPerpetualOrderOrder, _feeTbps uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.RebatePostingFee(&_IPerpetualManager.TransactOpts, _order, _feeTbps)
}

// ReduceMarginCollateral is a paid mutator transaction binding the contract method 0x796596fc.
//
// Solidity: function reduceMarginCollateral(uint24 _iPerpetualId, address _traderAddr, int128 _fAmountToWithdraw) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) ReduceMarginCollateral(opts *bind.TransactOpts, _iPerpetualId *big.Int, _traderAddr common.Address, _fAmountToWithdraw *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "reduceMarginCollateral", _iPerpetualId, _traderAddr, _fAmountToWithdraw)
}

// ReduceMarginCollateral is a paid mutator transaction binding the contract method 0x796596fc.
//
// Solidity: function reduceMarginCollateral(uint24 _iPerpetualId, address _traderAddr, int128 _fAmountToWithdraw) returns()
func (_IPerpetualManager *IPerpetualManagerSession) ReduceMarginCollateral(_iPerpetualId *big.Int, _traderAddr common.Address, _fAmountToWithdraw *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.ReduceMarginCollateral(&_IPerpetualManager.TransactOpts, _iPerpetualId, _traderAddr, _fAmountToWithdraw)
}

// ReduceMarginCollateral is a paid mutator transaction binding the contract method 0x796596fc.
//
// Solidity: function reduceMarginCollateral(uint24 _iPerpetualId, address _traderAddr, int128 _fAmountToWithdraw) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) ReduceMarginCollateral(_iPerpetualId *big.Int, _traderAddr common.Address, _fAmountToWithdraw *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.ReduceMarginCollateral(&_IPerpetualManager.TransactOpts, _iPerpetualId, _traderAddr, _fAmountToWithdraw)
}

// RunLiquidityPool is a paid mutator transaction binding the contract method 0x5e4555b8.
//
// Solidity: function runLiquidityPool(uint8 _liqPoolID) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) RunLiquidityPool(opts *bind.TransactOpts, _liqPoolID uint8) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "runLiquidityPool", _liqPoolID)
}

// RunLiquidityPool is a paid mutator transaction binding the contract method 0x5e4555b8.
//
// Solidity: function runLiquidityPool(uint8 _liqPoolID) returns()
func (_IPerpetualManager *IPerpetualManagerSession) RunLiquidityPool(_liqPoolID uint8) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.RunLiquidityPool(&_IPerpetualManager.TransactOpts, _liqPoolID)
}

// RunLiquidityPool is a paid mutator transaction binding the contract method 0x5e4555b8.
//
// Solidity: function runLiquidityPool(uint8 _liqPoolID) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) RunLiquidityPool(_liqPoolID uint8) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.RunLiquidityPool(&_IPerpetualManager.TransactOpts, _liqPoolID)
}

// SetAMMPerpLogic is a paid mutator transaction binding the contract method 0xfc67e5d9.
//
// Solidity: function setAMMPerpLogic(address _AMMPerpLogic) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetAMMPerpLogic(opts *bind.TransactOpts, _AMMPerpLogic common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setAMMPerpLogic", _AMMPerpLogic)
}

// SetAMMPerpLogic is a paid mutator transaction binding the contract method 0xfc67e5d9.
//
// Solidity: function setAMMPerpLogic(address _AMMPerpLogic) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetAMMPerpLogic(_AMMPerpLogic common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetAMMPerpLogic(&_IPerpetualManager.TransactOpts, _AMMPerpLogic)
}

// SetAMMPerpLogic is a paid mutator transaction binding the contract method 0xfc67e5d9.
//
// Solidity: function setAMMPerpLogic(address _AMMPerpLogic) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetAMMPerpLogic(_AMMPerpLogic common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetAMMPerpLogic(&_IPerpetualManager.TransactOpts, _AMMPerpLogic)
}

// SetBlockDelay is a paid mutator transaction binding the contract method 0xaac90c91.
//
// Solidity: function setBlockDelay(uint8 _delay) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetBlockDelay(opts *bind.TransactOpts, _delay uint8) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setBlockDelay", _delay)
}

// SetBlockDelay is a paid mutator transaction binding the contract method 0xaac90c91.
//
// Solidity: function setBlockDelay(uint8 _delay) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetBlockDelay(_delay uint8) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetBlockDelay(&_IPerpetualManager.TransactOpts, _delay)
}

// SetBlockDelay is a paid mutator transaction binding the contract method 0xaac90c91.
//
// Solidity: function setBlockDelay(uint8 _delay) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetBlockDelay(_delay uint8) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetBlockDelay(&_IPerpetualManager.TransactOpts, _delay)
}

// SetBrokerTiers is a paid mutator transaction binding the contract method 0x37e76c90.
//
// Solidity: function setBrokerTiers(uint256[] _tiers, uint16[] _feesTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetBrokerTiers(opts *bind.TransactOpts, _tiers []*big.Int, _feesTbps []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setBrokerTiers", _tiers, _feesTbps)
}

// SetBrokerTiers is a paid mutator transaction binding the contract method 0x37e76c90.
//
// Solidity: function setBrokerTiers(uint256[] _tiers, uint16[] _feesTbps) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetBrokerTiers(_tiers []*big.Int, _feesTbps []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetBrokerTiers(&_IPerpetualManager.TransactOpts, _tiers, _feesTbps)
}

// SetBrokerTiers is a paid mutator transaction binding the contract method 0x37e76c90.
//
// Solidity: function setBrokerTiers(uint256[] _tiers, uint16[] _feesTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetBrokerTiers(_tiers []*big.Int, _feesTbps []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetBrokerTiers(&_IPerpetualManager.TransactOpts, _tiers, _feesTbps)
}

// SetBrokerVolumeTiers is a paid mutator transaction binding the contract method 0x88e27e2f.
//
// Solidity: function setBrokerVolumeTiers(uint256[] _tiers, uint16[] _feesTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetBrokerVolumeTiers(opts *bind.TransactOpts, _tiers []*big.Int, _feesTbps []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setBrokerVolumeTiers", _tiers, _feesTbps)
}

// SetBrokerVolumeTiers is a paid mutator transaction binding the contract method 0x88e27e2f.
//
// Solidity: function setBrokerVolumeTiers(uint256[] _tiers, uint16[] _feesTbps) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetBrokerVolumeTiers(_tiers []*big.Int, _feesTbps []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetBrokerVolumeTiers(&_IPerpetualManager.TransactOpts, _tiers, _feesTbps)
}

// SetBrokerVolumeTiers is a paid mutator transaction binding the contract method 0x88e27e2f.
//
// Solidity: function setBrokerVolumeTiers(uint256[] _tiers, uint16[] _feesTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetBrokerVolumeTiers(_tiers []*big.Int, _feesTbps []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetBrokerVolumeTiers(&_IPerpetualManager.TransactOpts, _tiers, _feesTbps)
}

// SetEmergencyState is a paid mutator transaction binding the contract method 0x79670ef2.
//
// Solidity: function setEmergencyState(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetEmergencyState(opts *bind.TransactOpts, _iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setEmergencyState", _iPerpetualId)
}

// SetEmergencyState is a paid mutator transaction binding the contract method 0x79670ef2.
//
// Solidity: function setEmergencyState(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetEmergencyState(_iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetEmergencyState(&_IPerpetualManager.TransactOpts, _iPerpetualId)
}

// SetEmergencyState is a paid mutator transaction binding the contract method 0x79670ef2.
//
// Solidity: function setEmergencyState(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetEmergencyState(_iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetEmergencyState(&_IPerpetualManager.TransactOpts, _iPerpetualId)
}

// SetFeesForDesignation is a paid mutator transaction binding the contract method 0xea8e1d58.
//
// Solidity: function setFeesForDesignation(uint32[] _designations, uint16[] _fees) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetFeesForDesignation(opts *bind.TransactOpts, _designations []uint32, _fees []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setFeesForDesignation", _designations, _fees)
}

// SetFeesForDesignation is a paid mutator transaction binding the contract method 0xea8e1d58.
//
// Solidity: function setFeesForDesignation(uint32[] _designations, uint16[] _fees) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetFeesForDesignation(_designations []uint32, _fees []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetFeesForDesignation(&_IPerpetualManager.TransactOpts, _designations, _fees)
}

// SetFeesForDesignation is a paid mutator transaction binding the contract method 0xea8e1d58.
//
// Solidity: function setFeesForDesignation(uint32[] _designations, uint16[] _fees) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetFeesForDesignation(_designations []uint32, _fees []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetFeesForDesignation(&_IPerpetualManager.TransactOpts, _designations, _fees)
}

// SetInitialFundAllocationWeight is a paid mutator transaction binding the contract method 0x43d62514.
//
// Solidity: function setInitialFundAllocationWeight(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetInitialFundAllocationWeight(opts *bind.TransactOpts, _iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setInitialFundAllocationWeight", _iPerpetualId)
}

// SetInitialFundAllocationWeight is a paid mutator transaction binding the contract method 0x43d62514.
//
// Solidity: function setInitialFundAllocationWeight(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetInitialFundAllocationWeight(_iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetInitialFundAllocationWeight(&_IPerpetualManager.TransactOpts, _iPerpetualId)
}

// SetInitialFundAllocationWeight is a paid mutator transaction binding the contract method 0x43d62514.
//
// Solidity: function setInitialFundAllocationWeight(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetInitialFundAllocationWeight(_iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetInitialFundAllocationWeight(&_IPerpetualManager.TransactOpts, _iPerpetualId)
}

// SetInitialVolumeForFee is a paid mutator transaction binding the contract method 0x7db0cea7.
//
// Solidity: function setInitialVolumeForFee(uint8 _poolId, address _brokerAddr, uint16 _feeTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetInitialVolumeForFee(opts *bind.TransactOpts, _poolId uint8, _brokerAddr common.Address, _feeTbps uint16) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setInitialVolumeForFee", _poolId, _brokerAddr, _feeTbps)
}

// SetInitialVolumeForFee is a paid mutator transaction binding the contract method 0x7db0cea7.
//
// Solidity: function setInitialVolumeForFee(uint8 _poolId, address _brokerAddr, uint16 _feeTbps) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetInitialVolumeForFee(_poolId uint8, _brokerAddr common.Address, _feeTbps uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetInitialVolumeForFee(&_IPerpetualManager.TransactOpts, _poolId, _brokerAddr, _feeTbps)
}

// SetInitialVolumeForFee is a paid mutator transaction binding the contract method 0x7db0cea7.
//
// Solidity: function setInitialVolumeForFee(uint8 _poolId, address _brokerAddr, uint16 _feeTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetInitialVolumeForFee(_poolId uint8, _brokerAddr common.Address, _feeTbps uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetInitialVolumeForFee(&_IPerpetualManager.TransactOpts, _poolId, _brokerAddr, _feeTbps)
}

// SetNormalState is a paid mutator transaction binding the contract method 0x52bb1e76.
//
// Solidity: function setNormalState(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetNormalState(opts *bind.TransactOpts, _iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setNormalState", _iPerpetualId)
}

// SetNormalState is a paid mutator transaction binding the contract method 0x52bb1e76.
//
// Solidity: function setNormalState(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetNormalState(_iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetNormalState(&_IPerpetualManager.TransactOpts, _iPerpetualId)
}

// SetNormalState is a paid mutator transaction binding the contract method 0x52bb1e76.
//
// Solidity: function setNormalState(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetNormalState(_iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetNormalState(&_IPerpetualManager.TransactOpts, _iPerpetualId)
}

// SetOracleFactory is a paid mutator transaction binding the contract method 0x804bfdd0.
//
// Solidity: function setOracleFactory(address _oracleFactory) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetOracleFactory(opts *bind.TransactOpts, _oracleFactory common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setOracleFactory", _oracleFactory)
}

// SetOracleFactory is a paid mutator transaction binding the contract method 0x804bfdd0.
//
// Solidity: function setOracleFactory(address _oracleFactory) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetOracleFactory(_oracleFactory common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetOracleFactory(&_IPerpetualManager.TransactOpts, _oracleFactory)
}

// SetOracleFactory is a paid mutator transaction binding the contract method 0x804bfdd0.
//
// Solidity: function setOracleFactory(address _oracleFactory) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetOracleFactory(_oracleFactory common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetOracleFactory(&_IPerpetualManager.TransactOpts, _oracleFactory)
}

// SetOracleFactoryForPerpetual is a paid mutator transaction binding the contract method 0xfa9f4d5f.
//
// Solidity: function setOracleFactoryForPerpetual(uint24 _iPerpetualId, address _oracleAddr) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetOracleFactoryForPerpetual(opts *bind.TransactOpts, _iPerpetualId *big.Int, _oracleAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setOracleFactoryForPerpetual", _iPerpetualId, _oracleAddr)
}

// SetOracleFactoryForPerpetual is a paid mutator transaction binding the contract method 0xfa9f4d5f.
//
// Solidity: function setOracleFactoryForPerpetual(uint24 _iPerpetualId, address _oracleAddr) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetOracleFactoryForPerpetual(_iPerpetualId *big.Int, _oracleAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetOracleFactoryForPerpetual(&_IPerpetualManager.TransactOpts, _iPerpetualId, _oracleAddr)
}

// SetOracleFactoryForPerpetual is a paid mutator transaction binding the contract method 0xfa9f4d5f.
//
// Solidity: function setOracleFactoryForPerpetual(uint24 _iPerpetualId, address _oracleAddr) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetOracleFactoryForPerpetual(_iPerpetualId *big.Int, _oracleAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetOracleFactoryForPerpetual(&_IPerpetualManager.TransactOpts, _iPerpetualId, _oracleAddr)
}

// SetOrderBookFactory is a paid mutator transaction binding the contract method 0x850e74d6.
//
// Solidity: function setOrderBookFactory(address _orderBookFactory) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetOrderBookFactory(opts *bind.TransactOpts, _orderBookFactory common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setOrderBookFactory", _orderBookFactory)
}

// SetOrderBookFactory is a paid mutator transaction binding the contract method 0x850e74d6.
//
// Solidity: function setOrderBookFactory(address _orderBookFactory) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetOrderBookFactory(_orderBookFactory common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetOrderBookFactory(&_IPerpetualManager.TransactOpts, _orderBookFactory)
}

// SetOrderBookFactory is a paid mutator transaction binding the contract method 0x850e74d6.
//
// Solidity: function setOrderBookFactory(address _orderBookFactory) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetOrderBookFactory(_orderBookFactory common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetOrderBookFactory(&_IPerpetualManager.TransactOpts, _orderBookFactory)
}

// SetPerpetualBaseParams is a paid mutator transaction binding the contract method 0x40730ae3.
//
// Solidity: function setPerpetualBaseParams(uint24 _iPerpetualId, int128[7] _baseParams) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetPerpetualBaseParams(opts *bind.TransactOpts, _iPerpetualId *big.Int, _baseParams [7]*big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setPerpetualBaseParams", _iPerpetualId, _baseParams)
}

// SetPerpetualBaseParams is a paid mutator transaction binding the contract method 0x40730ae3.
//
// Solidity: function setPerpetualBaseParams(uint24 _iPerpetualId, int128[7] _baseParams) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetPerpetualBaseParams(_iPerpetualId *big.Int, _baseParams [7]*big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPerpetualBaseParams(&_IPerpetualManager.TransactOpts, _iPerpetualId, _baseParams)
}

// SetPerpetualBaseParams is a paid mutator transaction binding the contract method 0x40730ae3.
//
// Solidity: function setPerpetualBaseParams(uint24 _iPerpetualId, int128[7] _baseParams) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetPerpetualBaseParams(_iPerpetualId *big.Int, _baseParams [7]*big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPerpetualBaseParams(&_IPerpetualManager.TransactOpts, _iPerpetualId, _baseParams)
}

// SetPerpetualOracles is a paid mutator transaction binding the contract method 0x14c5ce6c.
//
// Solidity: function setPerpetualOracles(uint24 _iPerpetualId, bytes4[2] _baseQuoteS2, bytes4[2] _baseQuoteS3) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetPerpetualOracles(opts *bind.TransactOpts, _iPerpetualId *big.Int, _baseQuoteS2 [2][4]byte, _baseQuoteS3 [2][4]byte) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setPerpetualOracles", _iPerpetualId, _baseQuoteS2, _baseQuoteS3)
}

// SetPerpetualOracles is a paid mutator transaction binding the contract method 0x14c5ce6c.
//
// Solidity: function setPerpetualOracles(uint24 _iPerpetualId, bytes4[2] _baseQuoteS2, bytes4[2] _baseQuoteS3) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetPerpetualOracles(_iPerpetualId *big.Int, _baseQuoteS2 [2][4]byte, _baseQuoteS3 [2][4]byte) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPerpetualOracles(&_IPerpetualManager.TransactOpts, _iPerpetualId, _baseQuoteS2, _baseQuoteS3)
}

// SetPerpetualOracles is a paid mutator transaction binding the contract method 0x14c5ce6c.
//
// Solidity: function setPerpetualOracles(uint24 _iPerpetualId, bytes4[2] _baseQuoteS2, bytes4[2] _baseQuoteS3) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetPerpetualOracles(_iPerpetualId *big.Int, _baseQuoteS2 [2][4]byte, _baseQuoteS3 [2][4]byte) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPerpetualOracles(&_IPerpetualManager.TransactOpts, _iPerpetualId, _baseQuoteS2, _baseQuoteS3)
}

// SetPerpetualParam is a paid mutator transaction binding the contract method 0xee55ba31.
//
// Solidity: function setPerpetualParam(uint24 _iPerpetualId, string _varName, int128 _value) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetPerpetualParam(opts *bind.TransactOpts, _iPerpetualId *big.Int, _varName string, _value *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setPerpetualParam", _iPerpetualId, _varName, _value)
}

// SetPerpetualParam is a paid mutator transaction binding the contract method 0xee55ba31.
//
// Solidity: function setPerpetualParam(uint24 _iPerpetualId, string _varName, int128 _value) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetPerpetualParam(_iPerpetualId *big.Int, _varName string, _value *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPerpetualParam(&_IPerpetualManager.TransactOpts, _iPerpetualId, _varName, _value)
}

// SetPerpetualParam is a paid mutator transaction binding the contract method 0xee55ba31.
//
// Solidity: function setPerpetualParam(uint24 _iPerpetualId, string _varName, int128 _value) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetPerpetualParam(_iPerpetualId *big.Int, _varName string, _value *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPerpetualParam(&_IPerpetualManager.TransactOpts, _iPerpetualId, _varName, _value)
}

// SetPerpetualParamPair is a paid mutator transaction binding the contract method 0xa2dfde10.
//
// Solidity: function setPerpetualParamPair(uint24 _iPerpetualId, string _name, int128 _value1, int128 _value2) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetPerpetualParamPair(opts *bind.TransactOpts, _iPerpetualId *big.Int, _name string, _value1 *big.Int, _value2 *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setPerpetualParamPair", _iPerpetualId, _name, _value1, _value2)
}

// SetPerpetualParamPair is a paid mutator transaction binding the contract method 0xa2dfde10.
//
// Solidity: function setPerpetualParamPair(uint24 _iPerpetualId, string _name, int128 _value1, int128 _value2) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetPerpetualParamPair(_iPerpetualId *big.Int, _name string, _value1 *big.Int, _value2 *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPerpetualParamPair(&_IPerpetualManager.TransactOpts, _iPerpetualId, _name, _value1, _value2)
}

// SetPerpetualParamPair is a paid mutator transaction binding the contract method 0xa2dfde10.
//
// Solidity: function setPerpetualParamPair(uint24 _iPerpetualId, string _name, int128 _value1, int128 _value2) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetPerpetualParamPair(_iPerpetualId *big.Int, _name string, _value1 *big.Int, _value2 *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPerpetualParamPair(&_IPerpetualManager.TransactOpts, _iPerpetualId, _name, _value1, _value2)
}

// SetPerpetualPoolFactory is a paid mutator transaction binding the contract method 0xd56811ac.
//
// Solidity: function setPerpetualPoolFactory(address _shareTokenFactory) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetPerpetualPoolFactory(opts *bind.TransactOpts, _shareTokenFactory common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setPerpetualPoolFactory", _shareTokenFactory)
}

// SetPerpetualPoolFactory is a paid mutator transaction binding the contract method 0xd56811ac.
//
// Solidity: function setPerpetualPoolFactory(address _shareTokenFactory) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetPerpetualPoolFactory(_shareTokenFactory common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPerpetualPoolFactory(&_IPerpetualManager.TransactOpts, _shareTokenFactory)
}

// SetPerpetualPoolFactory is a paid mutator transaction binding the contract method 0xd56811ac.
//
// Solidity: function setPerpetualPoolFactory(address _shareTokenFactory) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetPerpetualPoolFactory(_shareTokenFactory common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPerpetualPoolFactory(&_IPerpetualManager.TransactOpts, _shareTokenFactory)
}

// SetPerpetualRiskParams is a paid mutator transaction binding the contract method 0xec22dc18.
//
// Solidity: function setPerpetualRiskParams(uint24 _iPerpetualId, int128[5] _underlyingRiskParams, int128[13] _defaultFundRiskParams) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetPerpetualRiskParams(opts *bind.TransactOpts, _iPerpetualId *big.Int, _underlyingRiskParams [5]*big.Int, _defaultFundRiskParams [13]*big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setPerpetualRiskParams", _iPerpetualId, _underlyingRiskParams, _defaultFundRiskParams)
}

// SetPerpetualRiskParams is a paid mutator transaction binding the contract method 0xec22dc18.
//
// Solidity: function setPerpetualRiskParams(uint24 _iPerpetualId, int128[5] _underlyingRiskParams, int128[13] _defaultFundRiskParams) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetPerpetualRiskParams(_iPerpetualId *big.Int, _underlyingRiskParams [5]*big.Int, _defaultFundRiskParams [13]*big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPerpetualRiskParams(&_IPerpetualManager.TransactOpts, _iPerpetualId, _underlyingRiskParams, _defaultFundRiskParams)
}

// SetPerpetualRiskParams is a paid mutator transaction binding the contract method 0xec22dc18.
//
// Solidity: function setPerpetualRiskParams(uint24 _iPerpetualId, int128[5] _underlyingRiskParams, int128[13] _defaultFundRiskParams) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetPerpetualRiskParams(_iPerpetualId *big.Int, _underlyingRiskParams [5]*big.Int, _defaultFundRiskParams [13]*big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPerpetualRiskParams(&_IPerpetualManager.TransactOpts, _iPerpetualId, _underlyingRiskParams, _defaultFundRiskParams)
}

// SetPoolParam is a paid mutator transaction binding the contract method 0x2017641c.
//
// Solidity: function setPoolParam(uint8 _poolId, string _name, int128 _value) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetPoolParam(opts *bind.TransactOpts, _poolId uint8, _name string, _value *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setPoolParam", _poolId, _name, _value)
}

// SetPoolParam is a paid mutator transaction binding the contract method 0x2017641c.
//
// Solidity: function setPoolParam(uint8 _poolId, string _name, int128 _value) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetPoolParam(_poolId uint8, _name string, _value *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPoolParam(&_IPerpetualManager.TransactOpts, _poolId, _name, _value)
}

// SetPoolParam is a paid mutator transaction binding the contract method 0x2017641c.
//
// Solidity: function setPoolParam(uint8 _poolId, string _name, int128 _value) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetPoolParam(_poolId uint8, _name string, _value *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetPoolParam(&_IPerpetualManager.TransactOpts, _poolId, _name, _value)
}

// SetTraderTiers is a paid mutator transaction binding the contract method 0xd3a29fea.
//
// Solidity: function setTraderTiers(uint256[] _tiers, uint16[] _feesTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetTraderTiers(opts *bind.TransactOpts, _tiers []*big.Int, _feesTbps []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setTraderTiers", _tiers, _feesTbps)
}

// SetTraderTiers is a paid mutator transaction binding the contract method 0xd3a29fea.
//
// Solidity: function setTraderTiers(uint256[] _tiers, uint16[] _feesTbps) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetTraderTiers(_tiers []*big.Int, _feesTbps []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetTraderTiers(&_IPerpetualManager.TransactOpts, _tiers, _feesTbps)
}

// SetTraderTiers is a paid mutator transaction binding the contract method 0xd3a29fea.
//
// Solidity: function setTraderTiers(uint256[] _tiers, uint16[] _feesTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetTraderTiers(_tiers []*big.Int, _feesTbps []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetTraderTiers(&_IPerpetualManager.TransactOpts, _tiers, _feesTbps)
}

// SetTraderVolumeTiers is a paid mutator transaction binding the contract method 0x077564d3.
//
// Solidity: function setTraderVolumeTiers(uint256[] _tiers, uint16[] _feesTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetTraderVolumeTiers(opts *bind.TransactOpts, _tiers []*big.Int, _feesTbps []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setTraderVolumeTiers", _tiers, _feesTbps)
}

// SetTraderVolumeTiers is a paid mutator transaction binding the contract method 0x077564d3.
//
// Solidity: function setTraderVolumeTiers(uint256[] _tiers, uint16[] _feesTbps) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetTraderVolumeTiers(_tiers []*big.Int, _feesTbps []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetTraderVolumeTiers(&_IPerpetualManager.TransactOpts, _tiers, _feesTbps)
}

// SetTraderVolumeTiers is a paid mutator transaction binding the contract method 0x077564d3.
//
// Solidity: function setTraderVolumeTiers(uint256[] _tiers, uint16[] _feesTbps) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetTraderVolumeTiers(_tiers []*big.Int, _feesTbps []uint16) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetTraderVolumeTiers(&_IPerpetualManager.TransactOpts, _tiers, _feesTbps)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetTreasury(&_IPerpetualManager.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetTreasury(&_IPerpetualManager.TransactOpts, _treasury)
}

// SetUtilityTokenAddr is a paid mutator transaction binding the contract method 0x011f690e.
//
// Solidity: function setUtilityTokenAddr(address tokenAddr) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) SetUtilityTokenAddr(opts *bind.TransactOpts, tokenAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "setUtilityTokenAddr", tokenAddr)
}

// SetUtilityTokenAddr is a paid mutator transaction binding the contract method 0x011f690e.
//
// Solidity: function setUtilityTokenAddr(address tokenAddr) returns()
func (_IPerpetualManager *IPerpetualManagerSession) SetUtilityTokenAddr(tokenAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetUtilityTokenAddr(&_IPerpetualManager.TransactOpts, tokenAddr)
}

// SetUtilityTokenAddr is a paid mutator transaction binding the contract method 0x011f690e.
//
// Solidity: function setUtilityTokenAddr(address tokenAddr) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SetUtilityTokenAddr(tokenAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SetUtilityTokenAddr(&_IPerpetualManager.TransactOpts, tokenAddr)
}

// Settle is a paid mutator transaction binding the contract method 0x9d7c4ea2.
//
// Solidity: function settle(uint24 _perpetualID, address _traderAddr) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) Settle(opts *bind.TransactOpts, _perpetualID *big.Int, _traderAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "settle", _perpetualID, _traderAddr)
}

// Settle is a paid mutator transaction binding the contract method 0x9d7c4ea2.
//
// Solidity: function settle(uint24 _perpetualID, address _traderAddr) returns()
func (_IPerpetualManager *IPerpetualManagerSession) Settle(_perpetualID *big.Int, _traderAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.Settle(&_IPerpetualManager.TransactOpts, _perpetualID, _traderAddr)
}

// Settle is a paid mutator transaction binding the contract method 0x9d7c4ea2.
//
// Solidity: function settle(uint24 _perpetualID, address _traderAddr) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) Settle(_perpetualID *big.Int, _traderAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.Settle(&_IPerpetualManager.TransactOpts, _perpetualID, _traderAddr)
}

// SettleNextTraderInPool is a paid mutator transaction binding the contract method 0x7624eb40.
//
// Solidity: function settleNextTraderInPool(uint8 _id) returns(bool)
func (_IPerpetualManager *IPerpetualManagerTransactor) SettleNextTraderInPool(opts *bind.TransactOpts, _id uint8) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "settleNextTraderInPool", _id)
}

// SettleNextTraderInPool is a paid mutator transaction binding the contract method 0x7624eb40.
//
// Solidity: function settleNextTraderInPool(uint8 _id) returns(bool)
func (_IPerpetualManager *IPerpetualManagerSession) SettleNextTraderInPool(_id uint8) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SettleNextTraderInPool(&_IPerpetualManager.TransactOpts, _id)
}

// SettleNextTraderInPool is a paid mutator transaction binding the contract method 0x7624eb40.
//
// Solidity: function settleNextTraderInPool(uint8 _id) returns(bool)
func (_IPerpetualManager *IPerpetualManagerTransactorSession) SettleNextTraderInPool(_id uint8) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.SettleNextTraderInPool(&_IPerpetualManager.TransactOpts, _id)
}

// TogglePerpEmergencyState is a paid mutator transaction binding the contract method 0xf2ede6ae.
//
// Solidity: function togglePerpEmergencyState(uint24 _perpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) TogglePerpEmergencyState(opts *bind.TransactOpts, _perpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "togglePerpEmergencyState", _perpetualId)
}

// TogglePerpEmergencyState is a paid mutator transaction binding the contract method 0xf2ede6ae.
//
// Solidity: function togglePerpEmergencyState(uint24 _perpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerSession) TogglePerpEmergencyState(_perpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.TogglePerpEmergencyState(&_IPerpetualManager.TransactOpts, _perpetualId)
}

// TogglePerpEmergencyState is a paid mutator transaction binding the contract method 0xf2ede6ae.
//
// Solidity: function togglePerpEmergencyState(uint24 _perpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) TogglePerpEmergencyState(_perpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.TogglePerpEmergencyState(&_IPerpetualManager.TransactOpts, _perpetualId)
}

// TradeViaOrderBook is a paid mutator transaction binding the contract method 0x5b4c592c.
//
// Solidity: function tradeViaOrderBook((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order) returns(bool)
func (_IPerpetualManager *IPerpetualManagerTransactor) TradeViaOrderBook(opts *bind.TransactOpts, _order IPerpetualOrderOrder) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "tradeViaOrderBook", _order)
}

// TradeViaOrderBook is a paid mutator transaction binding the contract method 0x5b4c592c.
//
// Solidity: function tradeViaOrderBook((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order) returns(bool)
func (_IPerpetualManager *IPerpetualManagerSession) TradeViaOrderBook(_order IPerpetualOrderOrder) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.TradeViaOrderBook(&_IPerpetualManager.TransactOpts, _order)
}

// TradeViaOrderBook is a paid mutator transaction binding the contract method 0x5b4c592c.
//
// Solidity: function tradeViaOrderBook((uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) _order) returns(bool)
func (_IPerpetualManager *IPerpetualManagerTransactorSession) TradeViaOrderBook(_order IPerpetualOrderOrder) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.TradeViaOrderBook(&_IPerpetualManager.TransactOpts, _order)
}

// TransferBrokerLots is a paid mutator transaction binding the contract method 0x647307ab.
//
// Solidity: function transferBrokerLots(uint8 _poolId, address _transferToAddr, uint32 _lots) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) TransferBrokerLots(opts *bind.TransactOpts, _poolId uint8, _transferToAddr common.Address, _lots uint32) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "transferBrokerLots", _poolId, _transferToAddr, _lots)
}

// TransferBrokerLots is a paid mutator transaction binding the contract method 0x647307ab.
//
// Solidity: function transferBrokerLots(uint8 _poolId, address _transferToAddr, uint32 _lots) returns()
func (_IPerpetualManager *IPerpetualManagerSession) TransferBrokerLots(_poolId uint8, _transferToAddr common.Address, _lots uint32) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.TransferBrokerLots(&_IPerpetualManager.TransactOpts, _poolId, _transferToAddr, _lots)
}

// TransferBrokerLots is a paid mutator transaction binding the contract method 0x647307ab.
//
// Solidity: function transferBrokerLots(uint8 _poolId, address _transferToAddr, uint32 _lots) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) TransferBrokerLots(_poolId uint8, _transferToAddr common.Address, _lots uint32) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.TransferBrokerLots(&_IPerpetualManager.TransactOpts, _poolId, _transferToAddr, _lots)
}

// TransferBrokerOwnership is a paid mutator transaction binding the contract method 0xf2494746.
//
// Solidity: function transferBrokerOwnership(uint8 _poolId, address _transferToAddr) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) TransferBrokerOwnership(opts *bind.TransactOpts, _poolId uint8, _transferToAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "transferBrokerOwnership", _poolId, _transferToAddr)
}

// TransferBrokerOwnership is a paid mutator transaction binding the contract method 0xf2494746.
//
// Solidity: function transferBrokerOwnership(uint8 _poolId, address _transferToAddr) returns()
func (_IPerpetualManager *IPerpetualManagerSession) TransferBrokerOwnership(_poolId uint8, _transferToAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.TransferBrokerOwnership(&_IPerpetualManager.TransactOpts, _poolId, _transferToAddr)
}

// TransferBrokerOwnership is a paid mutator transaction binding the contract method 0xf2494746.
//
// Solidity: function transferBrokerOwnership(uint8 _poolId, address _transferToAddr) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) TransferBrokerOwnership(_poolId uint8, _transferToAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.TransferBrokerOwnership(&_IPerpetualManager.TransactOpts, _poolId, _transferToAddr)
}

// TransferEarningsToTreasury is a paid mutator transaction binding the contract method 0xc0bb3834.
//
// Solidity: function transferEarningsToTreasury(uint8 _poolId, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) TransferEarningsToTreasury(opts *bind.TransactOpts, _poolId uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "transferEarningsToTreasury", _poolId, _fAmount)
}

// TransferEarningsToTreasury is a paid mutator transaction binding the contract method 0xc0bb3834.
//
// Solidity: function transferEarningsToTreasury(uint8 _poolId, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerSession) TransferEarningsToTreasury(_poolId uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.TransferEarningsToTreasury(&_IPerpetualManager.TransactOpts, _poolId, _fAmount)
}

// TransferEarningsToTreasury is a paid mutator transaction binding the contract method 0xc0bb3834.
//
// Solidity: function transferEarningsToTreasury(uint8 _poolId, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) TransferEarningsToTreasury(_poolId uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.TransferEarningsToTreasury(&_IPerpetualManager.TransactOpts, _poolId, _fAmount)
}

// TransferValueToTreasury is a paid mutator transaction binding the contract method 0x8d89cf61.
//
// Solidity: function transferValueToTreasury() returns(bool success)
func (_IPerpetualManager *IPerpetualManagerTransactor) TransferValueToTreasury(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "transferValueToTreasury")
}

// TransferValueToTreasury is a paid mutator transaction binding the contract method 0x8d89cf61.
//
// Solidity: function transferValueToTreasury() returns(bool success)
func (_IPerpetualManager *IPerpetualManagerSession) TransferValueToTreasury() (*types.Transaction, error) {
	return _IPerpetualManager.Contract.TransferValueToTreasury(&_IPerpetualManager.TransactOpts)
}

// TransferValueToTreasury is a paid mutator transaction binding the contract method 0x8d89cf61.
//
// Solidity: function transferValueToTreasury() returns(bool success)
func (_IPerpetualManager *IPerpetualManagerTransactorSession) TransferValueToTreasury() (*types.Transaction, error) {
	return _IPerpetualManager.Contract.TransferValueToTreasury(&_IPerpetualManager.TransactOpts)
}

// UpdateAMMTargetFundSize is a paid mutator transaction binding the contract method 0x4b9fe63e.
//
// Solidity: function updateAMMTargetFundSize(uint24 _iPerpetualId, int128 fTargetFundSize) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) UpdateAMMTargetFundSize(opts *bind.TransactOpts, _iPerpetualId *big.Int, fTargetFundSize *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "updateAMMTargetFundSize", _iPerpetualId, fTargetFundSize)
}

// UpdateAMMTargetFundSize is a paid mutator transaction binding the contract method 0x4b9fe63e.
//
// Solidity: function updateAMMTargetFundSize(uint24 _iPerpetualId, int128 fTargetFundSize) returns()
func (_IPerpetualManager *IPerpetualManagerSession) UpdateAMMTargetFundSize(_iPerpetualId *big.Int, fTargetFundSize *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdateAMMTargetFundSize(&_IPerpetualManager.TransactOpts, _iPerpetualId, fTargetFundSize)
}

// UpdateAMMTargetFundSize is a paid mutator transaction binding the contract method 0x4b9fe63e.
//
// Solidity: function updateAMMTargetFundSize(uint24 _iPerpetualId, int128 fTargetFundSize) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) UpdateAMMTargetFundSize(_iPerpetualId *big.Int, fTargetFundSize *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdateAMMTargetFundSize(&_IPerpetualManager.TransactOpts, _iPerpetualId, fTargetFundSize)
}

// UpdateDefaultFundTargetSize is a paid mutator transaction binding the contract method 0x420e0b0e.
//
// Solidity: function updateDefaultFundTargetSize(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) UpdateDefaultFundTargetSize(opts *bind.TransactOpts, _iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "updateDefaultFundTargetSize", _iPerpetualId)
}

// UpdateDefaultFundTargetSize is a paid mutator transaction binding the contract method 0x420e0b0e.
//
// Solidity: function updateDefaultFundTargetSize(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerSession) UpdateDefaultFundTargetSize(_iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdateDefaultFundTargetSize(&_IPerpetualManager.TransactOpts, _iPerpetualId)
}

// UpdateDefaultFundTargetSize is a paid mutator transaction binding the contract method 0x420e0b0e.
//
// Solidity: function updateDefaultFundTargetSize(uint24 _iPerpetualId) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) UpdateDefaultFundTargetSize(_iPerpetualId *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdateDefaultFundTargetSize(&_IPerpetualManager.TransactOpts, _iPerpetualId)
}

// UpdateDefaultFundTargetSizeRandom is a paid mutator transaction binding the contract method 0xfb276d10.
//
// Solidity: function updateDefaultFundTargetSizeRandom(uint8 _iPoolIndex) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) UpdateDefaultFundTargetSizeRandom(opts *bind.TransactOpts, _iPoolIndex uint8) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "updateDefaultFundTargetSizeRandom", _iPoolIndex)
}

// UpdateDefaultFundTargetSizeRandom is a paid mutator transaction binding the contract method 0xfb276d10.
//
// Solidity: function updateDefaultFundTargetSizeRandom(uint8 _iPoolIndex) returns()
func (_IPerpetualManager *IPerpetualManagerSession) UpdateDefaultFundTargetSizeRandom(_iPoolIndex uint8) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdateDefaultFundTargetSizeRandom(&_IPerpetualManager.TransactOpts, _iPoolIndex)
}

// UpdateDefaultFundTargetSizeRandom is a paid mutator transaction binding the contract method 0xfb276d10.
//
// Solidity: function updateDefaultFundTargetSizeRandom(uint8 _iPoolIndex) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) UpdateDefaultFundTargetSizeRandom(_iPoolIndex uint8) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdateDefaultFundTargetSizeRandom(&_IPerpetualManager.TransactOpts, _iPoolIndex)
}

// UpdateFundingAndPricesAfter is a paid mutator transaction binding the contract method 0xb4cde753.
//
// Solidity: function updateFundingAndPricesAfter(uint24 _iPerpetualId0, uint24 _iPerpetualId1) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) UpdateFundingAndPricesAfter(opts *bind.TransactOpts, _iPerpetualId0 *big.Int, _iPerpetualId1 *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "updateFundingAndPricesAfter", _iPerpetualId0, _iPerpetualId1)
}

// UpdateFundingAndPricesAfter is a paid mutator transaction binding the contract method 0xb4cde753.
//
// Solidity: function updateFundingAndPricesAfter(uint24 _iPerpetualId0, uint24 _iPerpetualId1) returns()
func (_IPerpetualManager *IPerpetualManagerSession) UpdateFundingAndPricesAfter(_iPerpetualId0 *big.Int, _iPerpetualId1 *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdateFundingAndPricesAfter(&_IPerpetualManager.TransactOpts, _iPerpetualId0, _iPerpetualId1)
}

// UpdateFundingAndPricesAfter is a paid mutator transaction binding the contract method 0xb4cde753.
//
// Solidity: function updateFundingAndPricesAfter(uint24 _iPerpetualId0, uint24 _iPerpetualId1) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) UpdateFundingAndPricesAfter(_iPerpetualId0 *big.Int, _iPerpetualId1 *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdateFundingAndPricesAfter(&_IPerpetualManager.TransactOpts, _iPerpetualId0, _iPerpetualId1)
}

// UpdateFundingAndPricesBefore is a paid mutator transaction binding the contract method 0x5f8cded9.
//
// Solidity: function updateFundingAndPricesBefore(uint24 _iPerpetualId0, uint24 _iPerpetualId1) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) UpdateFundingAndPricesBefore(opts *bind.TransactOpts, _iPerpetualId0 *big.Int, _iPerpetualId1 *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "updateFundingAndPricesBefore", _iPerpetualId0, _iPerpetualId1)
}

// UpdateFundingAndPricesBefore is a paid mutator transaction binding the contract method 0x5f8cded9.
//
// Solidity: function updateFundingAndPricesBefore(uint24 _iPerpetualId0, uint24 _iPerpetualId1) returns()
func (_IPerpetualManager *IPerpetualManagerSession) UpdateFundingAndPricesBefore(_iPerpetualId0 *big.Int, _iPerpetualId1 *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdateFundingAndPricesBefore(&_IPerpetualManager.TransactOpts, _iPerpetualId0, _iPerpetualId1)
}

// UpdateFundingAndPricesBefore is a paid mutator transaction binding the contract method 0x5f8cded9.
//
// Solidity: function updateFundingAndPricesBefore(uint24 _iPerpetualId0, uint24 _iPerpetualId1) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) UpdateFundingAndPricesBefore(_iPerpetualId0 *big.Int, _iPerpetualId1 *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdateFundingAndPricesBefore(&_IPerpetualManager.TransactOpts, _iPerpetualId0, _iPerpetualId1)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0x3bb510c9.
//
// Solidity: function updatePriceFeeds(uint24 _perpetualId, bytes[] _updateData, uint64[] _publishTimes, uint256 _maxAcceptableFeedAge) payable returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) UpdatePriceFeeds(opts *bind.TransactOpts, _perpetualId *big.Int, _updateData [][]byte, _publishTimes []uint64, _maxAcceptableFeedAge *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "updatePriceFeeds", _perpetualId, _updateData, _publishTimes, _maxAcceptableFeedAge)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0x3bb510c9.
//
// Solidity: function updatePriceFeeds(uint24 _perpetualId, bytes[] _updateData, uint64[] _publishTimes, uint256 _maxAcceptableFeedAge) payable returns()
func (_IPerpetualManager *IPerpetualManagerSession) UpdatePriceFeeds(_perpetualId *big.Int, _updateData [][]byte, _publishTimes []uint64, _maxAcceptableFeedAge *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdatePriceFeeds(&_IPerpetualManager.TransactOpts, _perpetualId, _updateData, _publishTimes, _maxAcceptableFeedAge)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0x3bb510c9.
//
// Solidity: function updatePriceFeeds(uint24 _perpetualId, bytes[] _updateData, uint64[] _publishTimes, uint256 _maxAcceptableFeedAge) payable returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) UpdatePriceFeeds(_perpetualId *big.Int, _updateData [][]byte, _publishTimes []uint64, _maxAcceptableFeedAge *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdatePriceFeeds(&_IPerpetualManager.TransactOpts, _perpetualId, _updateData, _publishTimes, _maxAcceptableFeedAge)
}

// UpdateVolumeEMAOnNewTrade is a paid mutator transaction binding the contract method 0x6f5a4537.
//
// Solidity: function updateVolumeEMAOnNewTrade(uint24 _iPerpetualId, address _traderAddr, address _brokerAddr, int128 _tradeAmountBC) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) UpdateVolumeEMAOnNewTrade(opts *bind.TransactOpts, _iPerpetualId *big.Int, _traderAddr common.Address, _brokerAddr common.Address, _tradeAmountBC *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "updateVolumeEMAOnNewTrade", _iPerpetualId, _traderAddr, _brokerAddr, _tradeAmountBC)
}

// UpdateVolumeEMAOnNewTrade is a paid mutator transaction binding the contract method 0x6f5a4537.
//
// Solidity: function updateVolumeEMAOnNewTrade(uint24 _iPerpetualId, address _traderAddr, address _brokerAddr, int128 _tradeAmountBC) returns()
func (_IPerpetualManager *IPerpetualManagerSession) UpdateVolumeEMAOnNewTrade(_iPerpetualId *big.Int, _traderAddr common.Address, _brokerAddr common.Address, _tradeAmountBC *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdateVolumeEMAOnNewTrade(&_IPerpetualManager.TransactOpts, _iPerpetualId, _traderAddr, _brokerAddr, _tradeAmountBC)
}

// UpdateVolumeEMAOnNewTrade is a paid mutator transaction binding the contract method 0x6f5a4537.
//
// Solidity: function updateVolumeEMAOnNewTrade(uint24 _iPerpetualId, address _traderAddr, address _brokerAddr, int128 _tradeAmountBC) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) UpdateVolumeEMAOnNewTrade(_iPerpetualId *big.Int, _traderAddr common.Address, _brokerAddr common.Address, _tradeAmountBC *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.UpdateVolumeEMAOnNewTrade(&_IPerpetualManager.TransactOpts, _iPerpetualId, _traderAddr, _brokerAddr, _tradeAmountBC)
}

// Withdraw is a paid mutator transaction binding the contract method 0x344818ee.
//
// Solidity: function withdraw(uint24 _iPerpetualId, int128 _fAmount, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) Withdraw(opts *bind.TransactOpts, _iPerpetualId *big.Int, _fAmount *big.Int, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "withdraw", _iPerpetualId, _fAmount, _updateData, _publishTimes)
}

// Withdraw is a paid mutator transaction binding the contract method 0x344818ee.
//
// Solidity: function withdraw(uint24 _iPerpetualId, int128 _fAmount, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_IPerpetualManager *IPerpetualManagerSession) Withdraw(_iPerpetualId *big.Int, _fAmount *big.Int, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.Withdraw(&_IPerpetualManager.TransactOpts, _iPerpetualId, _fAmount, _updateData, _publishTimes)
}

// Withdraw is a paid mutator transaction binding the contract method 0x344818ee.
//
// Solidity: function withdraw(uint24 _iPerpetualId, int128 _fAmount, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) Withdraw(_iPerpetualId *big.Int, _fAmount *big.Int, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.Withdraw(&_IPerpetualManager.TransactOpts, _iPerpetualId, _fAmount, _updateData, _publishTimes)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0xd20edbdd.
//
// Solidity: function withdrawAll(uint24 _iPerpetualId, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) WithdrawAll(opts *bind.TransactOpts, _iPerpetualId *big.Int, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "withdrawAll", _iPerpetualId, _updateData, _publishTimes)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0xd20edbdd.
//
// Solidity: function withdrawAll(uint24 _iPerpetualId, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_IPerpetualManager *IPerpetualManagerSession) WithdrawAll(_iPerpetualId *big.Int, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.WithdrawAll(&_IPerpetualManager.TransactOpts, _iPerpetualId, _updateData, _publishTimes)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0xd20edbdd.
//
// Solidity: function withdrawAll(uint24 _iPerpetualId, bytes[] _updateData, uint64[] _publishTimes) payable returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) WithdrawAll(_iPerpetualId *big.Int, _updateData [][]byte, _publishTimes []uint64) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.WithdrawAll(&_IPerpetualManager.TransactOpts, _iPerpetualId, _updateData, _publishTimes)
}

// WithdrawDepositFromMarginAccount is a paid mutator transaction binding the contract method 0x961881b4.
//
// Solidity: function withdrawDepositFromMarginAccount(uint24 _iPerpetualId, address _traderAddr) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) WithdrawDepositFromMarginAccount(opts *bind.TransactOpts, _iPerpetualId *big.Int, _traderAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "withdrawDepositFromMarginAccount", _iPerpetualId, _traderAddr)
}

// WithdrawDepositFromMarginAccount is a paid mutator transaction binding the contract method 0x961881b4.
//
// Solidity: function withdrawDepositFromMarginAccount(uint24 _iPerpetualId, address _traderAddr) returns()
func (_IPerpetualManager *IPerpetualManagerSession) WithdrawDepositFromMarginAccount(_iPerpetualId *big.Int, _traderAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.WithdrawDepositFromMarginAccount(&_IPerpetualManager.TransactOpts, _iPerpetualId, _traderAddr)
}

// WithdrawDepositFromMarginAccount is a paid mutator transaction binding the contract method 0x961881b4.
//
// Solidity: function withdrawDepositFromMarginAccount(uint24 _iPerpetualId, address _traderAddr) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) WithdrawDepositFromMarginAccount(_iPerpetualId *big.Int, _traderAddr common.Address) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.WithdrawDepositFromMarginAccount(&_IPerpetualManager.TransactOpts, _iPerpetualId, _traderAddr)
}

// WithdrawFromDefaultFund is a paid mutator transaction binding the contract method 0x1624866d.
//
// Solidity: function withdrawFromDefaultFund(uint8 _poolId, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) WithdrawFromDefaultFund(opts *bind.TransactOpts, _poolId uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "withdrawFromDefaultFund", _poolId, _fAmount)
}

// WithdrawFromDefaultFund is a paid mutator transaction binding the contract method 0x1624866d.
//
// Solidity: function withdrawFromDefaultFund(uint8 _poolId, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerSession) WithdrawFromDefaultFund(_poolId uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.WithdrawFromDefaultFund(&_IPerpetualManager.TransactOpts, _poolId, _fAmount)
}

// WithdrawFromDefaultFund is a paid mutator transaction binding the contract method 0x1624866d.
//
// Solidity: function withdrawFromDefaultFund(uint8 _poolId, int128 _fAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) WithdrawFromDefaultFund(_poolId uint8, _fAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.WithdrawFromDefaultFund(&_IPerpetualManager.TransactOpts, _poolId, _fAmount)
}

// WithdrawLiquidity is a paid mutator transaction binding the contract method 0x78bc212d.
//
// Solidity: function withdrawLiquidity(uint8 _iPoolIndex, uint256 _shareAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactor) WithdrawLiquidity(opts *bind.TransactOpts, _iPoolIndex uint8, _shareAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.contract.Transact(opts, "withdrawLiquidity", _iPoolIndex, _shareAmount)
}

// WithdrawLiquidity is a paid mutator transaction binding the contract method 0x78bc212d.
//
// Solidity: function withdrawLiquidity(uint8 _iPoolIndex, uint256 _shareAmount) returns()
func (_IPerpetualManager *IPerpetualManagerSession) WithdrawLiquidity(_iPoolIndex uint8, _shareAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.WithdrawLiquidity(&_IPerpetualManager.TransactOpts, _iPoolIndex, _shareAmount)
}

// WithdrawLiquidity is a paid mutator transaction binding the contract method 0x78bc212d.
//
// Solidity: function withdrawLiquidity(uint8 _iPoolIndex, uint256 _shareAmount) returns()
func (_IPerpetualManager *IPerpetualManagerTransactorSession) WithdrawLiquidity(_iPoolIndex uint8, _shareAmount *big.Int) (*types.Transaction, error) {
	return _IPerpetualManager.Contract.WithdrawLiquidity(&_IPerpetualManager.TransactOpts, _iPoolIndex, _shareAmount)
}

// IPerpetualManagerBrokerLotsTransferredIterator is returned from FilterBrokerLotsTransferred and is used to iterate over the raw logs and unpacked data for BrokerLotsTransferred events raised by the IPerpetualManager contract.
type IPerpetualManagerBrokerLotsTransferredIterator struct {
	Event *IPerpetualManagerBrokerLotsTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerBrokerLotsTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerBrokerLotsTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerBrokerLotsTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerBrokerLotsTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerBrokerLotsTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerBrokerLotsTransferred represents a BrokerLotsTransferred event raised by the IPerpetualManager contract.
type IPerpetualManagerBrokerLotsTransferred struct {
	PoolId   uint8
	OldOwner common.Address
	NewOwner common.Address
	NumLots  uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBrokerLotsTransferred is a free log retrieval operation binding the contract event 0xce7ce31ba06dfdd70dd9c75f94082acfc55ad4dc2d185dcc2902b4cb4fee9b5c.
//
// Solidity: event BrokerLotsTransferred(uint8 indexed poolId, address oldOwner, address newOwner, uint32 numLots)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterBrokerLotsTransferred(opts *bind.FilterOpts, poolId []uint8) (*IPerpetualManagerBrokerLotsTransferredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "BrokerLotsTransferred", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerBrokerLotsTransferredIterator{contract: _IPerpetualManager.contract, event: "BrokerLotsTransferred", logs: logs, sub: sub}, nil
}

// WatchBrokerLotsTransferred is a free log subscription operation binding the contract event 0xce7ce31ba06dfdd70dd9c75f94082acfc55ad4dc2d185dcc2902b4cb4fee9b5c.
//
// Solidity: event BrokerLotsTransferred(uint8 indexed poolId, address oldOwner, address newOwner, uint32 numLots)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchBrokerLotsTransferred(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerBrokerLotsTransferred, poolId []uint8) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "BrokerLotsTransferred", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerBrokerLotsTransferred)
				if err := _IPerpetualManager.contract.UnpackLog(event, "BrokerLotsTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBrokerLotsTransferred is a log parse operation binding the contract event 0xce7ce31ba06dfdd70dd9c75f94082acfc55ad4dc2d185dcc2902b4cb4fee9b5c.
//
// Solidity: event BrokerLotsTransferred(uint8 indexed poolId, address oldOwner, address newOwner, uint32 numLots)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseBrokerLotsTransferred(log types.Log) (*IPerpetualManagerBrokerLotsTransferred, error) {
	event := new(IPerpetualManagerBrokerLotsTransferred)
	if err := _IPerpetualManager.contract.UnpackLog(event, "BrokerLotsTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerBrokerVolumeTransferredIterator is returned from FilterBrokerVolumeTransferred and is used to iterate over the raw logs and unpacked data for BrokerVolumeTransferred events raised by the IPerpetualManager contract.
type IPerpetualManagerBrokerVolumeTransferredIterator struct {
	Event *IPerpetualManagerBrokerVolumeTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerBrokerVolumeTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerBrokerVolumeTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerBrokerVolumeTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerBrokerVolumeTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerBrokerVolumeTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerBrokerVolumeTransferred represents a BrokerVolumeTransferred event raised by the IPerpetualManager contract.
type IPerpetualManagerBrokerVolumeTransferred struct {
	PoolId   uint8
	OldOwner common.Address
	NewOwner common.Address
	FVolume  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBrokerVolumeTransferred is a free log retrieval operation binding the contract event 0xbd513eef6c389e97ad68326e2ee5b8d23788c432aa41c01c291c800cc9715521.
//
// Solidity: event BrokerVolumeTransferred(uint8 indexed poolId, address oldOwner, address newOwner, int128 fVolume)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterBrokerVolumeTransferred(opts *bind.FilterOpts, poolId []uint8) (*IPerpetualManagerBrokerVolumeTransferredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "BrokerVolumeTransferred", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerBrokerVolumeTransferredIterator{contract: _IPerpetualManager.contract, event: "BrokerVolumeTransferred", logs: logs, sub: sub}, nil
}

// WatchBrokerVolumeTransferred is a free log subscription operation binding the contract event 0xbd513eef6c389e97ad68326e2ee5b8d23788c432aa41c01c291c800cc9715521.
//
// Solidity: event BrokerVolumeTransferred(uint8 indexed poolId, address oldOwner, address newOwner, int128 fVolume)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchBrokerVolumeTransferred(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerBrokerVolumeTransferred, poolId []uint8) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "BrokerVolumeTransferred", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerBrokerVolumeTransferred)
				if err := _IPerpetualManager.contract.UnpackLog(event, "BrokerVolumeTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBrokerVolumeTransferred is a log parse operation binding the contract event 0xbd513eef6c389e97ad68326e2ee5b8d23788c432aa41c01c291c800cc9715521.
//
// Solidity: event BrokerVolumeTransferred(uint8 indexed poolId, address oldOwner, address newOwner, int128 fVolume)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseBrokerVolumeTransferred(log types.Log) (*IPerpetualManagerBrokerVolumeTransferred, error) {
	event := new(IPerpetualManagerBrokerVolumeTransferred)
	if err := _IPerpetualManager.contract.UnpackLog(event, "BrokerVolumeTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerClearIterator is returned from FilterClear and is used to iterate over the raw logs and unpacked data for Clear events raised by the IPerpetualManager contract.
type IPerpetualManagerClearIterator struct {
	Event *IPerpetualManagerClear // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerClearIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerClear)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerClear)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerClearIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerClearIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerClear represents a Clear event raised by the IPerpetualManager contract.
type IPerpetualManagerClear struct {
	PerpetualId *big.Int
	Trader      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClear is a free log retrieval operation binding the contract event 0x5654917936696ceefa513c46724a958dd400e01e0db9bdf37e06a26ce0772a5c.
//
// Solidity: event Clear(uint24 indexed perpetualId, address indexed trader)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterClear(opts *bind.FilterOpts, perpetualId []*big.Int, trader []common.Address) (*IPerpetualManagerClearIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "Clear", perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerClearIterator{contract: _IPerpetualManager.contract, event: "Clear", logs: logs, sub: sub}, nil
}

// WatchClear is a free log subscription operation binding the contract event 0x5654917936696ceefa513c46724a958dd400e01e0db9bdf37e06a26ce0772a5c.
//
// Solidity: event Clear(uint24 indexed perpetualId, address indexed trader)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchClear(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerClear, perpetualId []*big.Int, trader []common.Address) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "Clear", perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerClear)
				if err := _IPerpetualManager.contract.UnpackLog(event, "Clear", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClear is a log parse operation binding the contract event 0x5654917936696ceefa513c46724a958dd400e01e0db9bdf37e06a26ce0772a5c.
//
// Solidity: event Clear(uint24 indexed perpetualId, address indexed trader)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseClear(log types.Log) (*IPerpetualManagerClear, error) {
	event := new(IPerpetualManagerClear)
	if err := _IPerpetualManager.contract.UnpackLog(event, "Clear", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerDistributeFeesIterator is returned from FilterDistributeFees and is used to iterate over the raw logs and unpacked data for DistributeFees events raised by the IPerpetualManager contract.
type IPerpetualManagerDistributeFeesIterator struct {
	Event *IPerpetualManagerDistributeFees // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerDistributeFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerDistributeFees)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerDistributeFees)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerDistributeFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerDistributeFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerDistributeFees represents a DistributeFees event raised by the IPerpetualManager contract.
type IPerpetualManagerDistributeFees struct {
	PoolId                 uint8
	PerpetualId            *big.Int
	Trader                 common.Address
	ProtocolFeeCC          *big.Int
	ParticipationFundFeeCC *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterDistributeFees is a free log retrieval operation binding the contract event 0x0361b6dfff94227bf8c480cdbc3b0ad383860b1f7c4cd535c8c3202f43c6f4d9.
//
// Solidity: event DistributeFees(uint8 indexed poolId, uint24 indexed perpetualId, address indexed trader, int128 protocolFeeCC, int128 participationFundFeeCC)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterDistributeFees(opts *bind.FilterOpts, poolId []uint8, perpetualId []*big.Int, trader []common.Address) (*IPerpetualManagerDistributeFeesIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "DistributeFees", poolIdRule, perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerDistributeFeesIterator{contract: _IPerpetualManager.contract, event: "DistributeFees", logs: logs, sub: sub}, nil
}

// WatchDistributeFees is a free log subscription operation binding the contract event 0x0361b6dfff94227bf8c480cdbc3b0ad383860b1f7c4cd535c8c3202f43c6f4d9.
//
// Solidity: event DistributeFees(uint8 indexed poolId, uint24 indexed perpetualId, address indexed trader, int128 protocolFeeCC, int128 participationFundFeeCC)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchDistributeFees(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerDistributeFees, poolId []uint8, perpetualId []*big.Int, trader []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "DistributeFees", poolIdRule, perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerDistributeFees)
				if err := _IPerpetualManager.contract.UnpackLog(event, "DistributeFees", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributeFees is a log parse operation binding the contract event 0x0361b6dfff94227bf8c480cdbc3b0ad383860b1f7c4cd535c8c3202f43c6f4d9.
//
// Solidity: event DistributeFees(uint8 indexed poolId, uint24 indexed perpetualId, address indexed trader, int128 protocolFeeCC, int128 participationFundFeeCC)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseDistributeFees(log types.Log) (*IPerpetualManagerDistributeFees, error) {
	event := new(IPerpetualManagerDistributeFees)
	if err := _IPerpetualManager.contract.UnpackLog(event, "DistributeFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the IPerpetualManager contract.
type IPerpetualManagerLiquidateIterator struct {
	Event *IPerpetualManagerLiquidate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerLiquidate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerLiquidate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerLiquidate represents a Liquidate event raised by the IPerpetualManager contract.
type IPerpetualManagerLiquidate struct {
	PerpetualId        *big.Int
	Liquidator         common.Address
	Trader             common.Address
	PositionId         [16]byte
	AmountLiquidatedBC *big.Int
	LiquidationPrice   *big.Int
	NewPositionSizeBC  *big.Int
	FFeeCC             *big.Int
	FPnlCC             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0x3d975b766e6428036dad815db81148b8ea14d0f4255df3fd44304abdf9ff87b0.
//
// Solidity: event Liquidate(uint24 perpetualId, address indexed liquidator, address indexed trader, bytes16 indexed positionId, int128 amountLiquidatedBC, int128 liquidationPrice, int128 newPositionSizeBC, int128 fFeeCC, int128 fPnlCC)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterLiquidate(opts *bind.FilterOpts, liquidator []common.Address, trader []common.Address, positionId [][16]byte) (*IPerpetualManagerLiquidateIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "Liquidate", liquidatorRule, traderRule, positionIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerLiquidateIterator{contract: _IPerpetualManager.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x3d975b766e6428036dad815db81148b8ea14d0f4255df3fd44304abdf9ff87b0.
//
// Solidity: event Liquidate(uint24 perpetualId, address indexed liquidator, address indexed trader, bytes16 indexed positionId, int128 amountLiquidatedBC, int128 liquidationPrice, int128 newPositionSizeBC, int128 fFeeCC, int128 fPnlCC)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerLiquidate, liquidator []common.Address, trader []common.Address, positionId [][16]byte) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "Liquidate", liquidatorRule, traderRule, positionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerLiquidate)
				if err := _IPerpetualManager.contract.UnpackLog(event, "Liquidate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidate is a log parse operation binding the contract event 0x3d975b766e6428036dad815db81148b8ea14d0f4255df3fd44304abdf9ff87b0.
//
// Solidity: event Liquidate(uint24 perpetualId, address indexed liquidator, address indexed trader, bytes16 indexed positionId, int128 amountLiquidatedBC, int128 liquidationPrice, int128 newPositionSizeBC, int128 fFeeCC, int128 fPnlCC)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseLiquidate(log types.Log) (*IPerpetualManagerLiquidate, error) {
	event := new(IPerpetualManagerLiquidate)
	if err := _IPerpetualManager.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the IPerpetualManager contract.
type IPerpetualManagerLiquidityAddedIterator struct {
	Event *IPerpetualManagerLiquidityAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerLiquidityAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerLiquidityAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerLiquidityAdded represents a LiquidityAdded event raised by the IPerpetualManager contract.
type IPerpetualManagerLiquidityAdded struct {
	PoolId      uint64
	User        common.Address
	TokenAmount *big.Int
	ShareAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0x8e989dad15aaebf959f1f1144a5d3c187781c04e2c4f255ca3f304d4fda014ec.
//
// Solidity: event LiquidityAdded(uint64 indexed poolId, address indexed user, uint256 tokenAmount, uint256 shareAmount)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, poolId []uint64, user []common.Address) (*IPerpetualManagerLiquidityAddedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "LiquidityAdded", poolIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerLiquidityAddedIterator{contract: _IPerpetualManager.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0x8e989dad15aaebf959f1f1144a5d3c187781c04e2c4f255ca3f304d4fda014ec.
//
// Solidity: event LiquidityAdded(uint64 indexed poolId, address indexed user, uint256 tokenAmount, uint256 shareAmount)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerLiquidityAdded, poolId []uint64, user []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "LiquidityAdded", poolIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerLiquidityAdded)
				if err := _IPerpetualManager.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidityAdded is a log parse operation binding the contract event 0x8e989dad15aaebf959f1f1144a5d3c187781c04e2c4f255ca3f304d4fda014ec.
//
// Solidity: event LiquidityAdded(uint64 indexed poolId, address indexed user, uint256 tokenAmount, uint256 shareAmount)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseLiquidityAdded(log types.Log) (*IPerpetualManagerLiquidityAdded, error) {
	event := new(IPerpetualManagerLiquidityAdded)
	if err := _IPerpetualManager.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerLiquidityPoolCreatedIterator is returned from FilterLiquidityPoolCreated and is used to iterate over the raw logs and unpacked data for LiquidityPoolCreated events raised by the IPerpetualManager contract.
type IPerpetualManagerLiquidityPoolCreatedIterator struct {
	Event *IPerpetualManagerLiquidityPoolCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerLiquidityPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerLiquidityPoolCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerLiquidityPoolCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerLiquidityPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerLiquidityPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerLiquidityPoolCreated represents a LiquidityPoolCreated event raised by the IPerpetualManager contract.
type IPerpetualManagerLiquidityPoolCreated struct {
	Id                               uint8
	MarginTokenAddress               common.Address
	ShareTokenAddress                common.Address
	ITargetPoolSizeUpdateTime        uint16
	FMaxTransferPerConvergencePeriod *big.Int
	FBrokerCollateralLotSize         *big.Int
	Raw                              types.Log // Blockchain specific contextual infos
}

// FilterLiquidityPoolCreated is a free log retrieval operation binding the contract event 0x3b9d25c810fa75adcd19d4fee54d577d91f9d0158d255f57306a22aac88a463f.
//
// Solidity: event LiquidityPoolCreated(uint8 id, address marginTokenAddress, address shareTokenAddress, uint16 iTargetPoolSizeUpdateTime, int128 fMaxTransferPerConvergencePeriod, int128 fBrokerCollateralLotSize)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterLiquidityPoolCreated(opts *bind.FilterOpts) (*IPerpetualManagerLiquidityPoolCreatedIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "LiquidityPoolCreated")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerLiquidityPoolCreatedIterator{contract: _IPerpetualManager.contract, event: "LiquidityPoolCreated", logs: logs, sub: sub}, nil
}

// WatchLiquidityPoolCreated is a free log subscription operation binding the contract event 0x3b9d25c810fa75adcd19d4fee54d577d91f9d0158d255f57306a22aac88a463f.
//
// Solidity: event LiquidityPoolCreated(uint8 id, address marginTokenAddress, address shareTokenAddress, uint16 iTargetPoolSizeUpdateTime, int128 fMaxTransferPerConvergencePeriod, int128 fBrokerCollateralLotSize)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchLiquidityPoolCreated(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerLiquidityPoolCreated) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "LiquidityPoolCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerLiquidityPoolCreated)
				if err := _IPerpetualManager.contract.UnpackLog(event, "LiquidityPoolCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidityPoolCreated is a log parse operation binding the contract event 0x3b9d25c810fa75adcd19d4fee54d577d91f9d0158d255f57306a22aac88a463f.
//
// Solidity: event LiquidityPoolCreated(uint8 id, address marginTokenAddress, address shareTokenAddress, uint16 iTargetPoolSizeUpdateTime, int128 fMaxTransferPerConvergencePeriod, int128 fBrokerCollateralLotSize)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseLiquidityPoolCreated(log types.Log) (*IPerpetualManagerLiquidityPoolCreated, error) {
	event := new(IPerpetualManagerLiquidityPoolCreated)
	if err := _IPerpetualManager.contract.UnpackLog(event, "LiquidityPoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerLiquidityProvisionPausedIterator is returned from FilterLiquidityProvisionPaused and is used to iterate over the raw logs and unpacked data for LiquidityProvisionPaused events raised by the IPerpetualManager contract.
type IPerpetualManagerLiquidityProvisionPausedIterator struct {
	Event *IPerpetualManagerLiquidityProvisionPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerLiquidityProvisionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerLiquidityProvisionPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerLiquidityProvisionPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerLiquidityProvisionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerLiquidityProvisionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerLiquidityProvisionPaused represents a LiquidityProvisionPaused event raised by the IPerpetualManager contract.
type IPerpetualManagerLiquidityProvisionPaused struct {
	PauseOn bool
	PoolId  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLiquidityProvisionPaused is a free log retrieval operation binding the contract event 0xdd50900dcfecd1d3006ae75c734e1ea17e5f8a020b58d16fa488b8efb29257d3.
//
// Solidity: event LiquidityProvisionPaused(bool pauseOn, uint8 poolId)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterLiquidityProvisionPaused(opts *bind.FilterOpts) (*IPerpetualManagerLiquidityProvisionPausedIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "LiquidityProvisionPaused")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerLiquidityProvisionPausedIterator{contract: _IPerpetualManager.contract, event: "LiquidityProvisionPaused", logs: logs, sub: sub}, nil
}

// WatchLiquidityProvisionPaused is a free log subscription operation binding the contract event 0xdd50900dcfecd1d3006ae75c734e1ea17e5f8a020b58d16fa488b8efb29257d3.
//
// Solidity: event LiquidityProvisionPaused(bool pauseOn, uint8 poolId)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchLiquidityProvisionPaused(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerLiquidityProvisionPaused) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "LiquidityProvisionPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerLiquidityProvisionPaused)
				if err := _IPerpetualManager.contract.UnpackLog(event, "LiquidityProvisionPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidityProvisionPaused is a log parse operation binding the contract event 0xdd50900dcfecd1d3006ae75c734e1ea17e5f8a020b58d16fa488b8efb29257d3.
//
// Solidity: event LiquidityProvisionPaused(bool pauseOn, uint8 poolId)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseLiquidityProvisionPaused(log types.Log) (*IPerpetualManagerLiquidityProvisionPaused, error) {
	event := new(IPerpetualManagerLiquidityProvisionPaused)
	if err := _IPerpetualManager.contract.UnpackLog(event, "LiquidityProvisionPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerLiquidityRemovedIterator is returned from FilterLiquidityRemoved and is used to iterate over the raw logs and unpacked data for LiquidityRemoved events raised by the IPerpetualManager contract.
type IPerpetualManagerLiquidityRemovedIterator struct {
	Event *IPerpetualManagerLiquidityRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerLiquidityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerLiquidityRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerLiquidityRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerLiquidityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerLiquidityRemoved represents a LiquidityRemoved event raised by the IPerpetualManager contract.
type IPerpetualManagerLiquidityRemoved struct {
	PoolId      uint64
	User        common.Address
	TokenAmount *big.Int
	ShareAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLiquidityRemoved is a free log retrieval operation binding the contract event 0xea2f3527bbaa09dc11a4468db411f0b002ea84220b52fba48d9619acdade1b84.
//
// Solidity: event LiquidityRemoved(uint64 indexed poolId, address indexed user, uint256 tokenAmount, uint256 shareAmount)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, poolId []uint64, user []common.Address) (*IPerpetualManagerLiquidityRemovedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "LiquidityRemoved", poolIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerLiquidityRemovedIterator{contract: _IPerpetualManager.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemoved is a free log subscription operation binding the contract event 0xea2f3527bbaa09dc11a4468db411f0b002ea84220b52fba48d9619acdade1b84.
//
// Solidity: event LiquidityRemoved(uint64 indexed poolId, address indexed user, uint256 tokenAmount, uint256 shareAmount)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerLiquidityRemoved, poolId []uint64, user []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "LiquidityRemoved", poolIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerLiquidityRemoved)
				if err := _IPerpetualManager.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidityRemoved is a log parse operation binding the contract event 0xea2f3527bbaa09dc11a4468db411f0b002ea84220b52fba48d9619acdade1b84.
//
// Solidity: event LiquidityRemoved(uint64 indexed poolId, address indexed user, uint256 tokenAmount, uint256 shareAmount)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseLiquidityRemoved(log types.Log) (*IPerpetualManagerLiquidityRemoved, error) {
	event := new(IPerpetualManagerLiquidityRemoved)
	if err := _IPerpetualManager.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerLiquidityWithdrawalInitiatedIterator is returned from FilterLiquidityWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for LiquidityWithdrawalInitiated events raised by the IPerpetualManager contract.
type IPerpetualManagerLiquidityWithdrawalInitiatedIterator struct {
	Event *IPerpetualManagerLiquidityWithdrawalInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerLiquidityWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerLiquidityWithdrawalInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerLiquidityWithdrawalInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerLiquidityWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerLiquidityWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerLiquidityWithdrawalInitiated represents a LiquidityWithdrawalInitiated event raised by the IPerpetualManager contract.
type IPerpetualManagerLiquidityWithdrawalInitiated struct {
	PoolId      uint64
	User        common.Address
	ShareAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLiquidityWithdrawalInitiated is a free log retrieval operation binding the contract event 0x26713ef017cbdcd67db0af7af54e848e4e22e7bdb4d207c3fccf5375518c256f.
//
// Solidity: event LiquidityWithdrawalInitiated(uint64 indexed poolId, address indexed user, uint256 shareAmount)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterLiquidityWithdrawalInitiated(opts *bind.FilterOpts, poolId []uint64, user []common.Address) (*IPerpetualManagerLiquidityWithdrawalInitiatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "LiquidityWithdrawalInitiated", poolIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerLiquidityWithdrawalInitiatedIterator{contract: _IPerpetualManager.contract, event: "LiquidityWithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchLiquidityWithdrawalInitiated is a free log subscription operation binding the contract event 0x26713ef017cbdcd67db0af7af54e848e4e22e7bdb4d207c3fccf5375518c256f.
//
// Solidity: event LiquidityWithdrawalInitiated(uint64 indexed poolId, address indexed user, uint256 shareAmount)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchLiquidityWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerLiquidityWithdrawalInitiated, poolId []uint64, user []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "LiquidityWithdrawalInitiated", poolIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerLiquidityWithdrawalInitiated)
				if err := _IPerpetualManager.contract.UnpackLog(event, "LiquidityWithdrawalInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidityWithdrawalInitiated is a log parse operation binding the contract event 0x26713ef017cbdcd67db0af7af54e848e4e22e7bdb4d207c3fccf5375518c256f.
//
// Solidity: event LiquidityWithdrawalInitiated(uint64 indexed poolId, address indexed user, uint256 shareAmount)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseLiquidityWithdrawalInitiated(log types.Log) (*IPerpetualManagerLiquidityWithdrawalInitiated, error) {
	event := new(IPerpetualManagerLiquidityWithdrawalInitiated)
	if err := _IPerpetualManager.contract.UnpackLog(event, "LiquidityWithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerPerpetualCreatedIterator is returned from FilterPerpetualCreated and is used to iterate over the raw logs and unpacked data for PerpetualCreated events raised by the IPerpetualManager contract.
type IPerpetualManagerPerpetualCreatedIterator struct {
	Event *IPerpetualManagerPerpetualCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerPerpetualCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerPerpetualCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerPerpetualCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerPerpetualCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerPerpetualCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerPerpetualCreated represents a PerpetualCreated event raised by the IPerpetualManager contract.
type IPerpetualManagerPerpetualCreated struct {
	PoolId                uint8
	Id                    *big.Int
	BaseParams            [7]*big.Int
	UnderlyingRiskParams  [5]*big.Int
	DefaultFundRiskParams [13]*big.Int
	ECollateralCurrency   *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterPerpetualCreated is a free log retrieval operation binding the contract event 0xaec4bf6f6bb99f29b30f1d4172d15b864d60d638d4676a1547fc2e52e8410967.
//
// Solidity: event PerpetualCreated(uint8 poolId, uint24 id, int128[7] baseParams, int128[5] underlyingRiskParams, int128[13] defaultFundRiskParams, uint256 eCollateralCurrency)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterPerpetualCreated(opts *bind.FilterOpts) (*IPerpetualManagerPerpetualCreatedIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "PerpetualCreated")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerPerpetualCreatedIterator{contract: _IPerpetualManager.contract, event: "PerpetualCreated", logs: logs, sub: sub}, nil
}

// WatchPerpetualCreated is a free log subscription operation binding the contract event 0xaec4bf6f6bb99f29b30f1d4172d15b864d60d638d4676a1547fc2e52e8410967.
//
// Solidity: event PerpetualCreated(uint8 poolId, uint24 id, int128[7] baseParams, int128[5] underlyingRiskParams, int128[13] defaultFundRiskParams, uint256 eCollateralCurrency)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchPerpetualCreated(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerPerpetualCreated) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "PerpetualCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerPerpetualCreated)
				if err := _IPerpetualManager.contract.UnpackLog(event, "PerpetualCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePerpetualCreated is a log parse operation binding the contract event 0xaec4bf6f6bb99f29b30f1d4172d15b864d60d638d4676a1547fc2e52e8410967.
//
// Solidity: event PerpetualCreated(uint8 poolId, uint24 id, int128[7] baseParams, int128[5] underlyingRiskParams, int128[13] defaultFundRiskParams, uint256 eCollateralCurrency)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParsePerpetualCreated(log types.Log) (*IPerpetualManagerPerpetualCreated, error) {
	event := new(IPerpetualManagerPerpetualCreated)
	if err := _IPerpetualManager.contract.UnpackLog(event, "PerpetualCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerPerpetualLimitOrderCancelledIterator is returned from FilterPerpetualLimitOrderCancelled and is used to iterate over the raw logs and unpacked data for PerpetualLimitOrderCancelled events raised by the IPerpetualManager contract.
type IPerpetualManagerPerpetualLimitOrderCancelledIterator struct {
	Event *IPerpetualManagerPerpetualLimitOrderCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerPerpetualLimitOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerPerpetualLimitOrderCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerPerpetualLimitOrderCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerPerpetualLimitOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerPerpetualLimitOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerPerpetualLimitOrderCancelled represents a PerpetualLimitOrderCancelled event raised by the IPerpetualManager contract.
type IPerpetualManagerPerpetualLimitOrderCancelled struct {
	PerpetualId *big.Int
	OrderHash   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPerpetualLimitOrderCancelled is a free log retrieval operation binding the contract event 0x053ad2eb505ab7dee10747f2b9b1572c82648c0c6e471c09631c4a10c79bcf9e.
//
// Solidity: event PerpetualLimitOrderCancelled(uint24 indexed perpetualId, bytes32 indexed orderHash)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterPerpetualLimitOrderCancelled(opts *bind.FilterOpts, perpetualId []*big.Int, orderHash [][32]byte) (*IPerpetualManagerPerpetualLimitOrderCancelledIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "PerpetualLimitOrderCancelled", perpetualIdRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerPerpetualLimitOrderCancelledIterator{contract: _IPerpetualManager.contract, event: "PerpetualLimitOrderCancelled", logs: logs, sub: sub}, nil
}

// WatchPerpetualLimitOrderCancelled is a free log subscription operation binding the contract event 0x053ad2eb505ab7dee10747f2b9b1572c82648c0c6e471c09631c4a10c79bcf9e.
//
// Solidity: event PerpetualLimitOrderCancelled(uint24 indexed perpetualId, bytes32 indexed orderHash)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchPerpetualLimitOrderCancelled(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerPerpetualLimitOrderCancelled, perpetualId []*big.Int, orderHash [][32]byte) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "PerpetualLimitOrderCancelled", perpetualIdRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerPerpetualLimitOrderCancelled)
				if err := _IPerpetualManager.contract.UnpackLog(event, "PerpetualLimitOrderCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePerpetualLimitOrderCancelled is a log parse operation binding the contract event 0x053ad2eb505ab7dee10747f2b9b1572c82648c0c6e471c09631c4a10c79bcf9e.
//
// Solidity: event PerpetualLimitOrderCancelled(uint24 indexed perpetualId, bytes32 indexed orderHash)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParsePerpetualLimitOrderCancelled(log types.Log) (*IPerpetualManagerPerpetualLimitOrderCancelled, error) {
	event := new(IPerpetualManagerPerpetualLimitOrderCancelled)
	if err := _IPerpetualManager.contract.UnpackLog(event, "PerpetualLimitOrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerRunLiquidityPoolIterator is returned from FilterRunLiquidityPool and is used to iterate over the raw logs and unpacked data for RunLiquidityPool events raised by the IPerpetualManager contract.
type IPerpetualManagerRunLiquidityPoolIterator struct {
	Event *IPerpetualManagerRunLiquidityPool // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerRunLiquidityPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerRunLiquidityPool)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerRunLiquidityPool)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerRunLiquidityPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerRunLiquidityPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerRunLiquidityPool represents a RunLiquidityPool event raised by the IPerpetualManager contract.
type IPerpetualManagerRunLiquidityPool struct {
	LiqPoolID uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRunLiquidityPool is a free log retrieval operation binding the contract event 0x55ca3285289fcd3031b400400b5dfa175f6522e3343b00c093b7939ec6654030.
//
// Solidity: event RunLiquidityPool(uint8 _liqPoolID)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterRunLiquidityPool(opts *bind.FilterOpts) (*IPerpetualManagerRunLiquidityPoolIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "RunLiquidityPool")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerRunLiquidityPoolIterator{contract: _IPerpetualManager.contract, event: "RunLiquidityPool", logs: logs, sub: sub}, nil
}

// WatchRunLiquidityPool is a free log subscription operation binding the contract event 0x55ca3285289fcd3031b400400b5dfa175f6522e3343b00c093b7939ec6654030.
//
// Solidity: event RunLiquidityPool(uint8 _liqPoolID)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchRunLiquidityPool(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerRunLiquidityPool) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "RunLiquidityPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerRunLiquidityPool)
				if err := _IPerpetualManager.contract.UnpackLog(event, "RunLiquidityPool", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRunLiquidityPool is a log parse operation binding the contract event 0x55ca3285289fcd3031b400400b5dfa175f6522e3343b00c093b7939ec6654030.
//
// Solidity: event RunLiquidityPool(uint8 _liqPoolID)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseRunLiquidityPool(log types.Log) (*IPerpetualManagerRunLiquidityPool, error) {
	event := new(IPerpetualManagerRunLiquidityPool)
	if err := _IPerpetualManager.contract.UnpackLog(event, "RunLiquidityPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetBlockDelayIterator is returned from FilterSetBlockDelay and is used to iterate over the raw logs and unpacked data for SetBlockDelay events raised by the IPerpetualManager contract.
type IPerpetualManagerSetBlockDelayIterator struct {
	Event *IPerpetualManagerSetBlockDelay // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetBlockDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetBlockDelay)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetBlockDelay)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetBlockDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetBlockDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetBlockDelay represents a SetBlockDelay event raised by the IPerpetualManager contract.
type IPerpetualManagerSetBlockDelay struct {
	Delay uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetBlockDelay is a free log retrieval operation binding the contract event 0x3311843a3b80e236643be5f9a5369280c7c811b826a1716aecf61f5982ac6fcc.
//
// Solidity: event SetBlockDelay(uint8 delay)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetBlockDelay(opts *bind.FilterOpts) (*IPerpetualManagerSetBlockDelayIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetBlockDelay")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetBlockDelayIterator{contract: _IPerpetualManager.contract, event: "SetBlockDelay", logs: logs, sub: sub}, nil
}

// WatchSetBlockDelay is a free log subscription operation binding the contract event 0x3311843a3b80e236643be5f9a5369280c7c811b826a1716aecf61f5982ac6fcc.
//
// Solidity: event SetBlockDelay(uint8 delay)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetBlockDelay(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetBlockDelay) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetBlockDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetBlockDelay)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetBlockDelay", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBlockDelay is a log parse operation binding the contract event 0x3311843a3b80e236643be5f9a5369280c7c811b826a1716aecf61f5982ac6fcc.
//
// Solidity: event SetBlockDelay(uint8 delay)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetBlockDelay(log types.Log) (*IPerpetualManagerSetBlockDelay, error) {
	event := new(IPerpetualManagerSetBlockDelay)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetBlockDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetBrokerDesignationsIterator is returned from FilterSetBrokerDesignations and is used to iterate over the raw logs and unpacked data for SetBrokerDesignations events raised by the IPerpetualManager contract.
type IPerpetualManagerSetBrokerDesignationsIterator struct {
	Event *IPerpetualManagerSetBrokerDesignations // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetBrokerDesignationsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetBrokerDesignations)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetBrokerDesignations)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetBrokerDesignationsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetBrokerDesignationsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetBrokerDesignations represents a SetBrokerDesignations event raised by the IPerpetualManager contract.
type IPerpetualManagerSetBrokerDesignations struct {
	Designations []uint32
	Fees         []uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetBrokerDesignations is a free log retrieval operation binding the contract event 0x06bd6826cf18942e315836d65aa12a957ef2b9eacf9cd6461d69d9b2ecd96629.
//
// Solidity: event SetBrokerDesignations(uint32[] designations, uint16[] fees)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetBrokerDesignations(opts *bind.FilterOpts) (*IPerpetualManagerSetBrokerDesignationsIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetBrokerDesignations")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetBrokerDesignationsIterator{contract: _IPerpetualManager.contract, event: "SetBrokerDesignations", logs: logs, sub: sub}, nil
}

// WatchSetBrokerDesignations is a free log subscription operation binding the contract event 0x06bd6826cf18942e315836d65aa12a957ef2b9eacf9cd6461d69d9b2ecd96629.
//
// Solidity: event SetBrokerDesignations(uint32[] designations, uint16[] fees)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetBrokerDesignations(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetBrokerDesignations) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetBrokerDesignations")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetBrokerDesignations)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetBrokerDesignations", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBrokerDesignations is a log parse operation binding the contract event 0x06bd6826cf18942e315836d65aa12a957ef2b9eacf9cd6461d69d9b2ecd96629.
//
// Solidity: event SetBrokerDesignations(uint32[] designations, uint16[] fees)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetBrokerDesignations(log types.Log) (*IPerpetualManagerSetBrokerDesignations, error) {
	event := new(IPerpetualManagerSetBrokerDesignations)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetBrokerDesignations", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetBrokerTiersIterator is returned from FilterSetBrokerTiers and is used to iterate over the raw logs and unpacked data for SetBrokerTiers events raised by the IPerpetualManager contract.
type IPerpetualManagerSetBrokerTiersIterator struct {
	Event *IPerpetualManagerSetBrokerTiers // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetBrokerTiersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetBrokerTiers)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetBrokerTiers)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetBrokerTiersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetBrokerTiersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetBrokerTiers represents a SetBrokerTiers event raised by the IPerpetualManager contract.
type IPerpetualManagerSetBrokerTiers struct {
	Tiers    []*big.Int
	FeesTbps []uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetBrokerTiers is a free log retrieval operation binding the contract event 0xbbaed6a7880442c105c24e8c13293bfb0357ecffc3b8d23d1cd742a7617f3793.
//
// Solidity: event SetBrokerTiers(uint256[] tiers, uint16[] feesTbps)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetBrokerTiers(opts *bind.FilterOpts) (*IPerpetualManagerSetBrokerTiersIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetBrokerTiers")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetBrokerTiersIterator{contract: _IPerpetualManager.contract, event: "SetBrokerTiers", logs: logs, sub: sub}, nil
}

// WatchSetBrokerTiers is a free log subscription operation binding the contract event 0xbbaed6a7880442c105c24e8c13293bfb0357ecffc3b8d23d1cd742a7617f3793.
//
// Solidity: event SetBrokerTiers(uint256[] tiers, uint16[] feesTbps)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetBrokerTiers(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetBrokerTiers) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetBrokerTiers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetBrokerTiers)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetBrokerTiers", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBrokerTiers is a log parse operation binding the contract event 0xbbaed6a7880442c105c24e8c13293bfb0357ecffc3b8d23d1cd742a7617f3793.
//
// Solidity: event SetBrokerTiers(uint256[] tiers, uint16[] feesTbps)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetBrokerTiers(log types.Log) (*IPerpetualManagerSetBrokerTiers, error) {
	event := new(IPerpetualManagerSetBrokerTiers)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetBrokerTiers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetBrokerVolumeTiersIterator is returned from FilterSetBrokerVolumeTiers and is used to iterate over the raw logs and unpacked data for SetBrokerVolumeTiers events raised by the IPerpetualManager contract.
type IPerpetualManagerSetBrokerVolumeTiersIterator struct {
	Event *IPerpetualManagerSetBrokerVolumeTiers // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetBrokerVolumeTiersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetBrokerVolumeTiers)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetBrokerVolumeTiers)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetBrokerVolumeTiersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetBrokerVolumeTiersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetBrokerVolumeTiers represents a SetBrokerVolumeTiers event raised by the IPerpetualManager contract.
type IPerpetualManagerSetBrokerVolumeTiers struct {
	Tiers    []*big.Int
	FeesTbps []uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetBrokerVolumeTiers is a free log retrieval operation binding the contract event 0x82d2db5f3fde7918747b9806de47201b97a8e3f31e35f9f12b1d16d9fd98941e.
//
// Solidity: event SetBrokerVolumeTiers(uint256[] tiers, uint16[] feesTbps)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetBrokerVolumeTiers(opts *bind.FilterOpts) (*IPerpetualManagerSetBrokerVolumeTiersIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetBrokerVolumeTiers")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetBrokerVolumeTiersIterator{contract: _IPerpetualManager.contract, event: "SetBrokerVolumeTiers", logs: logs, sub: sub}, nil
}

// WatchSetBrokerVolumeTiers is a free log subscription operation binding the contract event 0x82d2db5f3fde7918747b9806de47201b97a8e3f31e35f9f12b1d16d9fd98941e.
//
// Solidity: event SetBrokerVolumeTiers(uint256[] tiers, uint16[] feesTbps)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetBrokerVolumeTiers(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetBrokerVolumeTiers) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetBrokerVolumeTiers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetBrokerVolumeTiers)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetBrokerVolumeTiers", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBrokerVolumeTiers is a log parse operation binding the contract event 0x82d2db5f3fde7918747b9806de47201b97a8e3f31e35f9f12b1d16d9fd98941e.
//
// Solidity: event SetBrokerVolumeTiers(uint256[] tiers, uint16[] feesTbps)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetBrokerVolumeTiers(log types.Log) (*IPerpetualManagerSetBrokerVolumeTiers, error) {
	event := new(IPerpetualManagerSetBrokerVolumeTiers)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetBrokerVolumeTiers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetClearedStateIterator is returned from FilterSetClearedState and is used to iterate over the raw logs and unpacked data for SetClearedState events raised by the IPerpetualManager contract.
type IPerpetualManagerSetClearedStateIterator struct {
	Event *IPerpetualManagerSetClearedState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetClearedStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetClearedState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetClearedState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetClearedStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetClearedStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetClearedState represents a SetClearedState event raised by the IPerpetualManager contract.
type IPerpetualManagerSetClearedState struct {
	PerpetualId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetClearedState is a free log retrieval operation binding the contract event 0x128d100a6c057af149e59f78a0bf028cca822f9cf40ebfda523346defa6ae080.
//
// Solidity: event SetClearedState(uint24 indexed perpetualId)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetClearedState(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerSetClearedStateIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetClearedState", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetClearedStateIterator{contract: _IPerpetualManager.contract, event: "SetClearedState", logs: logs, sub: sub}, nil
}

// WatchSetClearedState is a free log subscription operation binding the contract event 0x128d100a6c057af149e59f78a0bf028cca822f9cf40ebfda523346defa6ae080.
//
// Solidity: event SetClearedState(uint24 indexed perpetualId)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetClearedState(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetClearedState, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetClearedState", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetClearedState)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetClearedState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetClearedState is a log parse operation binding the contract event 0x128d100a6c057af149e59f78a0bf028cca822f9cf40ebfda523346defa6ae080.
//
// Solidity: event SetClearedState(uint24 indexed perpetualId)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetClearedState(log types.Log) (*IPerpetualManagerSetClearedState, error) {
	event := new(IPerpetualManagerSetClearedState)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetClearedState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetEmergencyStateIterator is returned from FilterSetEmergencyState and is used to iterate over the raw logs and unpacked data for SetEmergencyState events raised by the IPerpetualManager contract.
type IPerpetualManagerSetEmergencyStateIterator struct {
	Event *IPerpetualManagerSetEmergencyState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetEmergencyStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetEmergencyState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetEmergencyState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetEmergencyStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetEmergencyStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetEmergencyState represents a SetEmergencyState event raised by the IPerpetualManager contract.
type IPerpetualManagerSetEmergencyState struct {
	PerpetualId                *big.Int
	FSettlementMarkPremiumRate *big.Int
	FSettlementS2Price         *big.Int
	FSettlementS3Price         *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterSetEmergencyState is a free log retrieval operation binding the contract event 0x6dcc793202e7499de905b97ada46ead67481162a6dac20d8c013378d1da8886c.
//
// Solidity: event SetEmergencyState(uint24 indexed perpetualId, int128 fSettlementMarkPremiumRate, int128 fSettlementS2Price, int128 fSettlementS3Price)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetEmergencyState(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerSetEmergencyStateIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetEmergencyState", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetEmergencyStateIterator{contract: _IPerpetualManager.contract, event: "SetEmergencyState", logs: logs, sub: sub}, nil
}

// WatchSetEmergencyState is a free log subscription operation binding the contract event 0x6dcc793202e7499de905b97ada46ead67481162a6dac20d8c013378d1da8886c.
//
// Solidity: event SetEmergencyState(uint24 indexed perpetualId, int128 fSettlementMarkPremiumRate, int128 fSettlementS2Price, int128 fSettlementS3Price)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetEmergencyState(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetEmergencyState, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetEmergencyState", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetEmergencyState)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetEmergencyState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetEmergencyState is a log parse operation binding the contract event 0x6dcc793202e7499de905b97ada46ead67481162a6dac20d8c013378d1da8886c.
//
// Solidity: event SetEmergencyState(uint24 indexed perpetualId, int128 fSettlementMarkPremiumRate, int128 fSettlementS2Price, int128 fSettlementS3Price)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetEmergencyState(log types.Log) (*IPerpetualManagerSetEmergencyState, error) {
	event := new(IPerpetualManagerSetEmergencyState)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetEmergencyState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetNormalStateIterator is returned from FilterSetNormalState and is used to iterate over the raw logs and unpacked data for SetNormalState events raised by the IPerpetualManager contract.
type IPerpetualManagerSetNormalStateIterator struct {
	Event *IPerpetualManagerSetNormalState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetNormalStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetNormalState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetNormalState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetNormalStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetNormalStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetNormalState represents a SetNormalState event raised by the IPerpetualManager contract.
type IPerpetualManagerSetNormalState struct {
	PerpetualId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetNormalState is a free log retrieval operation binding the contract event 0x31030c3db0583ae649fddbe7c2cafb61a45d610d4a51e40a86a031f740794a80.
//
// Solidity: event SetNormalState(uint24 indexed perpetualId)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetNormalState(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerSetNormalStateIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetNormalState", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetNormalStateIterator{contract: _IPerpetualManager.contract, event: "SetNormalState", logs: logs, sub: sub}, nil
}

// WatchSetNormalState is a free log subscription operation binding the contract event 0x31030c3db0583ae649fddbe7c2cafb61a45d610d4a51e40a86a031f740794a80.
//
// Solidity: event SetNormalState(uint24 indexed perpetualId)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetNormalState(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetNormalState, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetNormalState", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetNormalState)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetNormalState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetNormalState is a log parse operation binding the contract event 0x31030c3db0583ae649fddbe7c2cafb61a45d610d4a51e40a86a031f740794a80.
//
// Solidity: event SetNormalState(uint24 indexed perpetualId)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetNormalState(log types.Log) (*IPerpetualManagerSetNormalState, error) {
	event := new(IPerpetualManagerSetNormalState)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetNormalState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetOraclesIterator is returned from FilterSetOracles and is used to iterate over the raw logs and unpacked data for SetOracles events raised by the IPerpetualManager contract.
type IPerpetualManagerSetOraclesIterator struct {
	Event *IPerpetualManagerSetOracles // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetOraclesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetOracles)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetOracles)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetOraclesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetOraclesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetOracles represents a SetOracles event raised by the IPerpetualManager contract.
type IPerpetualManagerSetOracles struct {
	PerpetualId *big.Int
	BaseQuoteS2 [2][4]byte
	BaseQuoteS3 [2][4]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetOracles is a free log retrieval operation binding the contract event 0xec2386af9dfc495fd639d1560f5d01c690e80a43868e8c20a4893b1d996c2c1f.
//
// Solidity: event SetOracles(uint24 indexed perpetualId, bytes4[2] baseQuoteS2, bytes4[2] baseQuoteS3)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetOracles(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerSetOraclesIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetOracles", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetOraclesIterator{contract: _IPerpetualManager.contract, event: "SetOracles", logs: logs, sub: sub}, nil
}

// WatchSetOracles is a free log subscription operation binding the contract event 0xec2386af9dfc495fd639d1560f5d01c690e80a43868e8c20a4893b1d996c2c1f.
//
// Solidity: event SetOracles(uint24 indexed perpetualId, bytes4[2] baseQuoteS2, bytes4[2] baseQuoteS3)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetOracles(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetOracles, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetOracles", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetOracles)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetOracles", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetOracles is a log parse operation binding the contract event 0xec2386af9dfc495fd639d1560f5d01c690e80a43868e8c20a4893b1d996c2c1f.
//
// Solidity: event SetOracles(uint24 indexed perpetualId, bytes4[2] baseQuoteS2, bytes4[2] baseQuoteS3)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetOracles(log types.Log) (*IPerpetualManagerSetOracles, error) {
	event := new(IPerpetualManagerSetOracles)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetOracles", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetParameterIterator is returned from FilterSetParameter and is used to iterate over the raw logs and unpacked data for SetParameter events raised by the IPerpetualManager contract.
type IPerpetualManagerSetParameterIterator struct {
	Event *IPerpetualManagerSetParameter // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetParameterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetParameter)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetParameter)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetParameterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetParameterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetParameter represents a SetParameter event raised by the IPerpetualManager contract.
type IPerpetualManagerSetParameter struct {
	PerpetualId *big.Int
	Name        string
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetParameter is a free log retrieval operation binding the contract event 0xe660b389c8472d3482cecc129c4bca39701d8af41f2a645d1c57beb4e7694a79.
//
// Solidity: event SetParameter(uint24 indexed perpetualId, string name, int128 value)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetParameter(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerSetParameterIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetParameter", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetParameterIterator{contract: _IPerpetualManager.contract, event: "SetParameter", logs: logs, sub: sub}, nil
}

// WatchSetParameter is a free log subscription operation binding the contract event 0xe660b389c8472d3482cecc129c4bca39701d8af41f2a645d1c57beb4e7694a79.
//
// Solidity: event SetParameter(uint24 indexed perpetualId, string name, int128 value)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetParameter(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetParameter, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetParameter", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetParameter)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetParameter", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetParameter is a log parse operation binding the contract event 0xe660b389c8472d3482cecc129c4bca39701d8af41f2a645d1c57beb4e7694a79.
//
// Solidity: event SetParameter(uint24 indexed perpetualId, string name, int128 value)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetParameter(log types.Log) (*IPerpetualManagerSetParameter, error) {
	event := new(IPerpetualManagerSetParameter)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetParameter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetParameterPairIterator is returned from FilterSetParameterPair and is used to iterate over the raw logs and unpacked data for SetParameterPair events raised by the IPerpetualManager contract.
type IPerpetualManagerSetParameterPairIterator struct {
	Event *IPerpetualManagerSetParameterPair // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetParameterPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetParameterPair)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetParameterPair)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetParameterPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetParameterPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetParameterPair represents a SetParameterPair event raised by the IPerpetualManager contract.
type IPerpetualManagerSetParameterPair struct {
	PerpetualId *big.Int
	Name        string
	Value1      *big.Int
	Value2      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetParameterPair is a free log retrieval operation binding the contract event 0xe39cb9c14d1c32d0da00001288d72d1326d520c068b01cbeed1bc3dee71c5264.
//
// Solidity: event SetParameterPair(uint24 indexed perpetualId, string name, int128 value1, int128 value2)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetParameterPair(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerSetParameterPairIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetParameterPair", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetParameterPairIterator{contract: _IPerpetualManager.contract, event: "SetParameterPair", logs: logs, sub: sub}, nil
}

// WatchSetParameterPair is a free log subscription operation binding the contract event 0xe39cb9c14d1c32d0da00001288d72d1326d520c068b01cbeed1bc3dee71c5264.
//
// Solidity: event SetParameterPair(uint24 indexed perpetualId, string name, int128 value1, int128 value2)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetParameterPair(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetParameterPair, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetParameterPair", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetParameterPair)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetParameterPair", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetParameterPair is a log parse operation binding the contract event 0xe39cb9c14d1c32d0da00001288d72d1326d520c068b01cbeed1bc3dee71c5264.
//
// Solidity: event SetParameterPair(uint24 indexed perpetualId, string name, int128 value1, int128 value2)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetParameterPair(log types.Log) (*IPerpetualManagerSetParameterPair, error) {
	event := new(IPerpetualManagerSetParameterPair)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetParameterPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetPerpetualBaseParametersIterator is returned from FilterSetPerpetualBaseParameters and is used to iterate over the raw logs and unpacked data for SetPerpetualBaseParameters events raised by the IPerpetualManager contract.
type IPerpetualManagerSetPerpetualBaseParametersIterator struct {
	Event *IPerpetualManagerSetPerpetualBaseParameters // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetPerpetualBaseParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetPerpetualBaseParameters)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetPerpetualBaseParameters)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetPerpetualBaseParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetPerpetualBaseParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetPerpetualBaseParameters represents a SetPerpetualBaseParameters event raised by the IPerpetualManager contract.
type IPerpetualManagerSetPerpetualBaseParameters struct {
	PerpetualId *big.Int
	BaseParams  [7]*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetPerpetualBaseParameters is a free log retrieval operation binding the contract event 0x78fdff43ac50b618922b36eaa7b8ca5f0877eaa7bff6e9ec687f752c20c08a4a.
//
// Solidity: event SetPerpetualBaseParameters(uint24 indexed perpetualId, int128[7] baseParams)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetPerpetualBaseParameters(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerSetPerpetualBaseParametersIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetPerpetualBaseParameters", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetPerpetualBaseParametersIterator{contract: _IPerpetualManager.contract, event: "SetPerpetualBaseParameters", logs: logs, sub: sub}, nil
}

// WatchSetPerpetualBaseParameters is a free log subscription operation binding the contract event 0x78fdff43ac50b618922b36eaa7b8ca5f0877eaa7bff6e9ec687f752c20c08a4a.
//
// Solidity: event SetPerpetualBaseParameters(uint24 indexed perpetualId, int128[7] baseParams)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetPerpetualBaseParameters(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetPerpetualBaseParameters, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetPerpetualBaseParameters", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetPerpetualBaseParameters)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetPerpetualBaseParameters", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPerpetualBaseParameters is a log parse operation binding the contract event 0x78fdff43ac50b618922b36eaa7b8ca5f0877eaa7bff6e9ec687f752c20c08a4a.
//
// Solidity: event SetPerpetualBaseParameters(uint24 indexed perpetualId, int128[7] baseParams)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetPerpetualBaseParameters(log types.Log) (*IPerpetualManagerSetPerpetualBaseParameters, error) {
	event := new(IPerpetualManagerSetPerpetualBaseParameters)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetPerpetualBaseParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetPerpetualRiskParametersIterator is returned from FilterSetPerpetualRiskParameters and is used to iterate over the raw logs and unpacked data for SetPerpetualRiskParameters events raised by the IPerpetualManager contract.
type IPerpetualManagerSetPerpetualRiskParametersIterator struct {
	Event *IPerpetualManagerSetPerpetualRiskParameters // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetPerpetualRiskParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetPerpetualRiskParameters)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetPerpetualRiskParameters)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetPerpetualRiskParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetPerpetualRiskParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetPerpetualRiskParameters represents a SetPerpetualRiskParameters event raised by the IPerpetualManager contract.
type IPerpetualManagerSetPerpetualRiskParameters struct {
	PerpetualId           *big.Int
	UnderlyingRiskParams  [5]*big.Int
	DefaultFundRiskParams [13]*big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetPerpetualRiskParameters is a free log retrieval operation binding the contract event 0xc33b635190db9337e21794049b85610a4d1b1b44f83fcec4ff4afd58f05bc0ba.
//
// Solidity: event SetPerpetualRiskParameters(uint24 indexed perpetualId, int128[5] underlyingRiskParams, int128[13] defaultFundRiskParams)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetPerpetualRiskParameters(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerSetPerpetualRiskParametersIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetPerpetualRiskParameters", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetPerpetualRiskParametersIterator{contract: _IPerpetualManager.contract, event: "SetPerpetualRiskParameters", logs: logs, sub: sub}, nil
}

// WatchSetPerpetualRiskParameters is a free log subscription operation binding the contract event 0xc33b635190db9337e21794049b85610a4d1b1b44f83fcec4ff4afd58f05bc0ba.
//
// Solidity: event SetPerpetualRiskParameters(uint24 indexed perpetualId, int128[5] underlyingRiskParams, int128[13] defaultFundRiskParams)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetPerpetualRiskParameters(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetPerpetualRiskParameters, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetPerpetualRiskParameters", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetPerpetualRiskParameters)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetPerpetualRiskParameters", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPerpetualRiskParameters is a log parse operation binding the contract event 0xc33b635190db9337e21794049b85610a4d1b1b44f83fcec4ff4afd58f05bc0ba.
//
// Solidity: event SetPerpetualRiskParameters(uint24 indexed perpetualId, int128[5] underlyingRiskParams, int128[13] defaultFundRiskParams)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetPerpetualRiskParameters(log types.Log) (*IPerpetualManagerSetPerpetualRiskParameters, error) {
	event := new(IPerpetualManagerSetPerpetualRiskParameters)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetPerpetualRiskParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetPoolParameterIterator is returned from FilterSetPoolParameter and is used to iterate over the raw logs and unpacked data for SetPoolParameter events raised by the IPerpetualManager contract.
type IPerpetualManagerSetPoolParameterIterator struct {
	Event *IPerpetualManagerSetPoolParameter // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetPoolParameterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetPoolParameter)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetPoolParameter)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetPoolParameterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetPoolParameterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetPoolParameter represents a SetPoolParameter event raised by the IPerpetualManager contract.
type IPerpetualManagerSetPoolParameter struct {
	PoolId uint8
	Name   string
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPoolParameter is a free log retrieval operation binding the contract event 0x6747c4de7169039c1c38954c282f79cfeb76f56f43763b9e964ef8867ce63b72.
//
// Solidity: event SetPoolParameter(uint8 indexed poolId, string name, int128 value)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetPoolParameter(opts *bind.FilterOpts, poolId []uint8) (*IPerpetualManagerSetPoolParameterIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetPoolParameter", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetPoolParameterIterator{contract: _IPerpetualManager.contract, event: "SetPoolParameter", logs: logs, sub: sub}, nil
}

// WatchSetPoolParameter is a free log subscription operation binding the contract event 0x6747c4de7169039c1c38954c282f79cfeb76f56f43763b9e964ef8867ce63b72.
//
// Solidity: event SetPoolParameter(uint8 indexed poolId, string name, int128 value)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetPoolParameter(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetPoolParameter, poolId []uint8) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetPoolParameter", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetPoolParameter)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetPoolParameter", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPoolParameter is a log parse operation binding the contract event 0x6747c4de7169039c1c38954c282f79cfeb76f56f43763b9e964ef8867ce63b72.
//
// Solidity: event SetPoolParameter(uint8 indexed poolId, string name, int128 value)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetPoolParameter(log types.Log) (*IPerpetualManagerSetPoolParameter, error) {
	event := new(IPerpetualManagerSetPoolParameter)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetPoolParameter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetTraderTiersIterator is returned from FilterSetTraderTiers and is used to iterate over the raw logs and unpacked data for SetTraderTiers events raised by the IPerpetualManager contract.
type IPerpetualManagerSetTraderTiersIterator struct {
	Event *IPerpetualManagerSetTraderTiers // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetTraderTiersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetTraderTiers)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetTraderTiers)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetTraderTiersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetTraderTiersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetTraderTiers represents a SetTraderTiers event raised by the IPerpetualManager contract.
type IPerpetualManagerSetTraderTiers struct {
	Tiers    []*big.Int
	FeesTbps []uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetTraderTiers is a free log retrieval operation binding the contract event 0xcf984eda28d9412e92f5da392de58eac4e0f88da0881e1c3fc905e3380652ea0.
//
// Solidity: event SetTraderTiers(uint256[] tiers, uint16[] feesTbps)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetTraderTiers(opts *bind.FilterOpts) (*IPerpetualManagerSetTraderTiersIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetTraderTiers")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetTraderTiersIterator{contract: _IPerpetualManager.contract, event: "SetTraderTiers", logs: logs, sub: sub}, nil
}

// WatchSetTraderTiers is a free log subscription operation binding the contract event 0xcf984eda28d9412e92f5da392de58eac4e0f88da0881e1c3fc905e3380652ea0.
//
// Solidity: event SetTraderTiers(uint256[] tiers, uint16[] feesTbps)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetTraderTiers(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetTraderTiers) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetTraderTiers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetTraderTiers)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetTraderTiers", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTraderTiers is a log parse operation binding the contract event 0xcf984eda28d9412e92f5da392de58eac4e0f88da0881e1c3fc905e3380652ea0.
//
// Solidity: event SetTraderTiers(uint256[] tiers, uint16[] feesTbps)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetTraderTiers(log types.Log) (*IPerpetualManagerSetTraderTiers, error) {
	event := new(IPerpetualManagerSetTraderTiers)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetTraderTiers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetTraderVolumeTiersIterator is returned from FilterSetTraderVolumeTiers and is used to iterate over the raw logs and unpacked data for SetTraderVolumeTiers events raised by the IPerpetualManager contract.
type IPerpetualManagerSetTraderVolumeTiersIterator struct {
	Event *IPerpetualManagerSetTraderVolumeTiers // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetTraderVolumeTiersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetTraderVolumeTiers)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetTraderVolumeTiers)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetTraderVolumeTiersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetTraderVolumeTiersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetTraderVolumeTiers represents a SetTraderVolumeTiers event raised by the IPerpetualManager contract.
type IPerpetualManagerSetTraderVolumeTiers struct {
	Tiers    []*big.Int
	FeesTbps []uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetTraderVolumeTiers is a free log retrieval operation binding the contract event 0x35b8209becc212ddb60a6d5072c76e6abc631d85ddef5cf9cb4fbb5a08ceb47f.
//
// Solidity: event SetTraderVolumeTiers(uint256[] tiers, uint16[] feesTbps)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetTraderVolumeTiers(opts *bind.FilterOpts) (*IPerpetualManagerSetTraderVolumeTiersIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetTraderVolumeTiers")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetTraderVolumeTiersIterator{contract: _IPerpetualManager.contract, event: "SetTraderVolumeTiers", logs: logs, sub: sub}, nil
}

// WatchSetTraderVolumeTiers is a free log subscription operation binding the contract event 0x35b8209becc212ddb60a6d5072c76e6abc631d85ddef5cf9cb4fbb5a08ceb47f.
//
// Solidity: event SetTraderVolumeTiers(uint256[] tiers, uint16[] feesTbps)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetTraderVolumeTiers(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetTraderVolumeTiers) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetTraderVolumeTiers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetTraderVolumeTiers)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetTraderVolumeTiers", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTraderVolumeTiers is a log parse operation binding the contract event 0x35b8209becc212ddb60a6d5072c76e6abc631d85ddef5cf9cb4fbb5a08ceb47f.
//
// Solidity: event SetTraderVolumeTiers(uint256[] tiers, uint16[] feesTbps)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetTraderVolumeTiers(log types.Log) (*IPerpetualManagerSetTraderVolumeTiers, error) {
	event := new(IPerpetualManagerSetTraderVolumeTiers)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetTraderVolumeTiers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSetUtilityTokenIterator is returned from FilterSetUtilityToken and is used to iterate over the raw logs and unpacked data for SetUtilityToken events raised by the IPerpetualManager contract.
type IPerpetualManagerSetUtilityTokenIterator struct {
	Event *IPerpetualManagerSetUtilityToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSetUtilityTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSetUtilityToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSetUtilityToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSetUtilityTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSetUtilityTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSetUtilityToken represents a SetUtilityToken event raised by the IPerpetualManager contract.
type IPerpetualManagerSetUtilityToken struct {
	TokenAddr common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetUtilityToken is a free log retrieval operation binding the contract event 0xc9cec33e5874737cd2e63540418a3c8e68d1893486e599aa85368e6ab56644a5.
//
// Solidity: event SetUtilityToken(address tokenAddr)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSetUtilityToken(opts *bind.FilterOpts) (*IPerpetualManagerSetUtilityTokenIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SetUtilityToken")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSetUtilityTokenIterator{contract: _IPerpetualManager.contract, event: "SetUtilityToken", logs: logs, sub: sub}, nil
}

// WatchSetUtilityToken is a free log subscription operation binding the contract event 0xc9cec33e5874737cd2e63540418a3c8e68d1893486e599aa85368e6ab56644a5.
//
// Solidity: event SetUtilityToken(address tokenAddr)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSetUtilityToken(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSetUtilityToken) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SetUtilityToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSetUtilityToken)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SetUtilityToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetUtilityToken is a log parse operation binding the contract event 0xc9cec33e5874737cd2e63540418a3c8e68d1893486e599aa85368e6ab56644a5.
//
// Solidity: event SetUtilityToken(address tokenAddr)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSetUtilityToken(log types.Log) (*IPerpetualManagerSetUtilityToken, error) {
	event := new(IPerpetualManagerSetUtilityToken)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SetUtilityToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSettleIterator is returned from FilterSettle and is used to iterate over the raw logs and unpacked data for Settle events raised by the IPerpetualManager contract.
type IPerpetualManagerSettleIterator struct {
	Event *IPerpetualManagerSettle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSettle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSettle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSettle represents a Settle event raised by the IPerpetualManager contract.
type IPerpetualManagerSettle struct {
	PerpetualId *big.Int
	Trader      common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0x4384b4ebb0b9033163d414511512e92dc0085472186f0a6b72b71503d6cca294.
//
// Solidity: event Settle(uint24 indexed perpetualId, address indexed trader, int256 amount)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSettle(opts *bind.FilterOpts, perpetualId []*big.Int, trader []common.Address) (*IPerpetualManagerSettleIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "Settle", perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSettleIterator{contract: _IPerpetualManager.contract, event: "Settle", logs: logs, sub: sub}, nil
}

// WatchSettle is a free log subscription operation binding the contract event 0x4384b4ebb0b9033163d414511512e92dc0085472186f0a6b72b71503d6cca294.
//
// Solidity: event Settle(uint24 indexed perpetualId, address indexed trader, int256 amount)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSettle(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSettle, perpetualId []*big.Int, trader []common.Address) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "Settle", perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSettle)
				if err := _IPerpetualManager.contract.UnpackLog(event, "Settle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettle is a log parse operation binding the contract event 0x4384b4ebb0b9033163d414511512e92dc0085472186f0a6b72b71503d6cca294.
//
// Solidity: event Settle(uint24 indexed perpetualId, address indexed trader, int256 amount)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSettle(log types.Log) (*IPerpetualManagerSettle, error) {
	event := new(IPerpetualManagerSettle)
	if err := _IPerpetualManager.contract.UnpackLog(event, "Settle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerSettleStateIterator is returned from FilterSettleState and is used to iterate over the raw logs and unpacked data for SettleState events raised by the IPerpetualManager contract.
type IPerpetualManagerSettleStateIterator struct {
	Event *IPerpetualManagerSettleState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerSettleStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerSettleState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerSettleState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerSettleStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerSettleStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerSettleState represents a SettleState event raised by the IPerpetualManager contract.
type IPerpetualManagerSettleState struct {
	PerpetualId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSettleState is a free log retrieval operation binding the contract event 0xfc78faee463ec43f0f34872b97e7cca76d94b67ad8dee00af55ab3965e8a434f.
//
// Solidity: event SettleState(uint24 indexed perpetualId)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterSettleState(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerSettleStateIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "SettleState", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerSettleStateIterator{contract: _IPerpetualManager.contract, event: "SettleState", logs: logs, sub: sub}, nil
}

// WatchSettleState is a free log subscription operation binding the contract event 0xfc78faee463ec43f0f34872b97e7cca76d94b67ad8dee00af55ab3965e8a434f.
//
// Solidity: event SettleState(uint24 indexed perpetualId)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchSettleState(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerSettleState, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "SettleState", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerSettleState)
				if err := _IPerpetualManager.contract.UnpackLog(event, "SettleState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettleState is a log parse operation binding the contract event 0xfc78faee463ec43f0f34872b97e7cca76d94b67ad8dee00af55ab3965e8a434f.
//
// Solidity: event SettleState(uint24 indexed perpetualId)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseSettleState(log types.Log) (*IPerpetualManagerSettleState, error) {
	event := new(IPerpetualManagerSettleState)
	if err := _IPerpetualManager.contract.UnpackLog(event, "SettleState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerTokensDepositedIterator is returned from FilterTokensDeposited and is used to iterate over the raw logs and unpacked data for TokensDeposited events raised by the IPerpetualManager contract.
type IPerpetualManagerTokensDepositedIterator struct {
	Event *IPerpetualManagerTokensDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerTokensDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerTokensDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerTokensDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerTokensDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerTokensDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerTokensDeposited represents a TokensDeposited event raised by the IPerpetualManager contract.
type IPerpetualManagerTokensDeposited struct {
	PerpetualId *big.Int
	Trader      common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensDeposited is a free log retrieval operation binding the contract event 0xc081bbcab6cd8e93398c2acb28d6801a5f6acae9e8c34bac7de54cc23b6cdbf7.
//
// Solidity: event TokensDeposited(uint24 indexed perpetualId, address indexed trader, int128 amount)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterTokensDeposited(opts *bind.FilterOpts, perpetualId []*big.Int, trader []common.Address) (*IPerpetualManagerTokensDepositedIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "TokensDeposited", perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerTokensDepositedIterator{contract: _IPerpetualManager.contract, event: "TokensDeposited", logs: logs, sub: sub}, nil
}

// WatchTokensDeposited is a free log subscription operation binding the contract event 0xc081bbcab6cd8e93398c2acb28d6801a5f6acae9e8c34bac7de54cc23b6cdbf7.
//
// Solidity: event TokensDeposited(uint24 indexed perpetualId, address indexed trader, int128 amount)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchTokensDeposited(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerTokensDeposited, perpetualId []*big.Int, trader []common.Address) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "TokensDeposited", perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerTokensDeposited)
				if err := _IPerpetualManager.contract.UnpackLog(event, "TokensDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensDeposited is a log parse operation binding the contract event 0xc081bbcab6cd8e93398c2acb28d6801a5f6acae9e8c34bac7de54cc23b6cdbf7.
//
// Solidity: event TokensDeposited(uint24 indexed perpetualId, address indexed trader, int128 amount)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseTokensDeposited(log types.Log) (*IPerpetualManagerTokensDeposited, error) {
	event := new(IPerpetualManagerTokensDeposited)
	if err := _IPerpetualManager.contract.UnpackLog(event, "TokensDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the IPerpetualManager contract.
type IPerpetualManagerTokensWithdrawnIterator struct {
	Event *IPerpetualManagerTokensWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerTokensWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerTokensWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerTokensWithdrawn represents a TokensWithdrawn event raised by the IPerpetualManager contract.
type IPerpetualManagerTokensWithdrawn struct {
	PerpetualId *big.Int
	Trader      common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x3163de37636170cf3c2fd98be539c57db8ccd0c53700e78b3526a55448be3c28.
//
// Solidity: event TokensWithdrawn(uint24 indexed perpetualId, address indexed trader, int128 amount)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, perpetualId []*big.Int, trader []common.Address) (*IPerpetualManagerTokensWithdrawnIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "TokensWithdrawn", perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerTokensWithdrawnIterator{contract: _IPerpetualManager.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x3163de37636170cf3c2fd98be539c57db8ccd0c53700e78b3526a55448be3c28.
//
// Solidity: event TokensWithdrawn(uint24 indexed perpetualId, address indexed trader, int128 amount)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerTokensWithdrawn, perpetualId []*big.Int, trader []common.Address) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "TokensWithdrawn", perpetualIdRule, traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerTokensWithdrawn)
				if err := _IPerpetualManager.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x3163de37636170cf3c2fd98be539c57db8ccd0c53700e78b3526a55448be3c28.
//
// Solidity: event TokensWithdrawn(uint24 indexed perpetualId, address indexed trader, int128 amount)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseTokensWithdrawn(log types.Log) (*IPerpetualManagerTokensWithdrawn, error) {
	event := new(IPerpetualManagerTokensWithdrawn)
	if err := _IPerpetualManager.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the IPerpetualManager contract.
type IPerpetualManagerTradeIterator struct {
	Event *IPerpetualManagerTrade // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerTrade)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerTrade)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerTrade represents a Trade event raised by the IPerpetualManager contract.
type IPerpetualManagerTrade struct {
	PerpetualId       *big.Int
	Trader            common.Address
	PositionId        [16]byte
	Order             IPerpetualOrderOrder
	OrderDigest       [32]byte
	NewPositionSizeBC *big.Int
	Price             *big.Int
	FFeeCC            *big.Int
	FPnlCC            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x0610199fca79ac7a0423c55f12765dac05aceeb3999e174c06fcc6303fe639b0.
//
// Solidity: event Trade(uint24 indexed perpetualId, address indexed trader, bytes16 indexed positionId, (uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) order, bytes32 orderDigest, int128 newPositionSizeBC, int128 price, int128 fFeeCC, int128 fPnlCC)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterTrade(opts *bind.FilterOpts, perpetualId []*big.Int, trader []common.Address, positionId [][16]byte) (*IPerpetualManagerTradeIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "Trade", perpetualIdRule, traderRule, positionIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerTradeIterator{contract: _IPerpetualManager.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x0610199fca79ac7a0423c55f12765dac05aceeb3999e174c06fcc6303fe639b0.
//
// Solidity: event Trade(uint24 indexed perpetualId, address indexed trader, bytes16 indexed positionId, (uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) order, bytes32 orderDigest, int128 newPositionSizeBC, int128 price, int128 fFeeCC, int128 fPnlCC)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerTrade, perpetualId []*big.Int, trader []common.Address, positionId [][16]byte) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "Trade", perpetualIdRule, traderRule, positionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerTrade)
				if err := _IPerpetualManager.contract.UnpackLog(event, "Trade", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTrade is a log parse operation binding the contract event 0x0610199fca79ac7a0423c55f12765dac05aceeb3999e174c06fcc6303fe639b0.
//
// Solidity: event Trade(uint24 indexed perpetualId, address indexed trader, bytes16 indexed positionId, (uint16,uint16,uint24,address,uint32,address,uint32,uint32,uint32,address,int128,int128,int128,bytes) order, bytes32 orderDigest, int128 newPositionSizeBC, int128 price, int128 fFeeCC, int128 fPnlCC)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseTrade(log types.Log) (*IPerpetualManagerTrade, error) {
	event := new(IPerpetualManagerTrade)
	if err := _IPerpetualManager.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerTransferAddressToIterator is returned from FilterTransferAddressTo and is used to iterate over the raw logs and unpacked data for TransferAddressTo events raised by the IPerpetualManager contract.
type IPerpetualManagerTransferAddressToIterator struct {
	Event *IPerpetualManagerTransferAddressTo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerTransferAddressToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerTransferAddressTo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerTransferAddressTo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerTransferAddressToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerTransferAddressToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerTransferAddressTo represents a TransferAddressTo event raised by the IPerpetualManager contract.
type IPerpetualManagerTransferAddressTo struct {
	Name         string
	OldOBFactory common.Address
	NewOBFactory common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferAddressTo is a free log retrieval operation binding the contract event 0xd9f95d40c0dcd3212ce867e5f948a5ff37256ec0ca671ec49afda2e1eeb99af7.
//
// Solidity: event TransferAddressTo(string name, address oldOBFactory, address newOBFactory)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterTransferAddressTo(opts *bind.FilterOpts) (*IPerpetualManagerTransferAddressToIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "TransferAddressTo")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerTransferAddressToIterator{contract: _IPerpetualManager.contract, event: "TransferAddressTo", logs: logs, sub: sub}, nil
}

// WatchTransferAddressTo is a free log subscription operation binding the contract event 0xd9f95d40c0dcd3212ce867e5f948a5ff37256ec0ca671ec49afda2e1eeb99af7.
//
// Solidity: event TransferAddressTo(string name, address oldOBFactory, address newOBFactory)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchTransferAddressTo(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerTransferAddressTo) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "TransferAddressTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerTransferAddressTo)
				if err := _IPerpetualManager.contract.UnpackLog(event, "TransferAddressTo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferAddressTo is a log parse operation binding the contract event 0xd9f95d40c0dcd3212ce867e5f948a5ff37256ec0ca671ec49afda2e1eeb99af7.
//
// Solidity: event TransferAddressTo(string name, address oldOBFactory, address newOBFactory)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseTransferAddressTo(log types.Log) (*IPerpetualManagerTransferAddressTo, error) {
	event := new(IPerpetualManagerTransferAddressTo)
	if err := _IPerpetualManager.contract.UnpackLog(event, "TransferAddressTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerTransferEarningsToTreasuryIterator is returned from FilterTransferEarningsToTreasury and is used to iterate over the raw logs and unpacked data for TransferEarningsToTreasury events raised by the IPerpetualManager contract.
type IPerpetualManagerTransferEarningsToTreasuryIterator struct {
	Event *IPerpetualManagerTransferEarningsToTreasury // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerTransferEarningsToTreasuryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerTransferEarningsToTreasury)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerTransferEarningsToTreasury)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerTransferEarningsToTreasuryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerTransferEarningsToTreasuryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerTransferEarningsToTreasury represents a TransferEarningsToTreasury event raised by the IPerpetualManager contract.
type IPerpetualManagerTransferEarningsToTreasury struct {
	PoolId             uint8
	FEarnings          *big.Int
	NewDefaultFundSize *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTransferEarningsToTreasury is a free log retrieval operation binding the contract event 0xeacf1cfe92dc410ce99fbaf80d0dd3d588660d669a4aabbb225cdeb4b71dba00.
//
// Solidity: event TransferEarningsToTreasury(uint8 _poolId, int128 fEarnings, int128 newDefaultFundSize)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterTransferEarningsToTreasury(opts *bind.FilterOpts) (*IPerpetualManagerTransferEarningsToTreasuryIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "TransferEarningsToTreasury")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerTransferEarningsToTreasuryIterator{contract: _IPerpetualManager.contract, event: "TransferEarningsToTreasury", logs: logs, sub: sub}, nil
}

// WatchTransferEarningsToTreasury is a free log subscription operation binding the contract event 0xeacf1cfe92dc410ce99fbaf80d0dd3d588660d669a4aabbb225cdeb4b71dba00.
//
// Solidity: event TransferEarningsToTreasury(uint8 _poolId, int128 fEarnings, int128 newDefaultFundSize)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchTransferEarningsToTreasury(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerTransferEarningsToTreasury) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "TransferEarningsToTreasury")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerTransferEarningsToTreasury)
				if err := _IPerpetualManager.contract.UnpackLog(event, "TransferEarningsToTreasury", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferEarningsToTreasury is a log parse operation binding the contract event 0xeacf1cfe92dc410ce99fbaf80d0dd3d588660d669a4aabbb225cdeb4b71dba00.
//
// Solidity: event TransferEarningsToTreasury(uint8 _poolId, int128 fEarnings, int128 newDefaultFundSize)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseTransferEarningsToTreasury(log types.Log) (*IPerpetualManagerTransferEarningsToTreasury, error) {
	event := new(IPerpetualManagerTransferEarningsToTreasury)
	if err := _IPerpetualManager.contract.UnpackLog(event, "TransferEarningsToTreasury", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerTransferFeeToBrokerIterator is returned from FilterTransferFeeToBroker and is used to iterate over the raw logs and unpacked data for TransferFeeToBroker events raised by the IPerpetualManager contract.
type IPerpetualManagerTransferFeeToBrokerIterator struct {
	Event *IPerpetualManagerTransferFeeToBroker // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerTransferFeeToBrokerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerTransferFeeToBroker)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerTransferFeeToBroker)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerTransferFeeToBrokerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerTransferFeeToBrokerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerTransferFeeToBroker represents a TransferFeeToBroker event raised by the IPerpetualManager contract.
type IPerpetualManagerTransferFeeToBroker struct {
	PerpetualId *big.Int
	Broker      common.Address
	FeeCC       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferFeeToBroker is a free log retrieval operation binding the contract event 0xc48ca9157dad0ede520b24a175f0d341c6dade287987db04b05d9bc0a08f1da3.
//
// Solidity: event TransferFeeToBroker(uint24 indexed perpetualId, address indexed broker, int128 feeCC)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterTransferFeeToBroker(opts *bind.FilterOpts, perpetualId []*big.Int, broker []common.Address) (*IPerpetualManagerTransferFeeToBrokerIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var brokerRule []interface{}
	for _, brokerItem := range broker {
		brokerRule = append(brokerRule, brokerItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "TransferFeeToBroker", perpetualIdRule, brokerRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerTransferFeeToBrokerIterator{contract: _IPerpetualManager.contract, event: "TransferFeeToBroker", logs: logs, sub: sub}, nil
}

// WatchTransferFeeToBroker is a free log subscription operation binding the contract event 0xc48ca9157dad0ede520b24a175f0d341c6dade287987db04b05d9bc0a08f1da3.
//
// Solidity: event TransferFeeToBroker(uint24 indexed perpetualId, address indexed broker, int128 feeCC)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchTransferFeeToBroker(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerTransferFeeToBroker, perpetualId []*big.Int, broker []common.Address) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var brokerRule []interface{}
	for _, brokerItem := range broker {
		brokerRule = append(brokerRule, brokerItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "TransferFeeToBroker", perpetualIdRule, brokerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerTransferFeeToBroker)
				if err := _IPerpetualManager.contract.UnpackLog(event, "TransferFeeToBroker", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferFeeToBroker is a log parse operation binding the contract event 0xc48ca9157dad0ede520b24a175f0d341c6dade287987db04b05d9bc0a08f1da3.
//
// Solidity: event TransferFeeToBroker(uint24 indexed perpetualId, address indexed broker, int128 feeCC)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseTransferFeeToBroker(log types.Log) (*IPerpetualManagerTransferFeeToBroker, error) {
	event := new(IPerpetualManagerTransferFeeToBroker)
	if err := _IPerpetualManager.contract.UnpackLog(event, "TransferFeeToBroker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerTransferFeeToReferrerIterator is returned from FilterTransferFeeToReferrer and is used to iterate over the raw logs and unpacked data for TransferFeeToReferrer events raised by the IPerpetualManager contract.
type IPerpetualManagerTransferFeeToReferrerIterator struct {
	Event *IPerpetualManagerTransferFeeToReferrer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerTransferFeeToReferrerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerTransferFeeToReferrer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerTransferFeeToReferrer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerTransferFeeToReferrerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerTransferFeeToReferrerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerTransferFeeToReferrer represents a TransferFeeToReferrer event raised by the IPerpetualManager contract.
type IPerpetualManagerTransferFeeToReferrer struct {
	PerpetualId    *big.Int
	Trader         common.Address
	Referrer       common.Address
	ReferralRebate *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransferFeeToReferrer is a free log retrieval operation binding the contract event 0xaf86c12bfbbd0990195198943563927f203e0fa650c743e8dcf57703a0d923b1.
//
// Solidity: event TransferFeeToReferrer(uint24 indexed perpetualId, address indexed trader, address indexed referrer, int128 referralRebate)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterTransferFeeToReferrer(opts *bind.FilterOpts, perpetualId []*big.Int, trader []common.Address, referrer []common.Address) (*IPerpetualManagerTransferFeeToReferrerIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "TransferFeeToReferrer", perpetualIdRule, traderRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerTransferFeeToReferrerIterator{contract: _IPerpetualManager.contract, event: "TransferFeeToReferrer", logs: logs, sub: sub}, nil
}

// WatchTransferFeeToReferrer is a free log subscription operation binding the contract event 0xaf86c12bfbbd0990195198943563927f203e0fa650c743e8dcf57703a0d923b1.
//
// Solidity: event TransferFeeToReferrer(uint24 indexed perpetualId, address indexed trader, address indexed referrer, int128 referralRebate)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchTransferFeeToReferrer(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerTransferFeeToReferrer, perpetualId []*big.Int, trader []common.Address, referrer []common.Address) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "TransferFeeToReferrer", perpetualIdRule, traderRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerTransferFeeToReferrer)
				if err := _IPerpetualManager.contract.UnpackLog(event, "TransferFeeToReferrer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferFeeToReferrer is a log parse operation binding the contract event 0xaf86c12bfbbd0990195198943563927f203e0fa650c743e8dcf57703a0d923b1.
//
// Solidity: event TransferFeeToReferrer(uint24 indexed perpetualId, address indexed trader, address indexed referrer, int128 referralRebate)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseTransferFeeToReferrer(log types.Log) (*IPerpetualManagerTransferFeeToReferrer, error) {
	event := new(IPerpetualManagerTransferFeeToReferrer)
	if err := _IPerpetualManager.contract.UnpackLog(event, "TransferFeeToReferrer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerUpdateAMMFundCashIterator is returned from FilterUpdateAMMFundCash and is used to iterate over the raw logs and unpacked data for UpdateAMMFundCash events raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateAMMFundCashIterator struct {
	Event *IPerpetualManagerUpdateAMMFundCash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerUpdateAMMFundCashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerUpdateAMMFundCash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerUpdateAMMFundCash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerUpdateAMMFundCashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerUpdateAMMFundCashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerUpdateAMMFundCash represents a UpdateAMMFundCash event raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateAMMFundCash struct {
	PerpetualId                  *big.Int
	FNewAMMFundCash              *big.Int
	FNewLiqPoolTotalAMMFundsCash *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterUpdateAMMFundCash is a free log retrieval operation binding the contract event 0x3be85c4c544a13727d4ea15a2b5774b1f6192e3f5a43b5bd5ae1737a395d15ec.
//
// Solidity: event UpdateAMMFundCash(uint24 indexed perpetualId, int128 fNewAMMFundCash, int128 fNewLiqPoolTotalAMMFundsCash)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterUpdateAMMFundCash(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerUpdateAMMFundCashIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "UpdateAMMFundCash", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerUpdateAMMFundCashIterator{contract: _IPerpetualManager.contract, event: "UpdateAMMFundCash", logs: logs, sub: sub}, nil
}

// WatchUpdateAMMFundCash is a free log subscription operation binding the contract event 0x3be85c4c544a13727d4ea15a2b5774b1f6192e3f5a43b5bd5ae1737a395d15ec.
//
// Solidity: event UpdateAMMFundCash(uint24 indexed perpetualId, int128 fNewAMMFundCash, int128 fNewLiqPoolTotalAMMFundsCash)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchUpdateAMMFundCash(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerUpdateAMMFundCash, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "UpdateAMMFundCash", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerUpdateAMMFundCash)
				if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateAMMFundCash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateAMMFundCash is a log parse operation binding the contract event 0x3be85c4c544a13727d4ea15a2b5774b1f6192e3f5a43b5bd5ae1737a395d15ec.
//
// Solidity: event UpdateAMMFundCash(uint24 indexed perpetualId, int128 fNewAMMFundCash, int128 fNewLiqPoolTotalAMMFundsCash)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseUpdateAMMFundCash(log types.Log) (*IPerpetualManagerUpdateAMMFundCash, error) {
	event := new(IPerpetualManagerUpdateAMMFundCash)
	if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateAMMFundCash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerUpdateAMMFundTargetSizeIterator is returned from FilterUpdateAMMFundTargetSize and is used to iterate over the raw logs and unpacked data for UpdateAMMFundTargetSize events raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateAMMFundTargetSizeIterator struct {
	Event *IPerpetualManagerUpdateAMMFundTargetSize // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerUpdateAMMFundTargetSizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerUpdateAMMFundTargetSize)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerUpdateAMMFundTargetSize)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerUpdateAMMFundTargetSizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerUpdateAMMFundTargetSizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerUpdateAMMFundTargetSize represents a UpdateAMMFundTargetSize event raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateAMMFundTargetSize struct {
	PerpetualId                   *big.Int
	LiquidityPoolId               uint8
	FAMMFundCashCCInPerpetual     *big.Int
	FTargetAMMFundSizeInPerpetual *big.Int
	FAMMFundCashCCInPool          *big.Int
	FTargetAMMFundSizeInPool      *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterUpdateAMMFundTargetSize is a free log retrieval operation binding the contract event 0xe5ba267b9d73c092e04e81fe6940e8bd657e463ba78b4c7ee2e9da3e3de563d5.
//
// Solidity: event UpdateAMMFundTargetSize(uint24 indexed perpetualId, uint8 indexed liquidityPoolId, int128 fAMMFundCashCCInPerpetual, int128 fTargetAMMFundSizeInPerpetual, int128 fAMMFundCashCCInPool, int128 fTargetAMMFundSizeInPool)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterUpdateAMMFundTargetSize(opts *bind.FilterOpts, perpetualId []*big.Int, liquidityPoolId []uint8) (*IPerpetualManagerUpdateAMMFundTargetSizeIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var liquidityPoolIdRule []interface{}
	for _, liquidityPoolIdItem := range liquidityPoolId {
		liquidityPoolIdRule = append(liquidityPoolIdRule, liquidityPoolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "UpdateAMMFundTargetSize", perpetualIdRule, liquidityPoolIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerUpdateAMMFundTargetSizeIterator{contract: _IPerpetualManager.contract, event: "UpdateAMMFundTargetSize", logs: logs, sub: sub}, nil
}

// WatchUpdateAMMFundTargetSize is a free log subscription operation binding the contract event 0xe5ba267b9d73c092e04e81fe6940e8bd657e463ba78b4c7ee2e9da3e3de563d5.
//
// Solidity: event UpdateAMMFundTargetSize(uint24 indexed perpetualId, uint8 indexed liquidityPoolId, int128 fAMMFundCashCCInPerpetual, int128 fTargetAMMFundSizeInPerpetual, int128 fAMMFundCashCCInPool, int128 fTargetAMMFundSizeInPool)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchUpdateAMMFundTargetSize(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerUpdateAMMFundTargetSize, perpetualId []*big.Int, liquidityPoolId []uint8) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var liquidityPoolIdRule []interface{}
	for _, liquidityPoolIdItem := range liquidityPoolId {
		liquidityPoolIdRule = append(liquidityPoolIdRule, liquidityPoolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "UpdateAMMFundTargetSize", perpetualIdRule, liquidityPoolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerUpdateAMMFundTargetSize)
				if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateAMMFundTargetSize", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateAMMFundTargetSize is a log parse operation binding the contract event 0xe5ba267b9d73c092e04e81fe6940e8bd657e463ba78b4c7ee2e9da3e3de563d5.
//
// Solidity: event UpdateAMMFundTargetSize(uint24 indexed perpetualId, uint8 indexed liquidityPoolId, int128 fAMMFundCashCCInPerpetual, int128 fTargetAMMFundSizeInPerpetual, int128 fAMMFundCashCCInPool, int128 fTargetAMMFundSizeInPool)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseUpdateAMMFundTargetSize(log types.Log) (*IPerpetualManagerUpdateAMMFundTargetSize, error) {
	event := new(IPerpetualManagerUpdateAMMFundTargetSize)
	if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateAMMFundTargetSize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerUpdateBrokerAddedCashIterator is returned from FilterUpdateBrokerAddedCash and is used to iterate over the raw logs and unpacked data for UpdateBrokerAddedCash events raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateBrokerAddedCashIterator struct {
	Event *IPerpetualManagerUpdateBrokerAddedCash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerUpdateBrokerAddedCashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerUpdateBrokerAddedCash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerUpdateBrokerAddedCash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerUpdateBrokerAddedCashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerUpdateBrokerAddedCashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerUpdateBrokerAddedCash represents a UpdateBrokerAddedCash event raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateBrokerAddedCash struct {
	PoolId         uint8
	ILots          uint32
	INewBrokerLots uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateBrokerAddedCash is a free log retrieval operation binding the contract event 0x073cf7dc9cc23ebbd73fcec084dc65394d232c4cd2d00105f59d838baf15fc05.
//
// Solidity: event UpdateBrokerAddedCash(uint8 indexed poolId, uint32 iLots, uint32 iNewBrokerLots)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterUpdateBrokerAddedCash(opts *bind.FilterOpts, poolId []uint8) (*IPerpetualManagerUpdateBrokerAddedCashIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "UpdateBrokerAddedCash", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerUpdateBrokerAddedCashIterator{contract: _IPerpetualManager.contract, event: "UpdateBrokerAddedCash", logs: logs, sub: sub}, nil
}

// WatchUpdateBrokerAddedCash is a free log subscription operation binding the contract event 0x073cf7dc9cc23ebbd73fcec084dc65394d232c4cd2d00105f59d838baf15fc05.
//
// Solidity: event UpdateBrokerAddedCash(uint8 indexed poolId, uint32 iLots, uint32 iNewBrokerLots)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchUpdateBrokerAddedCash(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerUpdateBrokerAddedCash, poolId []uint8) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "UpdateBrokerAddedCash", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerUpdateBrokerAddedCash)
				if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateBrokerAddedCash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateBrokerAddedCash is a log parse operation binding the contract event 0x073cf7dc9cc23ebbd73fcec084dc65394d232c4cd2d00105f59d838baf15fc05.
//
// Solidity: event UpdateBrokerAddedCash(uint8 indexed poolId, uint32 iLots, uint32 iNewBrokerLots)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseUpdateBrokerAddedCash(log types.Log) (*IPerpetualManagerUpdateBrokerAddedCash, error) {
	event := new(IPerpetualManagerUpdateBrokerAddedCash)
	if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateBrokerAddedCash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerUpdateDefaultFundCashIterator is returned from FilterUpdateDefaultFundCash and is used to iterate over the raw logs and unpacked data for UpdateDefaultFundCash events raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateDefaultFundCashIterator struct {
	Event *IPerpetualManagerUpdateDefaultFundCash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerUpdateDefaultFundCashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerUpdateDefaultFundCash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerUpdateDefaultFundCash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerUpdateDefaultFundCashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerUpdateDefaultFundCashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerUpdateDefaultFundCash represents a UpdateDefaultFundCash event raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateDefaultFundCash struct {
	PoolId         uint8
	FDeltaAmountCC *big.Int
	FNewFundCash   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultFundCash is a free log retrieval operation binding the contract event 0x38b3318863e4792ebc3bb07351f52bad5e84f517d77b302bead80374d06fdd04.
//
// Solidity: event UpdateDefaultFundCash(uint8 indexed poolId, int128 fDeltaAmountCC, int128 fNewFundCash)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterUpdateDefaultFundCash(opts *bind.FilterOpts, poolId []uint8) (*IPerpetualManagerUpdateDefaultFundCashIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "UpdateDefaultFundCash", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerUpdateDefaultFundCashIterator{contract: _IPerpetualManager.contract, event: "UpdateDefaultFundCash", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultFundCash is a free log subscription operation binding the contract event 0x38b3318863e4792ebc3bb07351f52bad5e84f517d77b302bead80374d06fdd04.
//
// Solidity: event UpdateDefaultFundCash(uint8 indexed poolId, int128 fDeltaAmountCC, int128 fNewFundCash)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchUpdateDefaultFundCash(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerUpdateDefaultFundCash, poolId []uint8) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "UpdateDefaultFundCash", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerUpdateDefaultFundCash)
				if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateDefaultFundCash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateDefaultFundCash is a log parse operation binding the contract event 0x38b3318863e4792ebc3bb07351f52bad5e84f517d77b302bead80374d06fdd04.
//
// Solidity: event UpdateDefaultFundCash(uint8 indexed poolId, int128 fDeltaAmountCC, int128 fNewFundCash)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseUpdateDefaultFundCash(log types.Log) (*IPerpetualManagerUpdateDefaultFundCash, error) {
	event := new(IPerpetualManagerUpdateDefaultFundCash)
	if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateDefaultFundCash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerUpdateDefaultFundTargetSizeIterator is returned from FilterUpdateDefaultFundTargetSize and is used to iterate over the raw logs and unpacked data for UpdateDefaultFundTargetSize events raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateDefaultFundTargetSizeIterator struct {
	Event *IPerpetualManagerUpdateDefaultFundTargetSize // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerUpdateDefaultFundTargetSizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerUpdateDefaultFundTargetSize)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerUpdateDefaultFundTargetSize)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerUpdateDefaultFundTargetSizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerUpdateDefaultFundTargetSizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerUpdateDefaultFundTargetSize represents a UpdateDefaultFundTargetSize event raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateDefaultFundTargetSize struct {
	LiquidityPoolId    uint8
	FDefaultFundCashCC *big.Int
	FTargetDFSize      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultFundTargetSize is a free log retrieval operation binding the contract event 0x5136c3531ecd05c0cae10a3a8a4c767b8b97d73d14f7febb510b826022988f1c.
//
// Solidity: event UpdateDefaultFundTargetSize(uint8 indexed liquidityPoolId, int128 fDefaultFundCashCC, int128 fTargetDFSize)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterUpdateDefaultFundTargetSize(opts *bind.FilterOpts, liquidityPoolId []uint8) (*IPerpetualManagerUpdateDefaultFundTargetSizeIterator, error) {

	var liquidityPoolIdRule []interface{}
	for _, liquidityPoolIdItem := range liquidityPoolId {
		liquidityPoolIdRule = append(liquidityPoolIdRule, liquidityPoolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "UpdateDefaultFundTargetSize", liquidityPoolIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerUpdateDefaultFundTargetSizeIterator{contract: _IPerpetualManager.contract, event: "UpdateDefaultFundTargetSize", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultFundTargetSize is a free log subscription operation binding the contract event 0x5136c3531ecd05c0cae10a3a8a4c767b8b97d73d14f7febb510b826022988f1c.
//
// Solidity: event UpdateDefaultFundTargetSize(uint8 indexed liquidityPoolId, int128 fDefaultFundCashCC, int128 fTargetDFSize)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchUpdateDefaultFundTargetSize(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerUpdateDefaultFundTargetSize, liquidityPoolId []uint8) (event.Subscription, error) {

	var liquidityPoolIdRule []interface{}
	for _, liquidityPoolIdItem := range liquidityPoolId {
		liquidityPoolIdRule = append(liquidityPoolIdRule, liquidityPoolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "UpdateDefaultFundTargetSize", liquidityPoolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerUpdateDefaultFundTargetSize)
				if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateDefaultFundTargetSize", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateDefaultFundTargetSize is a log parse operation binding the contract event 0x5136c3531ecd05c0cae10a3a8a4c767b8b97d73d14f7febb510b826022988f1c.
//
// Solidity: event UpdateDefaultFundTargetSize(uint8 indexed liquidityPoolId, int128 fDefaultFundCashCC, int128 fTargetDFSize)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseUpdateDefaultFundTargetSize(log types.Log) (*IPerpetualManagerUpdateDefaultFundTargetSize, error) {
	event := new(IPerpetualManagerUpdateDefaultFundTargetSize)
	if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateDefaultFundTargetSize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerUpdateFundAllocationWeightIterator is returned from FilterUpdateFundAllocationWeight and is used to iterate over the raw logs and unpacked data for UpdateFundAllocationWeight events raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateFundAllocationWeightIterator struct {
	Event *IPerpetualManagerUpdateFundAllocationWeight // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerUpdateFundAllocationWeightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerUpdateFundAllocationWeight)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerUpdateFundAllocationWeight)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerUpdateFundAllocationWeightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerUpdateFundAllocationWeightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerUpdateFundAllocationWeight represents a UpdateFundAllocationWeight event raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateFundAllocationWeight struct {
	PoolId                     uint8
	PerpetualId                *big.Int
	FNormalizedPerpetualWeight *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterUpdateFundAllocationWeight is a free log retrieval operation binding the contract event 0x0ee9fecde1957f72be053052fe8a6896acc39ab927d76e9a56e59b60aa3fa2c4.
//
// Solidity: event UpdateFundAllocationWeight(uint8 poolId, uint24 indexed perpetualId, int128 fNormalizedPerpetualWeight)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterUpdateFundAllocationWeight(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerUpdateFundAllocationWeightIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "UpdateFundAllocationWeight", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerUpdateFundAllocationWeightIterator{contract: _IPerpetualManager.contract, event: "UpdateFundAllocationWeight", logs: logs, sub: sub}, nil
}

// WatchUpdateFundAllocationWeight is a free log subscription operation binding the contract event 0x0ee9fecde1957f72be053052fe8a6896acc39ab927d76e9a56e59b60aa3fa2c4.
//
// Solidity: event UpdateFundAllocationWeight(uint8 poolId, uint24 indexed perpetualId, int128 fNormalizedPerpetualWeight)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchUpdateFundAllocationWeight(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerUpdateFundAllocationWeight, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "UpdateFundAllocationWeight", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerUpdateFundAllocationWeight)
				if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateFundAllocationWeight", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateFundAllocationWeight is a log parse operation binding the contract event 0x0ee9fecde1957f72be053052fe8a6896acc39ab927d76e9a56e59b60aa3fa2c4.
//
// Solidity: event UpdateFundAllocationWeight(uint8 poolId, uint24 indexed perpetualId, int128 fNormalizedPerpetualWeight)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseUpdateFundAllocationWeight(log types.Log) (*IPerpetualManagerUpdateFundAllocationWeight, error) {
	event := new(IPerpetualManagerUpdateFundAllocationWeight)
	if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateFundAllocationWeight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerUpdateFundingRateIterator is returned from FilterUpdateFundingRate and is used to iterate over the raw logs and unpacked data for UpdateFundingRate events raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateFundingRateIterator struct {
	Event *IPerpetualManagerUpdateFundingRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerUpdateFundingRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerUpdateFundingRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerUpdateFundingRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerUpdateFundingRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerUpdateFundingRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerUpdateFundingRate represents a UpdateFundingRate event raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateFundingRate struct {
	PerpetualId  *big.Int
	FFundingRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateFundingRate is a free log retrieval operation binding the contract event 0x84d63146d80c7c8804535f7c8128604790876a5de37ded2ad94ac0c848f0d4f7.
//
// Solidity: event UpdateFundingRate(uint24 indexed perpetualId, int128 fFundingRate)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterUpdateFundingRate(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerUpdateFundingRateIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "UpdateFundingRate", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerUpdateFundingRateIterator{contract: _IPerpetualManager.contract, event: "UpdateFundingRate", logs: logs, sub: sub}, nil
}

// WatchUpdateFundingRate is a free log subscription operation binding the contract event 0x84d63146d80c7c8804535f7c8128604790876a5de37ded2ad94ac0c848f0d4f7.
//
// Solidity: event UpdateFundingRate(uint24 indexed perpetualId, int128 fFundingRate)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchUpdateFundingRate(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerUpdateFundingRate, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "UpdateFundingRate", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerUpdateFundingRate)
				if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateFundingRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateFundingRate is a log parse operation binding the contract event 0x84d63146d80c7c8804535f7c8128604790876a5de37ded2ad94ac0c848f0d4f7.
//
// Solidity: event UpdateFundingRate(uint24 indexed perpetualId, int128 fFundingRate)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseUpdateFundingRate(log types.Log) (*IPerpetualManagerUpdateFundingRate, error) {
	event := new(IPerpetualManagerUpdateFundingRate)
	if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateFundingRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerUpdateMarginAccountIterator is returned from FilterUpdateMarginAccount and is used to iterate over the raw logs and unpacked data for UpdateMarginAccount events raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateMarginAccountIterator struct {
	Event *IPerpetualManagerUpdateMarginAccount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerUpdateMarginAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerUpdateMarginAccount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerUpdateMarginAccount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerUpdateMarginAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerUpdateMarginAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerUpdateMarginAccount represents a UpdateMarginAccount event raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateMarginAccount struct {
	PerpetualId       *big.Int
	Trader            common.Address
	PositionId        [16]byte
	FPositionBC       *big.Int
	FCashCC           *big.Int
	FLockedInValueQC  *big.Int
	FFundingPaymentCC *big.Int
	FOpenInterestBC   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateMarginAccount is a free log retrieval operation binding the contract event 0x391b02442f61d09af65ce4ec6b400a3e3b2b2e0f0bbe3be2cdc4c0c1cc3ab24e.
//
// Solidity: event UpdateMarginAccount(uint24 indexed perpetualId, address indexed trader, bytes16 indexed positionId, int128 fPositionBC, int128 fCashCC, int128 fLockedInValueQC, int128 fFundingPaymentCC, int128 fOpenInterestBC)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterUpdateMarginAccount(opts *bind.FilterOpts, perpetualId []*big.Int, trader []common.Address, positionId [][16]byte) (*IPerpetualManagerUpdateMarginAccountIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "UpdateMarginAccount", perpetualIdRule, traderRule, positionIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerUpdateMarginAccountIterator{contract: _IPerpetualManager.contract, event: "UpdateMarginAccount", logs: logs, sub: sub}, nil
}

// WatchUpdateMarginAccount is a free log subscription operation binding the contract event 0x391b02442f61d09af65ce4ec6b400a3e3b2b2e0f0bbe3be2cdc4c0c1cc3ab24e.
//
// Solidity: event UpdateMarginAccount(uint24 indexed perpetualId, address indexed trader, bytes16 indexed positionId, int128 fPositionBC, int128 fCashCC, int128 fLockedInValueQC, int128 fFundingPaymentCC, int128 fOpenInterestBC)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchUpdateMarginAccount(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerUpdateMarginAccount, perpetualId []*big.Int, trader []common.Address, positionId [][16]byte) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "UpdateMarginAccount", perpetualIdRule, traderRule, positionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerUpdateMarginAccount)
				if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateMarginAccount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateMarginAccount is a log parse operation binding the contract event 0x391b02442f61d09af65ce4ec6b400a3e3b2b2e0f0bbe3be2cdc4c0c1cc3ab24e.
//
// Solidity: event UpdateMarginAccount(uint24 indexed perpetualId, address indexed trader, bytes16 indexed positionId, int128 fPositionBC, int128 fCashCC, int128 fLockedInValueQC, int128 fFundingPaymentCC, int128 fOpenInterestBC)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseUpdateMarginAccount(log types.Log) (*IPerpetualManagerUpdateMarginAccount, error) {
	event := new(IPerpetualManagerUpdateMarginAccount)
	if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateMarginAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerUpdateMarkPriceIterator is returned from FilterUpdateMarkPrice and is used to iterate over the raw logs and unpacked data for UpdateMarkPrice events raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateMarkPriceIterator struct {
	Event *IPerpetualManagerUpdateMarkPrice // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerUpdateMarkPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerUpdateMarkPrice)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerUpdateMarkPrice)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerUpdateMarkPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerUpdateMarkPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerUpdateMarkPrice represents a UpdateMarkPrice event raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateMarkPrice struct {
	PerpetualId       *big.Int
	FMidPricePremium  *big.Int
	FMarkPricePremium *big.Int
	FSpotIndexPrice   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateMarkPrice is a free log retrieval operation binding the contract event 0xb1f269c39daf5ddf28d2612824a142cca4974e4dff0ef59fedd401c2c197120b.
//
// Solidity: event UpdateMarkPrice(uint24 indexed perpetualId, int128 fMidPricePremium, int128 fMarkPricePremium, int128 fSpotIndexPrice)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterUpdateMarkPrice(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerUpdateMarkPriceIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "UpdateMarkPrice", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerUpdateMarkPriceIterator{contract: _IPerpetualManager.contract, event: "UpdateMarkPrice", logs: logs, sub: sub}, nil
}

// WatchUpdateMarkPrice is a free log subscription operation binding the contract event 0xb1f269c39daf5ddf28d2612824a142cca4974e4dff0ef59fedd401c2c197120b.
//
// Solidity: event UpdateMarkPrice(uint24 indexed perpetualId, int128 fMidPricePremium, int128 fMarkPricePremium, int128 fSpotIndexPrice)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchUpdateMarkPrice(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerUpdateMarkPrice, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "UpdateMarkPrice", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerUpdateMarkPrice)
				if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateMarkPrice", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateMarkPrice is a log parse operation binding the contract event 0xb1f269c39daf5ddf28d2612824a142cca4974e4dff0ef59fedd401c2c197120b.
//
// Solidity: event UpdateMarkPrice(uint24 indexed perpetualId, int128 fMidPricePremium, int128 fMarkPricePremium, int128 fSpotIndexPrice)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseUpdateMarkPrice(log types.Log) (*IPerpetualManagerUpdateMarkPrice, error) {
	event := new(IPerpetualManagerUpdateMarkPrice)
	if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateMarkPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerUpdateParticipationFundCashIterator is returned from FilterUpdateParticipationFundCash and is used to iterate over the raw logs and unpacked data for UpdateParticipationFundCash events raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateParticipationFundCashIterator struct {
	Event *IPerpetualManagerUpdateParticipationFundCash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerUpdateParticipationFundCashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerUpdateParticipationFundCash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerUpdateParticipationFundCash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerUpdateParticipationFundCashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerUpdateParticipationFundCashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerUpdateParticipationFundCash represents a UpdateParticipationFundCash event raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateParticipationFundCash struct {
	PoolId         uint8
	FDeltaAmountCC *big.Int
	FNewFundCash   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateParticipationFundCash is a free log retrieval operation binding the contract event 0x5d881a46aafa213356fa06f9888716d00d0bcfa989a269425c25df13f5581a40.
//
// Solidity: event UpdateParticipationFundCash(uint8 indexed poolId, int128 fDeltaAmountCC, int128 fNewFundCash)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterUpdateParticipationFundCash(opts *bind.FilterOpts, poolId []uint8) (*IPerpetualManagerUpdateParticipationFundCashIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "UpdateParticipationFundCash", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerUpdateParticipationFundCashIterator{contract: _IPerpetualManager.contract, event: "UpdateParticipationFundCash", logs: logs, sub: sub}, nil
}

// WatchUpdateParticipationFundCash is a free log subscription operation binding the contract event 0x5d881a46aafa213356fa06f9888716d00d0bcfa989a269425c25df13f5581a40.
//
// Solidity: event UpdateParticipationFundCash(uint8 indexed poolId, int128 fDeltaAmountCC, int128 fNewFundCash)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchUpdateParticipationFundCash(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerUpdateParticipationFundCash, poolId []uint8) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "UpdateParticipationFundCash", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerUpdateParticipationFundCash)
				if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateParticipationFundCash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateParticipationFundCash is a log parse operation binding the contract event 0x5d881a46aafa213356fa06f9888716d00d0bcfa989a269425c25df13f5581a40.
//
// Solidity: event UpdateParticipationFundCash(uint8 indexed poolId, int128 fDeltaAmountCC, int128 fNewFundCash)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseUpdateParticipationFundCash(log types.Log) (*IPerpetualManagerUpdateParticipationFundCash, error) {
	event := new(IPerpetualManagerUpdateParticipationFundCash)
	if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateParticipationFundCash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerUpdateReprTradeSizesIterator is returned from FilterUpdateReprTradeSizes and is used to iterate over the raw logs and unpacked data for UpdateReprTradeSizes events raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateReprTradeSizesIterator struct {
	Event *IPerpetualManagerUpdateReprTradeSizes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerUpdateReprTradeSizesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerUpdateReprTradeSizes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerUpdateReprTradeSizes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerUpdateReprTradeSizesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerUpdateReprTradeSizesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerUpdateReprTradeSizes represents a UpdateReprTradeSizes event raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateReprTradeSizes struct {
	PerpetualId                 *big.Int
	FCurrentTraderExposureEMA   *big.Int
	FCurrentAMMExposureEMAShort *big.Int
	FCurrentAMMExposureEMALong  *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterUpdateReprTradeSizes is a free log retrieval operation binding the contract event 0x5b19b87703a645e020f6ed2b9d71d1f5eeb5b5567aa1eb05fdd818505cc79d20.
//
// Solidity: event UpdateReprTradeSizes(uint24 indexed perpetualId, int128 fCurrentTraderExposureEMA, int128 fCurrentAMMExposureEMAShort, int128 fCurrentAMMExposureEMALong)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterUpdateReprTradeSizes(opts *bind.FilterOpts, perpetualId []*big.Int) (*IPerpetualManagerUpdateReprTradeSizesIterator, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "UpdateReprTradeSizes", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerUpdateReprTradeSizesIterator{contract: _IPerpetualManager.contract, event: "UpdateReprTradeSizes", logs: logs, sub: sub}, nil
}

// WatchUpdateReprTradeSizes is a free log subscription operation binding the contract event 0x5b19b87703a645e020f6ed2b9d71d1f5eeb5b5567aa1eb05fdd818505cc79d20.
//
// Solidity: event UpdateReprTradeSizes(uint24 indexed perpetualId, int128 fCurrentTraderExposureEMA, int128 fCurrentAMMExposureEMAShort, int128 fCurrentAMMExposureEMALong)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchUpdateReprTradeSizes(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerUpdateReprTradeSizes, perpetualId []*big.Int) (event.Subscription, error) {

	var perpetualIdRule []interface{}
	for _, perpetualIdItem := range perpetualId {
		perpetualIdRule = append(perpetualIdRule, perpetualIdItem)
	}

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "UpdateReprTradeSizes", perpetualIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerUpdateReprTradeSizes)
				if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateReprTradeSizes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateReprTradeSizes is a log parse operation binding the contract event 0x5b19b87703a645e020f6ed2b9d71d1f5eeb5b5567aa1eb05fdd818505cc79d20.
//
// Solidity: event UpdateReprTradeSizes(uint24 indexed perpetualId, int128 fCurrentTraderExposureEMA, int128 fCurrentAMMExposureEMAShort, int128 fCurrentAMMExposureEMALong)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseUpdateReprTradeSizes(log types.Log) (*IPerpetualManagerUpdateReprTradeSizes, error) {
	event := new(IPerpetualManagerUpdateReprTradeSizes)
	if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateReprTradeSizes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPerpetualManagerUpdateUnitAccumulatedFundingIterator is returned from FilterUpdateUnitAccumulatedFunding and is used to iterate over the raw logs and unpacked data for UpdateUnitAccumulatedFunding events raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateUnitAccumulatedFundingIterator struct {
	Event *IPerpetualManagerUpdateUnitAccumulatedFunding // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPerpetualManagerUpdateUnitAccumulatedFundingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPerpetualManagerUpdateUnitAccumulatedFunding)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPerpetualManagerUpdateUnitAccumulatedFunding)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPerpetualManagerUpdateUnitAccumulatedFundingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPerpetualManagerUpdateUnitAccumulatedFundingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPerpetualManagerUpdateUnitAccumulatedFunding represents a UpdateUnitAccumulatedFunding event raised by the IPerpetualManager contract.
type IPerpetualManagerUpdateUnitAccumulatedFunding struct {
	PerpetualId             *big.Int
	UnitAccumulativeFunding *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterUpdateUnitAccumulatedFunding is a free log retrieval operation binding the contract event 0x506271456eb07872643b2726800d02892abf6c47ea9213ef161282b1d0c99ef6.
//
// Solidity: event UpdateUnitAccumulatedFunding(uint24 perpetualId, int128 unitAccumulativeFunding)
func (_IPerpetualManager *IPerpetualManagerFilterer) FilterUpdateUnitAccumulatedFunding(opts *bind.FilterOpts) (*IPerpetualManagerUpdateUnitAccumulatedFundingIterator, error) {

	logs, sub, err := _IPerpetualManager.contract.FilterLogs(opts, "UpdateUnitAccumulatedFunding")
	if err != nil {
		return nil, err
	}
	return &IPerpetualManagerUpdateUnitAccumulatedFundingIterator{contract: _IPerpetualManager.contract, event: "UpdateUnitAccumulatedFunding", logs: logs, sub: sub}, nil
}

// WatchUpdateUnitAccumulatedFunding is a free log subscription operation binding the contract event 0x506271456eb07872643b2726800d02892abf6c47ea9213ef161282b1d0c99ef6.
//
// Solidity: event UpdateUnitAccumulatedFunding(uint24 perpetualId, int128 unitAccumulativeFunding)
func (_IPerpetualManager *IPerpetualManagerFilterer) WatchUpdateUnitAccumulatedFunding(opts *bind.WatchOpts, sink chan<- *IPerpetualManagerUpdateUnitAccumulatedFunding) (event.Subscription, error) {

	logs, sub, err := _IPerpetualManager.contract.WatchLogs(opts, "UpdateUnitAccumulatedFunding")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPerpetualManagerUpdateUnitAccumulatedFunding)
				if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateUnitAccumulatedFunding", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateUnitAccumulatedFunding is a log parse operation binding the contract event 0x506271456eb07872643b2726800d02892abf6c47ea9213ef161282b1d0c99ef6.
//
// Solidity: event UpdateUnitAccumulatedFunding(uint24 perpetualId, int128 unitAccumulativeFunding)
func (_IPerpetualManager *IPerpetualManagerFilterer) ParseUpdateUnitAccumulatedFunding(log types.Log) (*IPerpetualManagerUpdateUnitAccumulatedFunding, error) {
	event := new(IPerpetualManagerUpdateUnitAccumulatedFunding)
	if err := _IPerpetualManager.contract.UnpackLog(event, "UpdateUnitAccumulatedFunding", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
