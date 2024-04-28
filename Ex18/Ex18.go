package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++{
		fmt.Printf("O resto da divisão de %v por 4 é igual a %v \n", i, i%4)	
	}
}