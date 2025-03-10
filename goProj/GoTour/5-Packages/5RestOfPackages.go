package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func printPi() {
	fmt.Println(math.Pi)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func varibles() {
	c, python, java := true, false, "No!"

	var i, j int = 1, 2

	k := 3

	fmt.Println(i, j, k, c, python, java)
}

func types() {
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func zeros() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

}

func conversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(x, y, z)
}

func typeInference() {
	v := 4.123 // Change me!
	fmt.Printf("%v is of Type %T\n", v, v)
}

func constants() {
	const Pi = 3.14
	const World = "世界"
	const Truth = true

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	fmt.Println("Go rules?", Truth)
}

func numcon() {
	// Create a huge number, shift 1 100 bits left
	const (
		Big   = 1 << 100
		Small = Big >> 99
	)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {

	printPi()

	fmt.Println(add(42, 13))

	a, b := swap("Hello", "Earl")
	fmt.Println(a, b)

	fmt.Println(split(17))

	varibles()

	types()

	zeros()

	conversion()

	typeInference()

	constants()

	numcon()
}
