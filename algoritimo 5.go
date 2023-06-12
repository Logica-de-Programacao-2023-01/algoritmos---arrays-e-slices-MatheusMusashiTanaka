package main

import "fmt"

func main() {
	mtz := [3][2]int{}
	for x := 0; x < 3; x++ {
		for y := 0; y < 2; y++ {
			fmt.Printf("Digite o valor de [%d] [%d]: ", x, y)
			fmt.Scan(&mtz[x][y])
		}
	}
	fmt.Printf("matriz : ")
	for x := 0; x < 3; x++ {
		for y := 0; y < 2; y++ {
			fmt.Printf("%d ", mtz[x][y])
		}
		fmt.Println()
	}
}
