package intcode

import (
	"errors"
	"fmt"
)

const (
	PositionMode  int = 0
	ImmediateMode     = 1
)

const (
	Add      int = 1
	Multiply     = 2
	Store        = 3
	Retrieve     = 4
	Halt         = 99
)

type Flags struct {
	firstMode, secondMode, thirdMode int
}

type Instruction struct {
	flags Flags
	value int
}

type IntCode struct {
	pc     int
	memory []int

	instr Instruction
}

func (this *IntCode) load(instructions []int) {
	this.memory = instructions
	this.pc = 0
}

func (this *IntCode) parseInstruction() Instruction {
	instructionSet := this.memory[this.pc]
	instruction := Instruction{Flags{0, 0, 0}, 0}
	if instructionSet >= 10000 {
		instruction.flags.firstMode = instructionSet / 10000
	}
	if instructionSet >= 1000 {
		instruction.flags.secondMode = (instructionSet - (instruction.flags.firstMode * 10000)) / 1000
	}
	if instructionSet >= 100 {
		instruction.flags.thirdMode = (instructionSet - (instruction.flags.firstMode * 10000) - (instruction.flags.secondMode * 1000)) / 100
	}
	instruction.value = (instructionSet - (instruction.flags.firstMode * 10000) - (instruction.flags.secondMode * 1000) - (instruction.flags.thirdMode * 100))
	return instruction
}

func (this *IntCode) modeCheck(value int, flag int) *int {
	switch flag {
	case PositionMode:
		return &(this.memory[value])
	case ImmediateMode:
		return &value
	default:
		return nil
	}
}

func (this *IntCode) parseArgs() (int, int, *int) {
	return *(this.modeCheck(this.memory[this.pc+1], this.flags.firstMode)), *(this.modeCheck(this.memory[this.pc+2], this.flags.secondMode)), this.modeCheck(this.memory[this.pc+3], this.flags.thirdMode)
}

func (this *IntCode) Add() {
	first, second, position := this.parseArgs()
	*(position) = first + second
	this.pc += 4
}

func (this *IntCode) Multiply() {
	first, second, position := this.parseArgs()
	*(position) = first * second
	this.pc += 4
}

func (this *IntCode) Store(input int) {
	position := this.memory[this.pc+1]
	this.memory[position] = input
	this.pc += 2
}

func (this *IntCode) Retrieve() int {
	position := this.memory[this.pc+1]
	this.pc += 2
	return this.memory[position]
}

func (this *IntCode) ReadAddress(address int) (int, error) {
	if address >= 0 && address < len(this.memory) {
		return this.memory[address], nil
	}
	return 0, errors.New("Bad address")
}

func (this *IntCode) Init(instructions []int) {
	this.load(instructions)
}

func (this *IntCode) Compute() {
	halt := false
	for halt == false {
		instruction := this.parseInstruction()
		this.flags = instruction.flags
		switch instruction.value {
		case Add:
			this.Add()
		case Multiply:
			this.Multiply()
		case Store:
			this.Store(1)
		case Retrieve:
			fmt.Println(this.Retrieve())
			halt = true
		case Halt:
			halt = true
		default:
			fmt.Println("Unsupported instruction")
		}
	}
}
