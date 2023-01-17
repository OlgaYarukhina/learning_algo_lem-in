package main

import (
	"fmt"

	"01.kood.tech/git/oyarukhi/lem-in/functions"
)

func main(){

	//1. create Cells with tunnels
	
	getData := functions.ReadFile() 
	n, start, end, cells:= functions.ParsData(getData)

	fmt.Println(n)
	fmt.Println(cells)	

	// 2. create Path

    var c []*functions.Cell
	path := functions.AllPaths(start, end, c)


	// 3. strategy
	
	allSollutions := functions.AllSolutions(path)
	fmt.Println("All strutegies:")
			fmt.Println(allSollutions)
}


/*
for _, r := range path {
		fmt.Println("next")
		for _, g := range r{
			fmt.Println(g)
		}
		
	}
	*/