package d8x_futures

import (
	"fmt"
	"testing"
)

func TestQueryNestedPerpetualInfo(t *testing.T) {
	config, err := loadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
	p := QueryNestedPerpetualInfo(conn)
	fmt.Println(p.PerpetualIds)
}

func TestReadSymbolList(t *testing.T) {
	symMap, err := readSymbolList()
	if err != nil {
		t.Logf(err.Error())
	}
	fmt.Println((*symMap)["MATC"])

}

func TestQueryPoolStaticInfo(t *testing.T) {
	config, err := loadConfig("testnet")
	if err != nil {
		t.Logf(err.Error())
	}
	conn := CreateBlockChainConnector(config)
	nest := QueryNestedPerpetualInfo(conn)
	info := QueryPoolStaticInfo(conn, nest)
	fmt.Println(info)
}
