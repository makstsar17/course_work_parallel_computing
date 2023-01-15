package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"server/invertedIndex"
	"time"
)

func main() {
	var threads int
	flag.IntVar(&threads, "thr", 10, "Number of threads")
	flag.Parse()

	files, err := ioutil.ReadDir("dataset")
	if err != nil {
		log.Fatal(err)
	}
	fileNames := make([]string, len(files))
	for i, f := range files {
		fileNames[i] = "dataset/" + f.Name()
	}

	tm := time.Now()

	ht := invertedIndex.IndexDocs(fileNames, threads)
	//ht := invertedIndex.ConIndexDocs(fileNames)

	fmt.Printf("[Time %.6fs] Threads number: %d\n", time.Since(tm).Seconds(), threads)

	res, err := ht.Get("the")
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range res {
		fmt.Printf("the: %v: %d\n", key, value)
	}
}
