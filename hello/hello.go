package main

import (
	"fmt"

	"github.com/Dinosky/gotut/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
