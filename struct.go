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

func newCircle3() {
    var c Circle = Circle {100, 100, 50}
    fmt.Printf("%+v\n", c)   
}

func newCircle4() {
    var c1 Circle = Circle {Radius: 50}
    var c2 Circle = c1
    fmt.Printf("%+v\n", c1)
    fmt.Printf("%+v\n", c2)
    c1.Radius = 100
    fmt.Printf("%+v\n", c1)
    fmt.Printf("%+v\n", c2)

    var c3 *Circle = &Circle {Radius: 50}
    var c4 *Circle = c3
    fmt.Printf("%+v\n", c3)
    fmt.Printf("%+v\n", c4)
    c3.Radius = 100
    fmt.Printf("%+v\n", c3)
    fmt.Printf("%+v\n", c4)
}

func testCircle()  {
    newCircle1();
    newCircle2();
    newCircle3();
    newCircle4();
}