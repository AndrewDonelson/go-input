// file: mouse/events_test.go
// description: Test cases for mouse event types and methods.
//
// Copyright 2025 Andrew Donelson. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package mouse

import (
	"strings"
	"testing"
	"time"
)

func TestButtonEvent(t *testing.T) {
	now := time.Now()
	event := ButtonEvent{
		T:      now,
		Button: Left,
		State:  Down,
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
	expected := []string{"ButtonEvent", "Button=One", "State=Down"}
	for _, exp := range expected {
		if !strings.Contains(str, exp) {
			t.Fatalf("ButtonEvent.String() should contain '%s', got: %s", exp, str)
		}
	}
}

func TestScrolled(t *testing.T) {
	now := time.Now()
	scrolled := Scrolled{
		T: now,
		X: 10.5,
		Y: -5.25,
	}

	// Test Time() method
	if !scrolled.Time().Equal(now) {
		t.Fatalf("Expected scrolled.Time() to return %v, got %v", now, scrolled.Time())
	}

	// Test String() method
	str := scrolled.String()
	if str == "" {
		t.Fatal("Scrolled.String() should not return empty string")
	}

	// Check for expected components in string representation
	expected := []string{"Scrolled", "X=10.500000", "Y=-5.250000"}
	for _, exp := range expected {
		if !strings.Contains(str, exp) {
			t.Fatalf("Scrolled.String() should contain '%s', got: %s", exp, str)
		}
	}
}
