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
	"os"
	"strconv"
	"strings"
  "math"
	//"sort"
)

type Process struct {
  position int
}

type System struct {
  lowerCyl int
  upperCyl int
  curCyl int
  traversed int
  accessed bool
}

// var input = os.Args[1]
var input = "fcfs01.txt"
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
      fmt.Println("Seek algorithm:",alg)
		case "lowerCYL":
			sys.lowerCyl, _ = strconv.Atoi(words[1])
      fmt.Printf("\tLower cylinder: %5d\n",sys.lowerCyl)
		case "upperCYL":
			sys.upperCyl, _ = strconv.Atoi(words[1])
      fmt.Printf("\tUpper cylinder: %5d\n",sys.upperCyl)
		case "initCYL":
			sys.curCyl, _ = strconv.Atoi(words[1])
      fmt.Printf("\tInit cylinder: %5d\n",sys.curCyl)
      fmt.Println("\tCylinder requests:")
		case "cylreq":
      var p Process
      p.position, _ = strconv.Atoi(words[1])
			procList = append(procList, p)
      fmt.Printf("\t\tCylinder: %5d\n",p.position)
		}

	}
  switch alg {
      case "fcfs":
        fcfs(procList,sys)
  }
}

func fcfs(procList []Process, sys System) {
  for _, p := range procList {

    if(p.position > sys.upperCyl || p.position < sys.lowerCyl) {
      fmt.Println("Out of bounds")
    } else {
      sys.traversed += int(math.Abs(float64(p.position - sys.curCyl)))
      sys.curCyl = p.position
      fmt.Printf("Servicing %5d\n",p.position)
    }
  }

  fmt.Print("FCFS traversal count = ", sys.traversed)
}

func sstf(proclist []Process, sys System) {

}
