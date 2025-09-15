package main

import "fmt"

func CountIf(f func(string) bool, a []string) int { //Compte le nombre de "true" dans le tableau
	count := 0
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			count++
		}
	}
	return count
}

func IsNumber(a string) bool { //Vérifie si le caractère est compris entre 0 et 9
	if a > "0" && a < "9" {
		return true
	}
	return false
}

func main() {
	a1 := []string{"Hello", "World", "1", "4", "Melvin le bg"}
	result1 := CountIf(IsNumber, a1)
	fmt.Println(result1)
}
