package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	numberNonBlankLinesFlag = 1
)

var (
	fatal = log.New(os.Stderr, "", 0)
)

func SimpleCat(reader io.Reader, writer *bufio.Writer) int {
	_, err := io.Copy(writer, reader)
	if err != nil {
		fatal.Fatalln(err)
	}
	writer.Flush()
	return 0
}

func Cat(reader io.Reader, writer *bufio.Writer, flags int) int {
	var line string
	var err error
	bufferedReader := bufio.NewReader(reader)
	nr := 0
	countNonBlank := flags & numberNonBlankLinesFlag
	for {
		line, err = bufferedReader.ReadString('\n')
		if err != nil {
			return 1
		}
		if countNonBlank == 1 && line == "\n" || line == "" {
			fmt.Fprint(writer, line)
		} else if countNonBlank == 1 {
			nr++
			fmt.Fprintf(writer, "%6d\t%s", nr, line)
		} else {
		}
	}
	return 0
}

func main() {
}
