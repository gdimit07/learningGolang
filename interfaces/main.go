package main

import "fmt"

type bot interface {
	getGreeting() string
	askForAssistance() string
	provideOrigin() string
}

type englishBot struct {
	origin string
}
type spanishBot struct {
	origin string
}

func (englishBot) getGreeting() string {
	return ("hello there")
}

func (spanishBot) getGreeting() string {
	return ("hola")
}

func (englishBot) askForAssistance() string {
	return ("do you need help?")
}

func (spanishBot) askForAssistance() string {
	return ("necesitas ayuda con algo?")
}

func (eb englishBot) provideOrigin() string {
	return eb.origin
}

func (sb spanishBot) provideOrigin() string {
	return sb.origin
}

func provideHelp(b bot) {
	fmt.Println(b.askForAssistance())
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func printOrigin(b bot) {
	fmt.Println(b.provideOrigin())
}

func main() {

	eb := englishBot{origin: "United Kingdom"}
	sb := spanishBot{origin: "Spain"}

	printGreeting(eb)
	printGreeting(sb)

	provideHelp(sb)

	printOrigin(eb)
}
