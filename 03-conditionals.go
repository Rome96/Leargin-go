package main

import (
	"fmt"
	"runtime"
	"time"
)

func conditionals() {
	t := time.Now()
	hour := t.Hour()

	if hour < 12 {
		fmt.Println("Good morning!")
	} else if hour < 17 {
		fmt.Println("Good afternoon!")
	} else {
		fmt.Println("Good evening!")
	}

	//SWITCH -> Validate OS
	os := runtime.GOOS

	switch os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> Other Os")
	}

	// for
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
			// break  // detiene la ejecucion
		}
		fmt.Println(i)
	}

}
