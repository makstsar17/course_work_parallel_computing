package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"server/hashTable"
	"server/invertedIndex"
	"syscall"
	"time"
)

const (
	HOST    = "127.0.0.1"
	PORT    = "8000"
	NETWORK = "tcp"
)

func main() {
	var threads int
	var parallelExecution bool
	flag.IntVar(&threads, "thr", 10, "Number of threads")
	flag.BoolVar(&parallelExecution, "pe", true, "Execute parallel or consecutively")
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

	var ht *hashTable.HashTable

	if parallelExecution {
		ht = invertedIndex.IndexDocs(fileNames, threads)
	} else {
		ht = invertedIndex.ConsIndexDocs(fileNames)
	}
	fmt.Printf("[Time %.6fs] Threads number: %d\n", time.Since(tm).Seconds(), threads)

	listener, err := net.Listen(NETWORK, HOST+":"+PORT)
	if err != nil {
		log.Fatalln(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		err := listener.Close()
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Close server")
		os.Exit(1)
	}()

	clientId := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("Client %d: connected to server\n", clientId)
		go handleClientRequest(conn, clientId, ht)
		clientId += 1
	}
}

func handleClientRequest(conn net.Conn, clientId int, ht *hashTable.HashTable) {
	clientReader := bufio.NewReader(conn)
	defer conn.Close()
	for {
		clientRequest := make([]byte, 1024)
		n, err := clientReader.Read(clientRequest)

		if _, ok := err.(*net.OpError); ok || err == io.EOF {
			log.Printf("[Client %d] closed the connection\n", clientId)
			break
		} else if err != nil {
			log.Printf("[Client %d] Error: %v\n", clientId, err)
		}

		log.Printf("[Client %d] search word: %v\n", clientId, string(clientRequest[:n]))
		var response string
		res, err := ht.Get(string(clientRequest[:n]))
		if err != nil {
			response = err.Error() + "\n"
			if _, err := conn.Write([]byte(err.Error() + "\n")); err != nil {
				log.Printf("[Client %d] failed to respond to client: %v\n", clientId, err)
			}
		} else {
			response = fmt.Sprintf("|%-30s|%-10s|\n", "File", "Frequency")
			if _, err := conn.Write([]byte(response)); err != nil {
				log.Printf("[Client %d] failed to respond to client: %v\n", clientId, err)
			}
			for file, frequency := range res {
				response = fmt.Sprintf("|%-30s|%-10d|\n", file, frequency)
				if _, err := conn.Write([]byte(response)); err != nil {
					log.Printf("[Client %d] failed to respond to client: %v\n", clientId, err)
				}
			}
		}
		if _, err := conn.Write([]byte("\n")); err != nil {
			log.Printf("[Client %d] failed to respond to client: %v\n", clientId, err)
		}
	}
}
