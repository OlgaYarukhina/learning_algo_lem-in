package functions

import (
	"fmt"
	"sort"
)


func AllPaths(s ,f *Cell, c []*Cell) ([][]*Cell){
	var p, allp [][]*Cell
	var eachP []*Cell
	eachP = c
	eachP = append(eachP, s)

	Label:
	for i:=0; i<len(s.Tunnels); i++ {
		if s.Tunnels[i] == f {
		eachP = append(eachP, s.Tunnels[i])
		allp = append(allp, eachP)
		continue
		} else {
		for j:= 0; j<len(eachP); j++ {   //check if we have already been in that cell
	    if eachP[j] == s.Tunnels[i] {
		continue Label
		} 
		}	
		nextStep := AllPaths(s.Tunnels[i], f, eachP)
		allp = append(allp, nextStep...)
		}
	}
	p = append(p, allp...)
	return p
}


func AllSolutions(p [][]*Cell) ([][][]*Cell) {
	var sol [][][]*Cell
	var ar [][]*Cell
	srtd := p
	
	sort.Slice(srtd, func(i, j int) bool {
		return len(srtd[i]) < len(srtd[j])
	})
	ar = append(ar, srtd[0])
	sol = append(sol, ar)

	Label:
	for i:=1; i<len(srtd); i++ {
		for j:=1; j<len(srtd[i]); j++{ //without start and finish
			for k:=0; k<len(ar); k++ {
				for _, cell := range ar[k]{
				if cell == srtd[i][j] && cell != srtd[i][0] && cell != srtd[i][len(srtd[i])-1] {
					continue Label
				} else {
					ar = append(ar, srtd[i])
				}	
			}
		}
	}
		

		

	}

	fmt.Println(ar)
	return sol
}

/*
func bildPath (s ,f *Cell, p []*Cell)[][]*Cell{
	eachP := p


return eachP
}
*/

