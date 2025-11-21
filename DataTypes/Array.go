package main

import (
	"fmt"
	"sort"
)

func main() {

	var arr [5]int
	fmt.Println(arr)

	arr[0] = 100
	fmt.Println(arr)

	arrNew := [4]int{1, 2, 3, 4}
	fmt.Println(arrNew)

	for i := 0; i < len(arrNew); i++ {
		fmt.Println(arrNew[i])
	}

	for _, v := range arrNew {
		fmt.Println(v)
	}

	fmt.Println(arrNew[0:2]) // [1 2]
	fmt.Println(arrNew[:2])  // [1 2]
	fmt.Println(arrNew[2:])  // [3 4]
	fmt.Println(arrNew[:])   // [1 2 3 4]

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [4]int{1, 2, 3, 5}
	fmt.Println(arr1 == arr2) // true
	fmt.Println(arr1 == arr3) // false

	arrSort := [4]int{4, 3, 2, 1}
	sort.Ints(arrSort[:])
	fmt.Println(arrSort)

	arr4 := [4]int{1, 2, 3, 4}
	arr4[1] = 0
	fmt.Println(arr4)
	fmt.Println(arr4[1:])

	//Search
	arr99 := [4]int{1, 2, 3, 4}
	i := sort.Search(len(arr99), func(i int) bool { return arr99[i] >= 3 })
	if i < len(arr99) && arr99[i] == 3 {
		fmt.Println("found 3 at index", i) // found 3 at index 2clear
	}

	// Array Reverse

	arr100 := [4]int{1, 2, 3, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(arr100[:])))
	fmt.Println(arr100)

	// Multi dimensional Array

	arr111 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr111[0][0]) // 1
	fmt.Println(arr111[0][1]) // 2
	fmt.Println(arr111[0][2]) // 3
	fmt.Println(arr111[1][0]) // 4
	fmt.Println(arr111[1][1]) // 5
	fmt.Println(arr111[1][2]) // 6

}
