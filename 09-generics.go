package main

import "fmt"

func Generics() {
	printList("Pepito", "Tury", "Isaac", 2, 4)

	fmt.Println("The sum2 is:", sum2(1, 2, 5))
	fmt.Println("The sum3 is:", sum3(1.3, 2.2, 4, 7))

}

// haciendo uso de any
func printList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

// generic functions
// [T int] se le dice que trabajara con tipos de datos enteros
// [T int | float64] se le dice que trabajara con tipos de datos enteros o flotantes
func sum2[T int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}

// Crear restricciones
// el ~ significa similar a, o mas cercano de
type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

func sum3[T Numbers](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}
