package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type ZHeader struct {
	version           uint8
	hiMemBase         uint16
	ip                uint16
	dictAddress       uint32
	objTableAddress   uint32
	globalVarAddress  uint32
	staticMemAddress  uint32
	abbreviationTable uint32
}

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

	OBJECT_ENTRY_SIZE    = 9
	OBJECT_PARENT_INDEX  = 4
	OBJECT_SIBLING_INDEX = 5
	OBJECT_CHILD_INDEX   = 6
	NULL_OBJECT_INDEX    = 0

	DICT_NOT_FOUND = 0
)

type ZStack struct {
	stack      []uint16
	top        int
	localFrame int
}

func NewStack() *ZStack {
	s := new(ZStack)
	s.stack = make([]uint16, MAX_STACK)
	s.top = MAX_STACK
	s.localFrame = s.top

	return s
}

type ZMachine struct {
	ip         uint32
	header     ZHeader
	buf        []uint8
	stack      *ZStack
	localFrame uint16
	done       bool
}

type ZFunction func(*ZMachine, []uint16, uint16)
type ZFunction1Op func(*ZMachine, uint16)
type ZFunction0Op func(*ZMachine)

func DebugPrintf(format string, v ...interface{}) {
	//fmt.Printf(format, v...)
}

var alphabets = []string{"abcdefghijklmnopqrstuvwxyz",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	" \n0123456789.,!?_#'\"/\\-:()"}

func (s *ZStack) Push(value uint16) {
	if s.top == 0 {
		panic("Stack overflow")
	}
	s.top--
	s.stack[s.top] = value
}

func (s *ZStack) Pop() uint16 {
	if s.top == MAX_STACK {
		panic("Trying to pop from empty stack")
	}
	retValue := s.stack[s.top]

	s.top++
	return retValue
}

func (s *ZStack) Reset(newTop int) {
	if newTop > MAX_STACK || newTop < 0 {
		panic("Invalid stack top value")
	}
	s.top = newTop
}

func (s *ZStack) GetTopItem() uint16 {
	return s.stack[s.top]
}

func (s *ZStack) SaveFrame() {
	s.Push(uint16(s.localFrame))
	s.localFrame = s.top
}

// Returns caller address (where to return to)
func (s *ZStack) RestoreFrame() uint32 {

	// Discard local frame
	s.top = s.localFrame
	// Restore previous frame
	s.localFrame = int(s.Pop())

	retLo := s.Pop()
	retHi := s.Pop()

	return (uint32(retHi) << 16) | uint32(retLo)
}

func (s *ZStack) ValidateLocalVarIndex(localVarIndex int) {
	if localVarIndex > 0xF {
		panic("Local var index out of bounds")
	}
	if s.localFrame < localVarIndex {
		panic("Stack underflow")
	}
}
func (s *ZStack) GetLocalVar(localVarIndex int) uint16 {
	s.ValidateLocalVarIndex(localVarIndex)
	stackIndex := (s.localFrame - localVarIndex) - 1
	r := s.stack[stackIndex]
	return r
}

func (s *ZStack) SetLocalVar(localVarIndex int, value uint16) {
	s.ValidateLocalVarIndex(localVarIndex)
	stackIndex := (s.localFrame - localVarIndex) - 1
	s.stack[stackIndex] = value
}

func (s *ZStack) Dump() {
	DebugPrintf("Top = %d, local frame = %d\n", s.top, s.localFrame)

	for i := MAX_STACK - 1; i >= s.top; i-- {
		if i == s.localFrame {
			DebugPrintf("0x%X: 0x%X <------ local frame\n", i, s.stack[i])
		} else {
			DebugPrintf("0x%X: 0x%X\n", i, s.stack[i])
		}
	}
}

func GetUint16(buf []byte, offset uint32) uint16 {
	return (uint16(buf[offset]) << 8) | (uint16)(buf[offset+1])
}

func GetUint32(buf []byte, offset uint32) uint32 {
	return (uint32(buf[offset]) << 24) | (uint32(buf[offset+1]) << 16) | (uint32(buf[offset+2]) << 8) | uint32(buf[offset+3])
}

func (h *ZHeader) read(buf []byte) {

	h.version = buf[0]
	h.hiMemBase = GetUint16(buf, 4)
	h.ip = GetUint16(buf, 6)
	h.dictAddress = uint32(GetUint16(buf, 0x8))
	h.objTableAddress = uint32(GetUint16(buf, 0xA))
	h.globalVarAddress = uint32(GetUint16(buf, 0xC))
	h.staticMemAddress = uint32(GetUint16(buf, 0xE))
	h.abbreviationTable = uint32(GetUint16(buf, 0x18))

	DebugPrintf("End of dyn mem: 0x%X\n", h.staticMemAddress)
	DebugPrintf("Global vars: 0x%X\n", h.globalVarAddress)
}

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

// We can only write to dynamic memory
func (zm *ZMachine) IsSafeToWrite(address uint32) bool {
	return address < zm.header.staticMemAddress
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
func PackedAddress(a uint32) uint32 {
	return a * 2
}

func (zm *ZMachine) ReadGlobal(x uint8) uint16 {
	if x < 0x10 {
		panic("Invalid global variable")
	}

	addr := PackedAddress(uint32(x) - 0x10)
	ret := zm.GetUint16(zm.header.globalVarAddress + addr)

	return ret
}

func (zm *ZMachine) SetGlobal(x uint16, v uint16) {
	if x < 0x10 {
		panic("Invalid global variable")
	}

	addr := PackedAddress(uint32(x) - 0x10)
	zm.SetUint16(zm.header.globalVarAddress+addr, v)
}

func (zm *ZMachine) GetObjectEntryAddress(objectIndex uint16) uint32 {
	if objectIndex > MAX_OBJECT || objectIndex == 0 {
		fmt.Printf("Index: %d\n", objectIndex)
		panic("Invalid object index")
	}

	// Convert from 1-based (0 = NULL = no object) to 0-based

	objectIndex--
	// Skip default props
	objectEntryAddress := zm.header.objTableAddress + (31 * 2) + uint32(objectIndex*OBJECT_ENTRY_SIZE)

	return uint32(objectEntryAddress)
}

func (zm *ZMachine) SetObjectProperty(objectIndex uint16, propertyId uint16, value uint16) {

	objectEntryAddress := uint32(zm.GetObjectEntryAddress(objectIndex))

	propertiesAddress := GetUint16(zm.buf, objectEntryAddress+7)
	nameLength := uint16(zm.buf[propertiesAddress]) * 2 // in 2-byte words

	// Find property
	found := false
	propData := uint32(propertiesAddress + nameLength + 1)

	for !found {
		propSize := zm.buf[propData]
		if propSize == 0 {
			break
		}
		propData++
		propNo := uint16(propSize & 0x1F)

		// Props are sorted
		if propNo < propertyId {
			break
		}

		numBytes := (propSize >> 5) + 1
		if propNo == propertyId {
			found = true

			if numBytes == 1 {
				zm.buf[propData] = uint8(value & 0xFF)
			} else if numBytes == 2 {
				zm.SetUint16(propData, value)
			} else {
				panic("SetObjectProperty only supports 1/2 byte properties")
			}
		}
		propData += uint32(numBytes)
	}
	if !found {
		panic("Property not found!")
	}
}

func (zm *ZMachine) GetFirstPropertyAddress(objectIndex uint16) uint16 {
	objectEntryAddress := uint32(zm.GetObjectEntryAddress(objectIndex))
	propertiesAddress := GetUint16(zm.buf, objectEntryAddress+7)
	nameLength := uint16(zm.buf[propertiesAddress]) * 2 // in 2-byte words
	propData := propertiesAddress + nameLength + 1

	return propData
}

// Returns prop data address, number of property bytes
// (0 if not found)
func (zm *ZMachine) GetObjectPropertyInfo(objectIndex uint16, propertyId uint16) (uint16, uint16) {

	propData := zm.GetFirstPropertyAddress(objectIndex)

	// Find property
	found := false

	for !found {
		propSize := zm.buf[propData]
		if propSize == 0 {
			break
		}
		propData++
		propNo := uint16(propSize & 0x1F)

		// Props are sorted
		if propNo < propertyId {
			break
		}

		numBytes := uint16(propSize>>5) + 1
		if propNo == propertyId {
			return propData, numBytes
		}
		propData += uint16(numBytes)
	}
	return uint16(0), uint16(0)
}

func (zm *ZMachine) GetObjectPropertyAddress(objectIndex uint16, propertyId uint16) uint16 {
	address, _ := zm.GetObjectPropertyInfo(objectIndex, propertyId)
	return address
}

func (zm *ZMachine) GetNextObjectProperty(objectIndex uint16, propertyId uint16) uint16 {

	nextPropSize := uint8(0)

	// " if called with zero, it gives the first property number present."
	if propertyId == 0 {
		propData := zm.GetFirstPropertyAddress(objectIndex)
		nextPropSize = zm.buf[propData]
	} else {
		propData, numBytes := zm.GetObjectPropertyInfo(objectIndex, propertyId)
		if propData == 0 {
			panic("GetNextObjectProperty - non existent property")
		}
		nextPropSize = zm.buf[propData+numBytes]
	}
	// "zero, indicating the end of the property list"
	if nextPropSize == 0 {
		return 0
	} else {
		return uint16(nextPropSize & 0x1F)
	}
}

func (zm *ZMachine) GetObjectProperty(objectIndex uint16, propertyId uint16) uint16 {

	propData, numBytes := zm.GetObjectPropertyInfo(objectIndex, propertyId)
	result := uint16(0)

	if propData == 0 {
		// Get a default one
		result = zm.GetPropertyDefault(propertyId)
		DebugPrintf("Default prop %d = 0x%X\n", propertyId, result)
	} else {
		if numBytes == 1 {
			result = uint16(zm.buf[propData])
		} else if numBytes == 2 {
			result = GetUint16(zm.buf, uint32(propData))
		} else {
			panic("GetObjectProperty only supports 1/2 byte properties")
		}
	}

	return result
}

// True if set
func (zm *ZMachine) TestObjectAttr(objectIndex uint16, attribute uint16) bool {

	if attribute > 31 {
		panic("Attribute out of bounds")
	}

	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)

	attribs := GetUint32(zm.buf, objectEntryAddress)
	// 0: top bit
	// 31: bottom bit
	mask := uint32(1 << (31 - attribute))

	return (attribs & mask) != 0
}

func (zm *ZMachine) SetObjectAttr(objectIndex uint16, attribute uint16) {

	if attribute > 31 {
		panic("Attribute out of bounds")
	}

	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	byteIndex := uint32(attribute >> 3)
	shift := 7 - (attribute & 0x7)

	zm.buf[objectEntryAddress+byteIndex] |= (1 << shift)
}

func (zm *ZMachine) ClearObjectAttr(objectIndex uint16, attribute uint16) {

	if attribute > 31 {
		panic("Attribute out of bounds")
	}

	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	byteIndex := uint32(attribute >> 3)
	shift := 7 - (attribute & 0x7)

	zm.buf[objectEntryAddress+byteIndex] &= ^(1 << shift)
}

func (zm *ZMachine) IsDirectParent(childIndex uint16, parentIndex uint16) bool {

	return zm.GetParentObject(childIndex) == parentIndex
}

func (zm *ZMachine) GetParentObject(objectIndex uint16) uint16 {
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)

	return uint16(zm.buf[objectEntryAddress+OBJECT_PARENT_INDEX])
}

// Unlink object from its parent
func (zm *ZMachine) UnlinkObject(objectIndex uint16) {
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	currentParentIndex := uint16(zm.buf[objectEntryAddress+OBJECT_PARENT_INDEX])

	// Unlink from current parent first
	if currentParentIndex != NULL_OBJECT_INDEX {
		curParentAddress := zm.GetObjectEntryAddress(currentParentIndex)
		// If we're the first child -> move to sibling
		if uint16(zm.buf[curParentAddress+OBJECT_CHILD_INDEX]) == objectIndex {
			zm.buf[curParentAddress+OBJECT_CHILD_INDEX] = zm.buf[objectEntryAddress+OBJECT_SIBLING_INDEX]
		} else {
			childIter := uint16(zm.buf[curParentAddress+OBJECT_CHILD_INDEX])
			prevChild := uint16(NULL_OBJECT_INDEX)
			for childIter != objectIndex && childIter != NULL_OBJECT_INDEX {
				prevChild = childIter
				childIter = zm.GetSibling(childIter)
			}
			// Sanity checks
			if childIter == NULL_OBJECT_INDEX {
				panic("Object not found on parent children list")
			}
			if prevChild == NULL_OBJECT_INDEX {
				panic("Corrupted data")
			}

			prevSiblingAddress := zm.GetObjectEntryAddress(prevChild)
			sibling := zm.buf[objectEntryAddress+OBJECT_SIBLING_INDEX]
			zm.buf[prevSiblingAddress+OBJECT_SIBLING_INDEX] = sibling
		}
		zm.buf[objectEntryAddress+OBJECT_PARENT_INDEX] = NULL_OBJECT_INDEX
	}
}

func (zm *ZMachine) ReparentObject(objectIndex uint16, newParentIndex uint16) {

	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	currentParentIndex := uint16(zm.buf[objectEntryAddress+OBJECT_PARENT_INDEX])

	if currentParentIndex == newParentIndex {
		return
	}

	zm.UnlinkObject(objectIndex)

	// Make the first child of our new parent
	newParentAddress := zm.GetObjectEntryAddress(newParentIndex)
	zm.buf[objectEntryAddress+OBJECT_SIBLING_INDEX] = zm.buf[newParentAddress+OBJECT_CHILD_INDEX]
	zm.buf[newParentAddress+OBJECT_CHILD_INDEX] = uint8(objectIndex)
	zm.buf[objectEntryAddress+OBJECT_PARENT_INDEX] = uint8(newParentIndex)
}

func (zm *ZMachine) GetFirstChild(objectIndex uint16) uint16 {
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)

	return uint16(zm.buf[objectEntryAddress+OBJECT_CHILD_INDEX])
}

func (zm *ZMachine) GetSibling(objectIndex uint16) uint16 {
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)

	return uint16(zm.buf[objectEntryAddress+OBJECT_SIBLING_INDEX])
}

func (zm *ZMachine) PrintObjectName(objectIndex uint16) {
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	propertiesAddress := uint32(GetUint16(zm.buf, objectEntryAddress+7))
	zm.DecodeZString(propertiesAddress + 1)
}

func ZCall(zm *ZMachine, args []uint16, numArgs uint16) {
	if numArgs == 0 {
		panic("Data corruption, call instruction requires at least 1 argument")
	}

	// Save return address
	zm.stack.Push(uint16(zm.ip>>16) & 0xFFFF)
	zm.stack.Push(uint16(zm.ip & 0xFFFF))

	functionAddress := PackedAddress(uint32(args[0]))
	DebugPrintf("Jumping to 0x%X [0x%X]\n", functionAddress, args[0])

	zm.ip = functionAddress

	// Save local frame (think EBP)
	zm.stack.SaveFrame()

	if zm.ip == 0 {
		ZReturnFalse(zm)

		return
	}

	// Local function variables on the stack
	numLocals := zm.ReadByte()

	// "When a routine is called, its local variables are created with initial values taken from the routine header.
	// Next, the arguments are written into the local variables (argument 1 into local 1 and so on)."
	numArgs-- // first argument is function address
	for i := 0; i < int(numLocals); i++ {
		localVar := zm.ReadUint16()

		if numArgs > 0 {
			localVar = args[i+1]
			numArgs--
		}
		zm.stack.Push(localVar)
	}
}

// storew array word-index value
func ZStoreW(zm *ZMachine, args []uint16, numArgs uint16) {

	address := uint32(args[0] + args[1]*2)
	if !zm.IsSafeToWrite(address) {
		panic("Access violation")
	}

	zm.SetUint16(address, args[2])
}

func ZStoreB(zm *ZMachine, args []uint16, numArgs uint16) {

	address := uint32(args[0] + args[1])
	if !zm.IsSafeToWrite(address) {
		panic("Access violation")
	}

	zm.buf[address] = uint8(args[2])
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

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

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
	PrintZChar(ch)
}

func ZPrintNum(zm *ZMachine, args []uint16, numArgs uint16) {
	fmt.Printf("%d", int16(args[0]))
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

func ZLoadB(zm *ZMachine, args []uint16, numArgs uint16) {

	address := args[0] + args[1]
	value := zm.buf[address]

	zm.StoreResult(uint16(value))
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

// array word-index -> (result)
func ZLoadW(zm *ZMachine, args []uint16, numArgs uint16) {
	address := uint32(args[0] + (args[1] * 2))
	value := GetUint16(zm.buf, address)

	zm.StoreResult(value)
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
	zm.DecodeZString(uint32(arg) * 2)
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

func ZRet(zm *ZMachine, arg uint16) {
	returnAddress := zm.stack.RestoreFrame()
	zm.ip = returnAddress
	DebugPrintf("Returning to 0x%X\n", zm.ip)

	zm.StoreResult(arg)
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
	fmt.Printf("\n")
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
	fmt.Printf("\n")
}

func ZNOP0(zm *ZMachine) {
	panic("NOP0")
}

var ZFunctions_VAR = []ZFunction{
	ZCall,
	ZStoreW,
	ZStoreB,
	ZPutProp,
	ZRead,
	ZPrintChar,
	ZPrintNum,
	ZRandom,
	ZPush,
	ZPull,
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
		retValue = GetUint16(zm.buf, zm.ip)
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

	DebugPrintf("IP: 0x%X - opcode: 0x%X\n", zm.ip, opcode)
	// Form is stored in top 2 bits
	// "If the top two bits of the opcode are $$11 the form is variable; if $$10, the form is short.
	// If the opcode is 190 ($BE in hexadecimal) and the version is 5 or later, the form is "extended".
	// Otherwise, the form is "long"."
	form := (opcode >> 6) & 0x3

	if form == 0x2 {
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

func (zm *ZMachine) Initialize(buffer []uint8, header ZHeader) {
	zm.buf = buffer
	zm.header = header
	zm.ip = uint32(header.ip)
	zm.stack = NewStack()

	//zm.TestDictionary()
}

// Return DICT_NOT_FOUND (= 0) if not found
// Address in dictionary otherwise
func (zm *ZMachine) FindInDictionary(str string) uint16 {

	numSeparators := uint32(zm.buf[zm.header.dictAddress])
	entryLength := uint16(zm.buf[zm.header.dictAddress+1+numSeparators])
	numEntries := GetUint16(zm.buf, zm.header.dictAddress+1+numSeparators+1)

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
	return GetUint16(zm.buf, zm.header.objTableAddress+uint32(propertyIndex*2))
}

func PrintZChar(ch uint16) {
	if ch == 13 {
		fmt.Printf("\n")
	} else if ch >= 32 && ch <= 126 { // ASCII
		fmt.Printf("%c", ch)
	} // else ... do not bother
}

// V3 only
// Returns offset pointing just after the string data
func (zm *ZMachine) DecodeZString(startOffset uint32) uint32 {

	done := false
	zchars := []uint8{}

	i := startOffset
	for !done {

		//--first byte-------   --second byte---
		//7    6 5 4 3 2  1 0   7 6 5  4 3 2 1 0
		//bit  --first--  --second---  --third--

		w16 := GetUint16(zm.buf, i)

		done = (w16 & 0x8000) != 0
		zchars = append(zchars, uint8((w16>>10)&0x1F), uint8((w16>>5)&0x1F), uint8(w16&0x1F))

		i += 2
	}

	alphabetType := 0

	for i := 0; i < len(zchars); i++ {
		zc := zchars[i]

		// Abbreviation
		if zc > 0 && zc < 4 {
			abbrevIndex := zchars[i+1]

			// "If z is the first Z-character (1, 2 or 3) and x the subsequent one,
			// then the interpreter must look up entry 32(z-1)+x in the abbreviations table"
			abbrevAddress := GetUint16(zm.buf, zm.header.abbreviationTable+uint32(32*(zc-1)+abbrevIndex)*2)
			zm.DecodeZString(PackedAddress(uint32(abbrevAddress)))

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
			PrintZChar(zc10)

			i += 2

			alphabetType = 0
			continue
		}

		if zc == 0 {
			fmt.Printf(" ")
		} else {
			// If we're here zc >= 6. Alphabet tables are indexed starting at 6
			aindex := zc - 6
			fmt.Printf("%c", alphabets[alphabetType][aindex])
		}

		alphabetType = 0
	}

	return i
}

func main() {
	filename := flag.String("file", "zork1.dat", "Z-Machine file to run")
	flag.Parse()
	buffer, err := os.ReadFile(*filename)
	if err != nil {
		panic(err)
	}

	var header ZHeader
	header.read(buffer)

	if header.version != 3 {
		panic("Only Version 3 files supported. But found version " + strconv.Itoa(int(header.version)))
	}

	var zm ZMachine
	zm.Initialize(buffer, header)

	for !zm.done {
		zm.InterpretInstruction()
	}
}
