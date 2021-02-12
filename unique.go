package main

import (
	"fmt"
)

func main(){
	var a= []int{1,1,9,9,5,7,7,5,8,3,3}
	fmt.Println(unique1(a))
}
func unique1(a []int) int{
	var unique = a[0]
	for i := 1; i < len(a); i++ {
		unique = unique ^ a[i]
	}
	return unique
}
