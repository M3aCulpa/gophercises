package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// type to hold quiz question instance
type quiz struct {
	question string
	answer string
	grade *bool // Optional bool to track correct (true) and incorrect (false) answers, initialized as <nil> for unanswered questions
}

// read and parse csv data and save to an array
func readCSV() []quiz {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var data []quiz
	for _, line := range lines {
		data = append(data, quiz{
			question: line[0],
			answer: line[1],
		})
	}

	return data
}

func main() {
	// initialize full quiz
	quizzes := readCSV()

	// iterate through array of quizzes type
	for _, quiz := range quizzes {
		fmt.Println(quiz)
	}
}