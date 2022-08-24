package types

const (
	// ModuleName defines the module name
	ModuleName = "cfeenergybank"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_energybank"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	EnergyTokenKey      = "EnergyToken-value-"
	EnergyTokenCountKey = "EnergyToken-count-"
)

const (
	TokensHistoryKey      = "TokensHistory-value-"
	TokensHistoryCountKey = "TokensHistory-count-"
)
