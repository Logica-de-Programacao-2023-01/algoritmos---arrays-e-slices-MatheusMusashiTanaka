package main

import "fmt"

func main() {
	var numeros = []int {}
	var inp int
	i:= 1

	for i <=5{
		fmt.Println("numero: ")
		fmt.Scan(&inp)
		numeros = append(numeros, inp)
		i++
	}
	sum := numeros[0] + numeros[1] + numeros[2] + numeros[3] + numeros[4]
	fmt.Println(numeros)
	fmt.Println("a soma dos numeros é : ", sum)

}
