// Open means we should have the ability to add new things
// But to add new things we should not modify the code itself, it should automatically work
// That means it should be colosed for modification, but must have the ability to add new things
// Without modifiying the internal code of the function

// According to the example below (AFTER CHANGING THE PREVIOUS CODE)
// Now if we want to show user a new type of quiz
// We can just create a new struct depending on the structure of the quiz
// And then implement th equiz interface
// This way we do not have to modify printquiz function every time we add a new
// Type of quize
// Hence the OPEN CLOSE principle holds in this code

package main

import "fmt"

// Quiz interface
type quiz interface {
	getDescription() string
	getOptions() []string
}

// Quiz function receives a quiz
// It knows anything which is a quiz must have the
// mthods getDescription and getOptions
// It need not know that are the options and what is the type of a particuar quiz
// As long as It has the getDescription and getOptions methods
// it can print a quiz
func printQuiz(q quiz) {

	fmt.Println(q.getDescription())
	options := q.getOptions()
	for i, op := range options {
		fmt.Printf("%d.   %s\n", i+1, op)
	}
	fmt.Println()
}

type multipleChoice struct {
	description string
	options     []string
}

// Lets implement the quiz interface by implementing the
// methods of the interface
// NOTE : Go interfaces are implemented implicitly
func (mc multipleChoice) getDescription() string {
	return mc.description
}

func (mc multipleChoice) getOptions() []string {
	return mc.options
}

// If we have another ques type such as boolean
// Create the stuct and implement the quiz interface
type booleanQues struct {
	description string
}

func (bq booleanQues) getDescription() string {
	return bq.description
}

func (bq booleanQues) getOptions() []string {
	return []string{"True", "False"}
}

func main() {

	ques := []quiz{
		booleanQues{
			// Boolean quiz
			description: "Earth is flat?",
		},
		multipleChoice{
			description: "What is your favourite language",
			options:     []string{"Go", "pyhton", "c++", "java"},
		},

		// Now if we want to show user a new type of quiz
		// We can just create a new struct depending on the structure of the quiz
		// And then implement the quiz interface
		// Then add the data of the quiz here
		// This way we do not have to modify printquiz function every time we add a new
		// Type of quize
		// Hence the OPEN CLOSE principle holds in this code
	}

	for _, q := range ques {
		printQuiz(q)
	}

}
