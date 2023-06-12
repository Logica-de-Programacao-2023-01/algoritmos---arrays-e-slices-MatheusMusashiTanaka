package main

import "fmt"

func main() {
	var numeros = []int{1 ,2, 3, 4, 5}
	var slice = append(numeros[0:2],numeros[3:]...)
	fmt.Println(slice)
}
