package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 21
	var j bool = true

	// menampilkan nilai i : 21 fmt.Printf("%T \n", i) // contoh : fmt.Printf("%v \n", i)
	fmt.Printf("%v \n", i)
	// menampilkan tipe data dari variabel i
	fmt.Printf("%T \n", i)
	// menampilkan tanda %
	fmt.Printf("%% \n")
	// menampilkan nilai boolean j : true
	fmt.Printf("%t\n", j)
	// menampilkan nilai boolean j : true
	fmt.Printf("%t\n", j)
	fmt.Printf("\n")
	// menampilkan unicode russia : Я (ya)
	fmt.Println("\u042f")

	// menampilkan nilai base 10 : 21
	// var num int64
	// num = 21
	// oct := strconv.FormatInt(num, 2)
	// fmt.Println(oct)
	// fmt.Printf("\n")

	num2 := int64(21)
	fmt.Println(strconv.FormatInt(num2, 2))

	num3 := "10101"
	fmt.Println(strconv.ParseInt(num3, 2, 64))

	// menampilkan nilai base 8 :25

	// menampilkan nilai base 16 : f
	// menampilkan nilai base 16 : F 13
	// menampilkan unicode karakter Я : U+042F var k float64 = 123.456;
	var k float64 = 123.456
	fmt.Printf("%U \n", 'Я')

	fmt.Println(" ")
	// menampilkan float : 123.456000
	fmt.Printf("%f \n", k)
	// menampilkan float scientific : 1.234560E+02
	fmt.Printf("%e \n", k)
}
