// file: keyboard/stringers_test.go
// description: Test cases for the string representations of State and Key types.
//
// Copyright 2025 Andrew Donelson. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package keyboard

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

func TestKeyString(t *testing.T) {
	testCases := []struct {
		key      Key
		expected string
	}{
		{Invalid, "Invalid"},
		{A, "A"},
		{Enter, "Enter"},
		{Space, "Space"},
		{F1, "F1"},
		{LeftShift, "LeftShift"},
		{RightCtrl, "RightCtrl"},
		{Key(-1), "Key(-1)"},     // Negative out of range
		{Key(1000), "Key(1000)"}, // Positive out of range
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			result := tc.key.String()
			if result != tc.expected {
				t.Errorf("Expected Key(%d).String() to return '%s', got '%s'", tc.key, tc.expected, result)
			}
		})
	}
}
