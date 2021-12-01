package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// TODO: connect to server on localhost port 8000
	con, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	mustCopy(os.Stdout, con)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
