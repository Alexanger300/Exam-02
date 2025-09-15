package main

import "fmt"

func IsSorted(f func(int, int) int, tab []int) bool {
	ascending := true
	descending := true
	for i := 0; i < len(tab)-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			ascending = false
		}
		if f(tab[i], tab[i+1]) < 0 {
			descending = false
		}
	}
	return ascending || descending
}

func f(a int, b int) int {
	if a-b > 0 {
		return 1
	}
	if a-b < 0 {
		return -1
	}
	return 0
}

func main() {
	a1 := []int{1, 2, 3, 4}
	a2 := []int{0, 2, 1, 3}
	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)
	fmt.Println(result1)
	fmt.Println(result2)
}
