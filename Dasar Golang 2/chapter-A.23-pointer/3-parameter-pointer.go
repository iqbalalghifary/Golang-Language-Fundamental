package main

import "fmt"

func main() {
	var number = 4
	fmt.Println("before :", number) // 4

	change(&number, 15)
	fmt.Println("after  :", number) // 10
}

func change(original *int, value int) {
	*original = value
}
