package main

import "fmt"

func makeSlice() {
	var s1 []int = make([]int, 5, 8)
	var s2 []int = make([]int, 8) // 满容切片
	fmt.Println(s1)
	fmt.Println(s2)
   }

func makeSlice2()  {
	s1 := make([]int, 5, 8)
	s2 := make([]int, 8)
	fmt.Println(s1)
	fmt.Println(s2)
}

func initSlice() {	
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1, len(s1), cap(s1))

}

func nullSlice() {	
	var s1 []int
	var s2 []int = []int{}
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	fmt.Println(len(s1), len(s2), len(s3))
	fmt.Println(cap(s1), cap(s2), cap(s3))	
}

func setSlice() {	
	var s1 = make([]int, 5, 8)
	// 切片的访问和数组差不多
	for i := 0; i < len(s1); i++ {
	 s1[i] = i + 1
	}
	var s2 = s1
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
   
	// 尝试修改切片内容
	s2[0] = 255
	fmt.Println(s1)
	fmt.Println(s2)
}