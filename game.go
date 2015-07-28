package main

import "flag"

var cardServer bool

func init() {
	flag.BoolVar(&cardServer, "server", false, "Run the cardgame server.")

}
func main() {
	flag.Parse()
	if cardServer {
		runServer()
	} else {
		runClient(clientName, serverHost)
	}
}
