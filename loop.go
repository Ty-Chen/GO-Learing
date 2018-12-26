package main

import "fmt"

func unlimitedLoop() {
    for {
        fmt.Println("hello world!")
    }
}

func unlimitedLoop2() {
    for true {
        fmt.Println("hello world!")
    }
}

func limitedLoop()  {
	for i := 1; i < 10; i++ {
		fmt.Println("hellow, world")
	}
}