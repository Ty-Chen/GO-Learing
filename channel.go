package main
import "fmt"

func testChannel() {
	message := make(chan string)
	
	go func() { message <- "ping"}()

	msg:= <- message
	fmt.Println(msg)
}