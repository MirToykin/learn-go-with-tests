package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "

	spanishLanguage = "Spanish"
	frenchLanguage  = "French"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefixLanguage(lang) + name
}

func greetingPrefixLanguage(lang string) (prefix string) {
	switch lang {
	case spanishLanguage:
		prefix = spanishHelloPrefix
	case frenchLanguage:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
