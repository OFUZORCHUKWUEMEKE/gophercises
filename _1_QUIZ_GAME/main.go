package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	// timeLimit := flag.Int("time limit", 25, "The time liit fot he quiz")
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

	fmt.Println(problem)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)


	correct := 0

	for i, p := range problem {
		fmt.Printf("Problem #%d: %s =\n", i+1, p.q)

		answerChan := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out %d.\n", correct, len(problem))
			return
		case answer := <-answerChan:
			if answer == p.a {
				correct++
			}
		}

	}
	fmt.Printf("You scored %d out %d.\n", correct, len(problem))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
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
