package functions

import "fmt"

func BildAllPaths(s ,f *Cell, c []*Cell) ([][]*Cell){
	var p [][]*Cell
	var eachP []*Cell
	var allp [][] *Cell

	eachP = c
	eachP = append(eachP, s)


	
Label:
		for i:=0; i<len(s.Tunnels); i++ {
			fmt.Println(i)
			if s.Tunnels[i] == f {
			eachP = append(eachP, s.Tunnels[i])
			
			break
		   } else {
			for j:= 0; j<len(eachP); j++ {   //check if we have already been in that cell
			fmt.Println("Here1")
	        if eachP[j] == s.Tunnels[i] {
				fmt.Println(eachP[j])
				fmt.Println(s.Tunnels[i])
				continue Label
			} 
			}	
			fmt.Println("Here2")
			fmt.Println(s.Tunnels[i])
			nextStep := BildAllPaths(s.Tunnels[i], f, eachP)
			
			fmt.Println("nextStep:")
			fmt.Println(nextStep)
			allp = append(allp, nextStep...)

		}
	}

	p = append(p, allp...)
	
	return p
}


/*
func bildPath (s ,f *Cell, p []*Cell)[][]*Cell{
	eachP := p


return eachP
}
*/

