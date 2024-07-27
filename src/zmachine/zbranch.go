package zmachine

func GenericBranch(zm *ZMachine, conditionSatisfied bool) {
	branchInfo := zm.ReadByte()

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

	}
	ip := int32(zm.ip)

	// "Otherwise, a branch moves execution to the instruction at address
	// Address after branch data + Offset - 2."
	jumpAddress = ip + int32(branchOffset) - 2

	// "If bit 7 of the first byte is 0, a branch occurs when the condition was false; if 1, then branch is on true"
	branchOnFalse := (branchInfo >> 7) == 0
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
	if len(args) < 2 {
		panic("ZJumpEqual requires at least 2 arguments")
	}
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

func ZTestAttr(zm *ZMachine, args []uint16, numArgs uint16) {
	GenericBranch(zm, zm.TestObjectAttr(args[0], args[1]))
}

func ZCheckArgCountArgumentNumber(zm *ZMachine, args []uint16, numargs uint16) { // check arg count
	//DebugPrintf("Arg count: %d %v\n", numargs, args)
	argumentNumber := args[0]
	GenericBranch(zm, argumentNumber <= zm.stack.stack[zm.stack.localFrame+1]) // localFrame points to the number of arguments. See ZCall
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

func ZJumpZero(zm *ZMachine, arg uint16) {
	GenericBranch(zm, arg == 0)
}
