package zmachine

import (
	"strconv"
)

func (zm *ZMachine) InterpretLongVARInstruction() {
	opcode := zm.ReadByte()
	//instruction := opcode & 0x1F
	specifier1 := zm.ReadByte()
	specifier2 := zm.ReadByte()

	opValues1 := make([]uint16, 8)
	opValues2 := make([]uint16, 4)
	numOperands1 := zm.GetOperands(specifier1, opValues1)
	numOperands2 := zm.GetOperands(specifier2, opValues2)
	// merge both operand lists
	for i := uint16(0); i < numOperands2; i++ {
		opValues1[numOperands1] = opValues2[i]
		numOperands1++
	}
	DebugPrintf("opValues %v\n", opValues1)
	if opcode == 0xec {
		ZCall(zm, opValues1, numOperands1, ZCallTypeStore)
	} else { // 0xfa 250
		ZCall(zm, opValues1, numOperands1, ZCallTypeN)
	}
	//fn := ZFunctions_VAR[opcode-0xc0]
	//fn(zm, opValues1, numOperands1)
}

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
		if fn == nil {
			panic("VAR Instruction" + strconv.Itoa(int(instruction)) + " not supported")
		}
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
		//fmt.Println("instruction: ", instruction, opcode)
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

func (zm *ZMachine) InterpretExtended() {
	_ = zm.ReadByte() // 0xBE
	opcode := zm.ReadByte()
	// "In variable form, if bit 5 is 0 then the count is 2OP; if it is 1, then the count is VAR.
	// The opcode number is given in the bottom 5 bits.
	opTypesByte := zm.ReadByte()

	opValues := make([]uint16, 4)
	numOperands := zm.GetOperands(opTypesByte, opValues)
	DebugPrintf("Extended %d %v %d\n", opcode, opValues, numOperands)

	switch opcode {
	case 1:
		ZRestart(zm)
		break
	case 9:
		// Save Undo
		// -1 doesn't support, otherwise returns the save return value
		zm.SaveUndo()
		zm.StoreResult(1)
		break
	case 10:
		// Restore Undo
		zm.Restore()
		zm.StoreResult(1)
		break
	case 11:
		// Print Unicode Character
		break

	default:
		panic("Extended " + strconv.Itoa(int(opcode)) + " not supported")
	}
}

var counter = 0

func (zm *ZMachine) InterpretInstruction() {
	opcode := zm.PeekByte()

	// Form is stored in top 2 bits
	// "If the top two bits of the opcode are $$11 the form is variable; if $$10, the form is short.
	// If the opcode is 190 ($BE in hexadecimal) and the version is 5 or later, the form is "extended".
	// Otherwise, the form is "long"."
	form := (opcode >> 6) & 0x3

	DebugPrintf("%4d: ip=0x%05x opcode=%d\n", counter, zm.ip, opcode)
	counter++

	if opcode == 0xec || opcode == 0xfa {
		zm.InterpretLongVARInstruction()
	} else if opcode == 0xBE && zm.header.Version >= 5 {
		zm.InterpretExtended()
	} else if form == 0x2 {
		zm.InterpretShortInstruction()
	} else if form == 0x3 {
		zm.InterpretVARInstruction()
	} else {
		zm.InterpretLongInstruction()
	}

	if zm.outputstream == 3 && zm.Output.Len() != 0 { // output to memory
		size := uint32(zm.GetUint16(zm.outputstreamtable))
		str := zm.Output.String()
		for i := 0; i < len(str); i++ {
			zm.buf[zm.outputstreamtable+2+size] = str[i]
			size++
			zm.SetUint16(zm.outputstreamtable, uint16(size))
		}
		zm.Output.Reset()
	}

}
