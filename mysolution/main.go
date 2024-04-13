package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	problemsFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("Error occurred while opening csv file")
	}

	defer problemsFile.Close()

	problems := csv.NewReader(problemsFile)

	correctAns := 0

	numProblems := 12

	fmt.Println(numProblems)

	for i := 0; i < numProblems; i++ {
		problem, err := problems.Read()

		if err != nil {
			log.Fatal("error", err)
		}

		ans, _ := strconv.Atoi(problem[1])
		var userInput int

		fmt.Printf("%v = ", problem[0])
		fmt.Scanln(&userInput)

		if userInput == ans {
			correctAns++
		}

	}

	fmt.Printf("You got %v correct", correctAns)

}
