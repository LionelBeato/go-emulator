package main

type register interface {
	test() string
}

type EightBitRegister struct {
	label string
	value byte
}

type SixteenBitRegister struct {
	label string
	value int16
}

func (e EightBitRegister) test() string {
	return "test"
}

func (e SixteenBitRegister) test() string {
	return "test"
}

/*
16 8-bit (one byte) general-purpose variable
registers numbered 0 through F hexadecimal,
ie. 0 through 15 in decimal, called V0 through `VF
*/
var indexRegister SixteenBitRegister
var V0 EightBitRegister
var V1 EightBitRegister
var V2 EightBitRegister
var V3 EightBitRegister
var V4 EightBitRegister
var V5 EightBitRegister
var V6 EightBitRegister
var V7 EightBitRegister
var V8 EightBitRegister
var V9 EightBitRegister
var VA EightBitRegister
var VB EightBitRegister
var VC EightBitRegister
var VD EightBitRegister
var VE EightBitRegister
var VF EightBitRegister

/*
	VF is also used as a flag register;
	many instructions will set it to
	either 1 or 0 based on some rule,
	for example using it as a carry flag
*/

var registers = []EightBitRegister{
	EightBitRegister{label: "V0"},
	EightBitRegister{label: "V1"},
	EightBitRegister{label: "V2"},
	EightBitRegister{label: "V3"},
	EightBitRegister{label: "V4"},
	EightBitRegister{label: "V5"},
	EightBitRegister{label: "V6"},
	EightBitRegister{label: "V7"},
	EightBitRegister{label: "V8"},
	EightBitRegister{label: "V9"},
	EightBitRegister{label: "VA"},
	EightBitRegister{label: "VB"},
	EightBitRegister{label: "VC"},
	EightBitRegister{label: "VD"},
	EightBitRegister{label: "VE"},
	EightBitRegister{label: "VF"},
}
