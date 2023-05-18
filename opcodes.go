package main

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
	nil,
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
