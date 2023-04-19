package main

// type register interface {
// 	test() string
// }

type EightBitRegister struct {
	label string
	value byte
}

type SixteenBitRegister struct {
	label string
	value uint16
}

// func (e EightBitRegister) test() string {
// 	return "test"
// }

// func (e SixteenBitRegister) test() string {
// 	return "test"
// }

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
	{label: "V0"},
	{label: "V1"},
	{label: "V2"},
	{label: "V3"},
	{label: "V4"},
	{label: "V5"},
	{label: "V6"},
	{label: "V7"},
	{label: "V8"},
	{label: "V9"},
	{label: "VA"},
	{label: "VB"},
	{label: "VC"},
	{label: "VD"},
	{label: "VE"},
	{label: "VF"},
}
