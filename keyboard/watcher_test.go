// file: keyboard/watcher_test.go
// description: Test cases for the keyboard watcher functionality.
//
// Copyright 2025 Andrew Donelson. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package keyboard

import (
	"strings"
	"testing"
)

func TestWatcher(t *testing.T) {
	m := NewWatcher()
	m.SetState(A, Down)
	m.SetState(ArrowLeft, Up)
	if !m.Down(A) {
		t.Fatal("expect keyboard.A in state keyboard.Down")
	}
	if !m.Up(ArrowLeft) {
		t.Fatal("expect keyboard.ArrowLeft in state keyboard.Up")
	}
	if !m.Up(Escape) {
		t.Fatal("expect keyboard.Escape in state keyboard.Up")
	}

	// Verify the state lookup table.
	want := map[Key]State{
		A:         Down,
		ArrowLeft: Up,
		Escape:    Up,
	}
	states := m.States()
	if len(states) != 2 { // Only 2 states were set explicitly, Escape is implicitly Up
		t.Fatalf("got %d states, want %d\n", len(states), 2)
	}
	for key, state := range states {
		wantState, exists := want[key]
		if !exists {
			t.Fatalf("unexpected key in states: %v\n", key)
		}
		if wantState != state {
			t.Fatalf("got %v=%v, want %v=%v\n", key, state, key, wantState)
		}
	}
}

func TestWatcherString(t *testing.T) {
	m := NewWatcher()
	m.SetState(A, Down)
	m.SetState(B, Up)

	str := m.String()
	if !strings.Contains(str, "keyboard.Watcher") {
		t.Fatal("String() should contain 'keyboard.Watcher'")
	}
	if !strings.Contains(str, "A: Down") {
		t.Fatal("String() should contain 'A: Down'")
	}
	if !strings.Contains(str, "B: Up") {
		t.Fatal("String() should contain 'B: Up'")
	}
}

func TestWatcherEachState(t *testing.T) {
	m := NewWatcher()
	m.SetState(A, Down)
	m.SetState(B, Up)
	m.SetState(C, Down)

	// Count the number of keys in Down state
	downCount := 0
	m.EachState(func(k Key, s State) bool {
		if s == Down {
			downCount++
		}
		return true
	})
	if downCount != 2 {
		t.Fatalf("Expected 2 keys in Down state, got %d", downCount)
	}

	// Test early termination
	keyCount := 0
	m.EachState(func(k Key, s State) bool {
		keyCount++
		return false // Stop after first key
	})
	if keyCount != 1 {
		t.Fatalf("Expected EachState to process 1 key before stopping, got %d", keyCount)
	}
}

func TestWatcherRawState(t *testing.T) {
	m := NewWatcher()

	// Test setting and retrieving raw states
	m.SetRawState(123, Down)
	m.SetRawState(456, Up)

	if !m.RawDown(123) {
		t.Fatal("Expected raw key 123 to be in Down state")
	}

	if !m.RawUp(456) {
		t.Fatal("Expected raw key 456 to be in Up state")
	}

	// Test default state for unknown raw keys
	if m.RawState(789) != Up {
		t.Fatal("Expected unknown raw key to be in Up state")
	}

	// Test RawStates
	rawStates := m.RawStates()
	if len(rawStates) != 2 {
		t.Fatalf("Expected 2 raw states, got %d", len(rawStates))
	}
	if rawStates[123] != Down {
		t.Fatalf("Expected raw key 123 to be Down in the map, got %v", rawStates[123])
	}
	if rawStates[456] != Up {
		t.Fatalf("Expected raw key 456 to be Up in the map, got %v", rawStates[456])
	}
}
