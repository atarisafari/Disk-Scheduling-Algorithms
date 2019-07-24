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
	//"sort"
)

type Process struct {
  position int
}

// var input = os.Args[1]
var input = "fcfs01.txt"
var in, err1 = os.Open(input)
var reader = bufio.NewReader(in)

func main() {
	processInput()
}

func processInput() {

  var lowerCyl, upperCyl, curCyl int
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
			lowerCyl, _ = strconv.Atoi(words[1])
		case "upperCYL":
			upperCyl, _ = strconv.Atoi(words[1])
		case "initCYL":
			curCyl, _ = strconv.Atoi(words[1])
		case "cylreq":
      var p Process
      p.position, _ = strconv.Atoi(words[1])
			procList = append(procList, p)
		}

	}

  switch alg {
      case "fcfs":
        fcfs(procList)
  }

	fmt.Println("Lower Cylinder:", lowerCyl)
	fmt.Println("Upper Cylinder:", upperCyl)
	fmt.Println("Initial Cylinder:", curCyl)

}

func fcfs(procList []Process) {

}
