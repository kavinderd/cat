package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

var (
	fatal = log.New(os.Stderr, "", 0)
)

func Cat(reader io.Reader, buf []byte, writer *bufio.Writer) int {
	_, err := io.Copy(writer, reader)
	if err != nil {
		fatal.Fatalln(err)
	}
	writer.Flush()
	return 0
}

func main() {
}
