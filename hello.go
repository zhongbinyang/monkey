package main

import (
	"example.com/hello/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)

}
