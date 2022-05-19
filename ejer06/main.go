package main

import "fmt"

var Calculo func(int,int) int = func(num1 int, num2 int) int{
	return num1 + num2
}

func main(){
	fmt.Printf("suma 5 + 7  = %d \n",Calculo(5,7))
	//resta
	Calculo = func (num1 int, num2 int) int{
		return num1 - num2

	}
	fmt.Printf("resta 6 - 4  = %d \n",Calculo(6,4))

	Calculo = func (num1 int, num2 int) int{
		return num1 / num2

	}
	fmt.Printf("resta 12 / 3  = %d \n",Calculo(12,3))

	operaciones()

	/* CLOSURES*/
	tablaDel:= 2
	Mitabla := Tabla(tablaDel)
	for i:=3; i < 11; i++{
		fmt.Println(Mitabla())
	}
}

func operaciones() {
	resultado := func() int{
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}
//Closiur que es para temas de segurida

func Tabla(valor int) func() int {
	numero := valor
	secuencia:=0
	return func() int {
		secuencia+=1
		return numero*secuencia
	}
}

