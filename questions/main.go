package questions

import (
	"fmt"

	"github.com/AndrewDonelson/go-input/keyboard"
	"github.com/AndrewDonelson/go-input/mouse"
)

type Ask struct {
	Keyboard *keyboard.Watcher
	Mouse    *mouse.Watcher
}

var DefaultAsk = Ask{}

func NewAsk(k *keyboard.Watcher, m *mouse.Watcher) *Ask {
	return &Ask{
		Keyboard: k,
		Mouse:    m,
	}
}

func (a *Ask) ForString(question string) string {

	// Display question string and wait for keyboard input.
	fmt.Printf("%s? ", question)
	n, err := fmt.Scanln()
	if err != nil {
		return ""
	}
	fmt.Println(n)

	return question
}

func (a *Ask) ForInt(question string) int {
	return 0
}

func (a *Ask) ForFloat(question string) float64 {
	return 0.0
}

func (a *Ask) ForBool(question string) bool {
	return false
}

func (a *Ask) ForChoice(question string, choices []string) string {
	return ""
}
