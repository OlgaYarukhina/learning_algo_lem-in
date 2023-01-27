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
	Name    string // 1234
	Tunnels []*Cell
}

//1. create Cells with tunnels
func ReadFile() []string {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run . example00.txt")
		os.Exit(0)
	}
	d, err := os.Open("examples/" + args[0])
	if err != nil {
		fmt.Println("Can not read file: no such file or directory")
		os.Exit(0)
	}
	defer d.Close()

	scanner := bufio.NewScanner(d)
	var scanLines []string
	for scanner.Scan() {
		scanLines = append(scanLines, scanner.Text())

	}
	return scanLines //get lines
}

func ParsData(scan []string) (int, *Cell, *Cell) {
	var allCells []Cell
	var start, end *Cell
	var cellsArr []Cell
	var tunArr []string
	var st, fn string
	nSt := scan[0]
	n, err := strconv.Atoi(nSt)
	if err != nil {
		fmt.Println("ERROR: Invalid file format", err)
		os.Exit(1)
	}

	regexpCells := regexp.MustCompile(`[\d\w]+\s*[\d\w]*\s*[\d\w]*`)
	regexpTunells := regexp.MustCompile(`(\S+)-(\S+)`)

	for i := 1; i < len(scan); i++ {
		foundCells := regexpCells.FindAllString(scan[i], -1)
		foundTunells := regexpTunells.FindAllString(scan[i], -1)
		if len(foundCells) > 0 && scan[i] == foundCells[0] {
			l := strings.Split(scan[i], " ")
			if len(l) != 3 {
				fmt.Println("ERROR: invalid data format, check each cell coordinates")
				os.Exit(0)
			} else {
				_, err1 := strconv.Atoi(l[1])
				if err1 != nil {
					fmt.Println("ERROR: invalid data format, check x coordinates")
					os.Exit(0)
				}
				_, err2 := strconv.Atoi(l[2])
				if err2 != nil {
					fmt.Println("ERROR: invalid data format, check y coordinates")
					os.Exit(0)
				}
			}
			cell := Cell{Name: l[0]}
			for _, check := range cellsArr {
				if cell.Name == check.Name {
					fmt.Println("ERROR: invalid data format, file contains duplicated cell")
					os.Exit(0)
				}
			}
			cellsArr = append(cellsArr, cell)

			if scan[i-1] == "##start" {
				l := strings.Split(scan[i], " ")
				st = l[0]
			} else if scan[i-1] == "##end" {
				l := strings.Split(scan[i], " ")
				fn = l[0]
			}
		} else if len(foundTunells) > 0 && scan[i] == foundTunells[0] {
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
	if n == 0 || n > 10000 {
		fmt.Println("ERROR: invalid data format, invalid number of Ants")
		os.Exit(0)
	} else if start == nil || end == nil {
		fmt.Println("ERROR: no start or end")
		os.Exit(0)
	}
	return n, start, end
}

func AddTunnels(c []Cell, t string) []Cell {
	tn := strings.Split(t, "-")
	w1, w2 := tn[0], tn[1]
	if w1 == w2 {
		fmt.Println("ERROR: invalid data format, invalid tunnel")
		os.Exit(0)
	}
	for i := 0; i < len(c); i++ {
		if c[i].Name == w1 {
			for j := 0; j < len(c); j++ {
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
