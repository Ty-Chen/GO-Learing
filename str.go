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

func testStr() {
	strForEachChar();
	strForEachByte();
}