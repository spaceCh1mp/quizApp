package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type quiz struct {
	question string
	answer   string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "A csv file in the format (question, answer)")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Could not open file: %s", *csvFile))
	}
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Could not read file: %s", *csvFile))
	}
	problem := parseRecords(records)
	score := startQuiz(problem)
	fmt.Printf("Your score is %d out of %d", score, len(problem))
}

func startQuiz(problems []quiz) int {
	var s int = 0
	var a string
	for i, p := range problems {
		fmt.Printf("Q%d: %s = \n", i+1, p.question)
		fmt.Scanf("%s\n", &a)
		if a == p.answer {
			s += 1
		}
	}
	return s
}
func parseRecords(records [][]string) []quiz {
	ret := make([]quiz, len(records))
	for i, record := range records {
		ret[i] = quiz{
			question: record[0],
			answer:   record[1],
		}
	}
	return ret
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
