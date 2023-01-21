package functions

import "fmt"



func AllPaths(s, f *Cell, c []*Cell) [][]*Cell {
	var p, allp, allPP [][]*Cell                      // allPP  all pathes of 1 way
	var eachP []*Cell
	eachP = c
	eachP = append(eachP, s)

	lable:
	for i := 0; i < len(s.Tunnels); i++ {                         
		eachPP := eachP

		if s.Tunnels[i] == f {
			
			fmt.Println(("here4"))
			fmt.Println(f)
			eachPP = append(eachPP, s.Tunnels[i])
		
			fmt.Println("Eaaaachhh")
			fmt.Println(eachPP)
			fmt.Println("ALLLLL BEFORE:")
			fmt.Println(allPP)
			allPP = append(allPP, eachPP)
			fmt.Println("ALLLLL:")
			fmt.Println(allPP)

		} else {
			for j := 0; j < len(eachP); j++ { //check if we have already been in that cell
				if eachP[j] == s.Tunnels[i] {
					fmt.Println(("here2"))
					fmt.Println(eachP[j])
				continue lable
				}
			}
			fmt.Println(("here3"))
			fmt.Println(s.Tunnels[i])
			nextStep := AllPaths(s.Tunnels[i], f, eachP)

			allp = append(allp, nextStep...)
		}

		allp = append(allp, allPP...)
	}

	p = append(p, allp...)
	return p
}

func AllSolutions(p [][]*Cell) [][][]*Cell {
	var sol [][][]*Cell
	var check bool

	allp := p

	for w := 0; w < (len(allp)); w++ {
		var passedCelles []*Cell
		var str [][]*Cell

		str = append(str, allp[w])
		sol = append(sol, str)

		for _, allcell := range allp[w] {
			passedCelles = append(passedCelles, allcell)
		}

	lable:
		for r := w + 1; r < len(allp); r++ {
			for k := 0; k < len(allp[r]); k++ {
				check = false
				for _, cell := range passedCelles {
					if cell == allp[r][k] && cell != allp[0][0] && cell != allp[0][len(allp[0])-1] { //without start and finish
						check = true
						continue lable
					}
				}
			}
			if !check {
				str = append(str, allp[r])
				sol = append(sol, str)
				for _, allcell := range allp[r] {
					passedCelles = append(passedCelles, allcell)
				}
			}
		}
	}
	return sol
}
