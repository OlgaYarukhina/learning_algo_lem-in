package functions

import (
	"fmt"
	"strconv"
)

type Ants struct {
	Name   string
	Position *Cell
}


func Print (n, h int, s [][]*Cell) {
	fmt.Println(h)
  //  var slicAnts [][]string

	nAnts := n

	var groups [][]string
	
	for g := 0; g < len(s); g++ {
		fmt.Println("Here")
		fmt.Println(g)
		var group []string           //group of ants in the same way
		if g == len(s)-1 {
			for a := 1; a <= nAnts; a++ {
				fmt.Println("Here2")
				aStr := strconv.Itoa(n-nAnts+a)
				str := "L" + aStr + "-" 
				group = append(group, str)
				//fmt.Println(group)
			}
		} else {
			for a := 1; a <= h-len(s[g]); a++ {
				fmt.Println("Here1")
				fmt.Println(nAnts-len(s[g]))
				aStr := strconv.Itoa(n-nAnts+a)
				str := "L" + aStr + "-" 
				group = append(group, str)
				fmt.Println(group)
			}
			nAnts = nAnts - (h-len(s[g]))
			fmt.Println(h-len(s[g]))
			
		}
		groups = append(groups, group)
	}

	fmt.Println(groups)

	// put ants into the start

	// start := new(Ants)

	 for i:=0; i < len(groups); i++ {
		for j := 0; j < len(groups[i]); j++ {

		}
	 }

}





/*


	//try 1

	// create slice of ants in each path

	sAnt := 1                                                     // number of ant in start of each way


	for i := 0; i < len(s); i++ {
		nAnts := h - len(s[i])                                   // number of ant in start of each way                                                           
		slicAnts = append(slicAnts, slice(sAnt, nAnts, s[i]))
		sAnt = h-len(s)                                          // number of ants in path 
	}

	fmt.Println(slicAnts)
} 


	// create slice of ants in each path

	func slice (s, n int, p []*Cell) []string {  
		var pStr [] string
		var str string                        // startN number of ant in start of each way
	if n > 0 {
		for i := 1; i < len(p); i++ {
			numNext := strconv.Itoa(s+i-1)
		str = "L" + numNext + "-" + *&p[i].Name + " "
		pStr = append(pStr, str)
		}
	}
	return pStr
}




for i := 0; i < len(s); i++ {
		nAnts := h - len(s[i])-1
		 
		
		for j := 1; i < len(s[i]); j++ {

			len(s[i])

			
			if nAnts > 0 {
			fmt.Print("L",n,"-",*&s[0][j].Name," ")
			}
			
		}
		
	}
	
	// try 1

	fmt.Print("L",1,"-",*&s[0][0].Name," ")


	for i := 1; i <= i; i++ {       // number of ants

		for j := 0; j < len(s); j++ {    // number of path (2)
			if  
			
			n1 := h - len(s[j])
		}
	
	}


	

	*/