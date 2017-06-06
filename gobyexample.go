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
	"sshlibs"
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

func Maps() {
	fmt.Println("**************** Maps ****************")

	// Create an empty map using the builtin make.
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Declare/initialize a new map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	fmt.Println("**************** Maps ****************")
}

func JsonParser() {
	fmt.Println("**************** JsonParser ****************")
	// atomic values
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	//JSON arrays
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

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
	Maps()
}
