// file: keyboard/events_test.go
// description: Test cases for keyboard event types and methods.
//
// Copyright 2025 Andrew Donelson. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package keyboard

import (
	"strings"
	"testing"
	"time"
)

func TestButtonEvent(t *testing.T) {
	now := time.Now()
	event := ButtonEvent{
		T:     now,
		Key:   A,
		State: Down,
		Raw:   123,
	}

	// Test Time() method
	if !event.Time().Equal(now) {
		t.Fatalf("Expected event.Time() to return %v, got %v", now, event.Time())
	}

	// Test String() method
	str := event.String()
	if str == "" {
		t.Fatal("ButtonEvent.String() should not return empty string")
	}

	// Check for expected components in string representation
	expected := []string{"ButtonEvent", "Key=A", "State=Down", "Raw=123"}
	for _, exp := range expected {
		if !strings.Contains(str, exp) {
			t.Fatalf("ButtonEvent.String() should contain '%s', got: %s", exp, str)
		}
	}
}

func TestTyped(t *testing.T) {
	now := time.Now()
	typed := Typed{
		T: now,
		S: "Hello, World!",
	}

	// Test Time() method
	if !typed.Time().Equal(now) {
		t.Fatalf("Expected typed.Time() to return %v, got %v", now, typed.Time())
	}

	// Test String() method
	str := typed.String()
	if str != "Hello, World!" {
		t.Fatalf("Expected typed.String() to return 'Hello, World!', got '%s'", str)
	}
}
