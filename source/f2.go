package prac

import "fmt"

func Closz() {
	message := "Hello, "

	greet := func(name string) {
		fmt.Println(message + name)
	}

	greet("Alice")
	greet("Bob")
}

// An anonymous function greet is created that takes a name parameter and appends it to the message.
// This function uses a variable message from the surrounding scope, which is accessible due to closure???
// The greet function is called twice with different names, and each time it prints a custom greeting
// combining the message and the provided name.
