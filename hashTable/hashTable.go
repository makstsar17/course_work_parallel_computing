package hashTable

import (
	"server/linkedList"
	"sync"
)

type Bucket struct {
	Key   string
	Value interface{}
}

type HashTable struct {
	buckets  []*linkedList.LinkedList
	capacity int
	length   int
	mu       sync.Mutex
}

func GetHashTable() *HashTable {
	ht := HashTable{
		buckets:  make([]*linkedList.LinkedList, 16),
		capacity: 16,
		length:   0,
	}
	for i := range ht.buckets {
		ht.buckets[i] = &linkedList.LinkedList{Head: nil, Length: 0}
	}
	return &ht
}
