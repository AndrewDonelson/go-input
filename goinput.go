// file: goinput.go
// description: Main package file containing global variables and polling constants

package goinput

import (
	"github.com/AndrewDonelson/go-input/keyboard"
	"github.com/AndrewDonelson/go-input/mouse"
	"github.com/AndrewDonelson/go-input/questions"
)

// Polling represents the update speed for detecting input.
type Polling uint8

// Polling mode constants, Eco will check the input ten times per second.
// Normal will check thirty times per seconds and Game will update at sixty
// times per second. Default is Eco which is suggested for standard console
// or user input that does not require a high update interval.
const (
	Eco    Polling = iota // 10x/sec
	Normal                // 30x/sec
	Game                  // 60x/sec
)

// InputState holds the current state of the keyboard and mouse input.
type InputState struct {
	Keyboard keyboard.State
	Mouse    mouse.State
}

var (
	// DefaultKeyboard allows immediate access to default keyboard input
	DefaultKeyboard = NewKeyboard()

	// DefaultMouse allows immediate access to default mouse input
	DefaultMouse = NewMouse()

	// DefaultAsk provides a default instance of the questions.Ask struct
	DefaultAsk = questions.New(DefaultKeyboard, DefaultMouse)
)

// NewKeyboard returns a new keyboard watcher that can be used to monitor keyboard input.
// The returned watcher is configured with the default polling mode of Eco.
func NewKeyboard() *keyboard.Watcher {
	return keyboard.NewWatcher()
}

// NewMouse returns a new mouse watcher that can be used to monitor mouse input.
// The returned watcher is configured with the default polling mode of Eco.
func NewMouse() *mouse.Watcher {
	return mouse.NewWatcher()
}
