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
