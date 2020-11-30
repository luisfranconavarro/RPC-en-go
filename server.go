package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"
)

type Server struct {
	Materias, Alumnos map[string]map[string]float64
	construir bool
}

func (this *Server) Constructor(s string, reply *string) error {
	if !this.construir {
		this.Materias = make(map[string]map[string]float64)
		this.Alumnos = make(map[string]map[string]float64)
		this.construir = true
	}

	*reply = "Se conecto al servidor"

	return nil
}

func (this *Server) AgregarCalificacion(s []string, reply *string) error {
	_, err := this.Materias[s[0]]
	*reply = "Se agrego correctamente"

	if err == false {
		alumno := make(map[string]float64)
		f2, _ := strconv.ParseFloat(s[2], 8)
		alumno[s[1]] = f2
		this.Materias[s[0]] = alumno
	} 

	_, err2 := this.Alumnos[s[1]]
	*reply = "Calificacion agregada"

	if err2 == false {
		clase := make(map[string]float64)
		f2, _ := strconv.ParseFloat(s[2], 8)
		clase[s[0]] = f2
		this.Alumnos[s[1]] = clase
	} 

	return nil
}

func (this *Server) PromedioAlumno(nombre string, reply *float64) error {
	var promedio float64
	var i int64
	promedio = 0
	i = 0

	for _, calificacion := range this.Alumnos[nombre] {
		promedio = promedio + calificacion
		i = i + 1
	}
	promedio = promedio / float64(i)
	*reply = promedio
	return nil
}

func (this *Server) PromedioMateria(nombre string, reply *float64) error {
	var promedio float64
	var i int64
	promedio = 0
	i = 0

	for _, calificacion := range this.Materias[nombre] {
		promedio = promedio + calificacion
		i = i + 1
	}
	promedio = promedio / float64(i)
	*reply = promedio
	return nil
}

func (this *Server) PromedioGeneral(f int64, reply *float64) error {
	var promedio float64
	var promedioGeneral float64
	var i int64
	var j int64
	i = 0
	j = 0
	promedio = 0
	promedioGeneral = 0

	for nombreAlumno := range this.Alumnos {
		i = 0
		promedio = 0
		for _, calificacion := range this.Alumnos[nombreAlumno] {
			promedio = promedio + float64(calificacion)
			i = i + 1
		}
		promedio = promedio / float64(i)
		promedioGeneral = promedioGeneral + promedio
		j = j + 1
	}
	promedioGeneral = promedioGeneral / float64(j)
	*reply = promedioGeneral
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	fmt.Println("Servidor encendido, presione enter para salir")
	var input string
	fmt.Scanln(&input)
}
