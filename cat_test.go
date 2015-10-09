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
	var routine func(reader io.ReadCloser, channel chan string)
	path := "test.txt"

	BeforeEach(func() {
		stdout = os.Stdout
		readFile, writeFile, err = os.Pipe()
		if err != nil {
			//This isn't correct
			Fail("Couldn't Create Pipe")
		}

		routine = func(reader io.ReadCloser, channel chan string) {
			var b bytes.Buffer
			_, err := io.Copy(&b, reader)
			reader.Close()
			if err != nil {
				Fail("Error in Channel")
			}
			channel <- b.String()
		}
		os.Stdout = writeFile
	})

	AfterEach(func() {
		buf.Reset()
	})

	cleanup := func() {
		writeFile.Close()
		os.Stdout = stdout
	}

	NewTempWriter := func(fileBlockSize int) *bufio.Writer {
		size := 20 + fileBlockSize*4
		return bufio.NewWriterSize(os.Stdout, size)
	}

	var _ = Describe("One argument without any flags", func() {
		It("Outputs the contents of the file", func() {
			file, err := os.Open(path)
			if err != nil {
				Fail("Couldn't Open File")
			}

			fileStat, err := file.Stat()
			if err != nil {
				Fail("Couldn't Stat File")
			}

			outBuf := NewTempWriter(int(fileStat.Sys().(*syscall.Stat_t).Blksize))

			_ = SimpleCat(file, outBuf)
			file.Close()

			outC := make(chan string)
			go routine(readFile, outC)

			args := []string{path}
			b, err := SystemCat(args)
			if err != nil {
				Fail("Error line 70")
			}

			buf.Write(b)
			cleanup()

			out := <-outC

			Expect(out).To(Equal(buf.String()))
		})
	})

	var _ = Describe("One argument & -b", func() {
		It("Outputs the contents of the file with line numbers", func() {
			file, err := os.Open(path)
			if err != nil {
				Fail("Couldn't Open File")
			}

			fileStat, err := file.Stat()
			if err != nil {
				Fail("Couldn't Stat File")
			}

			outBuf := NewTempWriter(int(fileStat.Sys().(*syscall.Stat_t).Blksize))

			flags := 1
			_ = Cat(file, outBuf, flags)
			outBuf.Flush()

			file.Close()

			outC := make(chan string)
			go routine(readFile, outC)

			args := []string{"-b", path}
			b, err := SystemCat(args)
			if err != nil {
				Fail("Error line 70")
			}
			buf.Write(b)

			cleanup()

			out := <-outC

			Expect(out).To(Equal(buf.String()))
		})
	})

})

func SystemCat(args []string) ([]byte, error) {
	cat := exec.Command("cat", args...)
	return cat.Output()
}
