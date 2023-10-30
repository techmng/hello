package main

import (
	"fmt"
	"log"

	"github.com/techmng/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	var name string

	/*name = ""
	msg1, err := greetings.Hello(name)
	if err != nil {
		//log.Fatal(err)   Fatal will stop execution because in call exit(0) at end
		log.Print(err)
	}
	fmt.Println(msg1)*/

	name = "Manoj"
	msg2, err := greetings.Hello(name)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(msg2)
}
