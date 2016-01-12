package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const destSocket string = "45.79.95.217"

func getHTTPHeader() {
	r, err := http.Get("http://www.nfl-lives.com")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()
	for i, v := range r.Header {
		fmt.Printf("%s : %s\n", i, v)
	}
}

func main() {
	// conn, err := net.Dial("tcp", destSocket)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// bs := make([]byte, 1024)
	// n, _ := conn.Read(bs)
	// fmt.Println(n)
	for _, v := range os.Args {
		if v == "http" {
			getHTTPHeader()
		}
	}
	getHTTPHeader()
}
