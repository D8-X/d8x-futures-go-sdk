# d8x-futures-go-sdk

Go SDK for D8X Perpetual Futures

# Read-Only

Create a read-only SDK instance for zkEVM testnet (no gas is spent)
```
var sdkRo SdkRO
err := sdkRo.New("testnet")
```
To use a custom blockchain node-RPC, add an argument `sdkRo.New("testnet", "https://polygon-zkevm-testnet.blockpi.network/v1/rpc/public")` 
to use a custom Pyth-Hermes server  `sdkRo.New("testnet", "", "https://https://hermes-beta.pyth.network/")`, keeping the second argument "" 
(or adding a custom RPC).

The most relevant functions for the read-only SDK are as follows, ignoring the error return value for simplicity:
```
trader := common.HexToAddress("0x9d5aaB428e98678d0E645ea4AeBd25f744341a05")
broker := common.HexToAddress("0xB0CBeeC370Af6ca2ed541F6a2264bc95b991F6E1")
pr, err := sdkRo.GetPositionRisk("BTC-USD-MATIC", trader)
perpState, err := sdkRo.QueryPerpetualState([]int32{100000, 100001, 200002})
poolState, err := sdkRo.QueryPoolStates()
oo, dgsts, err := sdkRo.QueryOpenOrders("BTC-USD-MATIC", trader) //([]Order, []string, error)
oo, dgsts, err := sdkRo.QueryAllOpenOrders("BTC-USDC-USDC") //([]Order, []string, error)
status, err := sdkRo.QueryOrderStatus("BTC-USD-MATIC", trader, "") // (string, error)
m, err := sdkRo.QueryMaxTradeAmount("BTC-USD-MATIC", 0, true) // (float64, error) {
vol, err := sdkRo.QueryTraderVolume(1, trader) //(float64, error) {
fee, err := sdkRo.QueryExchangeFeeTbpsForTrader(1, trader, broker) // (uint16, error) {
minpos, err := sdkRo.GetMinimalPositionSize("BTC-USD-MATIC") //(float64, error)
```
# Read-Write
Create a read-write SDK instance for zkEVM testnet (gas and collateral tokens are spent)
```
err := sdk.New(pk, "testnet")
tx, err := sdk.ApproveTknSpending("ETH-USD-MATIC", nil)
```
where `pk` is the private-key (string) of the wallet that is trading (or a broker depending on the functions used). Alternatively,
RPC and Pyth-server can be added as for the read-only sdk: `sdk.New(pk, "testnet", "", "https://mypythendpoint.com/api")`. All functions of the read-only SDK can be executed also on the read-write sdk.

Add (amountCC positive) or remove (amountCC negative) collateral:
```
err := sdk.AddCollateral(symbol string, amountCC float64) 
```

Example 1: Create a new order with minimal parameters plus a limit price:
```
order := NewOrder("ETH-USD-MATIC", SIDE_SELL, ORDER_TYPE_LIMIT, 0.1, 10, &OrderOptions{LimitPrice: 2240})
```

Example 2: Post the order
```
orderId, txHash, err := sdk.PostOrder(order)
```
Example 3: Query the order status
```
status, err := sdk.QueryOrderStatus("ETH-USD-MATIC", sdk.Wallet.Address, orderId)
```
Example 4: Query position risk
```
pr, err := sdk.GetPositionRisk("ETH-USD-MATIC", sdk.Wallet.Address)
```
Example 5: Cancel order

```
tx, err := sdk.CancelOrder("BTC-USDC-USDC", id)
```


### Dev
Generate the ABI for the network it should compile to:
`abigen --abi abi/IPerpetualManager.json --pkg contracts --type IPerpetualManager --out IPerpetualManager.go`

`abigen --abi abi/LimitOrderBook.json --pkg contracts --type LimitOrderBook --out LimitOrderBook.go`

`abigen --abi abi/LimitOrderBookFactory.json --pkg contracts --type LimitOrderBookFactory --out LimitOrderBookFactory.go`

`go build` -> Remove duplicated structs

