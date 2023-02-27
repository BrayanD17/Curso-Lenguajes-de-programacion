/*
Para dicho fin, escriba y use una función que reciba de parámetro
la cantidad de elementos de la línea del centro,
la cual debe ser impar positiva.
*/
package main

import (
	"fmt"
	"sort"
)

var filas []string

func imparPositivo(num int) bool {
	if num%2 == 0 && num > 1 {
		return true
	}
	return false
}

func estructura(num int, agregar int) {
	var centro int = (num - 1) / 2
	var patron string = ""
	if agregar == 0 {
		for i := 0; i < num; i++ {
			if i == centro {
				patron = patron + "*"
			} else {
				patron = patron + " "
			}
		}
		filas = append(filas, patron)
	} else {
		agregar = agregar + 1
		var vacios int = num - agregar
		var mitadVacios int = vacios / 2

		for i := 0; i < num; i++ {
			if i < mitadVacios && agregar != num {
				patron = patron + " "
			} else if agregar != 0 {
				patron = patron + "*"
				agregar = agregar - 1
			} else {
				patron = patron + " "
			}
		}
		filas = append(filas, patron)
	}

}
func dibujarFigura(num int) {
	var comprobacion bool = imparPositivo(num)
	if comprobacion == false {
		anadir := 0
		for anadir < num {
			estructura(num, anadir)
			anadir = anadir + 2
		}

	}
	for i := 0; i < len(filas); i++ {
		fmt.Println(filas[i])
	}
	sort.Sort(sort.Reverse(sort.StringSlice(filas)))
	for i := 1; i < len(filas); i++ {
		fmt.Println(filas[i])
	}
}
func main() {
	dibujarFigura(5)
	filas = nil
	fmt.Println("_____________")
	dibujarFigura(7)

}
