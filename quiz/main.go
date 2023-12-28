package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file " + filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + filePath, err)
	}

	return records
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	file_name := flag.String("file", "problems.csv", "Specify csv file to read questions from")
	time_limit := flag.Int("limit", 30, "Time limit in seconds")
	flag.Parse()
	records := readCsvFile(*file_name)

	timer := time.NewTimer(time.Duration(*time_limit) * time.Second)

	correct := 0
	problemloop:
	for _, v := range records {
		question, answer := v[0], v[1]
		fmt.Println("What is the answer to " + question + "?")
		answer_ch := make(chan string)
		go func()  {
			input := get_input(reader)		
			answer_ch <- input
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case input := <-answer_ch:
			if strings.Compare(input, answer) == 0 {
				correct++
			}
		}
	}
	fmt.Println("You got", correct, "out of", len(records))
}

func get_input(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		input = strings.TrimRight(input, "\r\n")
	} else {
		input = strings.TrimRight(input, "\n")
	}
	return input
}
