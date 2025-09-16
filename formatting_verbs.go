package main

import "fmt"

func main() {
	// General formatting verbs
	// %v - prints the value in the default format
	// %#v - prints the value in go-syntax format
	// %T - prints the type of the value
	// %% - prints the % character

	// i := 111_505.5 // use _ to make the number more readable
	// string := "Hello"

	// fmt.Printf("%v\n", i)
	// fmt.Printf("%#v\n", i) // same as %v but with go-syntax format
	// fmt.Printf("%T\n", i)
	// fmt.Printf("%v%%\n", i)

	// fmt.Printf("%v\n", string)
	// fmt.Printf("%#v\n", string) // double quotes are used to print the string
	// fmt.Printf("%T\n", string)


	// Integer formatting verbs
	// %b Base 2 - binary format
	// %d Base 10 - decimal format
	// %+d Base 10 and always show sign - decimal format and always show sign + or -
	// %o Base 8 - octal format
	// %O Base 8, with leading 0o - octal format with leading 0o
	// %x Base 16, lowercase - hexadecimal format in lowercase
	// %X Base 16, uppercase - hexadecimal format in uppercase
	// %#x Base 16, with leading 0x - hexadecimal format with leading 0x
	// %4d Pad with spaces (width 4, rigth justified)
	// %-4d Pad with spaces (width 4, left justified)
	// %04d Pad with 0s (width 4)

	// int := 255
	// fmt.Printf("%b\n", int)
	// fmt.Printf("%d\n", int)
	// fmt.Printf("%+d\n", int)
	// fmt.Printf("%o\n", int)
	// fmt.Printf("%O\n", int)
	// fmt.Printf("%x\n", int)
	// fmt.Printf("%X\n", int)
	// fmt.Printf("%#x\n", int)
	// fmt.Printf("%4d\n", int)
	// fmt.Printf("%-4d\n", int)
	// fmt.Printf("%04d\n", int)

	// String formatting verbs

	// %s Prints the value as plain string
	// %q Prints tge value as a double-quoted string
	// %8s Pad with spaces (width 8, rigth justified)
	// %-8s Pad with spaces (width 8, left justified)
	// %x prints the value as hex dump of byte values
	// % x prints the value as hex dump of byte values with spaces

	// string := "world"

	// fmt.Printf("%s\n", string)
	// fmt.Printf("%q\n", string)
	// fmt.Printf("%8s\n", string)
	// fmt.Printf("%-8s\n", string)
	// fmt.Printf("%x\n", string)
	// fmt.Printf("% x\n", string)

	// Boolean formatting verbs
	// %t value of the boolean operator in true or false format ( same as using %v )

	// t := true
	// f := false

	// fmt.Printf("%t\n", t)
	// fmt.Printf("%v\n", t) // default value of boolean type
	// fmt.Printf("%t\n", f)

	// Float formatting verbs
	// %e - Scientific notation with e-notation
	// %f - Decimal point, no exponent
	// %.2f - Default width, precision 2
	// %6.2f - Width 6, precision 2
	// %g Exponent as needed, only if necessary

	flt := 9.18

	fmt.Printf("%e\n", flt)
	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
	fmt.Printf("%6.2f\n", flt)
	fmt.Printf("%g\n", flt)


}