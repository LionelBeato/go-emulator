package main

// opcode with OR bitwise operation
var opCode int = mem[currentInstruction] | mem[currentInstruction+1]
var VX = (opCode & 0x0F00) >> 8
var VY = (opCode & 0x00F0) >> 4
var N = opCode & 0x00F

// Function to clear the screen.
func OP_00E0() {
	for row := range screen {
		for cell := range screen[row] {
			screen[row][cell] = 0
		}
	}
}

func OP_1NNN() {
	currentInstruction = opCode & 0xFFF
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
	indexRegister.value = uint16(opCode) & 0xFFF
}

// draw pixels
func OP_DXYN() {
	VX := (opCode & 0x0F00) >> 8
	VY := (opCode & 0x00F0) >> 4
	// modulo / remainder operator
	xPos := registers[VX].value % 64
	yPos := registers[VY].value % 32
	height := opCode & 0x000F

	registers[0xF].value = 0

	// for N rows
	for row := 0; row < height-1; row++ {
		spriteByte := mem[indexRegister.value+uint16(row)]
		for col := 0; col < 8; col++ {
			spritePixel := spriteByte & (0x80 >> col)
			r := (int(yPos) + row)
			c := (int(xPos) + col)

			screenPixel := screen[r][c]
			if spritePixel > 0 {
				if screenPixel > 0 {
					registers[0xF].value = uint8(1)
				}
				// ^ XOR
				screen[(int(yPos) + row)][int(xPos)+col] = screenPixel ^ 0xFFFFFF
			}
		}
	}
}

func OP_7xNN() {
	byt := opCode & 0x00FF
	VX := (opCode & 0xFFF) >> 8
	registers[VX].value = (registers[VX].value + byte(byt))
}
