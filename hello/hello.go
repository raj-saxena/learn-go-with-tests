package main

import "fmt"

func main() {
	fmt.Println(Hello("Raj", ""))
}

func Hello(name, language string) string {
	if len(name) == 0 {
		name = "World"
	}
	greeting := greeting(language)
	return fmt.Sprintf("%s, %s!", greeting, name)
}

func greeting(language string) (greeting string) {
	switch language {
	case "French":
		greeting = "Bonjour"
	default:
		greeting = "Hello"
	}
	return
}
