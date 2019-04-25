package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var sum int = 0

type quiz struct {
	question string
	answer   string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "A csv file in the format (question, answer)")
	limit := flag.Int("limit", 30, "The timer is in seconds")
	shuffle := flag.Bool("shuffle", false, "Quiz doesn't shuffle by default, pass \"true\" as a value to shuffle")
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
	if *shuffle {
		rand.Seed(time.Now().UTC().UnixNano())
		rand.Shuffle(len(problem), func(i, j int) {
			problem[i], problem[j] = problem[j], problem[i]
		})
	}
	startQuiz(problem, timer)
	fmt.Printf("Your score is %d out of %d", sum, len(problem))
}

func startQuiz(problems []quiz, limit *time.Timer) {
	c := make(chan bool)
	go func() {
		var a string
		for i, p := range problems {
			fmt.Printf("Q%d: %s = \n", i+1, p.question)
			fmt.Scanf("%s\n", &a)
			a = cleanup(a)
			ans := cleanup(p.answer)
			if a == ans {
				sum += 1
			}
		}
		c <- true
	}()
	var m string
	select {
	case <-limit.C:
		m = "\nYour time is up"
		break
	case <-c:
		m = "\nYou Finished!"
		break
	}
	fmt.Println(m)
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
func cleanup(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
