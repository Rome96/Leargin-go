package main

import (
	"fmt"
	"os"
	"strconv"
)

func Contacts() {
	str := "123"
	num, err := strconv.Atoi(str) // convert string to int
	if err != nil {
		fmt.Println("Err ->", err)
		return
	}
	fmt.Println("num:", num)

	//--------------------
	result, divErr := divide(12, 6)
	if divErr != nil {
		fmt.Println("divErr ->", divErr)
	} else {
		fmt.Println("result:", result)
	}

	//example defer
	exampleDefer()
}

func divide(dividend, divisor int) (int, error) {
	//if divisor == 0 {
	//	return divisor, errors.New("cannot divide by zero")
	//}
	validateZero(divisor)
	return dividend / divisor, nil
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("cannot divide by zero")
	}
}

func exampleDefer() {
	defer fmt.Println(1) // con defer hacemos que se ejecute de ultimo
	fmt.Println(2)
	fmt.Println(3)

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.WriteString("Hello World, tury the best")
	if err != nil {
		fmt.Println(err)
		return
	}
}
