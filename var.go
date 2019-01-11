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

	//指针
	var ptr *int = &s;
	fmt.Println("ptr =  ", *ptr)

	fmt.Println("global var = ", globali)
}


func variables() {
    // 有符号整数，可以表示正负
    var a int8 = 1 // 1 字节
    var b int16 = 2 // 2 字节
    var c int32 = 3 // 4 字节
    var d int64 = 4 // 8 字节
    fmt.Println(a, b, c, d)

    // 无符号整数，只能表示非负数
    var ua uint8 = 1
    var ub uint16 = 2
    var uc uint32 = 3
    var ud uint64 = 4
    fmt.Println(ua, ub, uc, ud)

    // int 类型，在32位机器上占4个字节，在64位机器上占8个字节
    var e int = 5
    var ue uint = 5
    fmt.Println(e, ue)

    // bool 类型
    var f bool = true
    fmt.Println(f)

    // 字节类型
    var j byte = 'a'
    fmt.Println(j)

    // 字符串类型
    var g string = "abcdefg"
    fmt.Println(g)

    // 浮点数
    var h float32 = 3.14
    var i float64 = 3.141592653
    fmt.Println(h, i)
}

func testVars() {
    vars();
    variables();
}