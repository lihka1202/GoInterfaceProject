package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}
func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string {
	// EMULATE CUSTOM LOGC
	return "Hello World!"
}
func (spanishBot) getGreeting() string {
	// EMULATE CUSTOM LOGC
	return "Hola!"
}
func main() {

}
