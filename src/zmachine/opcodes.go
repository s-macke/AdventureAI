package zmachine

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
	func(zm *ZMachine, args []uint16, numargs uint16) {
		// split window lines
	},
	func(zm *ZMachine, args []uint16, numargs uint16) {
		//fmt.Println("Set Window", args)
		zm.WindowId = int(args[0])
		// set window
	},
	func(zm *ZMachine, args []uint16, numargs uint16) {
		// command 0xec, call_vs2
		ZCall(zm, args, numargs, ZCallTypeStore)
	},
	func(zm *ZMachine, args []uint16, numargs uint16) {
		// erase window
	},
	nil,
	func(zm *ZMachine, args []uint16, numargs uint16) {
		// set cursor
	},
	nil, // 0x10 get cursor array
	func(zm *ZMachine, args []uint16, numargs uint16) {
		// set Text Style
	},
	func(zm *ZMachine, args []uint16, numargs uint16) {
		// z_buffer_mode TODO
	},
	ZOutputStream,
	nil,
	nil, // sound effect
	func(zm *ZMachine, args []uint16, numargs uint16) {
		zm.StoreResult(13) // read char
	},
	nil,
	func(zm *ZMachine, args []uint16, numargs uint16) {
		zm.StoreResult(^args[0]) // not value -> (result)
	},
	func(zm *ZMachine, args []uint16, numargs uint16) {
		// command 0xfa, call_vs2
		ZCall(zm, args, numargs, ZCallTypeN)
	},
	nil,
	ZTokenize,
	nil,
	nil,
	nil,
	ZCheckArgCountArgumentNumber,
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
	ZLoadB, // 0x10
	ZGetProp,
	ZGetPropAddr,
	ZGetNextProp,
	ZAdd,
	ZSub,
	ZMul,
	ZDiv,
	ZMod,
	func(zm *ZMachine, args []uint16, numargs uint16) {
		ZCall(zm, args, numargs, ZCallTypeStore)
	},
	func(zm *ZMachine, args []uint16, numargs uint16) {
		ZCall(zm, args, numargs, ZCallTypeN)
	},
	func(zm *ZMachine, args []uint16, numargs uint16) { // Set Color
		// ignore
	},
	func(zm *ZMachine, args []uint16, numargs uint16) { // Throw value stackframe
		panic("Call 28 illegal")
	},
	func(zm *ZMachine, args []uint16, numargs uint16) {
		panic("Call 29 illegal")
	},
	func(zm *ZMachine, args []uint16, numargs uint16) {
		panic("Call 30 illegal")
	},
	func(zm *ZMachine, args []uint16, numargs uint16) {
		panic("Call 31 illegal")
	},
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
	func(zm *ZMachine, arg uint16) {
		ZCall(zm, []uint16{arg}, 1, ZCallTypeStore)
	},
	ZRemoveObj,
	ZPrintObj,
	ZRet,
	ZJump,
	ZPrintPAddr,
	ZLoad,
	func(zm *ZMachine, arg uint16) {
		ZCall(zm, []uint16{arg}, 1, ZCallTypeN)
	},
}

var ZFunctions_0P = []ZFunction0Op{
	ZReturnTrue,
	ZReturnFalse,
	ZPrint,
	ZPrintRet,
	ZNOP0,
	ZNOP0,
	ZNOP0,
	ZRestart,
	ZRetPopped,
	ZPop,
	ZQuit,
	ZNewLine,
}
