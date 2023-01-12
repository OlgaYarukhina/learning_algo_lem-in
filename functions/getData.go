package functions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)



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

func ParsData (scan []string) (int, Cell, Cell, []Cell, []string) {
	var start Cell
	var end Cell
	var cellsArr []Cell
	var tunArr [] string
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
		if scan[i] == "##start"{
			start = CreateCells(scan[i+1])
		} else if scan[i] == "##end" {
			end = CreateCells(scan[i+1])
		} else if len(foundCells)>0 && scan[i] == foundCells[0] {
			eachCell := CreateCells(scan[i])
			cellsArr = append(cellsArr, eachCell)
		} else if len(foundTunells)>0 && scan[i] == foundTunells[0] {
			tunArr = append(tunArr, scan[i])
		} 
	}

	return n, start, end, cellsArr, tunArr

}
/*
func WriteToNewFile(d []byte) {
	newFile := ioutil.WriteFile(os.Args[2], d, 0600) //0777
	if newFile != nil {
		fmt.Println("Unable to create file:\n")
	}
}
*/
