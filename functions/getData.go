package functions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func ParsData (scan []string) (int, string, string, []string, []string) {
	var start string
	var end string
	var cellsArr []string 
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
			l := strings.Split(scan[i+1], " ")
			start = l[0]
		} else if scan[i] == "##end" {
			l := strings.Split(scan[i+1], " ")
			end = l[0]
		} else if len(foundCells)>0 && scan[i] == foundCells[0] {
			l := strings.Split(scan[i], " ")
			cellsArr = append(cellsArr, l[0])
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
