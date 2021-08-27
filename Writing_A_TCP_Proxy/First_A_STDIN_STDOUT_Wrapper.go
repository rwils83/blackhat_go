package main

import (
	"fmt"
	"log"
	"os"
)
//Define an io.Reader to read from STDIN
type FooReader struct{}
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

//Define an io.Writer to write to STDOUT
type FooWriter struct{}
func (fooWriter *FooWriter) Write(b []byte) (int,error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	//Instantiate reader and writer
	var (
		reader FooReader
		writer FooWriter
	)
	//Create buffer to hold input output
	input := make([]byte, 4096)
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	//Use writer to write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)

}