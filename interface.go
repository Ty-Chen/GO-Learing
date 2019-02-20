package main

import "fmt"

type Smellable interface {
	smell()
}

type Eatable interface {
	eat()
}

type Fruit struct {
    Name string  // 属性变量
    Eatable  // 匿名内嵌接口变量
}

func (f Fruit) want() {
    fmt.Printf("I like ")
    f.eat() // 外结构体会自动继承匿名内嵌变量的方法
}

type Apple struct {}

func (a Apple) smell() {
	fmt.Println("apple can smell")
}

func (a Apple) eat() {
	fmt.Println("apple can eat")
}

type Banana struct {}

func (b Banana) eat() {
    fmt.Println("eating banana")
}

type Flower struct {}

func (f Flower) smell() {
	fmt.Println("flower can only smell")
}

func testDiffInterface() {
	var s1 Smellable
	var s2 Eatable
	var apple = Apple{}
	var flower = Flower{}
	s1 = apple
	s1.smell()
	s1 = flower
	s1.smell()
	s2 = apple
	s2.eat()	
}

func testEmptyInterface() {
	
	var user = map[string]interface{}{
		"age": 30,
		"address": "Beijing Tongzhou",
		"married": true,
	}
	fmt.Println(user)
	
	var age = user["age"].(int)
	var address = user["address"].(string)
	var married = user["married"].(bool)
	fmt.Println(age, address, married)	
}

func testMultiDeplo() {
    var f1 = Fruit{"Apple", Apple{}}
    var f2 = Fruit{"Banana", Banana{}}
    f1.want()
    f2.want()	
}

func testInterface()  {
	testDiffInterface();
	testEmptyInterface();
	testMultiDeplo();
}
