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

func loopSlice() {	
	var s = []int{1, 2, 3, 4, 5}
	for index := range s {
		fmt.Println(index, s[index])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}

func appendSlice() {
	var s1 = []int{1, 2, 3, 4, 5}
	fmt.Println(s1, len(s1), cap(s1))

	 // 对满容的切片进行追加会分离底层数组
	 var s2 = append(s1, 6)
	 fmt.Println(s1, len(s1), cap(s1))
	 fmt.Println(s2, len(s2), cap(s2))
	
	 // 对非满容的切片进行追加会共享底层数组
	 var s3 = append(s2, 7)
	 fmt.Println(s2, len(s2), cap(s2))
	 fmt.Println(s3, len(s3), cap(s3))
}

func cutSlice() {
	var s1 = []int{1, 2, 3, 4, 5, 6, 7}
	var s2 = s1[:5]
	var s3 = s1[3:]
	var s4 = s1[:]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s4, len(s4), cap(s4))
}

func deepCopySlice() {
	var s = make([]int, 5, 8)
	for i := 0; i < len(s); i++ {
	 s[i] = i+1
	}
	fmt.Println(s)
	var d = make([]int, 2, 6)
	var n = copy(d, s)
	fmt.Println(n, d)
}