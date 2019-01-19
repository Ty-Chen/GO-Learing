package main

import "fmt"

type Circle struct{
	x int
	y int
	Radius int
}

func initCircle() {
	var c Circle = Circle {
        x: 100,
        y: 100,
        Radius: 50,  // 注意这里的逗号不能少
    }
    fmt.Printf("%+v\n", c)
}

func testCircle()  {
	
}