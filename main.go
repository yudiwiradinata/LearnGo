package main

import "fmt"

type bd interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(o bd) {
	fmt.Println(o.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi there"
}

func (sb spanishBot) getGreeting() string {
	return "hola"
}
