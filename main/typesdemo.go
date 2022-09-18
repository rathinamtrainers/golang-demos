package main

import "fmt"

var softwareVersion = 2.0 // Package variables (avoid package variables)

func main() {
	fmt.Println(softwareVersion)
	softwareVersion = 3.0
	fmt.Println(softwareVersion)

	var primaryPlatform string = "US-east"
	var serverCount int = 100
	var portStatus bool = true   // default boolean value is false
	networkType := "Traditional" // literal type is taken as variable type
	var x, y int = 10, 20

	var employeeID, employeName = 100, "Rajan"

	fmt.Println(primaryPlatform)
	fmt.Println(serverCount)
	fmt.Println(portStatus)
	fmt.Println(softwareVersion)
	fmt.Println(networkType)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(employeeID)
	fmt.Println(employeName)

	var (
		studentID           int    = 200
		studentName         string = "Ajay"
		present                    = true
		mark1, mark2, mark3 int    = 90, 80, 99
	)
	fmt.Println(studentID)
	fmt.Println(studentName)
	fmt.Println(present)
	fmt.Println(mark1, mark2, mark3)

	present = false
	fmt.Println(present)

	const pi float32 = 3.14
	fmt.Println(pi)

	const π float32 = 3.14 // Using unicode characters in variables
	fmt.Println(π)

	var flag string
	fmt.Println(flag)

	var a complex64 = complex(2.5, 3.1)
	fmt.Println("Complex value is ", a)
	fmt.Println("real value is ", real(a))
	fmt.Println("imaginary value is ", imag(a))

	fmt.Println(20 + 10) // 30
	fmt.Println(20 - 10) // 10
	fmt.Println(20 * 10) // 200
	fmt.Println(20 / 10) // 2
	fmt.Println(25 / 2)  // 12
	fmt.Println(25 % 2)  // 1

	fmt.Println(25.0 / 2.0) // 12

	var hexValue int = 0xF
	fmt.Println("Hex Value is ", hexValue)
	var octValue int = 0o755
	fmt.Println("Octal value is ", octValue)
	var binValue = 0b1010
	fmt.Println("Binary Value is ", binValue)
	binValue = binValue << 2
	fmt.Println("Binary Value after left shifting is ", binValue)

	mark1 = mark1 + 5 // 95
	mark1 += 5        // 100

	// Type Casting
	var n1 int = 100
	var n2 float64 = 30.2
	var n3 float64 = float64(n1) + n2
	fmt.Println(n3)
}
