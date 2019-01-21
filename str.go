package main

import "fmt"

func strForEachChar() {

    var s = "嘻哈china"
    for codepoint, runeValue := range s {
        fmt.Printf("%d %d ", codepoint, int32(runeValue))
    }

}

func strForEachByte() {

	var s = "嘻哈china"	
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }	
}

func cutStr() {
    var s1 = "hello world"
    var s2 = s1[3:8]
    fmt.Println(s2)
}

func changeStr() {
    var s1 = "hello world"
    var b = []byte(s1)  // 字符串转字节切片
    var s2 = string(b)  // 字节切片转字符串
    fmt.Println(b)
    fmt.Println(s2)
}

func testStr() {
	strForEachChar();
    strForEachByte();
    cutStr();
    changeStr();
}