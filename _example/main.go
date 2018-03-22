package main

import (
	"fmt"

	"github.com/nicolascb/sortstruct"
)

type Person struct {
	Name         string
	ValueInt64   int64
	ValueInt     int
	ValueFloat32 float32
	ValueFloat64 float64
}

func main() {
	// First set data
	p := setData()
	// Call function to sort (struct, field, asc)
	p = SortMyStruct(p, "Name", true)
	fmt.Println(p)
}

func setData() []Person {
	p := []Person{}

	p = append(p, Person{
		Name:         "Nicolas",
		ValueInt64:   25,
		ValueInt:     25,
		ValueFloat32: 25.32,
		ValueFloat64: 25.32,
	})
	p = append(p, Person{
		Name:         "João",
		ValueInt64:   3,
		ValueInt:     3,
		ValueFloat32: 3.02,
		ValueFloat64: 3.02,
	})
	p = append(p, Person{
		Name:         "Natally",
		ValueInt64:   63,
		ValueInt:     63,
		ValueFloat32: 100.32,
		ValueFloat64: 100.32,
	})
	return p
}

// Convert struct to interface
func StructToInterface(p []Person) []interface{} {
	var i []interface{}
	for _, a := range p {
		i = append(i, a)
	}

	return i
}

// Sort interface and return struct
func SortMyStruct(p []Person, field string, asc bool) []Person {
	// Convert struct
	i := StructToInterface(p)

	// Call package and sort
	sortstruct.By(sortstruct.Prop(field, asc)).Sort(i)

	// Convert interface to struct
	var px []Person
	for _, v := range i {
		px = append(px, v.(Person))
	}

	return px
}
