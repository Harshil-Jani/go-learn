/*
As a coder in a language named after fungus whose mascot is crab
I can confirm that I am having panic attacks for GoLang and its
variable scopes.
*/

package main

import (
	"fmt"
)

// Structs
type OpenSource struct {
	openMR      int
	programName string
	projectName []string
}

// Inheritence
type Software struct {
	OpenSource
	companyName string
}

// constant iota's are block specific
const (
	a1 = iota
	b1 = iota
	c1 = iota
)

const (
	a2 = iota
	b2
	c2
)

const (
	// This works too
	// _ = iota + 5
	// Dropping the iota value
	_ = iota

	a3 = iota
	// This don't work here tho
	_  = iota + 9
	b3 = iota
	// Works here
	c3 = iota + 28
)

func main() {
	// Declaring Variable with type
	var i int = 17
	// Declaring Variable with type assumption
	j := 43
	// Reassigning the variable
	i = 28

	// Variable Block to write cleaner code
	var (
		a string  = "harshil"
		b float64 = 56.19
		c bool    = true
	)

	// Printing
	fmt.Println(i, j, a, b, c)
	fmt.Printf("%v, %T\n%v, %T\n%v, %T\n%v, %T\n%v, %T\n", i, i, j, j, a, a, b, b, c, c)

	// Typecasting
	var x float64
	x = float64(i)
	fmt.Printf("%v, %T\n", x, x)

	// Converting String to bytes
	s := "Harshil Jani"
	h := []byte(s)
	fmt.Printf("%v, %T\n", h, h)

	// Declaring constant variables
	const PI float64 = 3.14
	fmt.Printf("%v, %T\n", PI, PI)

	// Scopes of iota in constants
	fmt.Println(a1, b1, c1, a2, b2, c2, a3, b3, c3)

	// Complex Numbers
	var cmplx complex64 = 1 + 2i
	var cmplx2 complex128 = complex(9, 10)
	fmt.Println(cmplx, real(cmplx), imag(cmplx), cmplx2)

	// Strings are UTF Encoded in Go.
	var name string = "Harshil"
	// name[1] = "J"
	fmt.Println(name[0], string(name[0]))

	// Array [size]data_type{data1,data2,data3}
	grades := [3]int{1, 2, 3}
	names := [...]int{1, 2, 3, 4, 5}
	grades[1] = 4

	// Array slicing
	name1 := names[:]
	name2 := names[1:]
	name3 := names[1:2]
	fmt.Println(name1, name2, name3)
	// Array of Array
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(grades, names, identityMatrix)

	// Map
	gradeSheet := map[string]int{
		"DIC": 8,
		"WMC": 24,
		"OFC": 24,
	}
	_, ok1 := gradeSheet["GMT"]
	_, ok2 := gradeSheet["DIC"]
	fmt.Println(ok1, ok2)
	fmt.Println(gradeSheet["DIC"])
	delete(gradeSheet, "DIC")
	fmt.Println(gradeSheet["DIC"])

	// Structs
	aOpenSource := OpenSource{
		openMR:      3,
		programName: "Google Summer of Code",
		projectName: []string{
			"Qaul",
			"Geant4",
			"KDE Project Health",
		},
	}

	bOpenSource := struct{ name string }{name: "Harshil Jani"}
	fmt.Println(aOpenSource, bOpenSource)
	fmt.Println("bfuf")

	// Inheritence
	aSoftware := Software{
		companyName: "Meta",
	}
	aSoftware.openMR = 12
	aSoftware.programName = "XROS"
	fmt.Println(aSoftware)

	// Conditionals
	if true {
		fmt.Println("Yes it works tho")
	}

	/*
	   		defer : Defer is used to ensure that a function call is performed
		   later in a program's execution, usually for purposes of cleanup.

		   It works in LIFO prinicple.
		   Expected output : [end, middle, start]
		*/
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

	// A good snippet to learn about magic with defer.
	magic := "start"
	defer fmt.Println(magic)
	magic = "end"
	
	// Panic
	panic("Something bad happened")
}
