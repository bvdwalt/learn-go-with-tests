package main

import "fmt"

const (
	english            = "English"
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	s := greetingPrefix(language)
	return s + name
}

func greetingPrefix(language string) string {
	if language == "" {
		language = english
	}

	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
