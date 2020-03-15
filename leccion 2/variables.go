package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("Villa de Alvarez" + ", Colima")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println("")
	fmt.Println("Tabla de Valores de && (and)")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	fmt.Println()
	fmt.Println("Tabla de Valores de || (or)")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	fmt.Println()
	fmt.Println(" ! (not)")
	fmt.Println(!true)
	fmt.Println(!false)

	var nombre string = "Allison"
	var paterno string = "Terrazas"
	var espacio string = " "

	fmt.Println(nombre + " " + paterno)
	fmt.Println(nombre + espacio + paterno)
	fmt.Println()

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)
	fmt.Println(!d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

	var var1, var2 int = 7, 3
	fmt.Println(var1 + var2)

	var3 := "7"
	var4 := "3"
	fmt.Println(var3 + var4)
	fmt.Println(var3 + espacio + var4)

}
