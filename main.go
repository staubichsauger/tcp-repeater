package main

import (
	"io"
	"log"
	"net"
)

func main() {
	stop := false

	send, err := net.Listen("tcp", ":9991")
	if err != nil {
		log.Fatal(err)
	}
	outConn, err := send.Accept()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("accepted out connection")

	recv, err := net.Listen("tcp", ":9990")
	if err != nil {
		log.Fatal(err)
	}
	inConn, err := recv.Accept()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("accepted in connection")

	buf := make([]byte, 1, 1)
	for !stop {
		_, err = io.CopyBuffer(outConn, inConn, buf)
		if err != nil {
			stop = true
		}
	}

	inConn.Close()
	outConn.Close()
}