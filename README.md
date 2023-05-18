# d8x-futures-go-sdk
Go SDK for D8X Perpetual Futures

Generate the ABI for the network it should compile to:
`abigen --abi abi/testnet/v7/IPerpetualManager.json --pkg d8x_futures --type IPerpetualManager --out IPerpetualManager.go`

`abigen --abi abi/testnet/v7/LimitOrderBook.json --pkg d8x_futures --type LimitOrderBook --out LimitOrderBook.go`

` abigen --abi abi/testnet/v7/LimitOrderBookFactory.json --pkg d8x_futures --type LimitOrderBookFactory --out LimitOrderBookFactory.go`

Remove duplicated structs

https://www.digitalocean.com/community/tutorials/how-to-use-a-private-go-module-in-your-own-project
