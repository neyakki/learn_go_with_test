package main

import (
	"fmt"
)

const (
	englishPrefix = "Hello"
	spanishPrefix = "Hola"
	frenchPrefix  = "Bonjour"
)

func Hello(name string, lang string) string {
	prefix := getPrefix(lang)

	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s, %s!", prefix, name)
}

func getPrefix(lang string) (prefix string) {
	switch lang {
	case "Spanish":
		prefix = spanishPrefix
	case "French":
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Neyakki", ""))
}
