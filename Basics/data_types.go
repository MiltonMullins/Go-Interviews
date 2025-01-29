package main

import "fmt"

const PI = 3.14

func dataTypesFun() {

	fmt.Println("\n-----DATA TYPES-----")
	//Constant - is unchangeable and read-only.
	fmt.Printf("CONSTANT PI: %f\n", PI)

	//Boolean
	fmt.Println("\nBoolean")
	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true

	//Integer
	fmt.Println("\nInteger")
	//The default type for integer is int. If you do not specify a type, the type will be int.
	/*
		Depends on platform:
		32 bits in 32 bit systems and
		64 bit in 64 bit systems
	*/
	var i1 int = 9223372036854775807 // -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems
	var i2 int8 = 127                // -128 to 127
	var i3 int32 = 32767             // -32768 to 32767
	var i4 int64 = 2147483647        // -2147483648 to 2147483647

	fmt.Printf("Type: %T, value: %v\n", i1, i1)
	fmt.Printf("Type: %T, value: %v\n", i2, i2)
	fmt.Printf("Type: %T, value: %v\n", i3, i3)
	fmt.Printf("Type: %T, value: %v\n", i4, i4)

	//Uint - can only store non-negative values
	fmt.Println("\nuint")
	var u1 int = 9223372036854775807 // 0 to 2147483647 in 32 bit systems and 0 to 9223372036854775807 in 64 bit systems
	var u2 int8 = 127                // 0 to 127
	var u3 int32 = 32767             // 0 to 32767
	var u4 int64 = 2147483647        // 0 to 2147483647

	fmt.Printf("Type: %T, value: %v\n", u1, u1)
	fmt.Printf("Type: %T, value: %v\n", u2, u2)
	fmt.Printf("Type: %T, value: %v\n", u3, u3)
	fmt.Printf("Type: %T, value: %v\n", u4, u4)

	//Float
	fmt.Println("\nFloat")
	var f1 float32 = 123.78 //-3.4e+38 to 3.4e+38.
	var f2 float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", f1, f1)
	fmt.Printf("Type: %T, value: %v\n", f2, f2)
	var f3 float64 = 1.7e+308 //-1.7e+308 to +1.7e+308.
	fmt.Printf("Type: %T, value: %v\n", f3, f3)

	//String
	fmt.Println("\nString")
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)

}
