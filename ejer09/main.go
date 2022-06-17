package main

import (
	"fmt"
	"time"

	us "./user"
)

/*type usuario struct {
	Id int
	Nombre string
	FechaAlta time.Time
	Status bool
	

}*/
type teo struct{
	us.Usuario
}

func main(){
	u := new(teo)
	u.AltaUsuario(1, "Mateo Narvaez", time.Now(), true)
	fmt.Println(u.Usuario)

}

// Para correr el comando llamando los otros modulos locales es con el comando GO111MODULE="auto" go run main.go
