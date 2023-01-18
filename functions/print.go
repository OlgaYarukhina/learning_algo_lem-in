package functions

import "fmt"


func Print (n int, s [][]*Cell) {
	fmt.Print(s)

	if len(s) == 1 {
		for _, cell := range s[0] {
			fmt.Print("Cell")
			fmt.Print(*&cell.Name)
		}
	} else {

	}

}