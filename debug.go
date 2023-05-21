package main

import "fmt"

func DebugPrintf(format string, v ...any) {
	fmt.Printf(format, v...)
}

func (zm *ZMachine) ListObjects() {
	zm.output.Reset()
	for i := uint16(1); i < 85; i++ {
		zm.PrintObjectName(i)
		propData := zm.GetFirstPropertyAddress(i)
		fmt.Printf("%3d %s %04x\n", i, zm.output.String(), propData)
		for j := uint16(1); j < 64; j++ {
			addr, size := zm.GetObjectPropertyInfo(i, j)
			if addr != 0 {
				fmt.Printf("    [%d] size %d bytes\n", j, size)
			}
			//fmt.Println(addr, size)
			//fmt.Printf("%0x\n", zm.GetUint16(uint32(addr)))

		}

		zm.output.Reset()
	}
}

// https://www.inform-fiction.org/zmachine/standards/z1point1/sect13.html
func (zm *ZMachine) ListDictionary() {

	numSeparators := uint32(zm.buf[zm.header.dictAddress])
	// followed by a list of keyboard input codes

	// The "entry length" is the length of each word's entry in the dictionary table.
	entryLength := uint16(zm.buf[zm.header.dictAddress+1+numSeparators])

	numEntries := uint32(zm.GetUint16(zm.header.dictAddress + 1 + numSeparators + 1))
	fmt.Printf("numSeparators: %d entryLength: %d numEntries: %d\n", numSeparators, entryLength, numEntries)
	//fmt.Println(numSeparators, entryLength, numEntries)

	entriesAddress := zm.header.dictAddress + 1 + numSeparators + 1 + 2
	zm.output.Reset()
	for i := uint32(0); i < numEntries; i++ {
		foundAddress := entriesAddress + i*uint32(entryLength)
		//fmt.Println(foundAddress)
		zm.DecodeZString(foundAddress)

		fmt.Printf("%3d %s\n", i+1, zm.output.String())
		zm.output.Reset()
		//fmt.Println(foundAddress)
	}
}

func (zm *ZMachine) ListAbbreviations() {

	//fmt.Println(numSeparators, entryLength, numEntries)
	zm.output.Reset()

	entriesAddress := zm.header.abbreviationTable
	for i := uint32(0); i < 96; i++ {
		foundAddress := entriesAddress + 2*i
		//fmt.Println(foundAddress)
		//zm.DecodeZString(zm.PackedAddress(uint32(zm.GetUint16(foundAddress))))
		zm.DecodeZString(uint32(zm.GetUint16(foundAddress)) * 2)

		fmt.Printf("%3d '%s'\n", i, zm.output.String())
		zm.output.Reset()
		//fmt.Println(foundAddress)
	}
}
