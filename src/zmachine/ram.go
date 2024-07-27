package zmachine

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

func (zm *ZMachine) GetUint8(offset uint32) uint8 {
	return zm.buf[offset]
}

func (zm *ZMachine) GetUint32(offset uint32) uint32 {
	return (uint32(zm.buf[offset]) << 24) | (uint32(zm.buf[offset+1]) << 16) | (uint32(zm.buf[offset+2]) << 8) | uint32(zm.buf[offset+3])
}

func (zm *ZMachine) SetUint16(offset uint32, v uint16) {
	zm.buf[offset] = uint8(v >> 8)
	zm.buf[offset+1] = uint8(v & 0xFF)
}

func (zm *ZMachine) SetUint8(offset uint32, v uint8) {
	zm.buf[offset] = v
}

// Given a packed address P, the formula to obtain the corresponding byte address B is:
//
//		2P           Versions 1, 2 and 3
//		4P           Versions 4, 5
//	    8P			 Version 8
func (zm *ZMachine) PackedAddress(a uint32) uint32 {
	if zm.header.Version <= 3 {
		return a * 2
	}
	if zm.header.Version <= 5 {
		return a * 4
	}
	return a * 8
}
