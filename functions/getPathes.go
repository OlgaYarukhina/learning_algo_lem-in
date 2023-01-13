package functions

import "fmt"

func BildAllPaths(s ,f *Cell)[][]*Cell{
	var p [][]*Cell
	var eachP []*Cell

	for i:=0; i<len(s.Tunnels); i++ {
		step := s.Tunnels[i]
		eachP = append(eachP, step)
		fmt.Println("Each:")
		fmt.Println(eachP)
		if s.Tunnels[i] == f {
			break
		} else {
		nextStep := bildPath(step, f)
		eachP = append(eachP, nextStep...)
		}
	}

	p = append(p, eachP)
	
	return p
}

func bildPath (s ,f *Cell)[]*Cell{
	var eachP []*Cell
	eachP = append(eachP, s) //+2
	fmt.Println("Each2:")
		fmt.Println(eachP)
	for i:=0; i<len(s.Tunnels); i++ {
		step := s.Tunnels[i] // +3
		eachP = append(eachP, step)
		if s.Tunnels[i] == f {
			eachP = append(eachP, step)
			break
		} else if s.Tunnels[i] == {
			break
		}else {
		nextStep := bildPath(step, f)
		eachP = append(eachP, nextStep...)
		}
	}
return eachP
}