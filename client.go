package main

import (
	"fmt"
	"net/rpc"
)

func client() {

	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	var op int64
	var result string

	err = c.Call("Server.Constructor", "Se hizo bien", &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	for {
		fmt.Println("\n1.- Agregar calificacion")
		fmt.Println("2.- Mostrar promedio de un alumno")
		fmt.Println("3.- Mostrar promedio general")
		fmt.Println("4.- Mostrar promedio de una materia")
		fmt.Println("0.- Salir")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var materia, nombre, cali, result string
			s := []string{}

			fmt.Print("\nMateria: ")
			fmt.Scanln(&materia)
			s = append(s, materia)
			fmt.Print("Nombre del alumno:")
			fmt.Scanln(&nombre)
			s = append(s, nombre)
			fmt.Print("Calificacion:")
			fmt.Scanln(&cali)
			s = append(s, cali)

			err = c.Call("Server.AgregarCalificacion", s, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}

		case 2:
			var nombre string
			var result float64

			fmt.Print("\nNombre del alumno: ")
			fmt.Scanln(&nombre)

			err = c.Call("Server.PromedioAlumno", nombre, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio de", nombre, ":", result)
			}

		case 3:
			var aux int64
			var result float64

			err = c.Call("Server.PromedioGeneral", aux, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("\nPromedio general:", result)
			}

		case 4:
			var materia string
			var result float64

			fmt.Print("\nMateria: ")
			fmt.Scanln(&materia)

			err = c.Call("Server.PromedioMateria", materia, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio en", materia, ":", result)
			}

		case 0:
			return
		}
	}
}

func main() {
	client()
}
