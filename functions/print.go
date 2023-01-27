package functions

import (
	
	"fmt"
	"strconv"
)

func Print(n, h int, s [][]*Cell) {
	nAnts := n
	groups := make([][]string, len(s))

	for a := 1; a <= nAnts; a++ {
		for i := 0; i < len(s); i++ {
			if nAnts > 0 && a+i <= n {
				numNext := strconv.Itoa(i+a)
				str := "L" + numNext + "-" 
				groups[i] = append(groups[i], str)
			} else {
				break
			}
		}
		a = a + len(s)-1
	}
	fmt.Println(groups)


	for i := 0; i < 3; i++ {
		for j := 0; j < len(groups); j++ {
			str := groups[j][i] + *&s[j][i+1].Name + " "
			fmt.Print(str)
		}
		
		fmt.Println()
	}
}

