package main

import "fmt"

/*
	in golang there is only one type of loop and this is "for"

	you can construct for loop with 3 parts

	1: initializator(optional) -> executed only once before loop,
	there for ex you can declare vars(variable will be block-scoped same as
	while declaring variables before if-else statments)

	2: condition -> executed before every iteration

	3: increment(optional) -> executed after every iteration


	for var := n; var <= n; var + n{
		...
	}

	if you will omit every part you will construct endless loop same as
	while in other languages
*/

func calculateSum(upTo int64) int64 {
	var i, result int64
	for ; i <= upTo; i++ {
		result += i
	}
	return result
}

/*
	if conditions are simple boolean expresions like in javascript
	the difference is you can omit () only brackets are required {}.

	if a > b{
		...
	}else{
		...
	}

	you can also declare a variable before statment

	var b int = 2
	if a := 1; a > b{
		not gonna happen but you have access to the "a" here
	}else{
		here also you can use "a"
	}

	variable declared before statment is limited only to the condition scope

*/

func alwaysTrue() bool {
	if a, b := 1, 2; a > b {
		// a && b is avaiable here
		return false
	}
	// a && b is not avaiable here
	return true
}

func main() {
	sumOfTen := calculateSum(10)
	sumOfOneHundred := calculateSum(100)
	sumOfOneThousand := calculateSum(1000)
	fmt.Println("sum of 10, 100 and 1000: ", sumOfTen, sumOfOneHundred, sumOfOneThousand)
	fmt.Println("always positive bool condition: ", alwaysTrue())
}
