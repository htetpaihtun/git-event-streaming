package main

import (
	"fmt"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func main() {
	message := Hello("August")
	print(message)
	print("We will have new commit")
}
