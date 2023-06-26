package main

import (
	// "log"

	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
)

/*
https://tobiasvl.github.io/blog/write-a-chip-8-emulator/
CHIP-8 has the following components:
- Memory: CHIP-8 has direct access to up to 4 kilobytes of RAM
- Display: 64 x 32 pixels (or 128 x 64 for SUPER-CHIP) monochrome, ie. black or white
- A program counter, often called just “PC”, which points at the current instruction in memory
- One 16-bit index register called “I” which is used to point at locations in memory
- A stack for 16-bit addresses, which is used to call subroutines/functions and return from them
- An 8-bit delay timer which is decremented at a rate of 60 Hz (60 times per second) until it reaches 0
- An 8-bit sound timer which functions like the delay timer, but which also gives off a beeping sound as long as it’s not 0
- 16 8-bit (one byte) general-purpose variable registers numbered 0 through F hexadecimal, ie. 0 through 15 in decimal, called V0 through `VF
- VF is also used as a flag register; many instructions will set it to either 1 or 0 based on some rule, for example using it as a carry flag
*/

// fun fetch() {
//     opCode = (memory[pc].toInt() shl 8) or (memory[pc + 1].toInt())
//     pc =(pc + 2)
// }

func fetch() {
	opCode = (mem[currentInstruction] << 8) | (mem[currentInstruction+1])
	currentInstruction = (currentInstruction + 2)
}

func decode() {
	code := (opCode & 0xF000)
	switch code >> 12 {
	case 0x0:
		OP_00E0()
	case 0x1:
		OP_1NNN()
	case 0x3:
		OP_3XNN()
	case 0x6:
		OP_6xNN()
	case 0x7:
		OP_7xNN()
	case 0x8:
		switch opCode & 0x000F {
		case 0x1:
			OP_8XY1()
		case 0x2:
			println("inside 2 case!")
			OP_8XY2()
		case 0x3:
			println("inside 3 case!")
			OP_8XY3()
		case 0x4:
			OP_8XY4()
		case 0x5:
			OP_8XY5()
		}
	case 0xA:
		OP_ANNN()
	case 0xD:
		OP_DXYN()
	}
}

// func execute(OP_CODE byte) {

// }

func main() {
	arg := os.Args[1]

	for i := range fontSet {
		mem[e] = i
		e++
	}
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(64*10), unit.Dp(32*10)),
		)
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	loadRom(arg)
	cycle()
	app.Main()
}

func cycle() {
	for i := 0; i < 100; i++ {
		fetch()
		decode()
	}
}

func loadRom(fileName string) {
	data, _ := os.ReadFile(fileName)
	for i, value := range data {
		mem[0x200+i] = int(value)
	}
}
