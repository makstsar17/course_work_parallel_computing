package invertedIndex

import (
	"io/ioutil"
	"log"
	"regexp"
	"server/hashTable"
	"sync"
)

func mappingFile(fileQueue <-chan string, ht *hashTable.HashTable) {
	for file := range fileQueue {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		validWord := regexp.MustCompile(`[a-zA-Z-']{2,}`)
		for _, word := range validWord.FindAllString(string(content), -1) {
			ht.Insert(word, file)
		}
	}
}

func IndexDocs(files []string, thr int) *hashTable.HashTable {
	var wg sync.WaitGroup
	wg.Add(thr)

	fileQueue := make(chan string, len(files))
	ht := hashTable.GetHashTable()
	for i := 0; i < thr; i++ {
		go func() {
			mappingFile(fileQueue, ht)
			wg.Done()
		}()
	}

	for _, file := range files {
		fileQueue <- file
	}

	close(fileQueue)
	wg.Wait()

	return ht
}
