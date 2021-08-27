package main

import (
	"flag"
	"io"
	"log"
	"net"
)



func handler (src net.Conn, outURL string) {

	dst, err := net.Dial("tcp", outURL)
	if err != nil {
		log.Fatalln("Unable to connect or unreachable host")
	}
	defer dst.Close()

	go func(){
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}
func main() {
	var send string
	var inurl string
	flag.StringVar(&inurl, "listen", "", "Enter the URL to listen for")
	flag.StringVar(&send, "send", "", "Enter the url to send to")
	flag.Parse()
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handler(conn, send)
	}
}