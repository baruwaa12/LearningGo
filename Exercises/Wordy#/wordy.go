package wordy


import (
	"fmt"
	"strconv"
	"strings"
)

var left int
var right int
var operator string
var answer int



func getOperator(question []string) (operator string) {
	if len(question) == 5 {
		operator = question[3]
		fmt.Println(operator)
	}
	if len(question) == 6 {
		operator = question[3]
		fmt.Println(operator)
	}
	return operator
}

func getLeft(question []string) (left int) {
	if len(question) == 5 || len(question) == 6 {	
		left, _ = strconv.Atoi(question[2])
		fmt.Println(left)
	}
	return left
}


func getRight(question []string) (right int) {
	if len(question) == 5 {
		right, _ = strconv.Atoi(question[4])
		fmt.Println(right)
	}
	if len(question) == 6 {
		right, _ = strconv.Atoi(question[5])
		fmt.Println(right)
	}
	return right
}


func calculate(questionStr string) (ok bool, answer int) {
	// Remove question mark from question if present
	questionStr = (strings.TrimRight(questionStr, "?"))

	// Split the question into an array by splitting by empty space
	question := strings.Split(questionStr, " ")

	// Use getOperator function to find the operator
	op := getOperator(question)

	// Find the number on the left hand side of operator
	left1 := getLeft(question)

	// Find the number on the right hand side of operator
	right1 := getRight(question)

	// Here are the possible ways the operator can be written
	// to perform the calculation
	if op == "multiplied" || op == "multiply" || op == "times"  {
		answer = (left1 * right1)
	}else if op == "plus" || op == "add" {
		answer = (left1 + right1)
	}else if op == "divide" || op == "divided" {
		answer = (left1 / right1)
	}else if op == "subtract" || op == "minus" || op == "subtracted" {
		answer = (left1 - right1)
	}else {
		return false, answer
	}
	
	return true, answer
} 

