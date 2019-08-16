package main

import "fmt"

// Hello function to return the greeting
func Hello(name string) string {
	return "Hello, " + name
}
func main() {
	fmt.Println(Hello("Jerry"))
}