package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
	var start string
	var tab []string
	for i := range s {
		start = start + string(s[i])
		if string(s[i]) == " " {
			tab = append(tab, start)
			start = string(s[i])
		}
	}
	return tab
}

func main() {
	fmt.Println(SplitWhiteSpaces("Hello World"))
}

//J'ai abandonné aussi
