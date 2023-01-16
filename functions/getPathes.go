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
		eachPP := eachP
		if s.Tunnels[i] == f {
		eachPP = append(eachPP, s.Tunnels[i])
		allp = append(allp, eachPP)
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
	fmt.Println(ar)
	fmt.Print("here")
	fmt.Println(srtd[0][1])


	Lable:
	for i:=1; i<len(srtd); i++ {
		//fmt.Println(i)
		//fmt.Println(srtd[i])
		check := false
		for j:=1; j<len(srtd[i]); j++{ 
			//fmt.Println(srtd[i][j])
			for k:=0; k<len(ar); k++ {		
				for _, cell := range ar[k]{
				if cell == srtd[i][j] && cell != srtd[0][0] && cell != srtd[0][len(srtd[0])-1] { //without start and finish
					fmt.Println("here1")
					fmt.Println(srtd[i][j])
					fmt.Println(srtd[0][0])
					fmt.Println(srtd[0][len(srtd[0])-1] )
					check = true
					continue Lable
					}		
				}
				}	
				if check == false {
					ar = append(ar, srtd[i])
					fmt.Println("arr")
						fmt.Println(ar)
			}
		

		}
	}	
	fmt.Println("Good:")
fmt.Println(ar) //ways without bottle
	return sol
}

/*
func bildPath (s ,f *Cell, p []*Cell)[][]*Cell{
	eachP := p


return eachP
 && cell != srtd[i][0] && cell != srtd[i][len(srtd[i])-1] 
}
*/

