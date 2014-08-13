package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
)

type Actor struct {
	Name string
	Age  int
}

func main() {
	actor := Actor{"Jose da Silva", 99}

	service := "localhost:1200"

	conn, err := net.Dial("tcp", service)
	if err != nil {
		fmt.Errorf("Error dialing to service")
		return
	}

	err = json.NewEncoder(conn).Encode(actor)
	if err != nil {
		fmt.Errorf("Error encoding JSON")
		return
	}

	var result []byte
	result, err = ioutil.ReadAll(conn)
	fmt.Println(string(result))
}
