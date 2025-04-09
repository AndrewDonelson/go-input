// file: mouse/stringers_test.go
// description: Test cases for the string representations of State and Button types.
//
// Copyright 2025 Andrew Donelson. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package mouse

import (
	"testing"
)

func TestStateString(t *testing.T) {
	testCases := []struct {
		state    State
		expected string
	}{
		{InvalidState, "InvalidState"},
		{Down, "Down"},
		{Up, "Up"},
		{State(99), "State(99)"}, // Out of range value
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			result := tc.state.String()
			if result != tc.expected {
				t.Errorf("Expected State(%d).String() to return '%s', got '%s'", tc.state, tc.expected, result)
			}
		})
	}
}

func TestButtonString(t *testing.T) {
	testCases := []struct {
		button   Button
		expected string
	}{
		{Invalid, "Invalid"},
		{One, "One"},
		{Two, "Two"},
		{Three, "Three"},
		{Left, "One"},              // Alias check
		{Right, "Two"},             // Alias check
		{Middle, "Three"},          // Alias check
		{Button(20), "Button(20)"}, // Out of range value
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			result := tc.button.String()
			if result != tc.expected {
				t.Errorf("Expected Button(%d).String() to return '%s', got '%s'", tc.button, tc.expected, result)
			}
		})
	}
}
