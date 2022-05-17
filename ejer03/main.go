package main

import "fmt"

var estado  bool

/*func main() {
	estado=true
	if estado= false; estado == true{
		fmt.Println(estado)
	}else{
		fmt.Println("Es falso")
	}
}*/

func main() {
	switch numero := 10; numero {
	case 1:
		fmt.Println(1)

	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	case 6:
		fmt.Println(6)
	default:
		fmt.Println("Mayo a 6")
	}
}

