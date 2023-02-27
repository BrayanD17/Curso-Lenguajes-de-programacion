package main

import (
	"fmt"
)

type listaN []int

var lista listaN
var cantidadD int = 30

func llenar(l *listaN) {
	for i := 1; i <= cantidadD; i++ {
		(*l) = append((*l), i)
	}
	fmt.Println("Lista original")
	fmt.Println(*l)
	fmt.Println("-------------------------------------------")
}
func mover(izq_der int, l *listaN, movimiento int) []int {
	var lista2 []int
	if izq_der == 0 {
		for i := 0; i < movimiento; i++ {
			if i == 0 {
				lista2 = append((*l)[1:], (*l)[0])
				fmt.Println("Primer movimiento a la izquierda")
				fmt.Println(lista2)
			} else {
				fmt.Println("Movimiento ", i+1, " a la izquierda")
				lista2 = append((lista2)[1:], (lista2)[0])
				fmt.Println(lista2)
			}
		}

	}
	if izq_der == 1 {
		for i := 0; i < movimiento; i++ {
			if i == 0 {
				lista2 = append((*l)[cantidadD-1:], (*l)[:cantidadD-1]...)
				fmt.Println("Primer movimiento a la dereccha")
				fmt.Println(lista2)
			} else {
				fmt.Println("Movimiento ", i+1, " a la derecha")
				lista2 = append(lista2[len(lista2)-1:], lista2[:len(lista2)-1]...)
				fmt.Println(lista2)
			}
		}

	}
	return lista2
}
func main() {
	llenar(&lista)
	mover(1, &lista, 4)
}
