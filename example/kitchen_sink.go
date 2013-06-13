package example

import (
	"fmt"
)

// KitchenSink demonstrates some other crap.
func KitchenSink() {
	defer func() { fmt.Println(recover()) }()

	values := []int32{5, 10, 15, 20, 25}
	firstThree := values[:3]
	lastThree := values[2:]
	lastThree[0] = 0
	fmt.Println(values, firstThree, lastThree)

	stringMap := make(map[int32]string)
	for _, v := range values {
		stringMap[v] = fmt.Sprintf("Number %d", v)
	}
	fmt.Println(stringMap)

	inputs := make(chan string)
	go func(m map[int]string) {
		fmt.Println("Sending values...")
		for _, v := range m {
			inputs <- v
		}
		fmt.Println("...finished sending values.")
	}(stringMap)

	fmt.Println(<-inputs)
	fmt.Println(<-inputs)

	panic("Recovered value.")
	fmt.Println("Unreachable code.")
}
