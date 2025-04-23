package main

import "fmt"

// Functios variadica -----------------------
// cons ...tipo - se le dice que reciba n cantidad de parametros y su tipo
func sum(name string, nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Hola %s, la suma es: %d\n", name, total)

	return total
}

// send multiple data
func returnData(data ...interface{}) {
	for _, dato := range data {
		fmt.Printf("%T - %v\n", dato, dato) //%T -> tipo de dato
	}
}

// Recursive functions -----------------
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func Functions() {
	fmt.Println(sum("pepito", 2, 3))
	returnData("Hola", 42, true, 3.14)

	//recursive
	fmt.Println("Recursive function ->", factorial(5))
}
