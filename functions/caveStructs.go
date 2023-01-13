package functions

import (
	"strings"
)

type Cell struct {
	Name   string
	Tunnels [] *Cell
}


type AntHill struct {
	Start *Cell
	End   *Cell
	Cells []Cell
}




/*
func BildCave (start string, end string) {
	Cave := AntHill {
		Start: start,
		End: *end,
		Cells: bildCell(),
	}
}

func bildCell(cellsName, tun) []Cell {
	
}
*/

func CreateCells (line string) *Cell {
	l := strings.Split(line, " ")		
	cell := Cell{Name: l[0]}
return &cell
}


	
func AddTunnels (c []Cell, t string) []Cell {
	tn := strings.Split(t, "-")
		w1, w2:= tn[0], tn[1]
	for i:=0; i<len(c); i++ {
		if c[i].Name == w1 {
		for j:=0; j<len(c); j++ {
			if c[j].Name == w2 {
				c[i].Tunnels = append(c[i].Tunnels, &c[j])
				c[j].Tunnels = append(c[j].Tunnels, &c[i])
			}
		}
	}
}
return c
}