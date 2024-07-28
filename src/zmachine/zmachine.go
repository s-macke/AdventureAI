package zmachine

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
)

// https://gitlab.com/DavidGriffith/frotz/
// https://www.inform-fiction.org/zmachine/standards/
// https://www.ifarchive.org/if-archive/infocom/interpreters/specification/zspec02/zmach06e.pdf

import (
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

type Checkpoint struct {
	ip                uint32
	stack             *ZStack
	localFrame        uint16
	buf               []uint8
	outputstreamtable uint32
	outputstream      int
	WindowId          int
}

type ZMachine struct {
	Name       string
	ip         uint32
	header     ZHeader
	backupBuf  []uint8 // the initial buffer
	buf        []uint8
	stack      *ZStack
	localFrame uint16
	Done       bool
	Output     strings.Builder
	Input      func() string

	outputstream      int // the id to where we are writing output
	outputstreamtable uint32
	WindowId          int // The selected window ID to print
	PrevUndo          Checkpoint
	Undo              Checkpoint
}

type ZFunction func(*ZMachine, []uint16, uint16)
type ZFunction1Op func(*ZMachine, uint16)
type ZFunction0Op func(*ZMachine)

var alphabets = []string{
	"abcdefghijklmnopqrstuvwxyz",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	" \n0123456789.,!?_#'\"/\\-:()"}

func (zm *ZMachine) SaveUndo() {
	zm.PrevUndo = zm.Undo
	zm.Undo.outputstream = zm.outputstream
	zm.Undo.outputstreamtable = zm.outputstreamtable
	zm.Undo.ip = zm.ip
	zm.Undo.localFrame = zm.localFrame
	zm.Undo.WindowId = zm.WindowId
	zm.Undo.buf = make([]uint8, len(zm.buf))
	copy(zm.Undo.buf, zm.buf)
	zm.Undo.stack = zm.stack.Clone()
}

func (zm *ZMachine) Restore() {
	zm.outputstream = zm.PrevUndo.outputstream
	zm.outputstreamtable = zm.PrevUndo.outputstreamtable
	zm.ip = zm.PrevUndo.ip
	zm.localFrame = zm.PrevUndo.localFrame
	zm.WindowId = zm.PrevUndo.WindowId
	copy(zm.buf, zm.PrevUndo.buf)
	zm.stack = zm.PrevUndo.stack.Clone()
}

func (zm *ZMachine) ReadGlobal(x uint8) uint16 {
	if x < 0x10 {
		panic("Invalid global variable")
	}

	addr := (uint32(x) - 0x10) * 2
	ret := zm.GetUint16(zm.header.globalVarAddress + addr)
	//fmt.Println("Read global", x, ret)
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
	PrintZChar(&zm.Output, ch)
}

func ZPrintNum(zm *ZMachine, args []uint16, numArgs uint16) {
	_, _ = fmt.Fprintf(&zm.Output, "%d", int16(args[0]))
}

// If range is positive, returns a uniformly random number between 1 and range.
// If range is negative, the random number generator is seeded to that value and the return value is 0.
// Most interpreters consider giving 0 as range illegal (because they attempt a division with remainder by the range),
// / but correct behaviour is to reseed the generator in as random a way as the interpreter can (e.g. by using the time
// in milliseconds).
func ZRandom(zm *ZMachine, args []uint16, numArgs uint16) {
	DebugPrintf("ZRandom\n")
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
	//fmt.Printf("IP=0x%X\n", zm.ip)
	//panic("NOP VAR")
}

func ZNOP(zm *ZMachine, args []uint16) {
	fmt.Printf("IP=0x%X\n", zm.ip)
	panic("NOP 2OP")
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

//	jin obj1 obj2 ?(label)
//
// Jump if object a is a direct child of b, i.e., if parent of a is b.
func ZJin(zm *ZMachine, args []uint16, numArgs uint16) {
	GenericBranch(zm, zm.IsDirectParent(args[0], args[1]))
}

func ZInsertObj(zm *ZMachine, args []uint16, numArgs uint16) {
	zm.ReparentObject(args[0], args[1])
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
	DebugPrintf("Get child: %d\n", arg)
	childIndex := uint16(0)
	if arg != 0 {
		childIndex = zm.GetChildIndex(arg)
	}
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
		if zm.header.Version <= 3 {
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

	_, _ = fmt.Fprintf(&zm.Output, "\n")
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
	zm.Done = true
}

func ZNewLine(zm *ZMachine) {
	_, _ = fmt.Fprintf(&zm.Output, "\n")
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

	if zm.header.Version <= 4 {
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
	// TODO: include other separators, not only spaces
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
	//Remove all tokens before inserting new ones
	//zm.buf[tokenaddress+1] = 0

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
		dictionaryAddress := zm.FindInDictionary(w)
		DebugPrintf("w = '%s' starts=%d address=%d\n", w, wordStarts[i], dictionaryAddress)

		zm.SetUint16(parseAddress, dictionaryAddress)
		zm.buf[parseAddress+2] = uint8(len(w))
		zm.buf[parseAddress+3] = uint8(wordStarts[i])
		parseAddress += 4
	}
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
	zm.Name = name
	zm.backupBuf = buffer
	zm.header = header
	ZRestart(zm)

	/* Adjust opcode tables */
	if zm.header.Version < 4 {
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

	zm.buf[1] = 156 // TODO: Set config of z-machine. See z-machine header flag at offset 1

	zm.buf[22] = 48  // screen rows
	zm.buf[33] = 228 // screen cols

	// Z-Machine standard 1.1
	zm.buf[50] = 1
	zm.buf[51] = 1

	zm.buf[16] = 0 // flags
	//zm.buf[17] = 16 // flags
	zm.buf[17] = 0 // flags
}

func (zm *ZMachine) GetPropertyDefault(propertyIndex uint16) uint16 {
	var maxProperty uint16 = 31
	if zm.header.Version >= 4 {
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
