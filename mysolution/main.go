package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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

func main() {

	problemsFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("Error occurred while opening csv file")
	}

	defer problemsFile.Close()

	problemsReader := csv.NewReader(problemsFile)

	lines, err := problemsReader.ReadAll()
	if err != nil {
		log.Fatal("error")
	}

	problems := ParseLines(lines)

	correctAns := 0

	for _, problem := range problems {
		var userInput int

		fmt.Printf("%v = ", problem.question)
		fmt.Scanln(&userInput)

		if userInput == problem.Answer {
			correctAns++
		}

	}

	fmt.Printf("You got %v correct", correctAns)

}
