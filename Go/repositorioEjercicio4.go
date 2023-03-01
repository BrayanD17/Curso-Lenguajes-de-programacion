package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Persona struct {
	nombre   string
	apellido string
	edad     int
}

type listaPersonas []Persona

var integrantes listaPersonas

func (lp *listaPersonas) agregarPersona(nombre string, apellido string, edad int) {
	nombrebool := slices.IndexFunc(*lp, func(per Persona) bool { return per.nombre == nombre })
	apellidobool := slices.IndexFunc(*lp, func(per Persona) bool { return per.apellido == apellido })
	if nombrebool == -1 && apellidobool == -1 {
		*lp = append(*lp, Persona{nombre: nombre, apellido: apellido, edad: edad})

	} else {
		fmt.Println("Persona repetida")
	}
}
func (lp *listaPersonas) imprimirLista() {
	for _, person := range *lp {
		fmt.Println(person.nombre)
		fmt.Println(person.apellido)
		fmt.Println(person.edad)
		fmt.Println("__________________")

	}
}
func (lp *listaPersonas) mayoresEdad() []Persona {
	listaMayores := filter2(*lp, func(person Persona) bool {
		return person.edad >= 18
	})
	return listaMayores
}

func map2[P, V any](l []P, f func(P) V) []V {
	mapped := make([]V, len(l))

	for i, e := range l {
		mapped[i] = f(e)
	}
	return mapped
}

func filter2[P any](l []P, f func(P) bool) []P {
	filtered := make([]P, 0)

	for _, element := range l {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered

}

func main() {
	integrantes.agregarPersona("Juan", "Lopez", 19)
	integrantes.agregarPersona("Ezequiel", "Duarte", 12)
	integrantes.agregarPersona("Kimberly", "Suarez", 20)
	integrantes.agregarPersona("Josue", "Rojas", 15)
	integrantes.agregarPersona("Kendall", "Salazar", 22)
	integrantes.agregarPersona("Nelsy", "Gomez", 48)
	integrantes.agregarPersona("Nelsy", "Gomez", 48)
	integrantes.imprimirLista()
	fmt.Println("Mayores de edad ingresados!: ")
	fmt.Println(integrantes.mayoresEdad())
}
