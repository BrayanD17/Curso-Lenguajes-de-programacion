/*
1)	Haga un programa que cuente e indique el número de caracteres,
el número de palabras y el número de líneas de un texto cualquiera
(obtenido de cualquier forma que considere).
Considere hacer el cálculo de las tres variables en el mismo proceso.
*/
package main

import (
	"fmt"
	"strings"
)

var txt1 string = "¿Cómo se simula un texto sin verdaderos contenidos? \nPara muchos profesionales del sector multimedia, es la primera opción."
var txt2 string = "Con el que nuestras células grises dejen de intentar comprender este texto. \nY es que el texto simulado, texto de relleno, \ntexto marcador de posición, texto comodín, texto falso, texto dummy, \no como se le quiera llamar, no tiene ningún sentido. \nY tiene sentido que así sea."
var txt3 string = "Normalmente, un texto simulado o texto de relleno se utiliza cuando no existe aún un texto real.\nDebe dar una idea del aspecto que tendrá una página impresa o una página web. \nEste puede ser el caso cuando el diseñador gráfico les muestra esbozos de un encargo a sus clientes \no el maquetador le especifica al redactor los límites disponibles para su artículo."
var txt4 string = "El texto original comienza así: Neque porro quisquam est, \nqui dolorem ipsum, quia dolor sit, amet, consectetur, adipisci velit. \nAhí sí: los latinistas respirarán aliviados, \npuesto que este pasaje sí puede traducirse si se cuenta con los debidos conocimientos."
var txt5 string = "Pero si se trata de nuevos diseños, \nuna presentación o un proyecto para un cliente, \nentonces en ocasiones tiene más sentido utilizar un texto en castellano, \npara dejar claro cuál es el efecto de conjunto en nuestro idioma; \nes decir, que la longitud de las palabras o pasajes de texto tenga un efecto realista \ny cómo se aprecia la fuente en nuestro idioma."

func contar(texto string) {
	fmt.Println("Cantidad de caracteres: ", len(texto))
	fmt.Println("Cantidad de palabras: ", len((strings.Split(texto, " "))))
	fmt.Println("Cantidad de lineas: ", len((strings.Split(texto, "\n"))))
}
func main() {
	fmt.Println("Primer texto: ")
	fmt.Println(txt1)
	contar(txt1)
	fmt.Println("Segundo texto: ")
	fmt.Println(txt2)
	contar(txt2)
	fmt.Println("Tercer texto: ")
	fmt.Println(txt3)
	contar(txt3)
	fmt.Println("Cuarto texto: ")
	fmt.Println(txt4)
	contar(txt4)
	fmt.Println("Quinto texto: ")
	fmt.Println(txt5)
	contar(txt5)
}
