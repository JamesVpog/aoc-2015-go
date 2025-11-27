package d6

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	command     string
	start_coord [2]int
	end_coord   [2]int
}

// how many lights are turned on after reading instructions?
// max number of lights is 1000x1000
// top left is (0,0), top right is (0, 999)
// bottom left is (999, 0), bot right is (999, 999)

// we are given instructions of coordinate pairs of lights to turn on, turn off, or toggle
// e.g turn on 0,0 through 2,2 would refer to 9 lights in a 3x3 square
// e.g  toggle 0,0 through 999,0 would toggle the first line of 1000 lights,
//
//	turning off the ones that were on, and turning on the ones that were off
func P1(content string) {

	side := 1000
	lights := make([][]int, side)
	for i := range lights {
		lights[i] = make([]int, side)
	}
	// create a 2d slice of size 1000x1000 called lights
	// 0 is off, 1 is on!

	instructions := strings.Split(content, "\n")
	for _, i := range instructions {
		// parse it
		instr := parseInstruction(i)
		// fmt.Printf("%s, start: [%d, %d] end: [%d, %d]\n", instr.command, instr.start_coord[0], instr.start_coord[1], instr.end_coord[0], instr.end_coord[1])

		// exeucte it
		switch instr.command {
		case "on":
			turnLightsOn(&lights, instr.start_coord, instr.end_coord)
		case "off":
			turnLightsOff(&lights, instr.start_coord, instr.end_coord)
		case "toggle":
			toggleLights(&lights, instr.start_coord, instr.end_coord)
		}
	}
	count := 0
	for _, row := range lights {
		for _, val := range row {
			if val == 1 {
				count++
			}
		}
	}
	fmt.Printf("%d\n", count)
}

func toggleLights(lights *[][]int, start [2]int, end [2]int) {
	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			(*lights)[i][j] = 1 - (*lights)[i][j]
		}
	}
}
func turnLightsOff(lights *[][]int, start [2]int, end [2]int) {
	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			(*lights)[i][j] = 0
		}
	}
}
func turnLightsOn(lights *[][]int, start [2]int, end [2]int) {
	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			(*lights)[i][j] = 1
		}
	}
}
func parseInstruction(instr string) Instruction {
	var instruction Instruction
	var startIndex [2]int
	var endIndex [2]int
	words := strings.Fields(instr)
	// for i, w := range words {
	// 	fmt.Printf("%d: %s\n", i, w)
	// }
	// fmt.Println()
	if len(words) == 4 {
		instruction.command = "toggle"
		populateIndex(&startIndex, words[1])
		populateIndex(&endIndex, words[3])
	} else if len(words) == 5 {
		instruction.command = words[1]
		populateIndex(&startIndex, words[2])
		populateIndex(&endIndex, words[4])
	}
	instruction.start_coord = startIndex
	instruction.end_coord = endIndex
	return instruction
}

// get the "num,num" into [num, num]
func populateIndex(index *[2]int, strIndex string) {
	for i, w := range strings.Split(strIndex, ",") {
		n, err := strconv.Atoi(w)
		if err != nil {
			panic(err)
		}
		index[i] = n
	}
}

func increaseLightsBy(num int, lights *[][]int, start [2]int, end [2]int) {
	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			(*lights)[i][j] += num
		}
	}
}

func decreaseLightsBy(num int, lights *[][]int, start [2]int, end [2]int) {
	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			if (*lights)[i][j] > 0 {
				(*lights)[i][j] -= num
			}
		}
	}
}
func P2(content string) {
	// parse instrucitons
	side := 1000
	lights := make([][]int, side)
	for i := range lights {
		lights[i] = make([]int, side)
	}
	instructions := strings.Split(content, "\n")
	for _, i := range instructions {
		// parse it
		instr := parseInstruction(i)
		switch instr.command {
		case "on":
			increaseLightsBy(1, &lights, instr.start_coord, instr.end_coord)
		case "off":
			decreaseLightsBy(1, &lights, instr.start_coord, instr.end_coord)
		case "toggle":
			increaseLightsBy(2, &lights, instr.start_coord, instr.end_coord)
		}
	}

	//count the total brightness of lights
	brightness := 0
	for _, row := range lights {
		for _, val := range row {
			brightness += val
		}
	}

	fmt.Printf("%d\n", brightness)
}
