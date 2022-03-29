package main

import "fmt"

/* functions construction:

func functionName(...args T) R  {
	...body
}

many arguments with same type can be shortened:

func add(x, y int) int {
	return x + y
}

you could return many values:

func multiply(number, factor int) (int, int){
	return number*factor, factor
}

if you will use return without any values you will return all declared
variables from the top of the functions (naked return)

func nakedReturn(sum  int) (x, y int){
	x = sum * 3 / 9
	y = sum - x
	return
}

nakedReturn(17) -> 5 12

*/

func add(x, y int) int {
	return x + y
}

func naked(number int) (x, y int) {
	x = number / 2
	y = number / 3
	return
}

func main() {
	fmt.Println("sum: ", add(15, 15))

	x, y := naked(60)

	fmt.Println("naked return:", x, y)
}
