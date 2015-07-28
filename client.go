package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type ChatClient struct {
	Name     string
	Incoming chan string
	Outgoing chan string
	Conn     net.Conn
	Quit     chan bool
}

func (c *ChatClient) Read(buffer []byte) bool {
	_, error := c.Conn.Read(buffer)
	if error != nil {
		c.Close()
		Log(error)
		return false
	}
	return true
}

// Close closes the client connection and send quit to channel
func (c *ChatClient) Close() {
	c.Quit <- true
	c.Conn.Close()
}

// ClientSender sends buffers to a clent
func ChatClientSender(client *ChatClient) {
	for {
		select {
		case buffer := <-client.Outgoing:
			//Log("ClientSender sending [", string(buffer), "] to Server")
			count := 0
			for i := 0; i < len(buffer); i++ {
				if buffer[i] == 0x00 {
					break
				}
				count++
			}
			client.Conn.Write([]byte(buffer)[0:count])
		case <-client.Quit:
			Log("Client ", client.Name, " quitting")
			client.Conn.Close()
			break
		}
	}
}
func ChatClientReader(client *ChatClient) {
	buffer := make([]byte, 2048)

	for client.Read(buffer) {
		Log(string(buffer))
		for i := 0; i < 2048; i++ {
			buffer[i] = 0x00
		}
	}
}

func ChatClientHandler(client *ChatClient) {
	// Send the name of the client to the server
	client.Conn.Write([]byte(client.Name))
	go ChatClientSender(client)
	go ChatClientReader(client)
}

func runClient(name string) {
	conn, err := net.Dial("tcp", "127.0.0.1:9988")
	checkError(err)
	defer conn.Close()

	myClient := &ChatClient{name, make(chan string), make(chan string), conn, make(chan bool)}
	go ChatClientHandler(myClient)

	//Start an input reader to type some things and stuff
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		myClient.Outgoing <- strings.TrimSpace(line)

		if err != nil {
			fmt.Println(err)
			conn.Close()
			break
		}
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}
