package main

import "fmt"

type Smellable interface {
	smell()
}

type Eatable interface {
	eat()
}

type Apple struct {}

func (a Apple) smell() {
	fmt.Println("apple can smell")
}

func testInterface()  {
	
}