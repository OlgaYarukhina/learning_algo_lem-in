package main

import (
	"fmt"

	"01.kood.tech/git/oyarukhi/lem-in/functions"
)

func main(){

	
	getData := functions.ReadFile() 
	n, start, end, allCells, tun := functions.ParsData(getData)
	//functions.BildCave(start, end, cellsName, tun)
	for _, eachCell := cellsName {
		cell := CreateCells (eachCell, tun)
	}


	fmt.Println(n)
	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(cells)
	
	
	

}