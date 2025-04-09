// file: mouse/watcher.go
// description: Implementation of the mouse watcher, which monitors mouse input.
//
// Copyright 2025 Andrew Donelson. All rights reserved.
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package mouse

import (
	"bytes"
	"fmt"
	"sync"
)

// Watcher watches the state of various mouse buttons and their states.
type Watcher struct {
	access sync.RWMutex

	// states is a (at max 8-bit) lookup table, where the indexes are literally
	// Button values.
	states []State
}

// String returns a multi-line string representation of this mouse watcher and
// it's associated states.
func (w *Watcher) String() string {
	w.access.RLock()
	defer w.access.RUnlock()

	bb := new(bytes.Buffer)
	fmt.Fprintf(bb, "mouse.Watcher(\n")
	for b, s := range w.states {
		if s == InvalidState {
			continue
		}
		fmt.Fprintf(bb, "\t%v: %v,\n", Button(b), s)
	}
	fmt.Fprintf(bb, ")")
	return bb.String()
}

// SetState specifies the current state of the specified mouse button.
func (w *Watcher) SetState(button Button, state State) {
	w.access.Lock()
	defer w.access.Unlock()

	// If the state lookup table is too small to contain the button, expand it.
	if len(w.states) < int(button)+1 {
		oldStates := w.states
		w.states = make([]State, int(button)+1)
		copy(w.states, oldStates)
	}

	w.states[button] = state
}

// States returns an copy of the internal mouse button state lookup table used
// by this watcher. The indices of the lookup table are literally Button
// values:
//
//	states := watcher.States()
//	leftState := states[mouse.Left]
//	if leftState != InvalidState {
//	    fmt.Println("The left mouse button state is", leftState)
//	}
//
// States for buttons not known to the watcher are equal to InvalidState.
//
// At max the lookup table will be of length 256 (as Button is declared as a
// uint8), but it may be less.
func (w *Watcher) States() []State {
	w.access.RLock()
	defer w.access.RUnlock()

	cpy := make([]State, len(w.states))
	copy(cpy, w.states)
	return cpy
}

// EachState calls f with each known button to this watcher and it's current
// button state. It does so until the function returns false or there are no
// more buttons known to the watcher.
func (w *Watcher) EachState(f func(b Button, s State) bool) {
	w.access.RLock()
	defer w.access.RUnlock()

	for b, state := range w.states {
		button := Button(b)
		// Skip the Invalid button and buttons with InvalidState
		if button == Invalid || state == InvalidState {
			continue
		}

		// Call the function without the lock being held, so they can access
		// methods on this watcher still.
		w.access.RUnlock()
		cont := f(button, state)
		w.access.RLock()

		if !cont {
			return
		}
	}
}

// State returns the current state of the specified mouse button.
func (w *Watcher) State(button Button) State {
	w.access.RLock()
	defer w.access.RUnlock()

	// If the lookup table isn't large enough to contain the button's state, we
	// are not aware of it so it's in the Up state.
	b := int(button)
	if b >= len(w.states) {
		return Up
	}

	// Return the actual state, including InvalidState if that's what's set
	return w.states[b]
}

// Down tells whether the specified mouse button is currently in the down
// state.
func (w *Watcher) Down(button Button) bool {
	return w.State(button) == Down
}

// Up tells whether the specified mouse button is currently in the up state.
func (w *Watcher) Up(button Button) bool {
	state := w.State(button)
	// InvalidState is treated as Up for the purpose of this method
	return state == Up || state == InvalidState
}

// NewWatcher returns a new, initialized, mouse watcher.
func NewWatcher() *Watcher {
	w := new(Watcher)
	w.states = make([]State, 8)
	return w
}
