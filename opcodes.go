package main

import "fmt"

// opcode with OR bitwise operation
var opCode = mem[currentInstruction] | mem[currentInstruction+1]
var VX = (opCode & 0x0F00) >> 8
var VY = (opCode & 0x00F0) >> 4
var N = opCode & 0x00F

// Function to clear the screen.
func OP_00E0() {
	for row := range screen {
		// println("in row")
		// println(row)
		for cell := range screen[row] {
			screen[row][cell] = 0
			// println("in cell")
		}
	}
}

func OP_1NNN() {
	currentInstruction = opCode & 0xFFF
	fmt.Printf("counter updated to %x\n", currentInstruction)
}

func OP_6xNN() {
	byt := opCode & 0x00FF
	VX = (opCode & 0xFFF) >> 8
	registers[VX].value = uint8(byt)
}

func OP_2NNN() {
	stack = append(stack, currentInstruction)
	currentInstruction = opCode & 0xFFF
}

func OP_00EE() {
	n := len(stack) - 1
	currentInstruction = 0
	stack = stack[:n]
}

func OP_ANNN() {
	indexRegister.value = int16(opCode) & 0xFFF
}

// draw pixels
func OP_DXYN() {
	VX := (opCode & 0x0F00) >> 8
	VY := (opCode & 0x0F00) >> 8
	// modulo / remainder operator
	xPos := registers[VX].value % 64
	yPos := registers[VY].value % 32
	height := opCode & 0x000F

	registers[0xF].value = 0

	// for N rows
	for row := 0; row < int(height); row++ {
		spriteByte := mem[indexRegister.value+int16(row)]
		for col := 0; col < 8; col++ {
			spritePixel := spriteByte & 0x80 >> col
			screenPixel := screen[(int(yPos) + row)][int(xPos)+col]
			if spritePixel > 0 {
				if screenPixel > 0 {
					registers[0xF].value = 1
				}
				// ^ is golangs's XOR for ints
				screen[(int(yPos) + row)][int(xPos)+col] = screenPixel ^ 0xFFFFFF
				println(screen[(int(yPos) + row)][int(xPos)+col])
			}
		}
	}
}

func OP_7xNN() {
	byt := opCode & 0x00FF
	VX := (opCode & 0xFFF) >> 8
	registers[VX].value = (registers[VX].value + byte(byt))
	printable := registers[VX]
	fmt.Printf("value at Register %v is now %v", VX, printable.value)
}
