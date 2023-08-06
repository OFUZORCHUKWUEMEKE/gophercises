package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	// _ = csvFilename

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
		// os.Exit(1)
	}
	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problem := parseLines(lines)

	for i, p := range problem {
		fmt.Printf("Problem #%d: %s =\n", i+1, p.q)
	}
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
