package functions

import (
	"fmt"
	"regexp"
	"strings"
)

type Cell struct {
	Name   string
	Tunels [] *Cell
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

func FindTunels (tunArr []string, names []string) {
	//var cell []Cell
	var s []*Cell
	
	cellList:= make (map[string][]*Cell)


	for i:=0; i<len(tunArr); i++ {
		t := strings.Split(tunArr[i], "-")
		w1, w2:= t[0], t[1]
		for j:=0; j<len(names); j++ {
			if names[i] == t[0] {
				s = append(s, &w2)
			}
			cellList[w1] = s
		}

	}
	fmt.Println(names)
	fmt.Println(tunArr)


}