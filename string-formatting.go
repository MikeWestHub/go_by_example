package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	fmt.Printf("%v\n", p)                      // instance of struct
	fmt.Printf("%+v\n", p)                     // include struct field names
	fmt.Printf("%#v\n", p)                     // source code snippet
	fmt.Printf("%T\n", p)                      // type of value
	fmt.Printf("%t\n", true)                   // boolean
	fmt.Printf("%d\n", 123)                    // base-10 formatting
	fmt.Printf("%b\n", 14)                     // binary
	fmt.Printf("%c\n", 33)                     // corresponding character
	fmt.Printf("%x\n", 456)                    // hex encoding
	fmt.Printf("%f\n", 78.9)                   // float
	fmt.Printf("%e\n", 12340000.0)             // scientific notation for float (e)
	fmt.Printf("%E\n", 12340000.0)             // variation of scientific notation (E)
	fmt.Printf("%s\n", "\"string\"")           // basic string printing
	fmt.Printf("%q\n", "\"string\"")           // double quote strings
	fmt.Printf("%x\n", "hex this")             // render in base-16 with two ouput characters per byte of input
	fmt.Printf("%p\n", &p)                     // representation of a pointer
	fmt.Printf("|%6d|%6d|\n", 12, 3.45)        // specify the width of an integer, right justified
	fmt.Printf("|%6.2d|%6.2d|\n", 12, 3.45)    // specify the width of an integer, width.precision syntax
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // left justify
	fmt.Printf("|%6s|%6s|\n", "foo", "b")      // specify the width of an string, right justified
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")    // specify the width of an string, left justified

	s := fmt.Sprintf("a %s", "string") // Sprintf returns a string without printing it anywhere (Printf prints to stdout)
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error") // specify different io.Writers
}
