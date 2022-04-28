package main

import (
	"fmt"
	"go-hello-tdd/di"
	"go-hello-tdd/mocking"
	"os"
)

const portuguese, french, spanish = "portuguese", "french", "spanish"
const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetigPrefix(language) + name + "!"
}

func greetigPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Thiago", portuguese))

	origArray := [2]int{1, 2}

	mySlice := origArray[:]
	mySlice = append(mySlice, 3)
	origArray[0] = 3
	fmt.Println(mySlice, origArray)

	di.GreetToConsole()
	mocking.Countdown(os.Stdout)
}
