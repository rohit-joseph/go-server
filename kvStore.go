package server

// ValVerPair stored as a struct
type ValVerPair struct {
	value   []byte
	version int32
}

var kvStore map[string]ValVerPair

func init() {
	kvStore = make(map[string]ValVerPair)
}

// PutKVP puts a value into the store
func PutKVP(key []byte, valVer ValVerPair) int {
	k := string(key)
	kvStore[k] = valVer
	return 0
}

// GetKVP fetchs a value/version pair from the store,
// Returns nil if there is no such key
func GetKVP(key []byte) ValVerPair {
	k := string(key)
	v := kvStore[k]
	return v
}

// RemoveKVP deletes a KVPair from the store
func RemoveKVP(key []byte) {
	k := string(key)
	delete(kvStore, k)
}

// ClearKVP removes all elemets from the store
func ClearKVStore() {
	kvStore = nil
	kvStore = make(map[string]ValVerPair)
}
