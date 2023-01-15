package main

import (
	"fmt"
	"sort"

	"01.kood.tech/git/oyarukhi/lem-in/functions"
)

func main(){

	

	//1. create Cells with tunnels
	
	getData := functions.ReadFile() 
	n, start, end, cells:= functions.ParsData(getData)

	fmt.Println(n)
	fmt.Println(cells[0].Tunnels[0])	

	fmt.Println(cells)


	// 2. create Path

    var c []*functions.Cell

	path := functions.AllPaths(start, end, c)
	fmt.Println("Arr of pathes:")
	fmt.Println(path)
	for _, r := range path {
		fmt.Println("next")
		for _, g := range r{
			fmt.Println(g)
		}
		
	}

	// 3. strategy
	
	allSollutions := functions.AllSolutions(path)
	
	sort.Slice(path, func(i, j int) bool {
		return len(path[i]) < len(path[j])
	})
		fmt.Println(allSollutions)
}