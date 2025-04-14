package main

import "fmt"

type Person struct {
	Name     string
	lastName string
	age      int8
	email    string
}

// esto es un metodo del struct person, y se le pasa como receptor p que es un puntero de tipo Person
func (p *Person) sayHello() {
	fmt.Println("Hello my name is -> " + p.Name)
}

func matrices() {
	//Matrices --------------##

	var a = [...]int{10, 40, 59, 28}

	//len nos da la cantidad de elementos dentro del array
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for i, value := range a {
		fmt.Printf("Indice => %d - value => %d\n", i, value)
	}

	//Matriz bidimencional
	var b = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(b)
	fmt.Println(b[2][1])

	//MAPS --------------##
	colors := map[string]string{
		"red":    "#FF0000",
		"green":  "#00FF00",
		"blue":   "#0000FF",
		"yellow": "#FF00FF",
	}
	fmt.Println(colors)
	fmt.Println(colors["red"])

	// agg value to map
	colors["black"] = "#0000"
	fmt.Println(colors)

	if value, ok := colors["gray"]; ok {
		fmt.Println("The value is -> ", value)
	} else {
		fmt.Println("The value not found")
	}

	// delete value in the map
	delete(colors, "red")
	fmt.Println(colors)

	//iteration to map
	for key, value := range colors {
		fmt.Printf("The key is -> %v, The value is -> %v\n ", key, value)
	}

	//STRUCTS --------------##
	//method 1
	//var p Person
	//p.Name = "Tury"
	//p.lastName = "Rome"
	//p.age = 29
	//p.email = "tury@gmail.com"

	//method 2
	p := Person{
		"Tury",
		"Rome",
		29,
		"tury@gmail.com",
	}

	fmt.Println(p)
	fmt.Println(p.Name)

	//POINTERS --------------##
	var x int = 10
	var y *int = &x // guarda le referencia de la memoria donde está x

	fmt.Println(x)
	edit(&x)
	fmt.Println(y)
	fmt.Println(x)

	//exaple whit struct and method
	per := Person{
		"pepito",
		"perez",
		20,
		"pepito@pepito.com",
	}

	per.sayHello() // de esta forma accedemos a los metodos de las structs

}

// POINTERS --------------##
func edit(x *int) {
	*x = 30 // accedemos a la reference en memoria a travez del puntero y le cambiamos su valor
}
