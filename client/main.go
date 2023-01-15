package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	HOST    = "127.0.0.1"
	PORT    = "8000"
	NETWORK = "tcp"
)

func main() {
	con, err := net.Dial(NETWORK, HOST+":"+PORT)
	if err != nil {
		log.Fatalln(err)
	}
	defer con.Close()

	clientReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(con)

	fmt.Println("___Inverted Index___")
	fmt.Println("enter '-e' to exit")

	for {
		fmt.Print("Enter word: ")
		clientRequest, err := clientReader.ReadString('\n')
		if err != nil {
			log.Printf("Client error: %v\n", err)
			continue
		}

		clientRequest = strings.TrimSpace(clientRequest)

		if clientRequest == "-e" {
			log.Println("Close connection")
			break
		}

		if len(strings.Fields(clientRequest)) != 1 {
			log.Println("Please, enter only one word...")
		}

		if _, err = con.Write([]byte(clientRequest)); err != nil {
			if _, ok := err.(*net.OpError); ok {
				log.Println("Server is closed")
				break
			}
			log.Printf("Failed to send the client request: %v\n Please try again...", err)
			continue
		}

		for {
			receivedStr, err := serverReader.ReadString('\n')
			if err != nil {
				log.Fatalln(err)
			}
			if len(receivedStr) == 1 {
				break
			}
			fmt.Print(receivedStr)
		}
	}
}
