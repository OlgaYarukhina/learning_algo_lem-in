package main

import (
	"fmt"

	"01.kood.tech/git/oyarukhi/lem-in/functions"
)

func main(){

	
	getData := functions.ReadFile() 
	n, start, end, cellsName, tun := functions.ParsData(getData)
	//functions.BildCave(start, end, cellsName, tun)
	cells := functions.FindTunels(tun, cellsName)


	fmt.Println(n)
	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(cells)
	
	
	

}