package main

import (
	"fmt"
	"unsafe"
)

// Constant
const PI = 3.1416

const (
	A int    = 1
	B        = 2
	C string = "jhlk"
)

func main() {
	fmt.Println("Hello worlds")
	//Comentario
	/*
		Comentario de multiples linea

	*/

	// Tipos de variables
	var entero = 55                      // de 32 - 64 bits dependiendo del sistema operativo
	var flotante float32 = 0.3           // Flotantes 6-9 accuracy y de
	var cadena string = "alksjfdalksdjf" //Cadena inmutable
	var booleano bool = true
	//Variable sin valor inicial
	var aaa string
	ket := "lsdkjf"
	//Variable de inferencia
	fmt.Println(entero, flotante, cadena, booleano, aaa, ket)

	var num int = 50 //Declaracion explicita de variable
	var num1 int     //Declaracion de variable sin valor
	var num2 = 50    //Declaracion de variable con inferencia de tipo se puede usar fuera de funciones
	num3 := 50       //Declaracion de variable con inferencia de tipo

	fmt.Println(num, num1, num2, num3)

	//Declarar multiples variables

	var a, b, c int = 2, 4, 5
	fmt.Println(a, b, c)

	var d, e = 2, "Hello"
	fmt.Println(d, e)

	var (
		aa int
		bb float32 = 12.5
		cc string  = "hello"
	)

	fmt.Println(aa, cc, bb)

	//Camel case
	var holaComoEstas string
	//Pascal cas
	var HolaComoEstas string
	//Snake case
	var hola_como_estas string
	fmt.Println(holaComoEstas, hola_como_estas, HolaComoEstas)

	//fmt print functions
	fmt.Print("@", "2") //Print sin salto de linea y no agrega espacio entre los argumentos
	fmt.Println()       //Print con salto de linea

	var firs = "hello"
	var second = 15

	//Formatea la salida segun le digamos
	fmt.Printf("primero %v type %T \n", firs, firs)
	fmt.Printf("sec %v type %T \n", second, second)

	//Formatting verbs
	//%v Print the value in the default formal
	//%#v print the value in Go-syntax format
	// %T print the value type
	// %% print % sign

	var i int = 15
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	//Integer formatting verbs
	i = 64

	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)

	txt = "Hello"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	var ja = 15.5
	fmt.Printf("%e\n", ja)
	fmt.Printf("%f\n", ja)
	fmt.Printf("%.2f\n", ja)
	fmt.Printf("%6.2f\n", ja)
	fmt.Printf("%g\n", ja)

	var prueba bool = false
	print(prueba)
	//Array
	var name = [4]int{1, 2, 3, 4} // Longitud definida
	var nam1 = [...]int{1, 23}    //Longiutd inferida
	name[0] = 55
	nam1[0] = 12
	var arr = [5]int{}
	arr[4] = 1
	fmt.Println("Value at index 4:", arr[4]) // Use the value at index 4
	//Iniciar determinados indices
	arr2 := [5]int{1: 20, 2: 20}
	fmt.Println(arr2)
	fmt.Println(len(arr))

	//Slices Su tamanho es variable

	slica := []int{1, 2}

	fmt.Println(slica)
	fmt.Println(len(slica))
	fmt.Println(cap(slica))

	//Tomar parte de un array

	arr12 := [6]int{10, 11, 12, 13, 14, 15}

	//Toma los valores de i hasta j-1
	sli12 := arr12[2:4]

	fmt.Println(sli12)

	//Puede crecer hasta 4 aunque actualmente solo tenga 2 valores
	fmt.Println(cap(sli12))

	//Make inicializa el objeto listo para usar se usa con tipos dinamicos
	// parameters objeto dinamico, tamaño inicial, hasta donde puede crecer
	slice := make([]int, 2, 4)
	fmt.Println(slice)

	//accder a aelemento
	fmt.Println(slice[0])

	slice[0] = 22

	fmt.Println(slice)
	fmt.Println(cap(slice))
	slice = append(slice, 1, 2, 2, 2, 3, 5, 2)
	fmt.Println(cap(slice))

	slice2 := []int{}

	slice2 = append(slice2, 1, 2, 3)

	slice = append(slice, slice2...)
	fmt.Println(slice)

	source := []int{1, 2, 3, 4, 5, 67}
	destin := make([]int, 6)

	copy(destin, source)
	fmt.Println(destin)
	//Operators

	var suma int = 10 + 11
	fmt.Println(suma)
	var x = 1
	x++
	x += 5

	if x > 0 {

	}

	fmt.Println(split(55))

	pointExercise(source, 2)

	var MaxInt uint64 = 1<<64 - 1
	var MinInt int64 = -1 << 31
	fmt.Println(MaxInt, MinInt)

}

func a(x, y int) int {
	return 1
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func pointExercise(arr []int, des uintptr) {
	ptr := &arr[0]

	address := uintptr(unsafe.Pointer(ptr))

	newDe := des * unsafe.Sizeof(int(0))

	address = address + newDe

	newValue := (*int)(unsafe.Pointer(address))

	fmt.Printf("Valor en la dirección: %d\n", *newValue)
}
