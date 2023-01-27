package functions

import (
	"fmt"
	"strconv"
)

type Ants struct {
	Id       int
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

	// 2. put ants into the start

	antsSprint := make([]Ants, n)
	q := 0

	for i := 0; i < len(groups); i++ {
		for j := 0; j < len(groups[i]); j++ {
			antsSprint[q].Id = i
			antsSprint[q].Name = groups[i][j]
			antsSprint[q].Path = s[i]
			antsSprint[q].Position = antsSprint[q].Path[0]
			q++
		}
	}

	// 3. run!

	for s := 0; s < h; s++ { // raw of printing
		pos := make([][]*Cell, len(groups)) // not empty cells
		var str, res string

	lable:
		for k := 0; k < len(antsSprint); k++ {
			for p := 0; p < len(antsSprint[k].Path); p++ {
				if antsSprint[k].Path[p] == antsSprint[k].Position && p != len(antsSprint[k].Path)-1 {
					for _, check := range pos[antsSprint[k].Id] {
						if check == antsSprint[k].Path[p+1] {
							continue lable
						}
					}
					antsSprint[k].Position = antsSprint[k].Path[p+1]
					pos[antsSprint[k].Id] = append(pos[antsSprint[k].Id], antsSprint[k].Position)
					str = antsSprint[k].Name + *&antsSprint[k].Position.Name + " "
					res = res + str
					continue lable

				} else if antsSprint[k].Path[p] == antsSprint[k].Position && p == len(antsSprint[k].Path)-1 {
					continue lable
				}
			}
		}
		if res != "" {
			fmt.Println(res)
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
