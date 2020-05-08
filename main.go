package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var status = "Play"
var command = make(chan string)

func main() {
	inputPath := flag.String("input", "", "a string")
	outputPath := flag.String("output", "", "a string")
	sleep := flag.String("sleep", "200ms", "a string")
	flag.Parse()
	if *inputPath == "" {
		log.Fatalln("input file is required")
	}
	if *outputPath == "" {
		log.Fatalln("output file is required")
	}

	duration, err := time.ParseDuration(*sleep)
	if err != nil {
		log.Fatalln(err)
	}
	go scannerCommands()
	writeLines(*inputPath, *outputPath, duration)
}

func writeLines(inputPath, outputPath string, sleep time.Duration) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileOutput, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer fileOutput.Close()
	r := bufio.NewReader(file)
	var wg sync.WaitGroup
	wg.Add(1)

	work := func(r *bufio.Reader, fileOutput *os.File) {
		bytes, err := r.ReadBytes('\n')
		if err == io.EOF {
			writeLineByLine(string(bytes), fileOutput, sleep)
			os.Exit(0)
		} else if err != nil {
			log.Fatal(err)
		}
		writeLineByLine(string(bytes), fileOutput, sleep)
	}

	go func(command <-chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case cmd := <-command:
				fmt.Println("Command -> ", cmd)
				switch cmd {
				case "STOP":
					return
				case "PAUSE":
					status = "PAUSE"
				default:
					status = "PLAY"
				}
			default:
				if status == "PLAY" {
					work(r, fileOutput)
				}

			}
		}
	}(command, &wg)

	command <- "PLAY"
	wg.Wait()
}

func writeLineByLine(line string, fileOutput *os.File, sleep time.Duration) {
	w := bufio.NewWriter(fileOutput)
	for _, strLine := range line {
		time.Sleep(sleep)
		if _, err := w.WriteString(string(strLine)); err != nil {
			panic(err)
		}
		if err := w.Flush(); err != nil {
			panic(err)
		}
		err := fileOutput.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func scannerCommands() {
	for {
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(char)
		switch char {
		case 13:
			if status == "PLAY" {
				command <- "PAUSE"
			} else {
				command <- "PLAY"
			}
			break
		case 'a':
			fmt.Println("a Key Pressed")
			break
		}
	}
}
