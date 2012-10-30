package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var cpu *CPU
var memory *Memory
var bus Bus
var pause *bool = flag.Bool("p", false, "machine should pause CPU after each cycle")
var dump *bool = flag.Bool("d", false, "machine should print CPU state after each cycle")

func Init() {
	memory = new(Memory)
	bus = Bus{memory}
	cpu = NewCPU()
}

func LoadROM(filename string) (byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}
	defer file.Close()

	bytes := make([]byte, 256)
	bufr := bufio.NewReader(file)
	n, err := bufr.Read(bytes)
	start := bytes[0]
	offset := start

	for _, b := range bytes[1:n] {
		bus.Write(offset, b)
		offset++
	}
	return start, nil
}

func main() {
	flag.Parse()
	var rom string = flag.Arg(0)

	Init()
	start, err := LoadROM(rom)
	if err != nil {
		fmt.Println(err)
		return
	}

	cpu.PC = start

	if *dump {
		fmt.Print(cpu)
	}

	for cpu.IsRunning {

		cpu.Step()
		if *dump {
			fmt.Print(cpu)
		}

		if *pause {
			var h int
			fmt.Scanf("%d\n", &h)
		}
	}

	if *dump {
		fmt.Print(bus.memory)
	}
}
