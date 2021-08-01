package main

/*
	== igual a
	!= diferente de
	< menor que
	> mayor que
	>= mayor o igual que
	<= menor o igual que
	&& AND (Y)
	|| OR (O)
*/

import (
	"fmt"
)
func main(){
	x:=10
	y:=5
	if x > y {
		fmt.Printf("%v es mayor que %v\n", x, y)
	} else if x == y {
		fmt.Printf("%v es mayor que %v\n", y, x)
	} else {
		fmt.Printf("%v y %v son iguales\n", x, y)
	}
}



