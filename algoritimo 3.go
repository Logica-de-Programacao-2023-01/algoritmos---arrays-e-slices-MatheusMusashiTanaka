package main

import "fmt"

func main() {
	var numeros = []int{2, 3, 4, 5}
	mult := numeros[0] * numeros[1] * numeros[2] *numeros[3]

	fmt.Printf("produto = %v" , mult)
