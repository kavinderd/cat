package main

import (
	"bufio"
	"io"
)

const ()

var (
	HorizTab = []byte("^I")
)

func Cat(reader io.Reader, buf []byte, writer *bufio.Writer) int {
	newlines := 0
	endOfBuffer := 0                     // end of buffer
	beginningOfBuffer := endOfBuffer + 1 //beginning of buffer
	ch := byte(0)                        // char in buffer
	size := len(buf) - 1                 // len of buffer with room for sentinel byte

	for {

		//For Loop for handling newline char
		for {
			if beginningOfBuffer > endOfBuffer {
				n, err := reader.Read(buf[:size])
				if err == io.EOF {
					//				totalNewLine = newlines
					writer.Flush()
					return 0
				}
				if err != nil {
					//				totalNewLine = newlines
					writer.Flush()
					return 1
				}

				beginningOfBuffer = 0 //Reset bpin to the beginning of the buffer
				endOfBuffer = n       //End of buffer is the number of bytes read
				buf[endOfBuffer] = 10 //Place a sentinel at the end of the buffer
			} else {
				newlines++
				//TODO: Logic for flags
			}

			ch = buf[beginningOfBuffer]
			beginningOfBuffer++
			if ch != 10 {
				break
			}
		}

		for {
			if ch != 10 {
				writer.WriteByte(ch)
			} else {
				writer.WriteByte(10)
				newlines = -1
				break
			}

			ch = buf[beginningOfBuffer]
			beginningOfBuffer++
		}
	}
}

func main() {
}
