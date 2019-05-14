package main

import (
	"github.com/AndrewDonelson/go-input/keyboard"
)

var DefaultKeyboard *keyboard.Watcher = NewKeyboard()

func NewKeyboard() *keyboard.Watcher {
	return NewWatcher()
}
