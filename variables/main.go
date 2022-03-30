package main

/*
there could be an group of imports ex:

import (
	"fmt"
	"math"
)

exported names starting from the uppercase letter ex:

from "math module" math.pi is incorrect,
correct way is math.Pi that means its exported.
*/

import (
	"fmt"
	"math"
)

/*
variables construction:

var variableName type <- initialization

variables can be initialized:

if variables has initializators type is optional
you can initialize variables in groups

var x, y, z int = 1, 2, 3

rules:

you cant redeclare them, but you can shadow them.
all of variables must be used!

visibility:

lower case first letter = package scope
upper case first letter = exported

variables could be declared with short declaration:

varName := value

type will be inferred from the value
*/

/* available types
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias dla uint8

rune // alias dla int32

float32 float64

complex64 complex128

default values:

0 for the the numeric types
false for the bools
"" for the striing

to convert variable to another type use assertion T(variable) ex:

var x int = 42
r := float64(x)

if you want to declare a constant value you can use "const"

const Factor = 30

consts could not be created using short assignment operator

const are also used to specify very precise numbers (if you wont specify type it will be inferred from context)
*/

const Pi = 3.14

func caclculateCircleArea(radius int) float64 {
	return Pi * math.Pow(float64(radius), 2)
}

func main() {
	var greeting string = "Hello World"
	fmt.Println(greeting)

	var x int
	var y float64
	var z bool
	var f string
	fmt.Printf("defaults %v %v %v %q\n", x, y, z, f)

	const circleRadius = 10
	circleArea := caclculateCircleArea(circleRadius)
	fmt.Println("circle area from circle of radius 10: ", circleArea)

}
