package main

import (
	"fmt"

	kalkulator "github.com/iqbaltk/go-kalkulator"
	kalkulatorv2 "github.com/iqbaltk/go-kalkulator/v2"
)

func main() {
	result := kalkulator.CekGanjilGenap(1)
	resultv2 := kalkulatorv2.CekGanjilGenap(1, 2, 3, 4, 5)

	fmt.Println("V1 : ", result)
	fmt.Println("V2 : ", resultv2)
}
