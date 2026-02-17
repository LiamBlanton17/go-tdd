package main

import "fmt"

const englishGreeting = "Hello"
const spanishGreeting = "Hola"
const frenchGreeting = "Bonjour"
const spanish = "Spanish"
const french = "French"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishGreeting
	switch language {
	case spanish:
		prefix = spanishGreeting
	case french:
		prefix = frenchGreeting
	}

	return prefix + ", " + name + "!"
}

func main() {
	fmt.Println(Hello("World", ""))
}
