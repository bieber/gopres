package example

import (
	"fmt"
)

// Divide divides seven by three and prints the result to the screen, twice.
func Divide() {
	fmt.Println(div1(7, 3))
	fmt.Println(div2(7, 3))
}

func div1(n int, d int) (int, int) {
	return n / d, n % d
}

func div2(n int, d int) (result int, remainder int) {
	result = n / d
	remainder = n % d
	return
}
