package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	//imprimir
	fmt.Println("imprimir en patalla")
	variable := 17
	fmt.Printf("imprimir variable: %d\n", variable) // Se usa printf, no println

	// leer
	var edad int
	fmt.Print("escribe tu edad:")
	fmt.Scanf("%d", &edad)
	fmt.Printf("tu edad es %v\n", edad)

	input := bufio.NewReader(os.Stdin)
	fmt.Print("ingresa tu nombre:")
	nombre,err := input.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Hola " +nombre)
	}
}
