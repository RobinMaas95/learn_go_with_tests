package main

import "fmt"

const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const germanHelloPrefix = "Hallo,"

func Hello(name, language string) string {
	prefix := greetingPrefix(language)
	if len(name) == 0 {
		name = "World"
	}
	return fmt.Sprintf("%s %s", prefix, name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "German":
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Robin", "English"))
}
