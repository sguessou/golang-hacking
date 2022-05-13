package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	IP := "scanme.nmap.org"
	Port := "80"

	address := IP + ":" + Port
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[+] Connection established...", conn.RemoteAddr().String())

}
