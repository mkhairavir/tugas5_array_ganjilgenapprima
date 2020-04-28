package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomArr := make([]int, 30)

	for i := 0; i < len(randomArr)-10; i++ {
		randomArr[i] = rand.Intn(50)
	}

	fmt.Println("Isi array: ", randomArr)

	// GENAP //
	jumlahGenap := 0
	var arrayGenap [20]int

	for i := 0; i < len(randomArr)-10; i++ {
		if randomArr[i]%2 == 0 {
			arrayGenap[i] = randomArr[i]
			jumlahGenap++
		}
	}

	fmt.Println("Isi array genap: ", arrayGenap)
	fmt.Println("Banyaknya array genap: ", jumlahGenap)
	fmt.Println("===========================================================================================")

	// GANJIL //

	jumlahGanjil := 0
	var arrayGanjil [20]int

	for i := 0; i < len(randomArr)-10; i++ {
		if randomArr[i]%2 != 0 {
			arrayGanjil[i] = randomArr[i]
			jumlahGanjil++
		}
	}
	fmt.Println("Isi array ganjil: ", arrayGanjil)
	fmt.Println("Banyaknya array ganjil: ", jumlahGanjil)
	fmt.Println("===========================================================================================")

	// PRIMA //
	jumlahPrima := 0
	var arrayPrima [20]int

	for i := 0; i < len(randomArr)-10; i++ {
		tanda := 0
		for j := 1; j <= randomArr[i]; j++ {
			if randomArr[i]%j == 0 {
				tanda++
			}
		}
		if (tanda == 2) && (randomArr[i] > 1) {
			arrayPrima[i] = randomArr[i]
			jumlahPrima++
		}
	}

	fmt.Println("Isi array prima: ", arrayPrima)
	fmt.Println("Banyaknya array prima: ", jumlahPrima)
	fmt.Println("===========================================================================================")

}
