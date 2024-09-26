package stork

const (
	MAX_STORK_FEED_AGE_SEC int64 = 60
)

// Signature holds the signature details (r, s, v)
type Signature struct {
	R string `json:"r"`
	S string `json:"s"`
	V string `json:"v"`
}

// TimestampedSignature holds the signature with its timestamp and message hash
type TimestampedSignature struct {
	Signature Signature `json:"signature"`
	Timestamp int64     `json:"timestamp"`
	MsgHash   string    `json:"msg_hash"`
}

// StorkSignedPrice holds the signed price details
type StorkSignedPrice struct {
	PublicKey       string               `json:"public_key"`
	EncodedAssetID  string               `json:"encoded_asset_id"`
	Price           string               `json:"price"`
	TimestampedSig  TimestampedSignature `json:"timestamped_signature"`
	PublisherMerkle string               `json:"publisher_merkle_root"`
	CalculationAlg  CalculationAlg       `json:"calculation_alg"`
}

// CalculationAlg holds the calculation algorithm details
type CalculationAlg struct {
	Type     string `json:"type"`
	Version  string `json:"version"`
	Checksum string `json:"checksum"`
}

// SignedPrice holds details of an individual signed price
type SignedPrice struct {
	PublisherKey    string               `json:"publisher_key"`
	ExternalAssetID string               `json:"external_asset_id"`
	SignatureType   string               `json:"signature_type"`
	Price           string               `json:"price"`
	TimestampedSig  TimestampedSignature `json:"timestamped_signature"`
}

// StorkPrice contains the price data for a specific asset
type StorkPrice struct {
	Timestamp        int64            `json:"timestamp"`
	AssetID          string           `json:"asset_id"`
	SignatureType    string           `json:"signature_type"`
	Trigger          string           `json:"trigger"`
	Price            string           `json:"price"`
	StorkSignedPrice StorkSignedPrice `json:"stork_signed_price"`
	SignedPrices     []SignedPrice    `json:"signed_prices"`
}

// ResponseData holds the overall response structure
type StorkResponseData struct {
	Data map[string]StorkPrice `json:"data"`
}
