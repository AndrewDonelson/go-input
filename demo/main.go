package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type UserInput struct {
	reader    *bufio.Reader
	falseAlts []string
	trueAlts  []string
}

// NewUserInput creates a new UserInput instance with a bufio.Reader
// for reading user input, and pre-defined lists of false and true
// alternative input values.
func NewUserInput() *UserInput {
	return &UserInput{
		reader:    bufio.NewReader(os.Stdin),
		falseAlts: []string{"n", "no", "f", "false", "0"},
		trueAlts:  []string{"y", "yes", "t", "true", "1"},
	}
}

// ReadString reads a string from the user input, prompting with the given string.
// It trims any leading or trailing whitespace from the input.
// If an error occurs while reading the input, it is returned.
func (ui *UserInput) ReadString(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := ui.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// ReadInt reads an integer value from the user, prompting with the given string.
// It returns the integer value read, or an error if the input could not be parsed.
func (ui *UserInput) ReadInt(prompt string, args ...int) (int, error) {

	var min, max, val int

	switch len(args) {
	case 2:
		// Only one value provided, treat it as max
		max = args[0]
	case 3:
		// Both min and max provided
		min = args[0]
		max = args[1]
	default:
		// No min and max provided, use default values
		min = math.MinInt32
		max = math.MaxInt32
	}

	// Loop until we get a valid integer value
	for {
		input, _ := ui.ReadString(prompt)

		val, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Invalid input, please enter an integer value between %d and %d\n", min, max)
			continue
		}
		if val < min || val > max {
			fmt.Printf("Invalid input, please enter an integer value between %d and %d\n", min, max)
			continue
		}

		break
	}

	return val, nil
}

// ReadFloat reads a float64 value from the user, prompting with the given string.
// It returns the float64 value read, or an error if there was a problem reading the input.
func (ui *UserInput) ReadFloat(prompt string) (float64, error) {
	input, err := ui.ReadString(prompt)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(input, 64)
}

// ReadBool reads a boolean value from the user input. It accepts a variety of common
// boolean string representations (e.g. "y", "yes", "t", "true", "1", "n", "no", "f", "false", "0")
// and returns the corresponding boolean value. If the input is not a valid boolean
// representation, it returns an error.
func (ui *UserInput) ReadBool(prompt string) (bool, error) {
	input, err := ui.ReadString(prompt)
	if err != nil {
		return false, err
	}

	input = strings.ToLower(input)
	for _, alt := range ui.falseAlts {
		if input == alt {
			return false, nil
		}
	}
	for _, alt := range ui.trueAlts {
		if input == alt {
			return true, nil
		}
	}
	return false, fmt.Errorf("invalid input: %v", input)
}

func main() {
	ui := NewUserInput()

	name, err := ui.ReadString("Enter your name: ")
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}
	fmt.Println("Hello,", name)

	age, err := ui.ReadInt("Enter your age: ")
	if err != nil {
		fmt.Println("Error reading age:", err)
		return
	}
	fmt.Println("Your age is:", age)

	height, err := ui.ReadFloat("Enter your height (in meters): ")
	if err != nil {
		fmt.Println("Error reading height:", err)
		return
	}
	fmt.Println("Your height is:", height, "meters")

	isStudent, err := ui.ReadBool("Are you a student? (true/false): ")
	if err != nil {
		fmt.Println("Error reading student status:", err)
		return
	}
	fmt.Println("Student status:", isStudent)
}
