package main

/* why go? -->
		- performance (runs fast)
		- concurrency
		- compiled language, garbage collector
		- clean syntax
		- powerful standard library
		- portable (compiles on many OS's)
   what is it used for?
		- web apps
		- android and iOS apps
		- scripts
		- image processing
*/

// packages, fmt (format)
import ("fmt"
		"math"
		"math/rand")

// structs (custom classes)
type car struct {
	gas_pedal      uint16 // min 0 max 65535
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

// to run the file, 'go run basic.go'
func main() {
	foo()
	createRand()
	dataTypes()
	w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1, w2))
	convertType()
	initializeClasses()
	cond()
	arrays()
	loops()
	maps()
}

const usixteenbitmax float64 = 65535
func initializeClasses() {
	a_car := car {
		65000,
		0,
		12561,
		225.0,
	}
	fmt.Println("Gas pedal:", a_car.gas_pedal) // reference attr in obj
	fmt.Println("KMH:", a_car.kmh())

	// pointer receiver
	a_car.new_top_speed(500)
	fmt.Println("KMH altered:", a_car.kmh())
}

// associated with car struct
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

// pointer receiver (*)--> pointer to change original struct, when you have large struct, use pointer receiver
// value receiver --> makes copy of struct and makes changes to that copy
func (c *car) new_top_speed (newspeed float64) {
	c.top_speed_kmh = newspeed
}

func cond() {
	myX := 10
	if myX < 12 { // standard if condition
		fmt.Println("myX is less than 12.")
	}
}

func arrays() {
	//var myA [5]int
	//myA := [5]int{5, 4, 3, 2, 1}
	myA := []int{5, 4, 3, 2, 1}
	myA = append(myA, 10)
	myA[2] = 4
	fmt.Println(myA)

	// iterate through array
	for index, value := range myA {
		fmt.Println("index:", index, "value:", value)
	}
}

func dataTypes() {
	// set variable to type (can't change this) -- MUST use outside of the function
	//var num1 float64 = 5.6
	//var num2 float64 = 9.5
	// or
	//var num1, num2 float64 = 5.6, 9.5
	// set variable to dynamic type:
	num1, num2 := 5.6, 9.5
	fmt.Println("Solution:", add(num1, num2))
}

// if the first letter is capitalized, then that function will be exported by Go, therefore all functions you use must be capitalized
func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4)) // how to print something
}

// declare a constant
const x int = 5

func createRand() {
	fmt.Println("A number from 1-100:", rand.Intn(99))
}

func add(x float64, y float64) float64 {
//this works --> func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a,b
}

func convertType() {
	var a int = 62
	var b float64 = float64(a)
	x := a // x will be type int
	fmt.Println("B (conversion type):", b)
	fmt.Println("X (conversion type):", x)
}

func loops() {
	sum := 0
	for i := 0; i < 10; i++ { //traditional for loop
		sum += i
	}

	whileLoop := 1
	for whileLoop < 1000 { // while loops
		whileLoop += whileLoop
	}
	fmt.Println("my sum:", sum)
	fmt.Println("my sum:", whileLoop)
}

func maps() {
	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices, "square")
	fmt.Println("Map:", vertices)
	// iterate through array
	for index, value := range vertices {
		fmt.Println("index:", index, "value:", value)
	}
}