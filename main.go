package main

import (
	"fmt"

	"01.kood.tech/git/oyarukhi/lem-in/functions"
)

func main(){

	
	getData := functions.ReadFile() 
	n, start, end, cellsName, tun := functions.ParsData(getData)


	fmt.Println(getData[2])

	fmt.Println(n)
	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(cellsName)
	fmt.Println(tun)
	
	

}