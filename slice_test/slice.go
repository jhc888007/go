package main

import "fmt"

func main() {
	a := make([]int, 10)

	for i := 0; i < 10; i++ {
		a[i] = i
	}

	for _, v := range a {
		fmt.Printf("%v", v)
	}
	fmt.Println()

	a = append(a[:1], a[2:]...)

	for _, v := range a {
		fmt.Printf("%v", v)
	}
	fmt.Println()

	fmt.Println("vim-go")
}
