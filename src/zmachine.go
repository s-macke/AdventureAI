package zmachine

// https://gitlab.com/DavidGriffith/frotz/
// https://www.inform-fiction.org/zmachine/standards/

import (
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
	name       string
	ip         uint32
	header     ZHeader
	backupBuf  []uint8 // the initial buffer
	buf        []uint8
	stack      *ZStack
	localFrame uint16
	done       bool
	output     strings.Builder
	input      func() string

	outputstream      int // the id to where we are writing output
	outputstreamtable uint32
	windowId          int // The selected window ID to print
}

type ZFunction func(*ZMachine, []uint16, uint16)
type ZFunction1Op func(*ZMachine, uint16)
type ZFunction0Op func(*ZMachine)

var alphabets = []string{
	"abcdefghijklmnopqrstuvwxyz",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	" \n0123456789.,!?_#'\"/\\-:()"}

func (zm *ZMachine) ReadGlobal(x uint8) uint16 {
	if x < 0x10 {
		panic("Invalid global variable")
	}

	addr := (uint32(x) - 0x10) * 2
	ret := zm.GetUint16(zm.header.globalVarAddress + addr)

	return ret
}

func (zm *ZMachine) SetGlobal(x uint16, v uint16) {
	if x < 0x10 {
		panic("Invalid global variable")
	}

	addr := (uint32(x) - 0x10) * 2
	zm.SetUint16(zm.header.globalVarAddress+addr, v)
}

func ZPutProp(zm *ZMachine, args []uint16, numArgs uint16) {
	zm.SetObjectProperty(args[0], args[1], args[2])
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
	//DebugPrintf("Popped %d 0x%X %d %d\n", r, zm.ip, numArgs, args[0])
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

		//DebugPrintf("Offset: 0x%X [%d]\n", branchOffset, branchOffset)
	}
	ip := int32(zm.ip)

	// "Otherwise, a branch moves execution to the instruction at address
	// Address after branch data + Offset - 2."
	jumpAddress = ip + int32(branchOffset) - 2

	//DebugPrintf("Jump address = 0x%X\n", jumpAddress)

	doJump := (conditionSatisfied != branchOnFalse)

	//DebugPrintf("Do jump: %t\n", doJump)

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
	zm.StoreAtLocation(args[0], args[1])
}

func ZTestAttr(zm *ZMachine, args []uint16, numArgs uint16) {
	GenericBranch(zm, zm.TestObjectAttr(args[0], args[1]))
}

func ZCheckArgCountArgumentNumber(zm *ZMachine, args []uint16, numargs uint16) { // check arg count
	//DebugPrintf("Arg count: %d %v\n", numargs, args)
	argumentNumber := args[0]
	GenericBranch(zm, argumentNumber <= zm.stack.stack[zm.stack.localFrame+1]) // localFrame points to the number of arguments. See ZCall
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
	sibling := zm.GetSiblingIndex(arg)
	zm.StoreResult(sibling)
	GenericBranch(zm, sibling != NULL_OBJECT_INDEX)
}

// get_child object -> (result) ?(label)
// Get first object contained in given object, branching if this exists, i.e. is not nothing (i.e., is not 0).
func ZGetChild(zm *ZMachine, arg uint16) {
	childIndex := zm.GetChildIndex(arg)
	zm.StoreResult(childIndex)
	GenericBranch(zm, childIndex != NULL_OBJECT_INDEX)
}

func ZGetParent(zm *ZMachine, arg uint16) {
	parent := zm.GetParentObjectIndex(arg)
	zm.StoreResult(parent)
}

func ZGetPropLen(zm *ZMachine, arg uint16) {
	if arg == 0 {
		zm.StoreResult(0)
	} else {
		// Arg = direct address of the property block
		// To get size, we need to go 1 byte back
		propSize := uint16(zm.buf[arg-1])
		numBytes := uint16(0)
		if zm.header.version <= 3 {
			numBytes = (propSize >> 5) + 1
		} else {
			if (propSize & 0x80) == 0 {
				numBytes = (propSize >> 6) + 1
			} else {
				numBytes = propSize & 0x3f
				if propSize == 0 {
					propSize = 64
				}
			}
		}
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
	//DebugPrintf("Jump address: 0x%X\n", jumpAddress)
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

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func TokenizeLine(zm *ZMachine, textaddress uint16, tokenaddress uint16, dct uint16, flag bool) {
	DebugPrintf("tokenise_line: text=%d token=%d dct=%d flag=%d\n", textaddress, tokenaddress, dct, Btoi(flag))
	text := ""

	if zm.header.version <= 4 {
		maxsize := int(zm.GetUint8(uint32(textaddress)))
		//fmt.Println(size)
		for i := 0; i < maxsize; i++ {
			char := zm.buf[textaddress+1+uint16(i)]
			if char == 0 {
				break
			}
			text += string(char)
		}
	} else {
		size := int(zm.GetUint8(uint32(textaddress + 1)))
		for i := 0; i < size; i++ {
			char := zm.buf[textaddress+2+uint16(i)]
			text += string(char)
		}
	}
	DebugPrintf("%s\n", text)
	words := strings.Split(text, " ")
	wordStarts := make([]uint16, len(words))
	wordStarts[0] = 1
	for i := 0; i < len(words)-1; i++ {
		wordStarts[i+1] = wordStarts[i] + uint16(len(words[i])) + 1
	}

	/*
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
	*/
	// TODO: include other separators, not only spaces

	parseAddress := uint32(tokenaddress)
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

		zm.SetUint16(parseAddress, dictionaryAddress)
		zm.buf[parseAddress+2] = uint8(len(w))
		zm.buf[parseAddress+3] = uint8(wordStarts[i])
		parseAddress += 4
	}
	//panic("Not implemented")
}

func ZTokenize(zm *ZMachine, args []uint16, numArgs uint16) {
	TokenizeLine(zm, args[0], args[1], args[2], args[3] != 0)
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
	DebugPrintf("Store %d - %d\n", storeLocation, v)
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

func NewZMachine(name string, buffer []uint8, header ZHeader) *ZMachine {
	zm := new(ZMachine)
	zm.name = name
	zm.backupBuf = buffer
	zm.header = header
	ZRestart(zm)

	/* Adjust opcode tables */
	if zm.header.version < 4 {
		ZFunctions_0P[0x09] = ZPop
		ZFunctions_1OP[0x0f] = nil // TODO: ZNot
	} else {
		ZFunctions_0P[0x09] = nil // TODO: ZCatch
		//ZFunctions_1OP[0x0f] = z_call_n; // already done
	}

	//zm.ListAbbreviations()
	//zm.ListDictionary()
	//zm.ListObjects()
	//os.Exit(1)

	return zm
}

func ZRestart(zm *ZMachine) {
	zm.buf = make([]uint8, len(zm.backupBuf))
	copy(zm.buf, zm.backupBuf)

	zm.ip = uint32(zm.header.ip)
	zm.stack = NewStack()
	zm.InitObjectsConstants()

	zm.buf[1] = 158 // TODO: why?

	zm.buf[22] = 25  // screen rows
	zm.buf[33] = 232 // screen cols

	// Z-Machine standard 1.1
	zm.buf[50] = 1
	zm.buf[51] = 1

	zm.buf[16] = 0  // flags
	zm.buf[17] = 16 // flags
}

func (zm *ZMachine) GetPropertyDefault(propertyIndex uint16) uint16 {
	var maxProperty uint16 = 31
	if zm.header.version >= 4 {
		maxProperty = 63
	}

	if propertyIndex < 1 || propertyIndex > maxProperty {
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
