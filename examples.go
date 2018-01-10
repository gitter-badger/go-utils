package main

import (
	"fmt"

	"github.com/mailoman/go-utils/examples"
	"github.com/mailoman/go-utils/mapping"
)

func main() {

	// Example 1, no strict mapping rules at all
	in := examples.InputExample1{
		Str: "1",
		I32: 32,
		I64: 64,
		Boo: true,
		F32: 32.23,
	}

	out := &examples.OutputExample1{}

	// Simple usage
	mapping.MapAllFields(in, out, nil)
	fmt.Printf("%+v ==> %+v\n", in, *out)
}
