package main

import (
	"io"
	"log"
	"net"
)

// echo is a handler function that echoes recieved data
func echo(conn net.Conn) {
	defer conn.Close()

	// create storage buffer. Known as a slice. Important: The size listed is the minimum size of memory taken, it will grow with input
	b := make([]byte, 512)
	for {
		//receive data via conn.Read to buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client Disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected Error.")
			break
		}
		log.Printf("Recieved %d bytes: %s\n", size, string(b))
		// Send data back with conn.Write
		log.Println("Writing Data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	// bind to TCP 20080 on all IFs
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to Port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		// wait for connection and create net.Conn on connection
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo(conn)
	}
}