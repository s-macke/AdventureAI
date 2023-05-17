package main

// https://gitlab.com/DavidGriffith/frotz/
// https://www.inform-fiction.org/zmachine/standards/

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

const (
	OPERAND_LARGE    = 0x0
	OPERAND_SMALL    = 0x1
	OPERAND_VARIABLE = 0x2
	OPERAND_OMITTED  = 0x3

	FORM_SHORT    = 0x0
	FORM_LONG     = 0x1
	FORM_VARIABLE = 0x2

	MAX_STACK  = 1024
	MAX_OBJECT = 256

	DICT_NOT_FOUND = 0
)

type ZCallType int

const (
	ZCallTypeStore ZCallType = 0 // call a subroutine and store its result
	ZCallTypeN               = 1 // call a subroutine and discard its result.
	//ZCallTypeDirect           = 2
)

type ZMachine struct {
	ip         uint32
	header     ZHeader
	buf        []uint8
	stack      *ZStack
	localFrame uint16
	done       bool
	output     strings.Builder
	input      func() string
}

type ZFunction func(*ZMachine, []uint16, uint16)
type ZFunction1Op func(*ZMachine, uint16)
type ZFunction0Op func(*ZMachine)

var alphabets = []string{
	"abcdefghijklmnopqrstuvwxyz",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	" \n0123456789.,!?_#'\"/\\-:()"}

// Doesn't modify IP
func (zm *ZMachine) PeekByte() uint8 {
	return zm.buf[zm.ip]
}

// Reads & moves to the next one (advances IP)
func (zm *ZMachine) ReadByte() uint8 {
	zm.ip++
	return zm.buf[zm.ip-1]
}

// Reads 2 bytes and advances IP
func (zm *ZMachine) ReadUint16() uint16 {
	retVal := zm.GetUint16(zm.ip)
	zm.ip += 2
	return retVal
}

func (zm *ZMachine) GetUint16(offset uint32) uint16 {
	return (uint16(zm.buf[offset]) << 8) | (uint16)(zm.buf[offset+1])
}

func (zm *ZMachine) SetUint16(offset uint32, v uint16) {
	zm.buf[offset] = uint8(v >> 8)
	zm.buf[offset+1] = uint8(v & 0xFF)
}

// " Given a packed address P, the formula to obtain the corresponding byte address B is:
//
//	2P           Versions 1, 2 and 3"
func (zm *ZMachine) PackedAddress(a uint32) uint32 {
	if zm.header.version <= 3 {
		return a * 2
	}
	return a * 4
}

func (zm *ZMachine) ReadGlobal(x uint8) uint16 {
	if x < 0x10 {
		panic("Invalid global variable")
	}

	addr := zm.PackedAddress(uint32(x) - 0x10)
	ret := zm.GetUint16(zm.header.globalVarAddress + addr)

	return ret
}

func (zm *ZMachine) SetGlobal(x uint16, v uint16) {
	if x < 0x10 {
		panic("Invalid global variable")
	}

	addr := zm.PackedAddress(uint32(x) - 0x10)
	zm.SetUint16(zm.header.globalVarAddress+addr, v)
}

func ZPutProp(zm *ZMachine, args []uint16, numArgs uint16) {
	zm.SetObjectProperty(args[0], args[1], args[2])
}

func ZRead(zm *ZMachine, args []uint16, numArgs uint16) {
	textAddress := args[0]
	maxChars := uint16(zm.buf[textAddress])
	if maxChars == 0 {
		panic("Invalid max chars")
	}
	maxChars--

	//reader := bufio.NewReader(os.Stdin)
	//input, _ := reader.ReadString('\n')
	input := zm.input()

	input = strings.ToLower(input)
	input = strings.Trim(input, "\r\n")

	copy(zm.buf[textAddress+1:textAddress+maxChars], input)
	zm.buf[textAddress+uint16(len(input))+1] = 0

	var words []string
	var wordStarts []uint16
	var stringBuffer bytes.Buffer
	prevWordStart := uint16(1)
	for i := uint16(1); zm.buf[textAddress+i] != 0; i++ {
		ch := zm.buf[textAddress+i]
		if ch == ' ' {
			if prevWordStart < 0xFFFF {
				words = append(words, stringBuffer.String())
				wordStarts = append(wordStarts, prevWordStart)
				stringBuffer.Truncate(0)
			}
			prevWordStart = 0xFFFF
		} else {
			stringBuffer.WriteByte(ch)
			if prevWordStart == 0xFFFF {
				prevWordStart = i
			}
		}
	}
	// Last word
	if prevWordStart < 0xFFFF {
		words = append(words, stringBuffer.String())
		wordStarts = append(wordStarts, prevWordStart)
	}

	// TODO: include other separators, not only spaces

	parseAddress := uint32(args[1])
	maxTokens := zm.buf[parseAddress]
	//DebugPrintf("Max tokens: %d\n", maxTokens)
	parseAddress++
	numTokens := uint8(len(words))
	if numTokens > maxTokens {
		numTokens = maxTokens
	}
	zm.buf[parseAddress] = numTokens
	parseAddress++

	// "Each block consists of the byte address of the word in the dictionary, if it is in the dictionary, or 0 if it isn't;
	// followed by a byte giving the number of letters in the word; and finally a byte giving the position in the text-buffer
	// of the first letter of the word.
	for i, w := range words {

		if uint8(i) >= maxTokens {
			break
		}

		DebugPrintf("w = %s, %d\n", w, wordStarts[i])
		dictionaryAddress := zm.FindInDictionary(w)
		DebugPrintf("Dictionary address: 0x%X\n", dictionaryAddress)

		zm.SetUint16(parseAddress, dictionaryAddress)
		zm.buf[parseAddress+2] = uint8(len(w))
		zm.buf[parseAddress+3] = uint8(wordStarts[i])
		parseAddress += 4
	}
}

func ZPrintChar(zm *ZMachine, args []uint16, numArgs uint16) {
	ch := args[0]
	PrintZChar(&zm.output, ch)
}

func ZPrintNum(zm *ZMachine, args []uint16, numArgs uint16) {
	_, _ = fmt.Fprintf(&zm.output, "%d", int16(args[0]))
}

// If range is positive, returns a uniformly random number between 1 and range.
// If range is negative, the random number generator is seeded to that value and the return value is 0.
// Most interpreters consider giving 0 as range illegal (because they attempt a division with remainder by the range),
// / but correct behaviour is to reseed the generator in as random a way as the interpreter can (e.g. by using the time
// in milliseconds).
func ZRandom(zm *ZMachine, args []uint16, numArgs uint16) {
	randRange := int16(args[0])

	if randRange > 0 {
		r := rand.Int31n(int32(randRange)) // [0, n]
		zm.StoreResult(uint16(r + 1))
	} else if randRange < 0 {
		rand.Seed(int64(randRange * -1))
		zm.StoreResult(0)
	} else {
		rand.Seed(time.Now().Unix())
		zm.StoreResult(0)
	}
}

func ZPush(zm *ZMachine, args []uint16, numArgs uint16) {
	zm.stack.Push(args[0])
}

func ZPull(zm *ZMachine, args []uint16, numArgs uint16) {
	r := zm.stack.Pop()
	DebugPrintf("Popped %d 0x%X %d %d\n", r, zm.ip, numArgs, args[0])
	zm.StoreAtLocation(args[0], r)
}

func ZNOP_VAR(zm *ZMachine, args []uint16, numArgs uint16) {
	fmt.Printf("IP=0x%X\n", zm.ip)
	panic("NOP VAR")
}

func ZNOP(zm *ZMachine, args []uint16) {
	fmt.Printf("IP=0x%X\n", zm.ip)
	panic("NOP 2OP")
}

func GenericBranch(zm *ZMachine, conditionSatisfied bool) {
	branchInfo := zm.ReadByte()

	// "If bit 7 of the first byte is 0, a branch occurs when the condition was false; if 1, then branch is on true"
	branchOnFalse := (branchInfo >> 7) == 0

	var jumpAddress int32
	var branchOffset int32
	// 0 = return false, 1 = return true, 2 = standard jump
	returnFromCurrent := int32(2)
	// "If bit 6 is set, then the branch occupies 1 byte only, and the "offset" is in the range 0 to 63, given in the bottom 6 bits"
	if (branchInfo & (1 << 6)) != 0 {
		branchOffset = int32(branchInfo & 0x3F)

		// "An offset of 0 means "return false from the current routine", and 1 means "return true from the current routine".
		if branchOffset <= 1 {
			returnFromCurrent = branchOffset
		}
	} else {
		// If bit 6 is clear, then the offset is a signed 14-bit number given in bits 0 to 5 of the first
		// byte followed by all 8 of the second.
		secondPart := zm.ReadByte()
		firstPart := uint16(branchInfo & 0x3F)
		// Propagate sign bit (2 complement)
		if (firstPart & 0x20) != 0 {
			firstPart |= (1 << 6) | (1 << 7)
		}

		branchOffset16 := int16(firstPart<<8) | int16(secondPart)
		branchOffset = int32(branchOffset16)

		DebugPrintf("Offset: 0x%X [%d]\n", branchOffset, branchOffset)
	}
	ip := int32(zm.ip)

	// "Otherwise, a branch moves execution to the instruction at address
	// Address after branch data + Offset - 2."
	jumpAddress = ip + int32(branchOffset) - 2

	DebugPrintf("Jump address = 0x%X\n", jumpAddress)

	doJump := (conditionSatisfied != branchOnFalse)

	DebugPrintf("Do jump: %t\n", doJump)

	if doJump {
		if returnFromCurrent != 2 {
			ZRet(zm, uint16(returnFromCurrent))
		} else {
			zm.ip = uint32(jumpAddress)
		}
	}
}

func ZJumpEqual(zm *ZMachine, args []uint16, numArgs uint16) {

	conditionSatisfied := (args[0] == args[1] ||
		(numArgs > 2 && args[0] == args[2]) || (numArgs > 3 && args[0] == args[3]))
	GenericBranch(zm, conditionSatisfied)
}

func ZJumpLess(zm *ZMachine, args []uint16, numArgs uint16) {
	conditionSatisfied := int16(args[0]) < int16(args[1])
	GenericBranch(zm, conditionSatisfied)
}

func ZJumpGreater(zm *ZMachine, args []uint16, numArgs uint16) {
	conditionSatisfied := int16(args[0]) > int16(args[1])
	GenericBranch(zm, conditionSatisfied)
}

func ZAdd(zm *ZMachine, args []uint16, numArgs uint16) {
	r := int16(args[0]) + int16(args[1])
	zm.StoreResult(uint16(r))
}

func ZSub(zm *ZMachine, args []uint16, numArgs uint16) {
	r := int16(args[0]) - int16(args[1])
	zm.StoreResult(uint16(r))
}

func ZMul(zm *ZMachine, args []uint16, numArgs uint16) {
	r := int16(args[0]) * int16(args[1])
	zm.StoreResult(uint16(r))
}

func ZDiv(zm *ZMachine, args []uint16, numArgs uint16) {
	if args[1] == 0 {
		panic("Division by zero")
	}

	r := int16(args[0]) / int16(args[1])
	zm.StoreResult(uint16(r))
}

func ZMod(zm *ZMachine, args []uint16, numArgs uint16) {
	if args[1] == 0 {
		panic("Division by zero (mod)")
	}

	r := int16(args[0]) % int16(args[1])
	zm.StoreResult(uint16(r))
}

func ZStore(zm *ZMachine, args []uint16, numArgs uint16) {
	DebugPrintf("%d - 0x%X\n", args[0], args[1])
	zm.StoreAtLocation(args[0], args[1])
}

func ZTestAttr(zm *ZMachine, args []uint16, numArgs uint16) {
	GenericBranch(zm, zm.TestObjectAttr(args[0], args[1]))
}

func ZOr(zm *ZMachine, args []uint16, numArgs uint16) {
	zm.StoreResult(args[0] | args[1])
}

func ZAnd(zm *ZMachine, args []uint16, numArgs uint16) {
	zm.StoreResult(args[0] & args[1])
}

func ZSetAttr(zm *ZMachine, args []uint16, numArgs uint16) {
	zm.SetObjectAttr(args[0], args[1])
}

func ZClearAttr(zm *ZMachine, args []uint16, numArgs uint16) {
	zm.ClearObjectAttr(args[0], args[1])
}

func ZGetProp(zm *ZMachine, args []uint16, numArgs uint16) {
	prop := zm.GetObjectProperty(args[0], args[1])
	zm.StoreResult(prop)
}

func ZGetPropAddr(zm *ZMachine, args []uint16, numArgs uint16) {
	addr := zm.GetObjectPropertyAddress(args[0], args[1])
	zm.StoreResult(addr)
}

func ZGetNextProp(zm *ZMachine, args []uint16, numArgs uint16) {
	addr := zm.GetNextObjectProperty(args[0], args[1])
	zm.StoreResult(addr)
}

// Returns new value.
func (zm *ZMachine) AddToVar(varType uint16, value int16) uint16 {
	retValue := uint16(0)
	if varType == 0 {
		zm.stack.stack[zm.stack.top] += uint16(value)
		retValue = zm.stack.GetTopItem()
	} else if varType < 0x10 {
		retValue = zm.stack.GetLocalVar((int)(varType - 1))
		retValue += uint16(value)
		zm.stack.SetLocalVar(int(varType-1), retValue)
	} else {
		retValue = zm.ReadGlobal(uint8(varType))
		retValue += uint16(value)
		zm.SetGlobal(varType, retValue)
	}
	return retValue
}

// dec_chk (variable) value ?(label)
// Decrement variable, and branch if it is now less than the given value.
func ZDecChk(zm *ZMachine, args []uint16, numArgs uint16) {
	newValue := zm.AddToVar(args[0], -1)
	GenericBranch(zm, int16(newValue) < int16(args[1]))
}

// inc_chk (variable) value ?(label)
// Increment variable, and branch if now greater than value.
func ZIncChk(zm *ZMachine, args []uint16, numArgs uint16) {
	newValue := zm.AddToVar(args[0], 1)
	GenericBranch(zm, int16(newValue) > int16(args[1]))
}

// test bitmap flags ?(label)
// Jump if all of the flags in bitmap are set (i.e. if bitmap & flags == flags).
func ZTest(zm *ZMachine, args []uint16, numArgs uint16) {
	bitmap := args[0]
	flags := args[1]
	GenericBranch(zm, (bitmap&flags) == flags)
}

//	jin obj1 obj2 ?(label)
//
// Jump if object a is a direct child of b, i.e., if parent of a is b.
func ZJin(zm *ZMachine, args []uint16, numArgs uint16) {
	GenericBranch(zm, zm.IsDirectParent(args[0], args[1]))
}

func ZInsertObj(zm *ZMachine, args []uint16, numArgs uint16) {
	zm.ReparentObject(args[0], args[1])
}

func ZJumpZero(zm *ZMachine, arg uint16) {
	GenericBranch(zm, arg == 0)
}

// get_sibling object -> (result) ?(label)
// Get next object in tree, branching if this exists, i.e. is not 0.
func ZGetSibling(zm *ZMachine, arg uint16) {
	sibling := zm.GetSibling(arg)
	zm.StoreResult(sibling)
	GenericBranch(zm, sibling != NULL_OBJECT_INDEX)
}

// get_child object -> (result) ?(label)
// Get first object contained in given object, branching if this exists, i.e. is not nothing (i.e., is not 0).
func ZGetChild(zm *ZMachine, arg uint16) {
	childIndex := zm.GetFirstChild(arg)
	zm.StoreResult(childIndex)
	GenericBranch(zm, childIndex != NULL_OBJECT_INDEX)
}

func ZGetParent(zm *ZMachine, arg uint16) {
	parent := zm.GetParentObject(arg)
	zm.StoreResult(parent)
}

func ZGetPropLen(zm *ZMachine, arg uint16) {
	if arg == 0 {
		zm.StoreResult(0)
	} else {
		// Arg = direct address of the property block
		// To get size, we need to go 1 byte back
		propSize := zm.buf[arg-1]
		numBytes := (propSize >> 5) + 1
		zm.StoreResult(uint16(numBytes))
	}
}

// print_paddr packed-address-of-string
func ZPrintPAddr(zm *ZMachine, arg uint16) {
	zm.DecodeZString(zm.PackedAddress(uint32(arg)))
}

func ZLoad(zm *ZMachine, arg uint16) {
	zm.StoreResult(arg)
}

func ZInc(zm *ZMachine, arg uint16) {
	zm.AddToVar(arg, 1)
}

func ZDec(zm *ZMachine, arg uint16) {
	zm.AddToVar(arg, -1)
}

func ZPrintAddr(zm *ZMachine, arg uint16) {
	zm.DecodeZString(uint32(arg))
}

func ZRemoveObj(zm *ZMachine, arg uint16) {
	zm.UnlinkObject(arg)
}

func ZPrintObj(zm *ZMachine, arg uint16) {
	zm.PrintObjectName(arg)
}

// Unconditional jump
func ZJump(zm *ZMachine, arg uint16) {
	jumpOffset := int16(arg)
	jumpAddress := int32(zm.ip) + int32(jumpOffset) - 2
	DebugPrintf("Jump address: 0x%X\n", jumpAddress)
	zm.ip = uint32(jumpAddress)
}

func ZNOP1(zm *ZMachine, arg uint16) {
	fmt.Printf("IP=0x%X\n", zm.ip)
	panic("NOP1")
}

func ZReturnTrue(zm *ZMachine) {
	ZRet(zm, uint16(1))
}

func ZReturnFalse(zm *ZMachine) {
	ZRet(zm, uint16(0))
}

func ZPrint(zm *ZMachine) {
	zm.ip = zm.DecodeZString(zm.ip)
}

func ZPrintRet(zm *ZMachine) {
	zm.ip = zm.DecodeZString(zm.ip)

	_, _ = fmt.Fprintf(&zm.output, "\n")
	ZRet(zm, 1)
}

func ZRetPopped(zm *ZMachine) {
	retValue := zm.stack.Pop()
	ZRet(zm, retValue)
}

func ZPop(zm *ZMachine) {
	zm.stack.Pop()
}

func ZQuit(zm *ZMachine) {
	zm.done = true
}

func ZNewLine(zm *ZMachine) {
	_, _ = fmt.Fprintf(&zm.output, "\n")
}

func ZNOP0(zm *ZMachine) {
	panic("NOP0")
}

var ZFunctions_VAR = []ZFunction{
	func(zm *ZMachine, args []uint16, numargs uint16) {
		ZCall(zm, args, numargs, ZCallTypeStore)
	},
	ZStoreW,
	ZStoreB,
	ZPutProp,
	ZRead,
	ZPrintChar,
	ZPrintNum,
	ZRandom,
	ZPush,
	ZPull,
	nil,
	nil,
	func(zm *ZMachine, args []uint16, numargs uint16) {
		ZCall(zm, args, numargs, ZCallTypeStore)
	},
	nil,
	nil,
	nil,
	nil, // get cursor array
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil, // not value -> (result)
	func(zm *ZMachine, args []uint16, numargs uint16) {
		ZCall(zm, args, numargs, ZCallTypeN)
	},
	nil,
}

var ZFunctions_2OP = []ZFunction{
	ZNOP_VAR,
	ZJumpEqual,
	ZJumpLess,
	ZJumpGreater,
	ZDecChk,
	ZIncChk,
	ZJin,
	ZTest,
	ZOr,
	ZAnd,
	ZTestAttr,
	ZSetAttr,
	ZClearAttr,
	ZStore,
	ZInsertObj,
	ZLoadW,
	ZLoadB,
	ZGetProp,
	ZGetPropAddr,
	ZGetNextProp,
	ZAdd,
	ZSub,
	ZMul,
	ZDiv,
	ZMod,
}

var ZFunctions_1OP = []ZFunction1Op{
	ZJumpZero,
	ZGetSibling,
	ZGetChild,
	ZGetParent,
	ZGetPropLen,
	ZInc,
	ZDec,
	ZPrintAddr,
	ZNOP1,
	ZRemoveObj,
	ZPrintObj,
	ZRet,
	ZJump,
	ZPrintPAddr,
	ZLoad,
	ZNOP1,
}

var ZFunctions_0P = []ZFunction0Op{
	ZReturnTrue,
	ZReturnFalse,
	ZPrint,
	ZPrintRet,
	ZNOP0,
	ZNOP0,
	ZNOP0,
	ZNOP0,
	ZRetPopped,
	ZPop,
	ZQuit,
	ZNewLine,
}

func (zm *ZMachine) GetOperand(operandType byte) uint16 {

	var retValue uint16

	switch operandType {
	case OPERAND_SMALL:
		retValue = uint16(zm.buf[zm.ip])
		zm.ip++
	case OPERAND_VARIABLE:
		varType := zm.buf[zm.ip]
		// 0 = top of the stack
		// 1 - 0xF = locals
		// 0x10 - 0xFF = globals
		if varType == 0 {
			retValue = zm.stack.Pop()
		} else if varType < 0x10 {
			retValue = zm.stack.GetLocalVar((int)(varType - 1))
		} else {
			retValue = zm.ReadGlobal(varType)
		}
		zm.ip++
	case OPERAND_LARGE:
		retValue = zm.GetUint16(zm.ip)
		zm.ip += 2
	case OPERAND_OMITTED:
		return 0
	default:
		panic("Unknown operand type")
	}

	return retValue
}

func (zm *ZMachine) GetOperands(opTypesByte uint8, operandValues []uint16) uint16 {
	numOperands := uint16(0)
	var shift uint8
	shift = 6

	for i := 0; i < 4; i++ {
		opType := (opTypesByte >> shift) & 0x3
		shift -= 2
		if opType == OPERAND_OMITTED {
			break
		}

		opValue := zm.GetOperand(opType)
		operandValues[numOperands] = opValue
		numOperands++
	}

	return numOperands
}

func (zm *ZMachine) StoreAtLocation(storeLocation uint16, v uint16) {
	// Same deal as read variable
	// 0 = top of the stack, 0x1-0xF = local var, 0x10 - 0xFF = global var

	if storeLocation == 0 {
		zm.stack.Push(v)
	} else if storeLocation < 0x10 {
		zm.stack.SetLocalVar((int)(storeLocation-1), v)
	} else {
		zm.SetGlobal(storeLocation, v)
	}
}

func (zm *ZMachine) StoreResult(v uint16) {
	storeLocation := zm.ReadByte()

	zm.StoreAtLocation(uint16(storeLocation), v)
}

func (zm *ZMachine) InterpretVARInstruction() {

	opcode := zm.ReadByte()
	// "In variable form, if bit 5 is 0 then the count is 2OP; if it is 1, then the count is VAR.
	// The opcode number is given in the bottom 5 bits.
	instruction := (opcode & 0x1F)
	twoOp := ((opcode >> 5) & 0x1) == 0

	// "In variable or extended forms, a byte of 4 operand types is given next.
	// This contains 4 2-bit fields: bits 6 and 7 are the first field, bits 0 and 1 the fourth."
	// "A value of 0 means a small constant and 1 means a variable."
	opTypesByte := zm.ReadByte()

	opValues := make([]uint16, 4)
	numOperands := zm.GetOperands(opTypesByte, opValues)

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
	instruction := (opcode & 0x0F)

	if opType != OPERAND_OMITTED {
		opValue := zm.GetOperand(opType)

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
	instruction := (opcode & 0x1F)

	// Operand types:
	// In long form, bit 6 of the opcode gives the type of the first operand, bit 5 of the second.
	// A value of 0 means a small constant and 1 means a variable.
	operandType0 := ((opcode & 0x40) >> 6) + 1
	operandType1 := ((opcode & 0x20) >> 5) + 1

	opValues := make([]uint16, 2)
	opValue0 := zm.GetOperand(operandType0)
	opValue1 := zm.GetOperand(operandType1)

	opValues[0] = opValue0
	opValues[1] = opValue1

	fn := ZFunctions_2OP[instruction]
	fn(zm, opValues, 2)
}

func (zm *ZMachine) InterpretInstruction() {
	opcode := zm.PeekByte()

	// Form is stored in top 2 bits
	// "If the top two bits of the opcode are $$11 the form is variable; if $$10, the form is short.
	// If the opcode is 190 ($BE in hexadecimal) and the version is 5 or later, the form is "extended".
	// Otherwise, the form is "long"."
	form := (opcode >> 6) & 0x3

	DebugPrintf("IP: 0x%X - opcode: 0x%X form: %02b\n", zm.ip, opcode, form)

	if opcode == 0xBE && zm.header.version >= 5 {
		panic("Extended not supported")
	} else if form == 0x2 {
		zm.InterpretShortInstruction()
	} else if form == 0x3 {
		zm.InterpretVARInstruction()
	} else {
		zm.InterpretLongInstruction()
	}
}

// NOTE: Doesn't support abbreviations.
func (zm *ZMachine) EncodeText(txt string) uint32 {

	encodedChars := make([]uint8, 12)
	encodedWords := make([]uint16, 2)
	padding := uint8(0x5)

	// Store 6 Z-chars. Clamp if longer, add padding if shorter
	i := 0
	j := 0
	for i < 6 {
		if j < len(txt) {
			c := txt[j]
			j++

			// See if we can find any alphabet
			ai := -1
			alphabetType := 0
			for a := 0; a < len(alphabets); a++ {
				index := strings.IndexByte(alphabets[a], c)
				if index >= 0 {
					ai = index
					alphabetType = a
					break
				}
			}
			if ai >= 0 {
				if alphabetType != 0 {
					// Alphabet change
					encodedChars[i] = uint8(alphabetType + 3)
					encodedChars[i+1] = uint8(ai + 6)
					i += 2
				} else {
					encodedChars[i] = uint8(ai + 6)
					i++
				}
			} else {
				// 10-bit ZC
				encodedChars[i] = 5
				encodedChars[i+1] = 6
				encodedChars[i+2] = (c >> 5)
				encodedChars[i+3] = (c & 0x1F)
				i += 4
			}
		} else {
			// Padding
			encodedChars[i] = padding
			i++
		}
	}

	for i := 0; i < 2; i++ {
		encodedWords[i] = (uint16(encodedChars[i*3+0]) << 10) | (uint16(encodedChars[i*3+1]) << 5) |
			uint16(encodedChars[i*3+2])
		if i == 1 {
			encodedWords[i] |= 0x8000
		}
	}

	return (uint32(encodedWords[0]) << 16) | uint32(encodedWords[1])
}

func NewZMachine(buffer []uint8, header ZHeader) *ZMachine {
	zm := new(ZMachine)
	zm.buf = buffer
	zm.header = header
	zm.ip = uint32(header.ip)
	zm.stack = NewStack()
	zm.InitObjectsConstants()
	//zm.ListAbbreviations()
	//zm.ListDictionary()
	//zm.ListObjects()
	//os.Exit(1)
	return zm
}

// Return DICT_NOT_FOUND (= 0) if not found
// Address in dictionary otherwise
func (zm *ZMachine) FindInDictionary(str string) uint16 {

	numSeparators := uint32(zm.buf[zm.header.dictAddress])
	entryLength := uint16(zm.buf[zm.header.dictAddress+1+numSeparators])
	numEntries := zm.GetUint16(zm.header.dictAddress + 1 + numSeparators + 1)

	entriesAddress := zm.header.dictAddress + 1 + numSeparators + 1 + 2

	// Dictionary entries are sorted, so we can use binary search
	lowerBound := uint16(0)
	upperBound := numEntries - 1

	encodedText := zm.EncodeText(str)

	foundAddress := uint16(DICT_NOT_FOUND)
	for lowerBound <= upperBound {

		currentIndex := lowerBound + (upperBound-lowerBound)/2
		dictValue := GetUint32(zm.buf, entriesAddress+uint32(currentIndex*entryLength))

		if encodedText < dictValue {
			upperBound = currentIndex - 1
		} else if encodedText > dictValue {
			lowerBound = currentIndex + 1
		} else {
			foundAddress = uint16(entriesAddress + uint32(currentIndex*entryLength))
			break
		}
	}

	return foundAddress
}

func (zm *ZMachine) GetPropertyDefault(propertyIndex uint16) uint16 {
	if propertyIndex < 1 || propertyIndex > 31 {
		panic("Invalid propertyIndex")
	}

	// 1-based -> 0-based
	propertyIndex--
	return zm.GetUint16(zm.header.objTableAddress + uint32(propertyIndex*2))
}

func PrintZChar(output io.Writer, ch uint16) {
	if ch == 13 {
		_, _ = fmt.Fprintf(output, "\n")
	} else if ch >= 32 && ch <= 126 { // ASCII
		_, _ = fmt.Fprintf(output, "%c", ch)
	} // else ... do not bother
}

// V3 only
// Returns offset pointing just after the string data
func (zm *ZMachine) DecodeZString(startOffset uint32) uint32 {

	done := false
	var zchars []uint8

	i := startOffset
	for !done {

		//--first byte-------   --second byte---
		//7    6 5 4 3 2  1 0   7 6 5  4 3 2 1 0
		//bit  --first--  --second---  --third--

		// Text in memory consists of a sequence of 2-byte words. Each word is divided into three 5-bit 'Z-characters', plus 1 bit left over, arranged as
		// --first byte-------   --second byte---
		// 7    6 5 4 3 2  1 0   7 6 5  4 3 2 1 0
		// bit  --first--  --second---  --third--
		w16 := zm.GetUint16(i)

		done = (w16 & 0x8000) != 0
		zchars = append(zchars, uint8((w16>>10)&0x1F), uint8((w16>>5)&0x1F), uint8(w16&0x1F))

		i += 2
	}

	alphabetType := 0

	for i := 0; i < len(zchars); i++ {
		zc := zchars[i]
		// z Characters 1,2,3 represent abbreviations, sometimes also called synonyms
		if zc > 0 && zc < 4 {
			//fmt.Println("Abbreviation", zc)
			// TODO, not sure if the if is correct
			var abbrevIndex uint8
			if i+1 >= len(zchars) {
				i++
				continue
			} else {
				abbrevIndex = zchars[i+1]
			}

			// "If z is the first Z-character (1, 2 or 3) and x the subsequent one,
			// then the interpreter must look up entry 32(z-1)+x in the abbreviations table"
			abbrevAddress := zm.GetUint16(zm.header.abbreviationTable + uint32(64*(zc-1)+abbrevIndex*2))
			zm.DecodeZString(zm.PackedAddress(uint32(abbrevAddress)))

			alphabetType = 0
			i++
			continue
		}
		if zc == 4 {
			alphabetType = 1
			continue
		} else if zc == 5 {
			alphabetType = 2
			continue
		}

		// Z-character 6 from A2 means that the two subsequent Z-characters specify a ten-bit ZSCII character code:
		// the next Z-character gives the top 5 bits and the one after the bottom 5.
		if alphabetType == 2 && zc == 6 {

			zc10 := (uint16(zchars[i+1]) << 5) | uint16(zchars[i+2])
			PrintZChar(&zm.output, zc10)

			i += 2

			alphabetType = 0
			continue
		}

		// z-character 0 is printed as a space
		if zc == 0 {
			_, _ = fmt.Fprintf(&zm.output, " ")
		} else {
			// If we're here zc >= 6. Alphabet tables are indexed starting at 6
			aindex := zc - 6
			_, _ = fmt.Fprintf(&zm.output, "%c", alphabets[alphabetType][aindex])
		}

		alphabetType = 0
	}

	return i
}
