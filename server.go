package main

import (
	"bytes"
	"container/list"
	"fmt"
	"net"
)

// Client is the struct for each client
type Client struct {
	Name       string
	Incoming   chan string
	Outgoing   chan string
	Conn       net.Conn
	Quit       chan bool
	ClientList *list.List
}

// Log ... is a loging function
func Log(v ...interface{}) {
	fmt.Println(v...)
}
func (c *Client) Read(buffer []byte) bool {
	bytesRead, error := c.Conn.Read(buffer)
	if error != nil {
		c.Close()
		Log(error)
		return false
	}
	Log("Read ", bytesRead, " bytes")
	return true
}

// Close closes the client connection and send quit to channel
func (c *Client) Close() {
	c.Quit <- true
	c.Conn.Close()
	c.RemoveMe()
}

// Equal is a compare on connections
func (c *Client) Equal(other *Client) bool {
	if bytes.Equal([]byte(c.Name), []byte(other.Name)) {
		if c.Conn == other.Conn {
			return true
		}
	}
	return false
}

// RemoveMe removes a client from a list of clients
func (c *Client) RemoveMe() {
	for entry := c.ClientList.Front(); entry != nil; entry = entry.Next() {
		client := entry.Value.(Client)
		if c.Equal(&client) {
			Log("RemoveMe: ", c.Name)
			c.ClientList.Remove(entry)
		}
	}
}

func ClientReader(client *Client) {
	buffer := make([]byte, 2048)

	for client.Read(buffer) {
		if bytes.Equal(buffer, []byte("/quit")) {
			client.Close()
			break
		}
		Log("ClientReader received ", client.Name, "> ", string(buffer))
		send := client.Name + "> " + string(buffer)
		client.Outgoing <- send
		for i := 0; i < 2048; i++ {
			buffer[i] = 0x00
		}
	}

	client.Outgoing <- client.Name + " has left chat"
	Log("ClientReader stopped for ", client.Name)
}

// ClientSender sends buffers to a clent
func ClientSender(client *Client) {
	for {
		select {
		case buffer := <-client.Incoming:
			Log("ClientSender sending [", string(buffer), "] to ", client.Name)
			count := 0
			for i := 0; i < len(buffer); i++ {
				if buffer[i] == 0x00 {
					break
				}
				count++
			}
			Log("Send size: ", count)
			client.Conn.Write([]byte(buffer)[0:count])
		case <-client.Quit:
			Log("Client ", client.Name, " quitting")
			client.Conn.Close()
			break
		}
	}
}
func ClientHandler(conn net.Conn, ch chan string, clientList *list.List) {
	buffer := make([]byte, 1024)
	bytesRead, error := conn.Read(buffer)
	if error != nil {
		Log("Client connection error: ", error)
	}

	name := string(buffer[0:bytesRead])
	newClient := &Client{name, make(chan string), ch, conn, make(chan bool), clientList}

	go ClientSender(newClient)
	go ClientReader(newClient)
	clientList.PushBack(*newClient)
	ch <- string(name + " has joined the chat")
}

func IOHandler(Incoming <-chan string, clientList *list.List) {
	for {
		Log("IOHandler: Waiting for input")
		input := <-Incoming
		Log("IOHandler: Handling ", input)
		for e := clientList.Front(); e != nil; e = e.Next() {
			client := e.Value.(Client)
			client.Incoming <- input
		}
	}
}

func runServer() {
	Log("Hello Server!")

	clientList := list.New()
	in := make(chan string)
	go IOHandler(in, clientList)

	service := ":9988"
	tcpAddr, error := net.ResolveTCPAddr("tcp", service)
	if error != nil {
		Log("Error: Could not resolve address")
	} else {
		netListen, error := net.Listen(tcpAddr.Network(), tcpAddr.String())
		if error != nil {
			Log(error)
		} else {
			defer netListen.Close()

			for {
				Log("Waiting for clients")
				connection, error := netListen.Accept()
				if error != nil {
					Log("Client error: ", error)
				} else {
					go ClientHandler(connection, in, clientList)
				}
			}
		}
	}
}
