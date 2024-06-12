package main

import (
	"encoding/csv"
	"math"
	"fmt"
	"log"
	"os"
)

// type to hold quiz question instance
type quiz struct {
	question string
	answer string
	grade *bool // optional bool to track correct (true) and incorrect (false) answers, initialized as <nil> for unanswered questions
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

// convert percentage to grade
func letterGrade(score float64) string {
    switch {
    case score >= 90:
        return "A"
    case score >= 80:
        return "B"
    case score >= 70:
        return "C"
    case score >= 60:
        return "D"
    default:
        return "F"
    }
}

func main() {
	// initialize full quiz
	quizzes := readCSV()

	print("Press enter to start quiz...")
	fmt.Scanln()
	fmt.Print("\033[H\033[2J") // clear screen


	// iterate through array of quizzes, prompt user for answers, and grade
	for i, quiz := range quizzes {
		var userAnswer string
		fmt.Printf("\r%s : ", quiz.question)
		fmt.Scan(&userAnswer)
		fmt.Print("\033[H\033[2J") // clear screen

		quiz.grade = new(bool)
		if userAnswer == quiz.answer {
			*quiz.grade = true
		} else {
			*quiz.grade = false
		}
		quizzes[i].grade = quiz.grade
	}

	// print results
	correct := 0.0
	length := float64(len(quizzes))

	// calculate correct answers
	for _, quiz := range quizzes {
		if *quiz.grade {
			correct++
		}
	}

	percentage := math.Round(correct / length * 100)

	fmt.Println("Quiz Results:     ", letterGrade(percentage))
	fmt.Printf("Grade Percentage:  %v%%\n\n", percentage)
	fmt.Println("Total questions:  ", length)
	fmt.Println("Correct answers:  ", correct)
}