package main

import (
	"fmt"

	"01.kood.tech/git/oyarukhi/lem-in/functions"
)

func main(){

	//1. create Cells with tunnels
	
	getData := functions.ReadFile() 
	n, start, end, cells:= functions.ParsData(getData)

	fmt.Println(cells)	

	// 2. create all possible pathes

    var c []*functions.Cell
	path := functions.AllPaths(start, end, c)


	// 3. strategies
	
	allSollutions := functions.AllSolutions(path)

	fmt.Println("All solution")
	fmt.Println(allSollutions)

	// 4. choose best solution

	solPath := functions.GetSolution(n,allSollutions)
	fmt.Println("Best solution")
	fmt.Println(solPath)

    // 5. Print

	functions.Print(n, solPath)

}





/*
for _, r := range path {
		fmt.Println("next")
		for _, g := range r{
			fmt.Println(g)
		}
		
	}
	*/