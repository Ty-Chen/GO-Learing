package main

import (   
     "fmt"
    "unsafe"
)

func dict() {
    var m map[int]string = make(map[int]string)
    fmt.Println(m, len(m))
}

func dict2() {
	var m map[int]string = map[int]string {
        90: "优秀",
        80: "良好",
        60: "及格",  // 注意这里逗号不可缺少，否则会报语法错误		
	}
	fmt.Println(m, len(m))
}

func dict3() {
    var m = map[int]string {
        90: "优秀",
        80: "良好",
        60: "及格",  // 注意这里逗号不可缺少，否则会报语法错误	        
    }
    fmt.Println(m, len(m))
}

func dictAdd() {
    var fruits = map[string]int {
        "apple": 2,
        "banana": 5,
        "orange": 8,       
    }

     // 读取元素   
    var score = fruits["banana"]
    fmt.Println(score)

    // 增加或修改元素
    fruits["pear"] = 3
    fmt.Println(fruits)

    // 删除元素
    delete(fruits, "pear")
    fmt.Println(fruits)    
}

func dictReturn() {
    var fruits = map[string]int {
        "apple": 2,
        "banana": 5,
        "orange": 8,
    }

    var score, ok = fruits["durin"]
    if ok {
        fmt.Println(score)
    } else {
        fmt.Println("durin not exists")
    }

    fruits["durin"] = 0
    score, ok = fruits["durin"]
    if ok {
        fmt.Println(score)
    } else {
        fmt.Println("durin still not exists")
    }
}

func dictForEach(){
    
    var fruits = map[string]int {
        "apple": 2,
        "banana": 5,
        "orange": 8,
    }

    for name, score := range fruits {
        fmt.Println(name, score)
    }

    for name := range fruits {
        fmt.Println(name)
    }   
}

func dictSize() {

    var fruits = map[string]int {
        "apple": 2,
        "banana": 5,
        "orange": 8,        
    }

    fmt.Println(unsafe.Sizeof(fruits))
}

func testDict()  {
    dict();
    dict2();
    dict3();
    dictAdd();
    dictReturn();
}