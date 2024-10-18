package main

import (
	"errors"
	"fmt" // to print
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello World!")

	/*CONSTANTS, VARIABLES, AND BASIC DATA TYPES*/

	var intNum int = 32767
	// variable name type, must assign right away, because don't want pointless vars,int default to 32, 64
	intNum = intNum + 1

	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	// float must be defined 32 or 64
	// cant do math on mixed types, must cast
	// integer devision = floor

	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)

	/// remainder = %

	var myString string = `Hello
	World`

	fmt.Println(myString)

	fmt.Println(len("ate"))
	// characters outside of regular ascii set = takes more than one byte

	// so to fix this

	fmt.Println(utf8.RuneCountInString("γ")) // have to import 	"unicode/utf8"

	// rune is another data type in go that represents characters

	var myRune rune = 'a'
	fmt.Println(myRune) // outputs 97, not a typo, runes are weird

	var myBoolean bool = false
	fmt.Println(myBoolean) // false

	var intNum3 int
	fmt.Println(intNum3)
	// variable does not need to be assigned to a value, but NEEDS to be used, go will set default values depending on type
	//  default for all (unsigned)int/rune = 0, string = '', bool = false

	var myVar = "text" // omit type if set value right away, because its inferred
	// shorthand myVar := "text"

	fmt.Println(myVar)

	// can initalize multible vars same way initialize one var

	var1, var2 := 1, 2      // do short hand only when assigning a value right a way, not fot passing., tbh od way is better with var and type info = READIBILTIY
	fmt.Println(var1, var2) // output = 1, 2

	const myConst string = "const value" // const = alt to vars, but cant change value when created and must initalize with value

	/* FUNCTIONS AND CONTROL STRUCS*/

	var printValue string = "i love you tracy mai"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var res, remain, errs = intDivision(numerator, denominator)

	/*

		if errs != nil {
			fmt.Printf(errs.Error()) // dont when error is not nill, so print error
		} else if remain == 0{
			fmt.Printf("The result of the integer division is %v", res)
		} else {
			fmt.Printf("The result of the integer division is %v with remainder %v", res, remain)
		// printf allows formatting
		}

	*/

	//WITh switch

	switch {
	case errs != nil:
		fmt.Printf(errs.Error()) // dont when error is not nill, so print error // then break because implied after each case
	case remain == 0:
		fmt.Printf("The result of the integer division is %v", res)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", res, remain)
		// printf allows formatting
	}

	switch remain { // conditional switch that checks the value of remain
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}

	/* ARRAYS, SLICES, MAPS, AND LOOPS*/

	var intArr [3]int32      // arrays = fixed ength, same type, indexable, contiguous in memory
	fmt.Println(intArr[0])   //zeroth index
	fmt.Println(intArr[1:3]) // first and second index

	fmt.Println(&intArr[2]) // prints out location of memory

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age:%v \n", name, age)
	}

	// loop over keys, then age = values thatasreturned if we want to,
	// normally loops over index, then like add age for value, add another variable name for value retrieval
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

// if function is returning a value, we must define the value, in this case int
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error // default value of nil
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero") // .Net and import "errors" allows to print new error message
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
