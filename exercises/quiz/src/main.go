package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func readCSV() []problem {
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

	var quiz []problem
	for _, line := range lines {
		quiz = append(quiz, problem{
			question: line[0],
			answer:   line[1],
		})
	}

	return quiz
}

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
	timeout := flag.Int("t", 30, "timeout in seconds")
	flag.Parse()

	timer := time.NewTimer(time.Duration(*timeout) * time.Second)
	quiz := readCSV()

	fmt.Scanln("Press enter to start the quiz...")
	fmt.Print("\033[H\033[2J") // clear screen

	correct := 0.0
	length := float64(len(quiz))

	// channel to signal that the time is up
	done := make(chan bool, 1)

	go func() {
		<-timer.C
		fmt.Print("\033[H\033[2J") // clear screen
		fmt.Print("Time's up! Press enter to see your results...")
		done <- true // timer expire signal
	}()

QuizLoop:
	for _, problem := range quiz {
		select {
		case <-done:
			break QuizLoop
		default:
			var userAnswer string
			fmt.Printf("%s : ", problem.question)
			fmt.Scanln(&userAnswer)
			fmt.Print("\033[H\033[2J") // clear screen

			isCorrect := userAnswer == problem.answer
			if isCorrect {
				correct++
			}
		}
	}

	timer.Stop()

	percentage := math.Round(correct / length * 100)
	grade := letterGrade(percentage)

	fmt.Println("Quiz Results: ", grade)
	fmt.Printf("Grade Percentage: %v%%\n", percentage)
	fmt.Println("Total questions: ", length)
	fmt.Println("Correct answers: ", correct)
}
