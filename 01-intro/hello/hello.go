package main

import "fmt"

var name = "Alexei"

func main() {
	greeting(name)
}

func greeting(msg string) {
	fmt.Printf("Hello, %s", msg)
}
