/* Exercise 4: Online Examination System
Topics Covered: Go Loops, Go Break and Continue, Go Constants, Go Strings, Go
Functions, Go Errors
Case Study:
	Develop an online examination system where users can take a quiz.
	1. Question Bank: Define a slice of structs to store questions. Each question
	should have a question string, options (array), and the correct answer.

	2. Take Quiz: Use loops to iterate over questions and display them one by one.
	Allow the user to select an answer by entering the option number.
		o Use continue to skip invalid inputs and prompt the user again.
		o Use break to exit the quiz early if the user enters a specific command
		(e.g., "exit").

	3. Score Calculation: After the quiz, calculate the user's score and display it. Use
	conditions to classify performance (e.g., "Excellent", "Good", "Needs
	Improvement").

	4. Error Handling: Handle errors like invalid input during the quiz (e.g., entering a
	non-integer value for an option).
	Bonus:
• Add a timer for the quiz, limiting each question to a fixed amount of time.
*/

package main

import (
	"fmt"
	"strconv"
	"time"
)

type Question struct {
	QuestionText  string
	Options       [4]string
	CorrectAnswer int
}

func main() {
	// Question Bank
	questions := []Question{
		{
			QuestionText:  "Who is Prime Minister of India?",
			Options:       [4]string{"1. Nithish", "2. Modi", "3. Rahul", "4. Amith"},
			CorrectAnswer: 2,
		},
		{
			QuestionText:  "Who is the captain of Indian cricket team in 2024 T20 World Cup?",
			Options:       [4]string{"1. Virat", "2. Hardik", "3. Rohit", "4. Dhoni"},
			CorrectAnswer: 3,
		},
		{
			QuestionText:  "What is the capital of India?",
			Options:       [4]string{"1. Delhi", "2. Mumbai", "3. Hyderabad", "4. Bangalore"},
			CorrectAnswer: 1,
		},
	}

	var score int
	var quit bool

	// Take the quiz
	for i, q := range questions {
		if quit {
			break
		}

		fmt.Printf("\nQuestion %d: %s\n", i+1, q.QuestionText)
		for _, option := range q.Options {
			fmt.Println(option)
		}

		timerExpired := make(chan bool)
		inputDone := make(chan bool)

		// Timer
		go func() {
			time.Sleep(10 * time.Second)
			timerExpired <- true
		}()

		var userInput string
		// Input
		go func() {
			for {
				fmt.Print("Enter your answer (1-4) or 'exit' to quit: ")
				fmt.Scanln(&userInput)

				if userInput == "exit" {
					fmt.Println("You exited the quiz.")
					quit = true
					close(inputDone)
					return
				}

				userAnswer, err := strconv.Atoi(userInput)
				if err != nil || userAnswer < 1 || userAnswer > 4 {
					fmt.Println("Invalid input, please enter a number between 1 and 4.")
					continue // Prompt the user again
				}

				// Valid input
				if userAnswer == q.CorrectAnswer {
					score++
				}
				close(inputDone)
				break
			}
		}()

		// Wait for either input or timer expiration
		select {
		case <-timerExpired:
			fmt.Println("\nTime's up for this question!")
			close(inputDone)
		case <-inputDone:
		}
	}

	// Score Calculation
	fmt.Printf("\nYour final score is: %d/%d\n", score, len(questions))

	// Performance classification
	switch {
	case score == len(questions):
		fmt.Println("Excellent!")
	case score >= len(questions)/2:
		fmt.Println("Good!")
	default:
		fmt.Println("Needs Improvement!")
	}
}
