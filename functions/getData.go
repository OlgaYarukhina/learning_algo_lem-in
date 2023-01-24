package functions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Cell struct {
	Name   string
	Tunnels [] *Cell
}


func ReadFile() []string {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run . [exampleXX.txt]")
		os.Exit(1)
	}
	d, err := os.Open("examples/" + args[0])
	if err != nil {
		fmt.Println("Can not read file", err)
	}
	defer d.Close()

	scanner := bufio.NewScanner(d)
	var scanLines []string
	for scanner.Scan(){
		scanLines = append(scanLines, scanner.Text())
	}
	return scanLines //get lines
}

func ParsData (scan []string) (int, *Cell, *Cell, []Cell) {
	var allCells []Cell
	var start, end *Cell
	var cellsArr []Cell
	var tunArr [] string
	var st, fn string
	nSt := scan[0]
	n, err := strconv.Atoi(nSt)
	if err != nil {
		fmt.Println("Invalid file format", err)
		os.Exit(1)
	}

	regexpCells := regexp.MustCompile(`(\S+)\s(\d+\s\d+)`)
	regexpTunells := regexp.MustCompile(`(\S+)-(\S+)`)
	

	for i:=1; i < len(scan); i++ {
		foundCells := regexpCells.FindAllString(scan[i], -1)
		foundTunells := regexpTunells.FindAllString(scan[i], -1)
		if len(foundCells)>0 && scan[i] == foundCells[0] {
			l := strings.Split(scan[i], " ")
			cell := Cell{Name: l[0]}
			cellsArr = append(cellsArr, cell)
			if scan[i-1] == "##start"{
				l := strings.Split(scan[i], " ")
				st = l[0]
			} else if scan[i-1] == "##end" {
				l := strings.Split(scan[i], " ")
				fn = l[0]
			} 
		} else if len(foundTunells)>0 && scan[i] == foundTunells[0] {
			tunArr = append(tunArr, scan[i])
		} 
	}
	for _, each := range tunArr { //add tunnels
		allCells = AddTunnels(cellsArr, each)
		cellsArr = allCells
	}

	for i, each := range cellsArr {
		if each.Name == st {
			start = &cellsArr[i]
		} else if each.Name == fn {
			end = &cellsArr[i]
		}
	}

	return n, start, end, cellsArr
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




/*
type AntHill struct {
	Start *Cell
	End   *Cell
	Cells []Cell
}*/