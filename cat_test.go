package main_test

import (
	"bufio"
	"bytes"
	. "cat"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"os"
	"os/exec"
	"syscall"
)

var _ = Describe("Cat", func() {

	var stdout *os.File
	var readFile *os.File
	var writeFile *os.File
	var err error
	var buf bytes.Buffer

	BeforeEach(func() {
		stdout = os.Stdout
		readFile, writeFile, err = os.Pipe()
		if err != nil {
			//This isn't correct
			Fail("Couldn't Create Pipe")
		}
		os.Stdout = writeFile
	})

	It("Outputs the contents of the file", func() {
		path := "test.txt"
		file, err := os.Open(path)
		if err != nil {
			Fail("Couldn't Open File")
		}

		fileStat, err := file.Stat()
		if err != nil {
			Fail("Couldn't Stat File")
		}

		inBsize := int(fileStat.Sys().(*syscall.Stat_t).Blksize)
		size := 20 + inBsize*4
		outBuf := bufio.NewWriterSize(os.Stdout, size)
		inBuf := make([]byte, inBsize+1)

		_ = Cat(file, inBuf, outBuf)
		file.Close()

		outC := make(chan string)
		go func() {
			var b bytes.Buffer
			_, err := io.Copy(&b, readFile)
			readFile.Close()
			if err != nil {
				Fail("Error in Channel")
			}
			outC <- b.String()
		}()

		cat := exec.Command("cat", path)
		b, err := cat.Output()
		if err != nil {
			Fail("Error line 70")
		}
		buf.Write(b)

		writeFile.Close()
		os.Stdout = stdout
		out := <-outC

		Expect(out).To(Equal(buf.String()))
	})
})
