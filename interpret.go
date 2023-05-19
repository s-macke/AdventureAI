package main

func (zm *ZMachine) InterpretVARInstruction() {

	opcode := zm.ReadByte()
	// "In variable form, if bit 5 is 0 then the count is 2OP; if it is 1, then the count is VAR.
	// The opcode number is given in the bottom 5 bits.
	instruction := opcode & 0x1F
	twoOp := ((opcode >> 5) & 0x1) == 0

	// "In variable or extended forms, a byte of 4 operand types is given next.
	// This contains 4 2-bit fields: bits 6 and 7 are the first field, bits 0 and 1 the fourth."
	// "A value of 0 means a small constant and 1 means a variable."
	opTypesByte := zm.ReadByte()

	opValues := make([]uint16, 4)
	numOperands := zm.GetOperands(opTypesByte, opValues)
	DebugPrintf("opValues %v\n", opValues)

	if twoOp {
		fn := ZFunctions_2OP[instruction]
		fn(zm, opValues, numOperands)
	} else {
		fn := ZFunctions_VAR[instruction]
		fn(zm, opValues, numOperands)
	}
}

func (zm *ZMachine) InterpretShortInstruction() {
	// "In short form, bits 4 and 5 of the opcode byte give an operand type.
	// If this is $11 then the operand count is 0OP; otherwise, 1OP. In either case the opcode number is given in the bottom 4 bits."

	opcode := zm.ReadByte()
	opType := (opcode >> 4) & 0x3
	instruction := opcode & 0x0F

	if opType != OPERAND_OMITTED {
		opValue := zm.GetOperand(opType)
		DebugPrintf("opValues [%d]\n", opValue)
		fn := ZFunctions_1OP[instruction]
		fn(zm, opValue)
	} else {
		fn := ZFunctions_0P[instruction]
		fn(zm)
	}
}

func (zm *ZMachine) InterpretLongInstruction() {

	opcode := zm.ReadByte()

	// In long form the operand count is always 2OP. The opcode number is given in the bottom 5 bits.
	instruction := opcode & 0x1F

	// Operand types:
	// In long form, bit 6 of the opcode gives the type of the first operand, bit 5 of the second.
	// A value of 0 means a small constant and 1 means a variable.
	operandType0 := ((opcode & 0x40) >> 6) + 1
	operandType1 := ((opcode & 0x20) >> 5) + 1

	opValues := make([]uint16, 2)
	opValues[0] = zm.GetOperand(operandType0)
	opValues[1] = zm.GetOperand(operandType1)
	DebugPrintf("opValues %v\n", opValues)
	fn := ZFunctions_2OP[instruction]
	fn(zm, opValues, 2)
}

var counter = 0

func (zm *ZMachine) InterpretInstruction() {
	opcode := zm.PeekByte()

	// Form is stored in top 2 bits
	// "If the top two bits of the opcode are $$11 the form is variable; if $$10, the form is short.
	// If the opcode is 190 ($BE in hexadecimal) and the version is 5 or later, the form is "extended".
	// Otherwise, the form is "long"."
	form := (opcode >> 6) & 0x3

	DebugPrintf("%4d: ip=0x%05x opcode=%d %d %d\n", counter, zm.ip, opcode, zm.buf[50], zm.buf[51])
	counter++

	if opcode == 0xCE || opcode == 0xfa {
		panic("Special VAR Not supported")
	}

	if opcode == 0xBE && zm.header.version >= 5 {
		panic("Extended not supported")
	}
	if form == 0x2 {
		zm.InterpretShortInstruction()
	} else if form == 0x3 {
		zm.InterpretVARInstruction()
	} else {
		zm.InterpretLongInstruction()
	}
}
