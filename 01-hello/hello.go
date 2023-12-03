package main

import arrays_and_slices "github.com/MirToykin/go-with-tests/04-arrays_and_slices"

func main() {
	//fmt.Println(Hello("world", ""))
	arrays_and_slices.AffectingSlicedArrayExample()
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
