# d8x-futures-go-sdk

Minimal Go SDK for D8X Perpetual Futures

Generate the ABI for the network it should compile to:
`abigen --abi abi/IPerpetualManager.json --pkg contracts --type IPerpetualManager --out IPerpetualManager.go`

`abigen --abi abi/LimitOrderBook.json --pkg contracts --type LimitOrderBook --out LimitOrderBook.go`

`abigen --abi abi/LimitOrderBookFactory.json --pkg contracts --type LimitOrderBookFactory --out LimitOrderBookFactory.go`

`go build` -> Remove duplicated structs

