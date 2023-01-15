package functions

import (
	
	"sort"
)


func BildAllPaths(s ,f *Cell, c []*Cell) ([][]*Cell){
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
			nextStep := BildAllPaths(s.Tunnels[i], f, eachP)
			allp = append(allp, nextStep...)
		}
	}
	p = append(p, allp...)
	return p
}


func BildAllSolution (p [][]*Cell) ([][]*Cell) {
	srtd := p
	sort.Slice(srtd, func(i, j int) bool {
		return len(srtd[i]) < len(srtd[j])
	})


	
	return s
}

/*
func bildPath (s ,f *Cell, p []*Cell)[][]*Cell{
	eachP := p


return eachP
}
*/

