package main

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

func (h *ZHeader) read(buf []byte) {
	h.version = buf[0]
	h.hiMemBase = GetUint16(buf, 4)
	h.ip = GetUint16(buf, 6)
	h.dictAddress = uint32(GetUint16(buf, 0x8))
	h.objTableAddress = uint32(GetUint16(buf, 0xA))
	h.globalVarAddress = uint32(GetUint16(buf, 0xC))
	h.staticMemAddress = uint32(GetUint16(buf, 0xE))
	h.abbreviationTable = uint32(GetUint16(buf, 0x18))

	DebugPrintf("Version: %d\n", h.version)
	DebugPrintf("Hi mem base: 0x%X\n", h.hiMemBase)
	DebugPrintf("IP: 0x%X\n", h.ip)
	DebugPrintf("Dict: 0x%X\n", h.dictAddress)
	DebugPrintf("Obj table: 0x%X\n", h.objTableAddress)
	DebugPrintf("End of dyn mem: 0x%X\n", h.staticMemAddress)
	DebugPrintf("Global vars: 0x%X\n", h.globalVarAddress)
	DebugPrintf("Abbrev table: 0x%X\n", h.abbreviationTable)
}
