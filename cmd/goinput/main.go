// file: cmd/goinput/main.go
// description: Main entry point for the go-input CLI application

package main

import (
	"fmt"
	"os"

	goinput "github.com/AndrewDonelson/go-input"
	"github.com/AndrewDonelson/go-input/userinput"
)

var (
	version = "1.0.0" // Set by build flags
)

func main() {
	fmt.Printf("GoInput CLI v%s\n", version)
	fmt.Println("A simple demonstration of input handling capabilities")
	fmt.Println("---------------------------------------------------")

	// Example using the default keyboard and mouse watchers
	fmt.Println("Default watchers available:")
	fmt.Printf("- Keyboard: %v\n", goinput.DefaultKeyboard != nil)
	fmt.Printf("- Mouse: %v\n", goinput.DefaultMouse != nil)
	fmt.Println()

	// Example using the user input functionality
	ui := userinput.New()

	name, err := ui.ReadString("Enter your name: ")
	if err != nil {
		fmt.Println("Error reading name:", err)
		os.Exit(1)
	}
	fmt.Println("Hello,", name)

	age, err := ui.ReadInt("Enter your age: ")
	if err != nil {
		fmt.Println("Error reading age:", err)
		os.Exit(1)
	}
	fmt.Println("Your age is:", age)

	height, err := ui.ReadFloat("Enter your height (in meters): ")
	if err != nil {
		fmt.Println("Error reading height:", err)
		os.Exit(1)
	}
	fmt.Println("Your height is:", height, "meters")

	isStudent, err := ui.ReadBool("Are you a student? (y/n): ")
	if err != nil {
		fmt.Println("Error reading student status:", err)
		os.Exit(1)
	}
	fmt.Println("Student status:", isStudent)
}
