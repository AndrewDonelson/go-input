package goinput

import (
	"testing"
)

func TestPollingModes(t *testing.T) {
	if Eco != 0 {
		t.Errorf("Expected Eco to be 0, got %d", Eco)
	}
	if Normal != 1 {
		t.Errorf("Expected Normal to be 1, got %d", Normal)
	}
	if Game != 2 {
		t.Errorf("Expected Game to be 2, got %d", Game)
	}
}

func TestNewKeyboard(t *testing.T) {
	kb := NewKeyboard()
	if kb == nil {
		t.Fatal("NewKeyboard() returned nil")
	}
}

func TestNewMouse(t *testing.T) {
	m := NewMouse()
	if m == nil {
		t.Fatal("NewMouse() returned nil")
	}
}

func TestDefaultKeyboardAndMouse(t *testing.T) {
	if DefaultKeyboard == nil {
		t.Fatal("DefaultKeyboard is nil")
	}
	if DefaultMouse == nil {
		t.Fatal("DefaultMouse is nil")
	}
}

func TestAskQuestionString(t *testing.T) {
	a := DefaultAsk
	answer := a.ForString("What is your name")
	t.Logf("answer = %v", answer)
}
