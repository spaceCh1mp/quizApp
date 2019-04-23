package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type quiz struct {
	question string
	answer   string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "A csv file in the format (question, answer)")
	limit := flag.Int("limit", 30, "The timer is in seconds")
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
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	fmt.Printf("Total number of questions to answer: %d \n", len(problem))
	fmt.Println("Time alloted: ", *limit)
	score := startQuiz(problem, timer)
	fmt.Printf("Your score is %d out of %d", score, len(problem))
}

func startQuiz(problems []quiz, limit *time.Timer) int {
	var s int = 0
	var a string
	go func() {
		for i, p := range problems {
			fmt.Printf("Q%d: %s = \n", i+1, p.question)
			fmt.Scanf("%s\n", &a)
			if a == p.answer {
				s += 1
			}
		}
	}()
	<-limit.C
	fmt.Printf("\nYour time is up\n")
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
