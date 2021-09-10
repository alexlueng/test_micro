package hello

import "fmt"

type Hello struct {
	Word string
	Name []Greeting
}

func SayHello() {
	fmt.Println("Hello!!!!!")
}
