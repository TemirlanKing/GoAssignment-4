package main

import "fmt"

func main() {
	slice := []int{5, 2, 9, 1, 7, 4}

	fmt.Println("Original slice:")
	PrintSlice(slice)

	fmt.Println("Sorted slice:")
	SortSlice(slice)
	PrintSlice(slice)

	fmt.Println("Incremented odd positions:")
	IncrementOdd(slice)
	PrintSlice(slice)

	fmt.Println("Reversed slice:")
	ReverseSlice(slice)
	PrintSlice(slice)

	fmt.Println("")

	sortFunc := func(slice []int) {
		fmt.Println("Sorting slice:", slice)
	}

	printFunc := func(slice []int) {
		fmt.Println("Printing slice:", slice)
	}

	dstFunc := func(slice []int) {
		fmt.Println("Original processing:", slice)
	}
	newFunc := appendFunc(dstFunc, sortFunc, printFunc)

	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	newFunc(nums)
}

func SortSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func IncrementOdd(slice []int) {
	for i := 1; i < len(slice); i += 2 {
		slice[i]++
	}
}

func PrintSlice(slice []int) {
	fmt.Println(slice)
}

func ReverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(slice []int) {
		dst(slice)
		for _, f := range src {
			f(slice)
		}
	}
}
