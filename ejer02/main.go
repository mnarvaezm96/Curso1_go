package main

import (
	"fmt"
	"strconv"


)
var numero int
var texto string
var status bool = true


func main (){
	//var numero2 int
	//numero3 := 4
	//numero3 = 15
	//var numero1, numero2, numero3 int
	numero1, numero2, numero3, texto, status := 2, 5, 67, "este es mi texto", false
	
	var moneda int64
	numero2 = int(moneda)
	//texto = fmt.Sprintf("%d", moneda)
	texto = Strconv.Itoa(int(moneda))

	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(texto)
	fmt.Println(status)

	mostrarstatus()
}

func Mostrarstatus(){
	fmt.Println(status)
}
