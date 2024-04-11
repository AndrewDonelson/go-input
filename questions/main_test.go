package questions

import (
	"os"
	"testing"
)

func TestForString(t *testing.T) {
	tests := []struct {
		name     string
		question string
		input    string
		want     string
	}{
		{"Empty input", "What is your name", "\n", ""},
		{"Valid input", "What is your name", "John\n", "What is your name"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}
			os.Stdin = r
			_, err = w.Write([]byte(tt.input))
			if err != nil {
				t.Fatal(err)
			}
			w.Close()

			//var buf bytes.Buffer

			// Execute
			got := (&Ask{}).ForString(tt.question)

			// Assert
			if got != tt.want {
				t.Errorf("ForString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
