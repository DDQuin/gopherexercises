package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
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
	records := readCsvFile("problems.csv")
	correct := 0

	for _, v := range records {
		question, answer := v[0], v[1]
		fmt.Println("What is the answer to " + question + "?")
		input, _ := reader.ReadString('\n')
		if runtime.GOOS == "windows" {
		  input = strings.TrimRight(input, "\r\n")
		} else {
		  input = strings.TrimRight(input, "\n")
		}

		if strings.Compare(input, answer) == 0 {
			correct++
		}
	}

	fmt.Println("You got", correct, "out of", len(records))


}
