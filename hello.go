package main

import "fmt"

const (
	englishGreeting = "Hello"
	spanishGreeting = "Hola"
	frenchGreeting  = "Bonjour"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greeting(language) + ", " + name + "!"
}

func greeting(language string) (salutation string) {

	switch language {
	case "English", "english", "":
		salutation = englishGreeting
	case "Spanish", "spanish":
		salutation = spanishGreeting
	case "French":
		salutation = frenchGreeting
	default:
		salutation = englishGreeting
	}

	return
}

func main() {
	fmt.Println(Hello("Ram", ""))
}
