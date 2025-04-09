// file: mouse/watcher_test.go
// description: Test cases for the mouse watcher functionality.
//
// Copyright 2025 Andrew Donelson. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package mouse

import (
	"strings"
	"testing"
)

var wantStr = `mouse.Watcher(
	One: Down,
	Two: Up,
	Button(255): Down,
)`

func TestWatcher(t *testing.T) {
	m := NewWatcher()
	m.SetState(Left, Down)
	m.SetState(Right, Up)
	if !m.Down(Left) {
		t.Fatal("expect mouse.Left in state mouse.Down")
	}
	if !m.Up(Right) {
		t.Fatal("expect mouse.Right in state mouse.Up")
	}
	if !m.Up(Wheel) {
		t.Fatal("expect mouse.Wheel in state mouse.Up")
	}

	// Verify the state lookup table.
	want := map[Button]State{
		Left:  Down,
		Right: Up,
		Wheel: InvalidState,
	}
	states := m.States()
	if len(states) != 8 {
		t.Fatalf("got %d states, want 8\n", len(states))
	}
	for b, s := range states {
		wantState := want[Button(b)]
		if b < 3 && wantState != s {
			t.Fatalf("got %v=%v, want %v=%v\n", Button(b), s, Button(b), wantState)
		}
	}

	// Verify that expansion on the lookup table works OK.
	m.SetState(255, Down)
	got := m.State(255)
	if got != Down {
		t.Fatalf("Wanted Button(255) == Down, got Button(255) == %v", got)
	}

	// Test when state isn't in the table
	// Button is defined as uint8, so let's use an index just outside the array bounds
	// but still valid for the type
	outsideButton := Button(254) // Use a valid but unset button
	m.SetState(outsideButton, InvalidState)
	outsideState := m.State(outsideButton)
	if outsideState != InvalidState {
		t.Fatalf("Expected Button(254) == InvalidState, got %v", outsideState)
	}

	// Check that Up() correctly handles buttons with InvalidState
	if !m.Up(outsideButton) {
		t.Fatal("Expected m.Up(Button(254)) to be true for InvalidState")
	}

	// Check behavior for a button index beyond the current array size
	largeButton := Button(200) // A large but valid button value
	largeState := m.State(largeButton)
	if largeState != Up {
		t.Fatalf("Expected Button(200) == Up (default), got %v", largeState)
	}

	str := m.String()
	if !strings.Contains(str, "One: Down") {
		t.Fatal("String() should contain 'One: Down'")
	}
	if !strings.Contains(str, "Two: Up") {
		t.Fatal("String() should contain 'Two: Up'")
	}
	if !strings.Contains(str, "Button(255): Down") {
		t.Fatal("String() should contain 'Button(255): Down'")
	}
}

func TestEachState(t *testing.T) {
	m := NewWatcher()
	// Set some states to Down
	m.SetState(Left, Down)  // Button 1
	m.SetState(Wheel, Down) // Button 3

	// Set one state to Up
	m.SetState(Right, Up) // Button 2

	// Set one state to InvalidState - this should be skipped by EachState
	m.SetState(Button(4), InvalidState)

	// Make sure the Invalid button state is handled correctly
	m.SetState(Invalid, Down) // This should be skipped by EachState

	// Count states
	buttonCount := 0
	downCount := 0

	m.EachState(func(b Button, s State) bool {
		buttonCount++
		if s == Down {
			downCount++
		}
		return true
	})

	// Should see 3 buttons (Invalid and InvalidState should be skipped)
	if buttonCount != 3 {
		t.Fatalf("Expected 3 buttons, got %d", buttonCount)
	}

	// Should see 2 down states (Left and Wheel)
	if downCount != 2 {
		t.Fatalf("Expected 2 down states, got %d", downCount)
	}

	// Test early termination
	earlyCount := 0
	m.EachState(func(b Button, s State) bool {
		earlyCount++
		return false // Stop after first button
	})

	if earlyCount != 1 {
		t.Fatalf("Expected early termination after 1 button, got %d", earlyCount)
	}
}
