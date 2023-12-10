package zmachine

// We can only write to dynamic memory
func (zm *ZMachine) IsSafeToWrite(address uint32) bool {
	return address < zm.header.staticMemAddress
}

func ZLoadB(zm *ZMachine, args []uint16, numArgs uint16) {
	address := args[0] + args[1]
	value := zm.buf[address]

	zm.StoreResult(uint16(value))
}

func ZStoreB(zm *ZMachine, args []uint16, numArgs uint16) {
	address := uint32(args[0] + args[1])
	if !zm.IsSafeToWrite(address) {
		panic("Access violation")
	}

	zm.buf[address] = uint8(args[2])
}

// array word-index -> (result)
func ZLoadW(zm *ZMachine, args []uint16, numArgs uint16) {
	address := uint32(args[0] + args[1]*2)
	value := zm.GetUint16(address)

	zm.StoreResult(value)
}

// storew array word-index value
func ZStoreW(zm *ZMachine, args []uint16, numArgs uint16) {

	address := uint32(args[0] + args[1]*2)
	if !zm.IsSafeToWrite(address) {
		panic("Access violation")
	}

	zm.SetUint16(address, args[2])
}
