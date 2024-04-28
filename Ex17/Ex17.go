
package main

import (
	"fmt"
)

func main() {
	ano := 2005
	anoFinal := 2024
	for {
		fmt.Println(ano)
		ano++	
		if(ano > anoFinal)	{
			break
		}
	}
}
