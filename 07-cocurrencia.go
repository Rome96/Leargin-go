package main

import (
	"fmt"
	"net/http"
)

// Crear un canal de tipo entero
//canal := make(chan int)

// Enviar un valor a través del canal
//canal <- 10

// Recibir un valor del canal
//valor := <-canal

func Concurrencia() {
	//start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	//Creamos un channel en go
	ch := make(chan string)

	// Recorreer los apis
	for _, api := range apis {
		go checkAPI(api, ch) // se ejecuta con goroutine
	}
	fmt.Print(<-ch)

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	//elapsed := time.Since(start)
	//fmt.Printf("¡Listo! ¡Tomó %v segundos!\n", elapsed.Seconds())
}

// Función que verifica los APIS
func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: ¡%s está caído!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: ¡%s está en funcionamiento!\n", api)
}
