
package main
import (
	"fmt"
	"strconv"
)

func main(){
	//conversion de tipos
	edad:=17
	edad_str:=strconv.Itoa(edad)
	edad_int,_:=strconv.Atoi(edad_str)
	fmt.Println("Mi edad es "+edad_str)
	fmt.Println("Mi edad no es", edad_int-1)
}
