package functions

import (
	"fmt"
)


func AllPaths(s ,f *Cell, c []*Cell) ([][]*Cell){
	var p, allp [][]*Cell
	var eachP []*Cell
	eachP = c
	eachP = append(eachP, s)

	lable:
	for i:=0; i<len(s.Tunnels); i++ {
		eachPP := eachP
		if s.Tunnels[i] == f {
		eachPP = append(eachPP, s.Tunnels[i])
		allp = append(allp, eachPP)
		continue
		} else {
		for j:= 0; j<len(eachP); j++ {   //check if we have already been in that cell
	    if eachP[j] == s.Tunnels[i] {
		continue lable
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
	
	var check bool
	allp := p

	for w := 0; w < (len(allp)); w++ {
		fmt.Print("w")
		fmt.Println(w)
		var passedCelles []*Cell 
		var str [][]*Cell

		str = append(str, allp[w])
		sol = append(sol, str)

		for _, allcell := range allp[w] {
			passedCelles = append(passedCelles, allcell)
		}
		fmt.Println(passedCelles)

     lable:
		for r := w+1; r<len(allp); r++{ 
			fmt.Println(r)
			for k:=0; k<len(allp[r]); k++ {		
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
			fmt.Println(passedCelles)
		}
	
	}
	}
	
	fmt.Println("Good:")
 //ways without bottle
	return sol
}

/*
func bildPath (s ,f *Cell, p []*Cell)[][]*Cell{
	eachP := p


return eachP
 && cell != srtd[i][0] && cell != srtd[i][len(srtd[i])-1] 
}


sort.Slice(srtd, func(i, j int) bool {
		return len(srtd[i]) < len(srtd[j])
	})
*/

