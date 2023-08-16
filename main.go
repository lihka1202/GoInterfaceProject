package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// EMULATE CUSTOM LOGC
	return "Hello World!"
}
func (spanishBot) getGreeting() string {
	// EMULATE CUSTOM LOGIC
	return "Hola!"
}
func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}
