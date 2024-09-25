package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	var target *string
	target = flag.String("target", "127.0.0.1:8080", "目标地址")
	flag.Parse()

	conn, err := net.Dial("tcp", *target)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Connected to", *target)
}
