package types

const (
	// ModuleName defines the module name
	ModuleName = "intravision"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_intravision"
)

var (
	ParamsKey = []byte("p_intravision")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
