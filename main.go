package main

import "fmt"

type bd interface {
	getBotVersion() float64
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
	fmt.Println(o.getBotVersion())
}
func (englishBot) getBotVersion() float64 {
	return 3.5
}

func (spanishBot) getBotVersion() float64 {
	return 3.6
}

func (eb englishBot) getGreeting() string {
	return "Hi there"
}

func (sb spanishBot) getGreeting() string {
	return "hola"
}
