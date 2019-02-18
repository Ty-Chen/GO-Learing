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

func (a Apple) eat() {
	fmt.Println("apple can eat")
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
	// 连续两个大括号，是不是看起来很别扭
	var user = map[string]interface{}{
		"age": 30,
		"address": "Beijing Tongzhou",
		"married": true,
	}
	fmt.Println(user)
	// 类型转换语法来了
	var age = user["age"].(int)
	var address = user["address"].(string)
	var married = user["married"].(bool)
	fmt.Println(age, address, married)	
}

func testInterface()  {
	testDiffInterface();
	testEmptyInterface();
}
