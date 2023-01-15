package hash

import "hash/fnv"

func GetHash64(word string) uint64 {
	h := fnv.New64()
	_, err := h.Write([]byte(word))
	if err != nil {
		panic(err)
	}
	return h.Sum64()
}
