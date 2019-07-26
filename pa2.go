/*    
/*
 “I Ethan Finlay (et428907) affirm that this program is entirely my own work and
that I have neither developed my code together with any another person, nor copied any code from any
other person, nor permitted my code to be copied or otherwise used by any other person, nor have I copied,
modified, or otherwise used programs created by others. I acknowledge that any violation of the above terms
will be treated as academic dishonesty.”
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Process struct {
	position   int
	accessed   bool
	difference int
}

type System struct {
	lowerCyl  int
	upperCyl  int
	curCyl    int
	traversed int
}

// var input = os.Args[1]
var input = "sstf01.txt"
var in, err1 = os.Open(input)
var reader = bufio.NewReader(in)

func main() {
	processInput()
}

func processInput() {

	var sys System
	var alg string
	var procList []Process

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		s := string(line)

		words := strings.Fields(s)
		instruction := words[0]

		switch instruction {
		case "use":
			alg = words[1]
		case "lowerCYL":
			sys.lowerCyl, _ = strconv.Atoi(words[1])
		case "upperCYL":
			sys.upperCyl, _ = strconv.Atoi(words[1])
      		if(sys.upperCyl < sys.lowerCyl) {
        		fmt.Printf("ABORT(13):upper (%d) < lower (%d)\n", sys.upperCyl, sys.lowerCyl)
        		return
      		}
		case "initCYL":
			sys.curCyl, _ = strconv.Atoi(words[1])
      		if(sys.curCyl > sys.upperCyl) {
        		fmt.Printf("ABORT(11):initial (%d) > upper (%d)\n", sys.curCyl, sys.upperCyl)
        		return
      		} else if(sys.curCyl < sys.lowerCyl) {
        		fmt.Printf("ABORT(12):initial (%d) < lower (%d)\n", sys.curCyl, sys.lowerCyl)
        		return
      		}
		case "cylreq":
			var p Process
			p.position, _ = strconv.Atoi(words[1])
      		if(p.position < sys.lowerCyl || p.position > sys.upperCyl) {
        		fmt.Printf("ERROR(15):Request out of bounds:  req (%d) > upper (%d) or < lower (%d)\n", p.position, sys.upperCyl, sys.lowerCyl)
      		} else { procList = append(procList, p) }
			}

	}

  fmt.Println("Seek algorithm:", strings.ToUpper(alg))
  fmt.Printf("\tLower cylinder: %5d\n", sys.lowerCyl)
  fmt.Printf("\tUpper cylinder: %5d\n", sys.upperCyl)
  fmt.Printf("\tInit cylinder:  %5d\n", sys.curCyl)
	fmt.Println("\tCylinder requests:")

  for _, p := range procList {
    fmt.Printf("\t\tCylinder: %5d\n", p.position)
  }

	switch alg {
		case "fcfs":
			fcfs(procList, sys)
		case "sstf":
			sstf(procList, sys)
	}
}

func fcfs(procList []Process, sys System) {
	for _, p := range procList {

		if p.position > sys.upperCyl || p.position < sys.lowerCyl {
			fmt.Println("Out of bounds")
		} else {
			sys.traversed += int(math.Abs(float64(p.position - sys.curCyl)))
			sys.curCyl = p.position
			fmt.Printf("Servicing %5d\n", p.position)
		}
	}

	fmt.Print("FCFS traversal count = ", sys.traversed)
}

func sstf(procList []Process, sys System) {

	var headIndex int

	//Sort by location
	sort.Slice(procList, func(i, j int) bool {
		return procList[i].position < procList[j].position
	})

	//Find first index with location value greater than sys.curCyl
	for i := 0; i < len(procList); i++ {
		if(procList[i].position > sys.curCyl) {
			headIndex = i
			break
		}
	}

	fmt.Println("Head is at", procList[headIndex].position)

	//Check whether that value or the one at the index to the left is closer
	//Go there
	//Mark it as visited and add distance
	//Set index to that location

	//Loop until all processes are visited
		//Check left/right of current index for closest unaccessed (return -1 if you reach an endpoint)
		//Go to closer location
		//Mark it as visited and add distance
		//Set index to that location 
		//Repeat
   

    //Maybe keep a counter for # procs visited, and if it's equal to proc len
    // then we break
}

func calcDiff(procList []Process, sys System) []Process {
	for _, p := range procList {
		p.difference = int(math.Abs(float64(p.position - sys.curCyl)))
	}

	return procList
}
