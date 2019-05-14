package goinput

import (
	"github.com/AndrewDonelson/go-input/keyboard"
	"github.com/AndrewDonelson/go-input/mouse"
)

var (
	// DefaultKeyboard allows immediate access to default keyboard input
	DefaultKeyboard = NewKeyboard()
	// Defaultmouse allows immediate access to default mouse input
	DefaultMouse = NewMouse()
)

func NewKeyboard() *keyboard.Watcher {
	return keyboard.NewWatcher()
}

func NewMouse() *mouse.Watcher {
	return mouse.NewWatcher()
}
