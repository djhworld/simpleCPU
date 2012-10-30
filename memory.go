package main

import (
	"fmt"
)

type Memory [256]byte

func (m *Memory) String() string {
	var cells string
	var spaces string

	cells += "Memory\n----------------------------------------------------------------\n"
	for j := 0; j < 256; j += 8 {
		for i, r := range m[j : j+8] {
			if r > 0xF {
				spaces = "  "
			} else {
				spaces = "   "
			}

			if i+j <= 0xF {
				cells += (fmt.Sprintf("0%X", i+j) + ": " + fmt.Sprintf("%X", r) + spaces)
			} else {
				cells += (fmt.Sprintf("%X", i+j) + ": " + fmt.Sprintf("%X", r) + spaces)
			}
		}
		cells += "\n"
	}
	cells += "----------------------------------------------------------------\n"

	return cells
}

type Bus struct {
	memory *Memory
}

func (b Bus) Read(addr byte) byte {
	return b.memory[addr]
}

func (b Bus) Write(addr byte, value byte) {
	b.memory[addr] = value
}
