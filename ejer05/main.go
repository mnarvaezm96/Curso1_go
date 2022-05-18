package main

import "fmt"

/*func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

}*/

/*func main() {
	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}

	}
}*/

/*func main() {
	var i = 0
	for i < 10 {
		fmt.Printf("\n Valor de i: %d", i)
		if i == 5 {
			fmt.Printf("Multiplicamos por 2 \n")
			i=i*2
			continue
		}
		fmt.Printf("	Pasó por aqui \n")
		i++
	}
}*/

func main() {
	var i int = 0

	RUTINA:
		for i < 10 {
			if i == 4 {
			i=i+2
			fmt.Println("voy a RUTINA suma de 2 a i")
			goto RUTINA
			}
			fmt.Printf("Valor de i: %d\n",i)
			i++
		}


}
