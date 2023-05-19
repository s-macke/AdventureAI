package main

func ZCall(zm *ZMachine, args []uint16, numArgs uint16, callType ZCallType) {
	DebugPrintf("ZCall with numArgs %d callType %d and args %v\n", numArgs, callType, args)
	if numArgs == 0 {
		panic("Data corruption, call instruction requires at least 1 argument")
	}

	// Save return address
	zm.stack.Push(uint16(zm.ip>>16) & 0xFFFF)
	zm.stack.Push(uint16(zm.ip & 0xFFFF))
	zm.stack.Push(uint16(callType))
	zm.stack.Push(numArgs - 1)

	functionAddress := zm.PackedAddress(uint32(args[0]))
	DebugPrintf("Jumping to 0x%X\n", functionAddress)

	zm.ip = functionAddress

	// Save local frame (think EBP)
	zm.stack.SaveFrame()

	if zm.ip == 0 {
		ZReturnFalse(zm)
		return
	}

	// Local function variables on the stack
	numLocals := zm.ReadByte()
	DebugPrintf("Number of local variables: %d\n", numLocals)
	// "When a routine is called, its local variables are created with initial values taken from the routine header.
	// Next, the arguments are written into the local variables (argument 1 into local 1 and so on)."
	numArgs-- // first argument is function address
	for i := 0; i < int(numLocals); i++ {
		localVar := uint16(0)
		if zm.header.version <= 3 { // older versions provide default. From version 4 onwards, local variables are initialized to 0
			localVar = zm.ReadUint16()
		}
		if numArgs > 0 {
			localVar = args[i+1]
			numArgs--
		}
		zm.stack.Push(localVar)
	}
}

// ZRet returns from the current subroutine and restore the previous stack
// frame
func ZRet(zm *ZMachine, arg uint16) {

	zm.stack.RestoreFrame()

	_ = zm.stack.Pop() // numArgs
	callType := ZCallType(zm.stack.Pop())
	retLo := zm.stack.Pop()
	retHi := zm.stack.Pop()
	returnAddress := (uint32(retHi) << 16) | uint32(retLo)

	//DebugPrintf("ZRet callType=%d\n", callType)

	zm.ip = returnAddress
	DebugPrintf("Returning to 0x%X\n", zm.ip)
	if callType == ZCallTypeStore {
		zm.StoreResult(arg)
	}
}
