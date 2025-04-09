# Go-Input

A Go package for handling user input from various sources (keyboard, mouse, console).

## Features

- Keyboard input monitoring
- Mouse input monitoring
- Console-based user input for strings, integers, floats, booleans, and choice selections
- Multiple polling rate options (Eco, Normal, Game) for different application needs

## Installation

```bash
go get github.com/AndrewDonelson/go-input
```

## Usage

### Basic Console Input

```go
package main

import (
    "fmt"
    "github.com/AndrewDonelson/go-input/userinput"
)

func main() {
    ui := userinput.New()
    
    name, err := ui.ReadString("Enter your name")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    age, err := ui.ReadInt("Enter your age")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
```

### Using Keyboard and Mouse Monitoring

```go
package main

import (
    "fmt"
    "time"
    "github.com/AndrewDonelson/go-input"
    "github.com/AndrewDonelson/go-input/keyboard"
    "github.com/AndrewDonelson/go-input/mouse"
)

func main() {
    kb := goinput.NewKeyboard()
    m := goinput.NewMouse()
    
    // Set up a key state
    kb.SetState(keyboard.A, keyboard.Down)
    
    // Check a key state
    if kb.Down(keyboard.A) {
        fmt.Println("Key A is down")
    }
    
    // Set up a mouse button state
    m.SetState(mouse.Left, mouse.Down)
    
    // Check a mouse button state
    if m.Down(mouse.Left) {
        fmt.Println("Left mouse button is down")
    }
}
```

## License

BSD-style - See the LICENSE file for details.