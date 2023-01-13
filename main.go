package main

import (
	"fmt"

	"01.kood.tech/git/oyarukhi/lem-in/functions"
)

func main(){

	var allCells []functions.Cell
	
	getData := functions.ReadFile() 
	n, start, end, cells, tun := functions.ParsData(getData)
	
	for _, each := range tun {
		allCells = functions.AddTunnels(cells, each)
		cells = allCells
	}
	


	fmt.Println(n)
	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(allCells)
	
	
	
	
	

}