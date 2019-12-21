package main

import(
	"fmt"
	"time"
)

func testTimer() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 expired")
	
	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2:= timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stop")
	}
}

func testTicker() {
	ticker:= time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
	}()
	
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Tickers stopped")
}