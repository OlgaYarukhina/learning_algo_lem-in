package functions

import (
	"fmt"
	"strconv"
)

type Ants struct {
	Name     string
	Path     []*Cell
	Position *Cell
}

func Run(n, h int, s [][]*Cell) {
	nAnts := n

	// 1. create teams in each path 
	var groups [][]string
	for g := 0; g < len(s); g++ {
		var group []string //group of ants in the same way
		if g == len(s)-1 {
			for a := 1; a <= nAnts; a++ {
				aStr := strconv.Itoa(n - nAnts + a)
				str := "L" + aStr + "-"
				group = append(group, str)
			}
		} else {
			for a := 1; a <= h-len(s[g]); a++ {
				aStr := strconv.Itoa(n - nAnts + a)
				str := "L" + aStr + "-"
				group = append(group, str)
			}
			nAnts = nAnts - (h - len(s[g]))
		}
		groups = append(groups, group)
	}
	fmt.Println(groups)

	// 2. put ants into the start

	antsSprint := make([]Ants, n)
	q := 0

	for i := 0; i < len(groups); i++ {
		for j := 0; j < len(groups[i]); j++ {
			antsSprint[q].Name = groups[i][j]
			antsSprint[q].Path = s[i]
			antsSprint[q].Position = antsSprint[q].Path[0]
			q++
		}
	}
	fmt.Println(antsSprint)

	// 3. run!
for s := 0; s < h; s++ {  // raw of printing

	var pos []*Cell
	for k := 0; k < len(antsSprint); k++ {
		fmt.Println(k)
		for l := 0; l < len(antsSprint[k].Path); l++ {
			fmt.Println("Here")
			fmt.Println(l)
			if antsSprint[k].Path[l] == antsSprint[k].Position && l != len(antsSprint[k].Path)-1 {
				antsSprint[k].Position = antsSprint[k].Path[l+1]
				pos = append(pos, antsSprint[k].Path[l+1])
			}
		}
		
///

	}
}

}

func print() {

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
