package main

import "fmt"

func main() {

	const name = "Onion"
	const port = 8000
	const host = "127.0.0.1"

	fmt.Printf("Starting %v on host:port %v:%v\n", name, host, port)
	server := Server{Name: name, Host: host, Port: port}
	server.start("Starting new server")
}
