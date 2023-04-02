package main

import "fmt"

func sum[T int | string](a, b T) T {
	return a + b
}

func main() {
	result := sum("1", "2")
	fmt.Println(result)
}
