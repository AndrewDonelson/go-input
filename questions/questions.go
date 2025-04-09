// file: questions/questions.go
// description: Package questions provides a simple way to ask questions and get user input

package questions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/AndrewDonelson/go-input/keyboard"
	"github.com/AndrewDonelson/go-input/mouse"
)

// Ask provides methods for asking questions and getting user input with
// keyboard and mouse support
type Ask struct {
	Keyboard *keyboard.Watcher
	Mouse    *mouse.Watcher
	reader   *bufio.Reader
}

// New creates a new Ask instance with the specified keyboard and mouse watchers
func New(k *keyboard.Watcher, m *mouse.Watcher) *Ask {
	return &Ask{
		Keyboard: k,
		Mouse:    m,
		reader:   bufio.NewReader(os.Stdin),
	}
}

// ForString prompts the user with a question and returns their string response
func (a *Ask) ForString(question string) string {
	fmt.Printf("%s? ", question)

	text, err := a.reader.ReadString('\n')
	if err != nil {
		return ""
	}

	// Trim newline character and any other whitespace
	return strings.TrimSpace(text)
}

// ForInt prompts the user with a question and returns their integer response
func (a *Ask) ForInt(question string) int {
	fmt.Printf("%s? ", question)

	text, err := a.reader.ReadString('\n')
	if err != nil {
		return 0
	}

	text = strings.TrimSpace(text)
	answer, err := strconv.Atoi(text)
	if err != nil {
		return 0
	}

	return answer
}

// ForFloat prompts the user with a question and returns their float response
func (a *Ask) ForFloat(question string) float64 {
	fmt.Printf("%s? ", question)

	text, err := a.reader.ReadString('\n')
	if err != nil {
		return 0
	}

	text = strings.TrimSpace(text)
	answer, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0
	}

	return answer
}

// ForBool prompts the user with a question and returns their boolean response
// It accepts "true", "false", "t", "f", "yes", "no", "y", "n", "1", "0" as valid inputs
func (a *Ask) ForBool(question string) bool {
	fmt.Printf("%s? ", question)

	text, err := a.reader.ReadString('\n')
	if err != nil {
		return false
	}

	text = strings.ToLower(strings.TrimSpace(text))

	// Check for truthy values
	switch text {
	case "true", "t", "yes", "y", "1":
		return true
	default:
		return false
	}
}

// ForChoice prompts the user with a question and a list of choices, and returns their selection
func (a *Ask) ForChoice(question string, choices []string) string {
	fmt.Printf("%s? Choose from: ", question)

	for i, choice := range choices {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%s", choice)
	}
	fmt.Println()

	text, err := a.reader.ReadString('\n')
	if err != nil {
		return ""
	}

	answer := strings.TrimSpace(text)

	// Check if answer is in choices
	for _, choice := range choices {
		if strings.EqualFold(answer, choice) {
			return choice
		}
	}

	// Return first choice if input not found
	if len(choices) > 0 {
		return choices[0]
	}

	return ""
}
