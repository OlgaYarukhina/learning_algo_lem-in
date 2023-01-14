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
	fmt.Println(cells[0].Tunnels[0])	


	// 2. create Path

 var c []*functions.Cell

	path := functions.BildAllPaths(start, end, c)
	fmt.Println("Arr of pathes:")
	fmt.Println(path)
	
	
	
	

}