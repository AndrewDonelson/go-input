// file: userinput/userinput_test.go
// description: Test cases for the userinput package.
//
// Copyright 2025 Andrew Donelson. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package userinput

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// Helper function to redirect standard input for testing
func withInput(input string, testFunc func()) {
	// Save the original stdin
	oldStdin := os.Stdin

	// Create a pipe and set up a reader and writer
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Write the input to the writer
	go func() {
		w.Write([]byte(input))
		w.Close()
	}()

	// Run the test
	testFunc()

	// Restore stdin
	os.Stdin = oldStdin
}

func TestNew(t *testing.T) {
	ui := New()

	if ui == nil {
		t.Fatal("New() returned nil")
	}

	if ui.reader == nil {
		t.Fatal("reader is nil")
	}

	if len(ui.falseAlts) != 5 {
		t.Fatalf("Expected 5 false alternatives, got %d", len(ui.falseAlts))
	}

	if len(ui.trueAlts) != 5 {
		t.Fatalf("Expected 5 true alternatives, got %d", len(ui.trueAlts))
	}
}

func TestReadString(t *testing.T) {
	ui := New()

	// Capture stdout to check the prompt
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test with simple input
	withInput("Hello\n", func() {
		result, err := ui.ReadString("Enter text")

		if err != nil {
			t.Fatalf("ReadString returned an error: %v", err)
		}

		if result != "Hello" {
			t.Fatalf("Expected 'Hello', got '%s'", result)
		}
	})

	// Test with input containing whitespace
	withInput("  Trimmed  \n", func() {
		result, err := ui.ReadString("Enter more text")

		if err != nil {
			t.Fatalf("ReadString returned an error: %v", err)
		}

		if result != "Trimmed" {
			t.Fatalf("Expected 'Trimmed', got '%s'", result)
		}
	})

	// Close the writer and restore stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Check that the prompts were displayed
	output := buf.String()
	if !strings.Contains(output, "Enter text") {
		t.Fatal("Prompt 'Enter text' not displayed")
	}
	if !strings.Contains(output, "Enter more text") {
		t.Fatal("Prompt 'Enter more text' not displayed")
	}
}

func TestReadInt(t *testing.T) {
	ui := New()

	// Test with valid integer
	withInput("42\n", func() {
		result, err := ui.ReadInt("Enter number")

		if err != nil {
			t.Fatalf("ReadInt returned an error: %v", err)
		}

		if result != 42 {
			t.Fatalf("Expected 42, got %d", result)
		}
	})

	// Test with min/max (valid)
	withInput("50\n", func() {
		result, err := ui.ReadInt("Enter number", 0, 100)

		if err != nil {
			t.Fatalf("ReadInt returned an error: %v", err)
		}

		if result != 50 {
			t.Fatalf("Expected 50, got %d", result)
		}
	})

	// Test with min/max (invalid - too low)
	withInput("-10\n", func() {
		_, err := ui.ReadInt("Enter number", 0, 100)

		if err == nil {
			t.Fatal("ReadInt should return an error for value below minimum")
		}
	})

	// Test with min/max (invalid - too high)
	withInput("200\n", func() {
		_, err := ui.ReadInt("Enter number", 0, 100)

		if err == nil {
			t.Fatal("ReadInt should return an error for value above maximum")
		}
	})

	// Test with invalid input
	withInput("not a number\n", func() {
		_, err := ui.ReadInt("Enter number")

		if err == nil {
			t.Fatal("ReadInt should return an error for non-numeric input")
		}
	})
}

func TestReadFloat(t *testing.T) {
	ui := New()

	// Test with valid float
	withInput("3.14\n", func() {
		result, err := ui.ReadFloat("Enter float")

		if err != nil {
			t.Fatalf("ReadFloat returned an error: %v", err)
		}

		if result != 3.14 {
			t.Fatalf("Expected 3.14, got %f", result)
		}
	})

	// Test with invalid input
	withInput("not a float\n", func() {
		_, err := ui.ReadFloat("Enter float")

		if err == nil {
			t.Fatal("ReadFloat should return an error for non-float input")
		}
	})
}

func TestReadBool(t *testing.T) {
	ui := New()

	// Test with true values
	trueInputs := []string{"y\n", "yes\n", "t\n", "true\n", "1\n", "Y\n", "YES\n", "True\n"}
	for _, input := range trueInputs {
		withInput(input, func() {
			result, err := ui.ReadBool("True value")

			if err != nil {
				t.Fatalf("ReadBool returned an error for '%s': %v", input, err)
			}

			if !result {
				t.Fatalf("Expected true for input '%s', got false", input)
			}
		})
	}

	// Test with false values
	falseInputs := []string{"n\n", "no\n", "f\n", "false\n", "0\n", "N\n", "NO\n", "False\n"}
	for _, input := range falseInputs {
		withInput(input, func() {
			result, err := ui.ReadBool("False value")

			if err != nil {
				t.Fatalf("ReadBool returned an error for '%s': %v", input, err)
			}

			if result {
				t.Fatalf("Expected false for input '%s', got true", input)
			}
		})
	}

	// Test with invalid input
	withInput("invalid\n", func() {
		_, err := ui.ReadBool("Invalid value")

		if err == nil {
			t.Fatal("ReadBool should return an error for invalid boolean input")
		}
	})
}
