package main

import (
	"Slice/function"
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("origin : %v\n", s)

	num, s := function.Pop(s)
	fmt.Println("after Pop")
	fmt.Printf("pop : %d, s : %v\n", num, s)

	s = function.Insert2(s, 3, 10, 11, 12, 13, 14)
	fmt.Println("after Insert2 ...  insert 10,11,12,13,14 in index 3")
	fmt.Println(s)

	s = function.Remove(s, 3)
	fmt.Println("after Remove index 3")
	fmt.Println(s)

	s1 := function.Copy(s)
	fmt.Println("copy s to s1 and s1[0] = 1000")
	s1[0] = 1000
	fmt.Printf("s : %v, s1 : %v", s, s1)

}
