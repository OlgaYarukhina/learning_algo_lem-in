package main

import (
	"fmt"

	"01.kood.tech/git/oyarukhi/lem-in/functions"
)

func main() {

	//1. create Cells with tunnels

	getData := functions.ReadFile()
	n, start, end, cells := functions.ParsData(getData)
	fmt.Println(cells)
	fmt.Println(start)

	// 2. create all possible pathes

	var c []*functions.Cell
	path := functions.AllPaths(start, end, c)

	// 3. strategies

	allSollutions := functions.AllSolutions(path)

	// 4. choose best solution

	solPath, height := functions.GetSolution(n, allSollutions)
	fmt.Println("Best solution:")
	//fmt.Println(solPath)

	for _, r := range solPath {
		fmt.Println("next")
		for _, g := range r {
			fmt.Println(g)
		}

	}

	// 5. Print

	functions.Run(n, height, solPath)

}

/*
for _, r := range path {
		fmt.Println("next")
		for _, g := range r{
			fmt.Println(g)
		}

	}
		fmt.Println("All sol:")

	for _, r := range allSollutions {
		fmt.Println("next")
		for _, g := range r {
			for _, K := range g {
			fmt.Println(K)
			}
		}
	}

*/
