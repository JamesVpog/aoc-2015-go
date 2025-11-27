package d7

import (
	"fmt"
	"strconv"
	"strings"
)

func P1(content string) {
	// each line is a signal and its circuit
	//
	// make a map of wires to instructino
	// wire is 1 or 2 letters, and instr is either 123 or some bitwise op (&, |,  !, >>, <<)
	wireToInstrMap := make(map[string]string)

	lines := strings.Split(content, "\n")
	for _, l := range lines {
		// split by the -> separator
		wireAndInstr := strings.Split(l, "->")
		if len(wireAndInstr) == 1 {
			continue
		}
		// wireAndInstr[0] will be the instr
		// wireAndInstr[1] will be the wire name
		wireToInstrMap[strings.TrimSpace(wireAndInstr[1])] = wireAndInstr[0]
	}

	cache := make(map[string]uint16)
	out := evaluate("a", wireToInstrMap, cache)
	fmt.Printf("%d\n", out)
}

// recursive function to calculate the wire for that instruction
func evaluate(wire string, wireMap map[string]string, cache map[string]uint16) uint16 {
	// also check if wire is an int, if its an int just return!
	val, err := strconv.Atoi(wire)
	if err == nil {
		uIntNum := uint16(val)
		return uIntNum
	}
	// check if the cache already has it!
	num, ok := cache[wire]
	if ok {
		return num
	}

	// grab the instruction, with the wire from the wireMap
	instr := wireMap[wire]

	// split by whiteSpace "a and b" = [a, and, b]
	parts := strings.Fields(instr)
	// fmt.Printf("%s is made of %v is of length: %d\n", wire, parts, len(parts))
	// each instruction can be a literal number, wire ref, unary op, or binary op

	// wire ref and literal number happen when len(parts) == 1
	var res uint16
	if len(parts) == 1 {
		num, err := strconv.Atoi(parts[0])
		numUint16 := uint16(num)
		if err != nil {
			// wire reference, so we call evaluate on parts[0]
			res = evaluate(parts[0], wireMap, cache)
			cache[wire] = res
			return res
		}
		// number so we cache it and return it...
		res = numUint16
		cache[wire] = res
		return res
	} else if len(parts) == 2 {
		// bitwise not ^ e.g "NOT a" = [not a]
		res = ^evaluate(parts[1], wireMap, cache)
		cache[wire] = res
		return res
	} else {
		// e.g "a AND b" = [a AND b]
		binaryOp := parts[1]
		firstOperand := evaluate(parts[0], wireMap, cache)
		secondOperand := evaluate(parts[2], wireMap, cache)
		switch binaryOp {
		case "AND":
			res = firstOperand & secondOperand
		case "OR":
			res = firstOperand | secondOperand
		case "LSHIFT":
			res = firstOperand << secondOperand
		case "RSHIFT":
			res = firstOperand >> secondOperand
		}
		cache[wire] = res
		return res
	}
}
func P2(content string) {
	// override wire b with result from a
	// 3176 -> b

	// each line is a signal and its circuit
	//
	// make a map of wires to instructino
	// wire is 1 or 2 letters, and instr is either 123 or some bitwise op (&, |,  !, >>, <<)
	wireToInstrMap := make(map[string]string)

	lines := strings.Split(content, "\n")
	for _, l := range lines {
		// split by the -> separator
		wireAndInstr := strings.Split(l, "->")
		if len(wireAndInstr) == 1 {
			continue
		}
		// wireAndInstr[0] will be the instr
		// wireAndInstr[1] will be the wire name
		wireToInstrMap[strings.TrimSpace(wireAndInstr[1])] = wireAndInstr[0]
	}

	wireToInstrMap["b"] = "3176"
	cache := make(map[string]uint16)
	out := evaluate("a", wireToInstrMap, cache)
	fmt.Printf("%d\n", out)
}
