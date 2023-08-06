package main

import (
	"dsa/datastructures"
	"fmt"
)

func main() {
	fmt.Println("hello word")
	arr1 := datastructures.NewMapArrayFormMap[int](
		10,
		1, 2, 3, 4, 5, 5, 6, 3,
	)
	arr2 := datastructures.NewMapArray[int](10)
	arr1.Push(1)
	arr2.Push(1)
	fmt.Println(arr1.ToString(), arr2.ToString())
	list := datastructures.NewMapArrayFormMap[int](
		10,
		1, 2, 3, 4, 5, 6, 7, 2, 6, 8, 7, 7, 2, 6, 8, 7,
	)
	fmt.Println(list.Capacity())
}
