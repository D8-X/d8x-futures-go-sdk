package d8x_futures

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"
)

const (
	BETTING_LIFECYCLE_URL = "https://sports.quantena.org/"
)

type SlotAssignment struct {
	ContractID string    `json:"contract_id"`
	YYMMDD     string    `json:"yymmdd"`
	Expiry     time.Time `json:"expiry"`
	ChainID    int       `json:"chain_id"`
	PoolSym    string    `json:"pool_sym"`
	Slot       string    `json:"slot"`
}

// SportsPrefix returns the 'slots' for sports perpetuals.
// CFB, MLB, NFL, NHL are legacy, SP is generic
func SportsPrefix() []string {
	return []string{"SP0", "SP1", "SP2", "CFB", "MLB", "NFL", "NHL"}
}

func IsSportsSymbol(sym string) bool {
	return len(sym) > 4 && sym[3] == '_'
}

// SportSlotAssignment takes a slot (SP01, MLB0,...) and
// returns the perpetual of the assigned slot and a boolean
// whether there is an assignment.
// For regular contracts (e.g. BTC) the function returns the
// input and TRUE
func (sdk *SdkRO) SportSlotAssignment(slot string) (string, bool) {
	if len(slot) < 4 {
		return slot, true
	}
	prfx := slot[0:3]
	sportP := SportsPrefix()
	if !slices.Contains(sportP, prfx) {
		return slot, true
	}
	sym, err := sdk.internalToSymbol(slot)
	if err != nil {
		return "", false
	}
	return sym, true
}

// InternalToSymbol is the inversion to SymbolToInternal
// Converts an internal symbol of the form MLB0 to its current
// contract id -USD-<collateral>
func (sdk *SdkRO) internalToSymbol(symInt string) (string, error) {
	league := symInt[0:3]
	leagues := SportsPrefix()
	if !slices.Contains(leagues, league) {
		return symInt, nil
	}
	now := time.Now().Unix()
	sdk.Sport.SlotsMux.RLock()
	ts := sdk.Sport.LastUpdateTs
	symInt, _ = strings.CutSuffix(symInt, "-USD")
	contractId, exists := sdk.Sport.SlotnameToCtrct[symInt]
	slot := sdk.Sport.Slots[contractId]
	sdk.Sport.SlotsMux.RUnlock()
	if (!exists && now-ts > 5*60) || slot.Expiry.Before(time.Now()) {
		if err := sdk.refreshSlotAssignment(); err != nil {
			return "", fmt.Errorf("unable to refresh slot assignment: %w", err)
		}
		sdk.Sport.SlotsMux.RLock()
		contractId, exists = sdk.Sport.SlotnameToCtrct[symInt]
		sdk.Sport.SlotsMux.RUnlock()
	}
	if !exists {
		return "", fmt.Errorf("no such symbol conversion: %s", symInt)
	}
	return contractId + "-USD-" + slot.PoolSym, nil
}

// SportSlot checks whether the contractId (NFL_ATL_SF_251019)
// is assigned to a perpetual slot. Returns the slot name
// (e.g. NFL0-USD-PUSD) and true if found, "" and false otherwise
func (sdk *SdkRO) SportSlot(contractId string) (string, bool) {
	if contractId[3] != '_' {
		return "", false
	}
	sym, err := sdk.symbolToInternal(contractId)
	if err != nil {
		return "", false
	}
	return strings.Split(sym, "-")[0], true
}

// symbolToInternal converts sports symbols
// to the currently assigned perpetual symbol
// E.g., MLB_TOR_NYY_251009 -> MLB1-USD-<poolsym>
// Returns the symbol unchanged if not a long-form
// sports symbol
func (sdk *SdkRO) symbolToInternal(sym string) (string, error) {
	if sym[3] != '_' {
		return sym, nil
	}
	sym = strings.Split(sym, "-")[0]
	sdk.Sport.SlotsMux.RLock()
	slot, exists := sdk.Sport.Slots[sym]
	ts := sdk.Sport.LastUpdateTs
	sdk.Sport.SlotsMux.RUnlock()

	now := time.Now().Unix()
	if (!exists && now-ts > 5*60) || slot.Expiry.Before(time.Now()) {
		if err := sdk.refreshSlotAssignment(); err != nil {
			return "", fmt.Errorf("unable to refresh slot assignment: %w", err)
		}
		sdk.Sport.SlotsMux.RLock()
		slot, exists = sdk.Sport.Slots[sym]
		sdk.Sport.SlotsMux.RUnlock()
	}
	if !exists {
		return "", fmt.Errorf("no such symbol %s", sym)
	}
	return slot.Slot + "-USD-" + slot.PoolSym, nil
}

// refreshSlotAssignment updates sdk.Sport.Slots from
// the sport api
func (sdk *SdkRO) refreshSlotAssignment() error {
	info, err := fetchSlotsInfo(int(sdk.ChainConfig.ChainId))
	if err != nil {
		return err
	}
	assign := make(map[string]SlotAssignment)
	slotName := make(map[string]string)
	for _, slot := range info {
		assign[slot.ContractID] = slot
		slotName[slot.Slot] = slot.ContractID
	}
	sdk.Sport.SlotsMux.Lock()
	defer sdk.Sport.SlotsMux.Unlock()
	sdk.Sport.LastUpdateTs = time.Now().Unix()
	sdk.Sport.Slots = assign
	sdk.Sport.SlotnameToCtrct = slotName
	return nil
}

// fetchSlotsInfo queries the betting-lifecycle API to get
// current slot assignments
func fetchSlotsInfo(chainId int) ([]SlotAssignment, error) {
	url := fmt.Sprintf("%s%s/%d", BETTING_LIFECYCLE_URL, "slots-info", chainId)
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var slots []SlotAssignment
	if err := json.NewDecoder(resp.Body).Decode(&slots); err != nil {
		return nil, fmt.Errorf("decode failed: %w", err)
	}
	return slots, nil
}
