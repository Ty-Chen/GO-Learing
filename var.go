package main

import ("fmt")

var globali int = 24

func vars(){	

	//显示赋值，用于强调
	var s int = 42
	fmt.Println("s = ", s)

	//去掉int自动推导
	var d = 42
	fmt.Println("d = ", d)

	//自动类型推导+赋值
	e := 42
	fmt.Println("e = ", e)

	fmt.Println("global var = ", globali)
}
