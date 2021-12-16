package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

//TODO: create pool of bytes.Buffers which can be reused.
var pool = sync.Pool{
	New: func() interface{} {
		fmt.Println("Creating new buffer")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, val string) {
	b := pool.Get().(*bytes.Buffer)

	b.Reset()

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	w.Write(b.Bytes())

	pool.Put(b)
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
	log(os.Stdout, "debug-string3")
}
