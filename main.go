package main

import "os"
import "time"
import "bufio"
import "log"
import "io"
import "flag"

//var command = make(chan string)
//var status = "Play"

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
  for {
    bytes, err := r.ReadBytes('\n')
    if err == io.EOF {
      writeLineByLine(string(bytes), fileOutput, sleep)
      break
    } else if err != nil {
      log.Fatal(err)
    }
    writeLineByLine(string(bytes), fileOutput, sleep)
  }
}

func writeLineByLine(line string, fileOutput *os.File, sleep time.Duration) {
  for _, strLine := range line {
    time.Sleep(sleep)
    fileOutput.WriteString(string(strLine))
  }
}
