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
	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(cells)
	fmt.Println(cells[0].Tunnels[0])	
	fmt.Println(cells[0].Tunnels[0].Tunnels[1])	
	fmt.Println(cells[0].Tunnels[0].Tunnels[1].Tunnels[1])	


	// 2. 
	
	
	
	
	

}