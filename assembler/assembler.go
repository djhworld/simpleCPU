package main

/*
	Really basic, fragile assembler for our CPU
*/
import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var input *string = flag.String("i", "", "The input file")
var output *string = flag.String("o", "", "The output file")

var InstructionSet map[string]byte = map[string]byte{
	"LDD": 0x1,
	"LDV": 0x2,
	"STO": 0x3,
	"MOV": 0x4,
	"ADD": 0x5,
	"ORR": 0x7,
	"AND": 0x8,
	"XOR": 0x9,
	"JMP": 0xB,
	"HLT": 0xC,
}

var StrToHex map[string]byte = map[string]byte{
	"0": 0x0,
	"1": 0x1,
	"2": 0x2,
	"3": 0x3,
	"4": 0x4,
	"5": 0x5,
	"6": 0x6,
	"7": 0x7,
	"8": 0x8,
	"9": 0x9,
	"A": 0xA,
	"B": 0xB,
	"C": 0xC,
	"D": 0xD,
	"E": 0xE,
	"F": 0xF,
}

//unforgiving
func Emit(line string) (byte, byte, error) {
	var result [4]byte
	parts := strings.Split(line, " ")

	if len(parts) != 2 {
		return 0, 0, errors.New("Instruction must come in the format <OPERATION> <OPERAND>")
	}

	firstHalf := parts[0]

	if v, ok := InstructionSet[firstHalf]; !ok {
		return 0, 0, errors.New("Unknown instruction: " + firstHalf)
	} else {
		result[0] = v
	}

	secondHalf := strings.Split(parts[1], "")

	if len(secondHalf) != 3 {
		return 0, 0, errors.New("Operand must be 3 digits in length")
	}

	for i, v := range secondHalf {
		v = strings.ToUpper(v)
		if digit, ok := StrToHex[v]; !ok {
			return 0, 0, errors.New("Operand must contain hex digits only")
		} else {
			result[i+1] = digit
		}
	}

	var b1 byte = 0
	var b2 byte = 0

	if result[0] == 0x0 {
		b1 = 0x00  ^ result[1]
	} else {
		b1 = (result[0] << 4) ^ result[1]
	}


	if result[1] == 0x0 {
		b2 = 0x00 ^ result[3]
	} else {
		b2 = (result[2] << 4) ^ result[3]
	}
	return b1, b2, nil
}

func main() {
	var lines int = 0
	flag.Parse()

	if *input == "" {
		fmt.Println("Please provide an input file")
		return
	}

	if *output == "" {
		fmt.Println("Please provide output filename")
		return
	}

	file, err := os.Open(*input)

	if err != nil {
		fmt.Println("Encountered error attempting to open the file", *input, "\n\t", err)
		return
	}

	defer file.Close()

	buf := bufio.NewReader(file)

	var Program []byte = make([]byte, 0)
	Program = append(Program, 0x00)

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err != io.EOF || len(line) > 0 {
				fmt.Println(err)
				return
			}
			break
		}

		lines += 1

		line = strings.TrimRight(line, "\n")

		if line != "" {
			b1, b2, err := Emit(line)
			if err != nil {
				fmt.Printf("Parse error on line %d (%s).\n\t%v\n", lines, line, err)
				return
			}
			Program = append(Program, b1, b2)
		}
	}

	for i, b := range Program {
		fmt.Printf("%X: %X\n", uint8(i), b)
	}

	writeErr := WriteBinaryFile(*output, Program)
	if writeErr != nil {
		fmt.Println("Encountered error attempting to write output file", *output, "\n\t", writeErr)
	}
}

func WriteBinaryFile(filename string, bytes []byte) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	bufr := bufio.NewWriter(file)
	bufr.Write(bytes)
	bufr.Flush()
	return nil
}
