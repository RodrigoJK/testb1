package main

import "github.com/01-edu/z01"

func main() {
	var a int = 97
	for i := 0; i < 26; i++ {
		z01.PrintRune(rune(a + i))
	}
	z01.PrintRune('\n')
}
