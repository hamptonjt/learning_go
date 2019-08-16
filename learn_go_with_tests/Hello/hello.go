package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello function to return the greeting
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("Jerry"))
}
