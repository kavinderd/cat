package main

import (
	"bufio"
	"io"
)

const (
	ShowTabs           = 1
	ShowAllLineNumbers = 2
)

var (
	HorizTab = []byte("^I")

	LineLen = 20
	LineBuf = []byte{
		' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', ' ', ' ',
		' ', ' ', ' ', '0', '\t',
	}
	LinePrint = LineLen - 7
	LineStart = LineLen - 2
	LineEnd   = LineLen - 2
)

func nextLineNum() {
	ep := LineEnd

	for {
		if LineBuf[ep] < '9' {
			LineBuf[ep]++
			return
		}

		LineBuf[ep] = '0'
		ep--

		if ep < LineStart {
			break
		}
	}

	LineStart--
	LineBuf[LineStart] = '1'

	if LineStart < LinePrint {
		LinePrint--
	}
}

func Cat(reader io.Reader, buf []byte, writer *bufio.Writer, flags int) int {
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
					writer.Flush()
					return 0
				}
				if err != nil {
					writer.Flush()
					return 1
				}

				beginningOfBuffer = 0 //Reset bpin to the beginning of the buffer
				endOfBuffer = n       //End of buffer is the number of bytes read
				buf[endOfBuffer] = 10 //Place a sentinel at the end of the buffer
			} else {
				newlines++
				if newlines > 0 {
					if (flags & ShowAllLineNumbers) == 2 {
						nextLineNum()
						writer.Write(LineBuf[LinePrint:])
					}
				}
				writer.WriteByte(10)
			}

			ch = buf[beginningOfBuffer]
			beginningOfBuffer++
			if ch != 10 {
				break
			}
		}

		if newlines >= 0 && (flags&ShowAllLineNumbers) == 2 {
			nextLineNum()
			writer.Write(LineBuf[LinePrint:])
		}

		for {
			if ch == 9 && (flags&ShowTabs) == 1 {
				writer.Write(HorizTab)
			} else if ch != 10 {
				writer.WriteByte(ch)
			} else {
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
