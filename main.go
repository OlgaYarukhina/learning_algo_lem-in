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
	fmt.Print("start:")
	fmt.Println(start)
	fmt.Print("end:")
	fmt.Println(end)
	fmt.Println(cells)
	fmt.Println(cells[0].Tunnels[0])	
	fmt.Println(cells[0].Tunnels[0].Tunnels[1])	


	// 2. create Path


	path := functions.BildAllPaths(start, end)
	fmt.Println(path)
	
	
	
	

}