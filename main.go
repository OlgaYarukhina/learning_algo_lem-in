package main

import (
	"fmt"

	"01.kood.tech/git/oyarukhi/lem-in/functions"
)

func main() {

	//1. create Cells with tunnels

	getData := functions.ReadFile()
	n, start, end := functions.ParsData(getData)

	// 2. create all possible pathes

	var c []*functions.Cell
	path := functions.AllPaths(start, end, c)

	// 3. all possible pathes

	allSollutions := functions.AllSolutions(path)

	// 4. choose best solution

	solPath, height := functions.GetSolution(n, allSollutions)

	// 5. print

	//5.1. print file data
	for i := 0; i < len(getData); i++ {
		fmt.Println(getData[i])
	}
	fmt.Println()

	//5.2. print results
	functions.Run(n, height, solPath)
}

/*
	for _, r := range allSollutions {
		fmt.Println("next")
		for _, g := range r {
			for _, K := range g {
			fmt.Println(K)
			}
		}
	}

*/
