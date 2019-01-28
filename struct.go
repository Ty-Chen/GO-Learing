package main

import (
    "fmt"
    "math"
)

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

// 面积
func (c Circle) Area() float64 {
    return math.Pi * float64(c.Radius) * float64(c.Radius)
   }
   
// 周长
func (c Circle) Circumference() float64 {
    return 2 * math.Pi * float64(c.Radius)
}

func testFunc() {
    var c = Circle {Radius: 50}
    fmt.Println(c.Area(), c.Circumference())
    // 指针变量调用方法形式上是一样的
    var pc = &c
    fmt.Println(pc.Area(), pc.Circumference())
}

func testCircle()  {
    newCircle1();
    newCircle2();
    newCircle3();
    newCircle4();
    testFunc();
}