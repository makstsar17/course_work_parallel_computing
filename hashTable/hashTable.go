package hashTable

import (
	"server/hash"
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

func (ht *HashTable) rehashing() {
	ht.capacity *= 2
	newBuckets := make([]*linkedList.LinkedList, ht.capacity)
	for i := range newBuckets {
		newBuckets[i] = &linkedList.LinkedList{Head: nil, Length: 0}
	}

	for _, bucket := range ht.buckets {
		if bucket.Head == nil {
			continue
		}
		node := bucket.Head
		for true {
			hashValue := hash.GetHash64(node.Word)
			keyId := hashValue % uint64(ht.capacity)

			newBuckets[keyId].InsertNode(&linkedList.Node{Word: node.Word, DocIds: node.DocIds})

			if node.Next == nil {
				break
			}
			node = node.Next
		}
	}
	ht.buckets = newBuckets
}

func (ht *HashTable) Insert(key string, docId string) {
	hashValue := hash.GetHash64(key)
	keyId := hashValue % uint64(ht.capacity)

	ht.mu.Lock()
	lenLinkList := ht.buckets[keyId].Length
	ht.buckets[keyId].Insert(key, docId)

	if lenLinkList != ht.buckets[keyId].Length {
		ht.length += 1
	}

	loadFactor := float64(ht.length) / float64(ht.capacity)
	if loadFactor > 0.75 {
		ht.rehashing()
	}
	ht.mu.Unlock()
}
