/* - Escreva um programa que mostre um número em decimal, binário, e hexadecimal.*/
package main

import(
	"fmt"
)

func main(){
    x := 100
	fmt.Printf("Variable X = %v \nDecimal = %d \nBinary = %b \nHexadecimal = %#x", x, x,x,x)
}