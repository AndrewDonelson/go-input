// file: userinput/userinput.go
// description: Package for handling user input from the console with various types

package userinput

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// UserInput provides methods for reading user input from the console
type UserInput struct {
	reader    *bufio.Reader
	falseAlts []string
	trueAlts  []string
}

// New creates a new UserInput instance with a bufio.Reader
// for reading user input, and pre-defined lists of false and true
// alternative input values.
func New() *UserInput {
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
	min := math.MinInt32
	max := math.MaxInt32

	switch len(args) {
	case 1:
		// Only one value provided, treat it as max
		max = args[0]
	case 2:
		// Both min and max provided
		min = args[0]
		max = args[1]
	}

	input, err := ui.ReadString(prompt)
	if err != nil {
		return 0, err
	}

	val, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid integer: %w", err)
	}

	if val < min || val > max {
		return 0, fmt.Errorf("value %d is outside allowed range [%d, %d]", val, min, max)
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

	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid floating point number: %w", err)
	}

	return val, nil
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
	return false, fmt.Errorf("invalid boolean input: %v (expected one of: %v or %v)",
		input, strings.Join(ui.trueAlts, ", "), strings.Join(ui.falseAlts, ", "))
}
