package main

import (
	"io"
	"log"
	"os"
)
type FooReaderImp struct{}
func (fooReader *FooReaderImp) Read(b []byte) (int, error) {
	// fmt.Print("in > ") Is skipped due to using the io.Copy method, but it did save several lines of code
	return os.Stdin.Read(b)
}

//Define an io.Writer to write to STDOUT

type FooWriterImp struct{}
func (fooWriter *FooWriterImp) Write(b []byte) (int,error) {
	//fmt.Print("out> ") See FooReader for explanation
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReaderImp
		writer FooWriterImp
	)
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}