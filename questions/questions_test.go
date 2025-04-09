// file: questions/questions_test.go
// description: Tests for the questions package

package questions

import (
	"os"
	"testing"

	"github.com/AndrewDonelson/go-input/keyboard"
	"github.com/AndrewDonelson/go-input/mouse"
)

func TestForString(t *testing.T) {
	tests := []struct {
		name     string
		question string
		input    string
		want     string
	}{
		{"Empty input", "What is your name", "\n", ""},
		{"Valid input", "What is your name", "John\n", "John"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original stdin and create pipe
			oldStdin := os.Stdin
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}
			os.Stdin = r

			// Write test input to pipe
			_, err = w.Write([]byte(tt.input))
			if err != nil {
				t.Fatal(err)
			}
			w.Close()

			// Create Ask with default keyboard and mouse
			ask := New(keyboard.NewWatcher(), mouse.NewWatcher())

			// Call ForString and capture output
			got := ask.ForString(tt.question)

			// Restore original stdin
			os.Stdin = oldStdin

			// Assert
			if got != tt.want {
				t.Errorf("ForString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForInt(t *testing.T) {
	t.Skip("Skipping test for ForInt as it requires user input")
}

func TestForFloat(t *testing.T) {
	t.Skip("Skipping test for ForFloat as it requires user input")
}

func TestForBool(t *testing.T) {
	t.Skip("Skipping test for ForBool as it requires user input")
}

func TestForChoice(t *testing.T) {
	t.Skip("Skipping test for ForChoice as it requires user input")
}
