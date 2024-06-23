package zmachine

import "fmt"

func DebugPrintf(format string, v ...any) {
	//fmt.Printf(format, v...)
}

func (zm *ZMachine) ListObjects() {
	fmt.Println("Objects:")
	zm.Output.Reset()
	for i := uint16(1); i < MAX_OBJECT; i++ {
		zm.PrintObjectName(i)
		propData := zm.GetFirstPropertyAddress(i)
		fmt.Printf("%3d '%s' offset=0x%04x\n", i, zm.Output.String(), propData)
		for j := uint16(1); j < 64; j++ {
			addr, size := zm.GetObjectPropertyInfo(i, j)
			if addr != 0 {
				if size == 1 {
					fmt.Printf("    [%d] size %d bytes = 0x%2x\n", j, size, zm.GetUint8(uint32(addr)))
				} else if size == 2 {
					fmt.Printf("    [%d] size %d bytes = 0x%04x\n", j, size, zm.GetUint16(uint32(addr)))
				} else {
					fmt.Printf("    [%d] size %d bytes\n", j, size)
				}
			}
			//fmt.Println(addr, size)
			//fmt.Printf("%0x\n", zm.GetUint16(uint32(addr)))
		}

		zm.Output.Reset()
	}
	fmt.Println("Objects end")
}

// https://www.inform-fiction.org/zmachine/standards/z1point1/sect13.html
func (zm *ZMachine) ListDictionary() {
	fmt.Println("Dictionary:")
	numSeparators := uint32(zm.buf[zm.header.dictAddress])
	// followed by a list of keyboard input codes

	// The "entry length" is the length of each word's entry in the dictionary table.
	entryLength := uint16(zm.buf[zm.header.dictAddress+1+numSeparators])

	numEntries := uint32(zm.GetUint16(zm.header.dictAddress + 1 + numSeparators + 1))
	fmt.Printf("numSeparators: %d entryLength: %d numEntries: %d\n", numSeparators, entryLength, numEntries)
	//fmt.Println(numSeparators, entryLength, numEntries)

	entriesAddress := zm.header.dictAddress + 1 + numSeparators + 1 + 2
	zm.Output.Reset()
	for i := uint32(0); i < numEntries; i++ {
		foundAddress := entriesAddress + i*uint32(entryLength)
		//fmt.Println(foundAddress)
		zm.DecodeZString(foundAddress)

		fmt.Printf("%3d %s\n", i+1, zm.Output.String())
		zm.Output.Reset()
		//fmt.Println(foundAddress)
	}
	fmt.Println("Dictionary end")
}

func (zm *ZMachine) ListAbbreviations() {
	fmt.Println("Abbreviations:")
	//fmt.Println(numSeparators, entryLength, numEntries)
	zm.Output.Reset()

	entriesAddress := zm.header.abbreviationTable
	for i := uint32(0); i < 96; i++ { // 96 in V3+
		foundAddress := entriesAddress + 2*i
		//fmt.Println(foundAddress)
		//zm.DecodeZString(zm.PackedAddress(uint32(zm.GetUint16(foundAddress))))
		zm.DecodeZString(uint32(zm.GetUint16(foundAddress)) * 2)

		fmt.Printf("%3d '%s'\n", i, zm.Output.String())
		zm.Output.Reset()
		//fmt.Println(foundAddress)
	}
	fmt.Println("Abbreviations end")
}
