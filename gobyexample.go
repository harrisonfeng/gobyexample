package main

/*

Go is an open source programming language designed for building simple,
fast, and reliable software.

This program is borrowed from Go by Example (https://gobyexample.com)
which is modified by Harrison Feng <feng.harrison@gmail.com>.

*/

import (
	"encoding/json"
	"fmt"
	// "os"
)

// For StringFormat
type point struct {
	x, y int
}

// For JsonParser
type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func JsonParser() {
	fmt.Println("**************** JsonParser ****************")
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	fmt.Println("**************** JsonParser ****************")
}

func StringFormat() {
	p := point{1, 2}
	fmt.Println("**************** StringFormat ****************")
	fmt.Printf("%v\n", p)    // Prints an instance of point struct.
	fmt.Printf("%+v\n", p)   // Prints an instance of point struct including the struct's field names.
	fmt.Printf("%#v\n", p)   // Prints syntax representation of the value.
	fmt.Printf("%T\n", p)    // Prints the type of the value.
	fmt.Printf("%t\n", true) // Formats boolean
	fmt.Printf("%d\n", 123)  // Formats the standard, base-10 formatting.
	fmt.Printf("%b\n", 14)   // Prints a binary representation.
	fmt.Printf("%c\n", 65)   // Prints the characer corresponding to the given integer.
	fmt.Println("**************** StringFormat ****************")
}

func main() {
	StringFormat()
	JsonParser()
}
