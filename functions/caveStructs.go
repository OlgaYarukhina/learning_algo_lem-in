package functions

import (
	"fmt"
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

var Data Cell

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

	
func CreateTunn () {
	cellList:= make (map[string][]*Cell)


	for i:=0; i<len(tunArr); i++ {
		t := strings.Split(tunArr[i], "-")
		w1, w2:= t[0], t[1]
		for j:=0; j<len(names); j++ {
			if names[i] == t[0] {
				var s []*Cell
				s = append(s, *w2)
			}
			cellList[w1] = s
		}

	}
	fmt.Println(names)
	fmt.Println(tunArr)


}