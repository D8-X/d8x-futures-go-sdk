package d8x_futures

import "testing"

func TestQueryPoolStaticInfo(t *testing.T) {
	config, err := loadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	proxy := CreatePerpetualManagerInstance(config)
	queryPoolStaticInfo(proxy, config)
}
