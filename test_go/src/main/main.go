package main

import (
	"data"
	"log"
	simplifier "github.com/yrsh/simplify-go"
)

func main() {
	orig, simpl_5f, simpl_3t := data.Get()

	simpl_5f_test := simplifier.Simplify(orig, 5, false)
	simpl_3t_test := simplifier.Simplify(orig, 3, true)

	if simplifier.CompareSlices(simpl_5f_test, simpl_5f) {
		log.Print("5 false | Test Ok")
	} else {
		log.Print(" 5 false | Something went wrong")
	}

	if simplifier.CompareSlices(simpl_3t_test, simpl_3t) {
		log.Print("3 true | Test Ok")
	} else {
		log.Print("3 true | Something went wrong")
	}

}
