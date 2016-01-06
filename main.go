package main

import (
	"fmt"
	"log"
	"net"
)

const destSocket string = "45.79.95.217"

func main() {
	conn, err := net.Dial("tcp", destSocket)
	if err != nil {
		log.Fatalln(err)
	}
	bs := make([]byte, 1024)
	n, _ := conn.Read(bs)
	fmt.Println(n)
}
