package zmachine

// https://www.inform-fiction.org/zmachine/standards/z1point1/sect11.html
type ZHeader struct {
	Version           uint8
	flags             uint8
	flags2            uint8
	hiMemBase         uint16
	ip                uint16
	dictAddress       uint32
	objTableAddress   uint32
	globalVarAddress  uint32
	staticMemAddress  uint32
	abbreviationTable uint32
}

func GetUint16(buf []byte, offset uint32) uint16 {
	return (uint16(buf[offset]) << 8) | (uint16)(buf[offset+1])
}

func (h *ZHeader) Read(buf []byte) {
	h.Version = buf[0]
	h.flags = buf[1]
	h.hiMemBase = GetUint16(buf, 4)
	h.ip = GetUint16(buf, 6) // Initial value of program counter
	h.dictAddress = uint32(GetUint16(buf, 0x8))
	h.objTableAddress = uint32(GetUint16(buf, 0xA))
	h.globalVarAddress = uint32(GetUint16(buf, 0xC))
	h.staticMemAddress = uint32(GetUint16(buf, 0xE))
	h.flags2 = buf[0x10]
	h.abbreviationTable = uint32(GetUint16(buf, 0x18))

	DebugPrintf("Version: %d\n", h.Version)
	DebugPrintf("Flags: %08b\n", h.flags)
	DebugPrintf("Hi mem base: 0x%X\n", h.hiMemBase)
	DebugPrintf("IP: 0x%X\n", h.ip)
	DebugPrintf("Dict: 0x%X\n", h.dictAddress)
	DebugPrintf("Obj table: 0x%X\n", h.objTableAddress)
	DebugPrintf("End of dyn mem: 0x%X\n", h.staticMemAddress)
	DebugPrintf("Global vars: 0x%X\n", h.globalVarAddress)
	DebugPrintf("Flags2: %08b\n", h.flags2)
	DebugPrintf("Abbrev table: 0x%X\n", h.abbreviationTable)
}
