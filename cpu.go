package main

import (
	"fmt"
	"log"
)

//Representation of a machine instruction for our CPU
type MachineInstruction struct {
	opCode   byte
	operand1 byte
	operand2 byte
	operand3 byte
}

//Decodes two bytes into 4 nibbles 
func Decode(x, y byte) MachineInstruction {
	return MachineInstruction{
		x & 0xF0 >> 4,
		x & 0x0F,
		y & 0xF0 >> 4,
		y & 0x0F}
}

//String representation of a machine instruction
func (m MachineInstruction) String() string {
	return fmt.Sprintf("%X%X%X%X", m.opCode, m.operand1, m.operand2, m.operand3)
}

//The CPU
type CPU struct {
	PC        byte
	IR        MachineInstruction
	R         [16]byte
	IsRunning bool
	HasJumped bool
	Cycles    uint16
}

func NewCPU() *CPU {
	var c *CPU = new(CPU)
	c.Reset()
	return c
}

func (c *CPU) Reset() {
	c.PC = 0x00
	c.IR = Decode(0xC0, 0x00)
	c.R = *new([16]byte)
	c.IsRunning = true
	c.HasJumped = false
	c.Cycles = 0
}


func (c *CPU) LoadDirect(register byte, memoryAddr byte) {
	log.Printf("LOADD [Register: %X, Address: %X]", register, memoryAddr)
	c.R[register] = bus.Read(memoryAddr)
}

func (c *CPU) LoadValue(register byte, value byte) {
	log.Printf("LOADV [Register: %X, Value: %X]", register, value)
	c.R[register] = value
}

func (c *CPU) Store(register byte, memoryAddr byte) {
	log.Printf("STORE [Register: %X, Address: %X]", register, memoryAddr)
	var value byte = c.R[register]
	bus.Write(memoryAddr, value)
}

func (c *CPU) Move(register1 byte, register2 byte) {
	log.Printf("MOVE [From Register: %X, To Register: %X]", register1, register2)
	c.R[register2] = c.R[register1]
}

func (c *CPU) Or(leftRegister byte, rightRegister byte, resultRegister byte) {
	log.Printf("OR [Left Register: %X, Right Register:%X, Result Register: %X]", leftRegister, rightRegister, resultRegister)
	c.R[resultRegister] = c.R[leftRegister] | c.R[rightRegister]
}

func (c *CPU) And(leftRegister byte, rightRegister byte, resultRegister byte) {
	log.Printf("AND [Left Register: %X, Right Register:%X, Result Register: %X]", leftRegister, rightRegister, resultRegister)
	c.R[resultRegister] = c.R[leftRegister] & c.R[rightRegister]
}

func (c *CPU) Xor(leftRegister byte, rightRegister byte, resultRegister byte) {
	log.Printf("XOR [Left Register: %X, Right Register:%X, Result Register: %X]", leftRegister, rightRegister, resultRegister)
	c.R[resultRegister] = c.R[leftRegister] ^ c.R[rightRegister]
}

func (c *CPU) Add(leftRegister byte, rightRegister byte, resultRegister byte) {
	log.Printf("ADD [Left Register: %X, Right Register:%X, Result Register: %X]", leftRegister, rightRegister, resultRegister)
	// Limited in function as overflow is possible 
	c.R[resultRegister] = c.R[leftRegister] + c.R[rightRegister]
}

func (c *CPU) Jump(register byte, memoryAddr byte) {
	log.Printf("JMP [Register: %X, Address: %X]", register, memoryAddr)
	if c.R[0] == c.R[register] {
		c.PC = memoryAddr
		c.HasJumped = true
	}
}

func (c *CPU) Halt() {
	log.Printf("HALT machine")
	c.IsRunning = false
}

//Increment the Program Counter by an amount
func (c *CPU) IncrementPC(by byte) {
	c.PC += by
}

//Call this to "step" the CPU
func (c *CPU) Step() {
	c.Cycles++
	c.HasJumped = false

	//fetch
	p1, p2 := bus.Read(c.PC), bus.Read(c.PC+1)

	//decode 
	c.IR = Decode(p1, p2)

	//execute
	switch c.IR.opCode {
	case 0x1:
		c.LoadDirect(c.IR.operand1, (c.IR.operand2<<4)^c.IR.operand3)
	case 0x2:
		c.LoadValue(c.IR.operand1, (c.IR.operand2<<4)^c.IR.operand3)
	case 0x3:
		c.Store(c.IR.operand1, (c.IR.operand2<<4)^c.IR.operand3)
	case 0x4:
		c.Move(c.IR.operand2, c.IR.operand3)
	case 0x5:
		c.Add(c.IR.operand2, c.IR.operand3, c.IR.operand1)
	case 0x7:
		c.Or(c.IR.operand2, c.IR.operand3, c.IR.operand1)
	case 0x8:
		c.And(c.IR.operand2, c.IR.operand3, c.IR.operand1)
	case 0x9:
		c.Xor(c.IR.operand2, c.IR.operand3, c.IR.operand1)
	case 0xB:
		c.Jump(c.IR.operand1, (c.IR.operand2<<4)^c.IR.operand3)
		if c.HasJumped {
			return
		}
	case 0xC:
		c.Halt()
	default:
		log.Fatalf("Instruction %X not in instruction set", c.IR.opCode)
	}

	c.IncrementPC(2)
}

//dump fancy represenation of our CPU to a string
func (c *CPU) String() string {
	var registers string
	var spaces string

	for j := 0; j < 16; j += 4 {
		registers += "\t"
		for i, r := range c.R[j : j+4] {
			if r > 0xF {
				spaces = "  "
			} else {
				spaces = "   "
			}
			registers += (fmt.Sprintf("%X", i+j) + ": " + fmt.Sprintf("%X", r) + spaces)
		}
		registers += "\n"
	}

	return fmt.Sprintf("\nCPU\n") +
		fmt.Sprintf("--------------------------------------------------------\n") +
		fmt.Sprintf("\tPC       = %X\n", c.PC) +
		fmt.Sprintf("\tIR       = %v\n", c.IR) +
		fmt.Sprintf("\tRunning  = %v\n", c.IsRunning) +
		fmt.Sprintf("\tCycles   = %d\n", c.Cycles) +
		fmt.Sprintf("\tHasJumped= %v\n", c.HasJumped) +
		fmt.Sprintf("\n\tRegisters\n") +
		fmt.Sprintf("%v", registers) +
		fmt.Sprintf("--------------------------------------------------------\n\n")
}
