package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type SimpleReader struct{}

func (simpleReader *SimpleReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type SimpleWriter struct{}

func (simpleWriter *SimpleWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader SimpleReader
		writer SimpleWriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalf("unable to read/write data")
	}
}
