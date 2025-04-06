package main

import "fmt"

// se recomienda que las constantes inialicen con mayus
/*const Pi = 3.1415926
const (
	Email = "test@gmail.com"
	City  = "Barranquilla"
)*/

// iota inicia en 0 pero se le suma el 1 para que comience en 1
// las proximas constante se les va sumando +1 ejemplo: Lunes = 2, Martes = 3... etc
const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func varAndTypeData() {

	//VARIABLES --------------##

	//FORMA 1
	//var firstName, lastName string
	//var age int = 29
	//firstName = "Turiano"
	//lastName = "Romero"

	//FORMA 2
	//var (
	//	firstName string = "Turiano"
	//	lastName  string = "Romero"
	//	age       int    = 29
	//)

	//FORMA 3
	//var firstName, lastName, age = "Turiano", "Romero", 29

	//FORMA 4 -> solo sirve para dentro de las fn, con var para fuera y dentro de fn
	//firstName, lastName, age := "Turiano", "Romero ->", 29
	//fmt.Println("My name is:", firstName, lastName, "I am:", age)

	// CONSTANTS --------------##
	//fmt.Println("The value of pi is ->", Pi)
	//fmt.Println("The value of email is:", Email, "And city is:", City)
	fmt.Println("The values days ->", Domingo, Lunes, Martes, Miercoles, Jueves, Viernes, Sabado)
}
