package main

import (
	"ejercicios/cursoUdemy/handlers"
	"ejercicios/cursoUdemy/infraestructura"
	"fmt"
	"log"
)

func main() {
	if infraestructura.ChequeoConexion() == 1 {
		fmt.Println("Conexion exitosa")
	} else {
		log.Fatal("No hay conexion")
	}
	handlers.Manejadores()
}
