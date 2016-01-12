package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

const (
	ip = "45.79.93.134"
)

func getBanner(socket string) net.Conn {
	conn, err := net.Dial("tcp", socket)
	if err != nil {
		fmt.Printf("Error Connecting %s\n", err)
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return conn
}
func main() {
	socket := ip + ":21"
	conn := getBanner(socket)
	w := bufio.NewWriter(conn)
	fmt.Printf("%v", io.Copy(w, conn))

}
