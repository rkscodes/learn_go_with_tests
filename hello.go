package main

import "fmt"

const salutation = "Hello"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return salutation + ", " + name + "!"
}

func main() {
	fmt.Println(Hello("Ram"))
}
