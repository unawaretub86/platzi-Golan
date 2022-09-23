package main

import "fmt"

func main() {
	// DECLARANDO CONSTANTES
	const pi float64 = 3.14
	const pi1 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi1:", pi1)

	// DECLARANDO VARIABLES ENTERAS

	// con los := significa que tiene que crear la variable y asignarla , pero si es asi = significa que debio estar creada antes
	base := 12

	// aca estamos declarando la variable y la asignamos de una vez el valor y el tipo de variable
	var altura int = 14

	var area int

	fmt.Println(base, altura, area)

	// ZERO VALUES
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// AREA DE CUADRADO

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado:", areaCuadrado)

	// OPERADORES ARITMETICOS

	x := 10
	y := 50

	// SUMA

	result := x + y

	fmt.Println("SUMA:", result)

	// RESTA

	result = y - x

	fmt.Println("RESTA:", result)

	// MULTIPLICACION

	result = x * y

	fmt.Println("MULTIPLICACION:", result)

	// DIVISION

	result = y / x

	fmt.Println("DIVISION:", result)

	// MODULO

	result = y % x

	fmt.Println("MODULO:", result)

	// INCREMENTAL
	x++

	fmt.Println("INCREMENTAL:", x)

	// DECREMENTAL
	x--

	fmt.Println("DECREMENTAL:", x)

	// RETOS (CALCULAR AREA DE RECTANDULO, TRAPECIO, CIRCULO)

	// RECTANGULO
	const baseRectangulo int = 6
	const alturaRectangulo int = 10
	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("Area Rectangulo: ", areaRectangulo)

	// TRAPECIO
	const base1 int = 10
	const base2 int = 4
	const alturaT int = 4
	areaTrapecio := (base1 + base2) * alturaT / 2

	fmt.Println("Area TRAPECIO: ", areaTrapecio)

	// CIRCULO
	const radio float64 = 5

	var areaCirculo float64 = radio * radio * pi

	fmt.Println("Area CIRCULO", areaCirculo)

	// TIPOS DE DATOS

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i
}
