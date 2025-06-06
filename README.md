# d8x-futures-go-sdk

Go SDK for D8X Perpetual Futures
Fetches config files from `https://github.com/D8-X/sync-hub`

# Read-Only (Public)

Create a read-only SDK instance for X Layer testnet (no gas is spent)
```
var sdkRo SdkRO
err := sdkRo.New("195")
```

err := sdkRo.New("421614",
		WithPriceFeedEndpoint("https://hermes.pyth.network"),
		WithRpcUrl("https://rpc.ankr.com/arbitrum")) //arbitrum sepolia

To use a custom blockchain node-RPC, add `sdkRo.New("195", WithRpcUrl("https://testrpc.xlayer.tech"))` 
to use a custom Pyth-Hermes server `sdkRo.New("testnet", WithPriceFeedEndpoint("https://hermes.pyth.network/"))`,
or you can add both.

The most relevant functions for the read-only SDK are as follows, ignoring the error return value for simplicity:
```
trader := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
broker := common.HexToAddress("0xB0CBeeC370Af6ca2ed541F6a2264bc95b991F6E1")
pr, err := sdkRo.GetPositionRisk("BTC-USDT-USDT", trader)
symbols := []string{"BTC-USD-BUSD", "BERA-USD-BUSD", "BTC-USD-BUSD", "XAU-USD-BUSD", "ETH-USD-BUSD"}
pRisk, err := sdkRo.GetPositionRisks(symbols, trader, &OptEndPoints{Rpc: client})
pRisk, err = sdkRo.GetPositionRiskAll(traderAddr, &OptEndPoints{Rpc: client})
perpState, err := sdkRo.QueryPerpetualState([]int32{100000, 100001, 200002})
px, err := sdkRo.QueryPerpetualPrices("BTC-USDT-USDT", []float{0.01}, nil, nil)
poolState, err := sdkRo.QueryPoolStates(nil)
oo, dgsts, err := sdkRo.QueryOpenOrders("BTC-USDT-USDT", trader, nil) //([]Order, []string, error)
oo, dgsts, err := sdkRo.QueryAllOpenOrders("BTC-USDT-USDT", nil) //([]Order, []string, error)
status, err := sdkRo.QueryOrderStatus("BTC-USDT-USDT", trader, "", nil) // (string, error)
m, err := sdkRo.QueryMaxTradeAmount("BTC-USDT-USDT", 0, true, nil) // (float64, error)
vol, err := sdkRo.QueryTraderVolume(1, trader, nil) //(float64, error)
fee, err := sdkRo.QueryExchangeFeeTbpsForTrader(1, trader, broker, nil) // (uint16, error)
minpos, err := sdkRo.GetMinimalPositionSize("BTC-USDT-USDT") //(float64, error)
marginaccounts, err := sdkRo.QueryMarginAccounts("BTC-USDT-USDT", addresses, nil)
amt, err := sdkRo.GetPoolShareTknBalance(1, common.HexToAddress(addr), nil)
px, err := sdkRo.GetPoolShareTknPrice([]int{1, 2, 3}, nil)

// see if any liquidatable account in given perpetual with id 100000
acc, err := sdkRo.QueryLiquidatableAccounts(100000, nil)

// see if any liquidatable account in any of the perpetuals given by pool with id 1
acc2, err := sdkRo.QueryLiquidatableAccountsInPool(1, nil)

```
# Read-Write (Private: wallet required)
Create a read-write SDK instance for testnet (gas and collateral tokens are spent)
```
err := sdk.New(pk, "196")
tx, err := sdk.ApproveTknSpending("ETH-USDC-USDC", nil)
```
where `pk` is the private-key (string) of the wallet that is trading (or a broker depending on the functions used). Alternatively,
RPC and Pyth-server can be added as for the read-only sdk: `sdk.New(pk, "195", WithPriceFeedEndpoint("https://hermes.mine.network/"))`. 
All functions of the read-only SDK can be executed also on the read-write sdk.

Add (amountCC positive) or remove (amountCC negative) collateral:
```
err := sdk.AddCollateral(symbol string, amountCC float64) 
```

Example 1: Create a new order with minimal parameters plus a limit price:
```
order := NewOrder("ETH-USDC-USDC", SIDE_SELL, ORDER_TYPE_LIMIT, 0.1, 10, &OrderOptions{LimitPrice: 2240})
```

Example 2: Post the order
```
orderId, txHash, err := sdk.PostOrder(order, nil)
```
Example 3: Query the order status
```
status, err := sdk.QueryOrderStatus("ETH-USDC-USDC", sdk.Wallets[0].Address, orderId)
```
Example 4: Query position risk
```
pr, err := sdk.GetPositionRisk("ETH-USDC-USDT", sdk.Wallets[0].Address)
```
Example 5: Cancel order

```
tx, err := sdk.CancelOrder("BTC-USDT-USDT", id, nil)
```
Example 6: Execute order

```
tx, err := sdk.ExecuteOrders("ETH-USDC-USDC", orderIds, nil)
```

Example 7: Liquidate (liquidatable) order

```
acc2, err := sdk.QueryLiquidatableAccountsInPool(1, nil)
tx, err := sdk.LiquidatePosition(acc2[0].PerpId, &acc2[0].addr[0], nil, nil)
```


### Dev
Generate the ABI for the network it should compile to:
`abigen --abi abi/IPerpetualManager.json --pkg contracts --type IPerpetualManager --out IPerpetualManager.go`

`abigen --abi abi/LimitOrderBook.json --pkg contracts --type LimitOrderBook --out LimitOrderBook.go`

`abigen --abi abi/LimitOrderBookFactory.json --pkg contracts --type LimitOrderBookFactory --out LimitOrderBookFactory.go`
`abigen --abi abi/LimitOrderBookBeacon.json --pkg contracts --type LimitOrderBookBeacon --out LimitOrderBookBeacon.go`
`abigen --abi abi/IPyth.json --pkg contracts --type IPyth --out IPyth.go`

`abigen --abi abi/OracleFactory.json --pkg contracts --type OracleFactory --out OracleFactory.go`
`abigen --abi abi/CompositeToken.json --pkg contracts --type CompositeToken --out CompositeToken.go`

`go build` -> Remove duplicated structs

