package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"flag"
	"time"
)

type MathProblem struct {
	question string
	Answer   int
}

func ParseLines(lines [][]string) []MathProblem {
	problems := make([]MathProblem, 0)
	for _, line := range lines {
		a, _ := strconv.Atoi(line[1])
		p := MathProblem{
			question: line[0],
			Answer:   a,
		}

		problems = append(problems, p)
	}

	return problems
}

func StartTimer(duration *int, c chan bool) {
	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	isTimerStopped := timer.Stop()
	if isTimerStopped {
		c <- true
	}
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timerDuration := flag.Int("limit", 30, "the time for the quiz in seconds")
	flag.Parse()

	problemsFile, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatal("Error occurred while opening csv file")
	}

	defer problemsFile.Close()

	problemsReader := csv.NewReader(problemsFile)

	lines, err := problemsReader.ReadAll()
	if err != nil {
		log.Fatal("error")
	}

	fmt.Println(lines)

	problems := ParseLines(lines)

	correctAns := 0

	timerChan := make(chan bool)

	for _, problem := range problems {
		var userInput int

		fmt.Printf("%v = ", problem.question)
		go StartTimer(timerDuration, timerChan)
		if <-timerChan {
			fmt.Scan
			break
		}
		fmt.Scanln(&userInput)

		if userInput == problem.Answer {
			correctAns++
		}

	}

	fmt.Printf("You got %v correct out of 12", correctAns)

}
