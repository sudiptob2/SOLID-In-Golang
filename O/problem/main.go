// Open means we should have the ability to add new things
// But to add new things we should not modify the code itself, it should automatically work
// That means it should be colosed for modification, but must have the ability to add new things
// Without modifiying the internal code of the function

// According to the example below
// printQuiz func prints different types of quiz
// But if we want to add any new type of quiz such as fill in the blank,
// we have to add another else if statement and put the printing logic there
// That means we have to modify the internal code
// And it violates Open Closed principle
// What we should have been able to do is that,
// pass the new type of ques to the printQuiz func and it should autometically work

package main

import "fmt"

type quiz struct {
	qtype       string
	description string
	options     []string
}

func (q quiz) printQuiz() {
	fmt.Println(q.description)
	if q.qtype == "boolean" {
		fmt.Printf("1.True\n2.False\n")

	} else if q.qtype == "multiple choice" {

		fmt.Println("Choose the best option")
		for i, option := range q.options {
			fmt.Println(i+1, option)
		}
	}
	// else if q.qtype == "open ques" {
	// 	fmt.Println("Write your ans wer here_______________")
	// }
	fmt.Println()
}

func main() {

	ques := []quiz{

		{
			// Boolean quiz
			qtype:       "boolean",
			description: "Earth is flat?",
		},
		{
			// Multiple Choice quiz
			qtype:       "multiple choice",
			description: "What is your favourite language",
			options:     []string{"Go", "pyhton", "c++", "java"},
		},
		{
			// open Ques quiz
			// This will not work without making changes in the printQuiz func
			// We have to provide the input field to the user to get the ans
			// So violets Open Close principle
			// Because according to the open close principle we shlould be able
			// to just add quiz here and it should autometically work
			qtype:       "open ques",
			description: "What is your name?",
		},
	}

	for _, qz := range ques {
		qz.printQuiz()
	}

}
