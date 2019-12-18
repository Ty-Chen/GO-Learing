package main
import "fmt"


func testUnblocking(){
		messages := make(chan string)
		signals := make(chan bool)
	
		select {
		case msg := <-messages:
			fmt.Println("received message", msg)
		default:
			fmt.Println("no message received")
		}
	
		msg := "hi"
		select {
		case messages <- msg:
			fmt.Println("sent message", msg)
		default:
			fmt.Println("no message sent")
		}
	
		select {
		case msg := <-messages:
			fmt.Println("received message", msg)
		case sig := <-signals:
			fmt.Println("received signal", sig)
		default:
			fmt.Println("no activity")
		}
}


func testMake() {
	message := make(chan string)
	
	go func() { message <- "ping"}()

	msg:= <- message
	fmt.Println(msg)
}

func testChannel(){
	//testMake();
	testUnblocking();
}