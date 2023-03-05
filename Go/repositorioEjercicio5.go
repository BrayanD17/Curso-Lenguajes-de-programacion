package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 20

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	var existe bool = false
	for i := 0; i < len(*l); i++ {
		if strings.Compare(nombre, (*l)[i].nombre) == 0 {
			(*l)[i].cantidad = (*l)[i].cantidad + cantidad
			existe = true
		}
	}
	if existe == false {
		(*l) = append((*l), producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
}

func (l *listaProductos) buscarProducto(nombre string) int {
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cant int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 && cant > 0 {
		if (*l)[prod].cantidad >= cant {
			(*l)[prod].cantidad = (*l)[prod].cantidad - cant
		}
		if (*l)[prod].cantidad == 0 {
			var i int = 0
			for i = 0; i < len(*l); i++ {
				if (*l)[i].nombre == nombre {
					(*l) = append((*l)[:i], (*l)[i+1:]...)
					fmt.Println("Producto eliminado!!")
				}
			}
		}

		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista"
	}
}

func llenarDatos() {
	datosArchvio, err := ioutil.ReadFile("clase3.txt")
	if err != nil {
		log.Fatal(err)
	}
	var textoComas []string = strings.Split(string(datosArchvio), ",")
	var listaLimpia []string
	var datos string
	for i := 0; i < len(textoComas); i++ {
		datos = strings.TrimSpace(textoComas[i]) //quita todos los espacios junto con los saltos de linea
		listaLimpia = append(listaLimpia, datos)
	}
	var nombreP string
	i := 0
	for i < len(listaLimpia) {
		nombreP = listaLimpia[i]
		cantidadP, err := strconv.Atoi(listaLimpia[i+1])
		precioP, err := strconv.Atoi(listaLimpia[i+2])
		if err != nil {
			fmt.Println("Error en la conversion de String a Int")
			return
		} else {
			lProductos.agregarProducto(nombreP, cantidadP, precioP)
		}
		i = i + 3
	}
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	var i int = 0
	var lminimos listaProductos
	for i = 0; i < len(*l); i++ {
		if (*l)[i].cantidad <= existenciaMinima {
			lminimos = append(lminimos, (*l)[i])
		}
	}
	// debe retornar una nueva lista con productos con existencia mínima
	return lminimos
}

func aumentarInventarioDeMinimos(l listaProductos) listaProductos {
	var i int = 0
	var Alminimos listaProductos
	for i = 0; i < len(l); i++ {
		Alminimos = append(Alminimos, producto{nombre: (l)[i].nombre, cantidad: existenciaMinima - (l)[i].cantidad, precio: (l)[i].precio})
	}
	return Alminimos
}

type NameSorter []producto

func (a NameSorter) Len() int {
	return len(a)
}
func (a NameSorter) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a NameSorter) Less(i, j int) bool {
	return a[i].nombre < a[j].nombre
}

func (l *listaProductos) ordenamientoNombre() {
	sort.Sort(NameSorter(*l))
	fmt.Println(*l)
}

func main() {
	llenarDatos()
	fmt.Println(lProductos)
	lProductos.venderProducto("arroz", 4)
	fmt.Println(lProductos)
	lProductos.agregarProducto("azucar", 20, 1500)
	fmt.Println(lProductos)
	lProductos.venderProducto("frijoles", 4)
	fmt.Println(lProductos)
	lProductos.venderProducto("leche", 10)
	lProductos.agregarProducto("azucar", 20, 1500)
	fmt.Println(lProductos)
	fmt.Println("Lista con minimos")
	fmt.Println(lProductos.listarProductosMínimos())

	/* Ejercicios para repositorio*/

	fmt.Println("Productos a los cuales se les amplio el inventario en: ")
	var lminimos listaProductos = lProductos.listarProductosMínimos()
	fmt.Println(aumentarInventarioDeMinimos(lminimos))

	sort.SliceStable(lProductos, func(i, j int) bool { return lProductos[i].cantidad < lProductos[j].cantidad })
	fmt.Println("Lista ordenada por cantidad disponible 0 - N+")
	fmt.Println(lProductos)
	fmt.Println("______________________")
	fmt.Println("Lista ordenada por nombre A-Z")
	lProductos.ordenamientoNombre()

}
