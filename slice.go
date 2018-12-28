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