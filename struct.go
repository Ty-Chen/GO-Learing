package main

import "fmt"

type Circle struct{
	x int
	y int
	Radius int
}

func newCircle1() {
	var c Circle = Circle {
        x: 100,
        y: 100,
        Radius: 50,  // 注意这里的逗号不能少
    }
    fmt.Printf("%+v\n", c)
}

func newCircle2() {
    var c1 Circle = Circle {
    Radius: 50,
}
    var c2 Circle = Circle {}
    fmt.Printf("%+v\n", c1)
    fmt.Printf("%+v\n", c2) 
}

func testCircle()  {
    newCircle1();
    newCircle2();
}