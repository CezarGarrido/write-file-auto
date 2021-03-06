package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	// open input file
	fi, err := os.Open("my-script.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	r := bufio.NewReader(fi)

	// open output file
	fo, err := os.Create("main.js")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		//fmt.Println("buf[:n]->", string(buf[:n]))
		for _, str := range strings.Split(string(buf[:n]), "") {
			// write a chunk
			fmt.Println(str)
			time.Sleep(200 * time.Millisecond)
			if _, err := w.WriteString(str); err != nil {
				panic(err)
			}
			if err = w.Flush(); err != nil {
				panic(err)
			}

		}
	}
	//buf[:n]
}
