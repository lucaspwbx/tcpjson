package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Actor struct {
	Name string
	Age  int
}

func main() {
	service := ":1200"
	addr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		fmt.Errorf("Error resolving tcp addr")
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Errorf("Error to get listener")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		var actor Actor
		err = json.NewDecoder(conn).Decode(&actor)
		if err != nil {
			fmt.Println("Server: Error decoding JSON")
			return
		}
		fmt.Println(actor)
		fmt.Println("Actor name: ", actor.Name)
		response := fmt.Sprintf("Received data from: %s", actor.Name)
		conn.Write([]byte(response))
		conn.Close()
	}
}
