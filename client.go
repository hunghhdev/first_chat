package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Your name:")
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')

	nameInput = nameInput[:len(nameInput) - 1]

	fmt.Println("*********** Messages *************")
	go _onMessage(connection)
	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, _ := msgReader.ReadString('\n')
		if err != nil {
			break
		}

		msg = fmt.Sprintf("%s: %s\n", nameInput,
			msg[:len(msg) - 1])

		connection.Write([]byte(msg))
	}
	connection.Close()
}

func _onMessage(conn net.Conn)  {
	for {
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')
		fmt.Print(msg)
	}
}
