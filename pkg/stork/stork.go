package stork

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/D8-X/d8x-futures-go-sdk/utils"
)

type Stork struct {
	Endpoint string // stork endpoint
	Token    string // access token
}

func NewStork(endpoint, credentials string) *Stork {
	token := base64.StdEncoding.EncodeToString([]byte(credentials))
	endpoint = strings.TrimSuffix(endpoint, "/")
	s := Stork{
		Endpoint: endpoint,
		// Encode the credentials string to base64
		Token: token,
	}
	return &s
}

func (strk *Stork) RestFetchStorkPrice(storkSym string) (float64, error) {
	url := strk.Endpoint + "/v1/prices/latest?assets=" + storkSym
	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("RestFetchStorkPrice request formation: %v", err)
	}

	// Set the Authorization header
	tokenCode := fmt.Sprintf("Basic %s", strk.Token)
	req.Header.Set("Authorization", tokenCode)

	// Create a new HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("RestFetchPrice request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response body: %v", err)
	}
	// Unmarshal the JSON data
	var response StorkResponseData
	err = json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		return 0, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	ts := response.Data[storkSym].Timestamp
	nowTs := time.Now().Unix()
	if nowTs-ts > MAX_STORK_FEED_AGE_SEC {
		return 0, fmt.Errorf("stork feed too old (%d seconds)", nowTs-ts)
	}
	price, err := storkPriceToFloat(response.Data[storkSym].Price)
	if err != nil {
		return 0, err
	}
	return price, nil
}

func storkPriceToFloat(price string) (float64, error) {
	dec18, ok := new(big.Int).SetString(price, 10)
	if !ok {
		return 0, fmt.Errorf("could not convert stork price %s to float", price)
	}
	return utils.DecNToFloat(dec18, 18), nil
}
