// file: keyboard/events.go
// description: This file contains the definition of the ButtonEvent and Typed types, which represent keyboard events.
//
// Copyright 2025 Andrew Donelson. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package keyboard

import (
	"fmt"
	"time"
)

// Event represents a general keyboard event interface
type Event interface {
	// Time returns the time at which this event occurred
	Time() time.Time

	// String returns a string representation of this event
	String() string
}

// ButtonEvent represents an event when a keyboard button changes state (i.e.
// being pushed down when it was previously up, or being toggled on when it was
// previously off, etc)
//
// If Key == Invalid then the key may not be known, but it can still be
// uniquely identified and it's state watched via the Raw member (e.g. for
// special or non-US keys).
//
// The Raw member must uniquely identify the keyboard button whose state is
// changing, and must always be present regardless of whether or not Key ==
// Invalid. It could (but does not have to be) e.g. the scancode of the key.
type ButtonEvent struct {
	T     time.Time
	Key   Key
	State State
	Raw   uint64
}

// Time returns the time at which this event occurred.
func (b ButtonEvent) Time() time.Time {
	return b.T
}

// String returns a string representation of this event.
func (b ButtonEvent) String() string {
	return fmt.Sprintf("ButtonEvent(Key=%v, State=%v, Raw=%v, Time=%v)", b.Key, b.State, b.Raw, b.T)
}

// Typed represents an event where some sort of user input has generated a
// string of text which should be considered as user input.
type Typed struct {
	T time.Time
	S string
}

// Time returns the time at which this event occurred.
func (t Typed) Time() time.Time {
	return t.T
}

// String simply returns the user input string.
func (t Typed) String() string {
	return t.S
}

// IsModifierKey returns true if the key is a modifier key (Shift, Ctrl, Alt, Super)
func IsModifierKey(k Key) bool {
	switch k {
	case LeftShift, RightShift, LeftCtrl, RightCtrl, LeftAlt, RightAlt, LeftSuper, RightSuper:
		return true
	default:
		return false
	}
}

// IsNavigationKey returns true if the key is used for navigation (arrows, page up/down, etc.)
func IsNavigationKey(k Key) bool {
	switch k {
	case ArrowUp, ArrowDown, ArrowLeft, ArrowRight, Home, End, PageUp, PageDown:
		return true
	default:
		return false
	}
}

// IsAlphaNumeric returns true if the key is a letter or number
func IsAlphaNumeric(k Key) bool {
	// Check if key is a letter
	if k >= A && k <= Z {
		return true
	}

	// Check if key is a number
	if k >= Zero && k <= Nine {
		return true
	}

	return false
}

// IsFunctionKey returns true if the key is a function key (F1-F25)
func IsFunctionKey(k Key) bool {
	return k >= F1 && k <= F25
}

// NewButtonEvent creates a new ButtonEvent with the current time
func NewButtonEvent(key Key, state State, raw uint64) ButtonEvent {
	return ButtonEvent{
		T:     time.Now(),
		Key:   key,
		State: state,
		Raw:   raw,
	}
}

// NewTyped creates a new Typed event with the current time
func NewTyped(s string) Typed {
	return Typed{
		T: time.Now(),
		S: s,
	}
}
