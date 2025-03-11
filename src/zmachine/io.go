package zmachine

import (
	"strings"
)

func CallInput(zm *ZMachine) string {
	key := 13 // CR .Terminating key or special key such as mouse click or timeout
	input := zm.Input()

	if key == 0x7f {
		return ""
	}

	input = strings.ToLower(input)
	input = strings.Trim(input, "\r\n")
	input = strings.Trim(input, "\n")
	input = strings.TrimSpace(input)

	DebugPrintf("Key: %d\n", key)
	DebugPrintf("Input: %s\n", input)
	return input
}

func ZRead(zm *ZMachine, args []uint16, numArgs uint16) {
	const INPUT_BUFFER_SIZE = 200
	textAddress := args[0]
	maxChars := uint16(zm.buf[textAddress])
	if maxChars == 0 {
		panic("Invalid max chars")
	}
	if zm.header.Version <= 4 {
		maxChars--
	} else {
		if maxChars >= INPUT_BUFFER_SIZE {
			maxChars = INPUT_BUFFER_SIZE - 1
		}
	}
	DebugPrintf("Max chars: %d\n", maxChars)

	// Get initial input size
	size := uint16(0)
	if zm.header.Version >= 5 {
		size = uint16(zm.buf[textAddress+1])
	}
	DebugPrintf("Size: %d\n", size)

	input := CallInput(zm)
	if HookCommand(zm, input) { // redo command
		zm.ip = zm.opcodeip
		return
	}

	// copy text and terminate with 0
	//copy(zm.buf[textAddress+1:textAddress+maxChars], input)
	//zm.buf[textAddress+uint16(len(input))+1] = 0
	if len(input) >= int(maxChars) {
		input = input[:maxChars-1]
	}
	if zm.header.Version >= 5 {
		for i := 0; i < len(input); i++ {
			zm.buf[textAddress+2+uint16(i)] = unicode_to_zscii(input[i])
		}
		zm.buf[textAddress+1] = uint8(len(input))
	} else {
		for i := 0; i < len(input); i++ {
			zm.buf[textAddress+1+uint16(i)] = unicode_to_zscii(input[i])
		}
		zm.buf[textAddress+1+uint16(len(input))] = 0
	}

	TokenizeLine(zm, args[0], args[1], args[2], args[3] != 0)

	/* Store key */
	if zm.header.Version >= 5 {
		c := unicode_to_zscii(13)
		if c == 0 {
			c = '?'
		}
		zm.StoreResult(uint16(c))
	}
}

func (zm *ZMachine) CompareWord(str1 string, str2 string) bool {
	/*
		resolution := 4
		if zm.header.Version > 3 {
			resolution = 6
		}

		resolution = min(len(str2), resolution)
		fmt.Println(str1, str2)
		return str1 == str2[0:resolution]
	*/
	if str1 == str2 {
		return true
	}
	resolution := min(len(str2), len(str1))
	if resolution < 6 {
		return false
	}
	return str1[0:resolution] == str2[0:resolution]

}

// Return DICT_NOT_FOUND (= 0) if not found
// Address in dictionary otherwise
func (zm *ZMachine) FindInDictionary(str string) uint16 {
	numSeparators := uint32(zm.buf[zm.header.dictAddress])
	entryLength := uint16(zm.buf[zm.header.dictAddress+1+numSeparators])
	numEntries := zm.GetUint16(zm.header.dictAddress + 1 + numSeparators + 1)

	entriesAddress := zm.header.dictAddress + 1 + numSeparators + 1 + 2

	// Dictionary entries are sorted, so we can use binary search
	//lowerBound := uint16(0)
	//upperBound := numEntries - 1

	//ncodedText := zm.EncodeText(str)

	zm.Output.Reset()
	for i := uint16(0); i < numEntries; i++ {
		foundAddress := entriesAddress + uint32(i)*uint32(entryLength)
		zm.DecodeZString(foundAddress)

		if zm.CompareWord(zm.Output.String(), str) {
			zm.Output.Reset()
			return uint16(foundAddress)
		}
		zm.Output.Reset()
	}
	return uint16(DICT_NOT_FOUND)
	/*
		for lowerBound <= upperBound {

			currentIndex := lowerBound + (upperBound-lowerBound)/2
			// TODO Probably wrong for V5
			dictValue := zm.GetUint32(entriesAddress + uint32(currentIndex*entryLength))

			if encodedText < dictValue {
				upperBound = currentIndex - 1
			} else if encodedText > dictValue {
				lowerBound = currentIndex + 1
			} else {
				foundAddress = uint16(entriesAddress + uint32(currentIndex*entryLength))
				break
			}
		}
	*/
	//return foundAddress
}

/*
 * According to Matteo De Luigi <matteo.de.luigi@libero.it>,
 * 0xab and 0xbb were in each other's proper positions.
 *   Sat Apr 21, 2001
 */
var zsciiToLatin1 = [...]uint16{
	0x0e4, 0x0f6, 0x0fc, 0x0c4, 0x0d6, 0x0dc, 0x0df, 0x0bb,
	0x0ab, 0x0eb, 0x0ef, 0x0ff, 0x0cb, 0x0cf, 0x0e1, 0x0e9,
	0x0ed, 0x0f3, 0x0fa, 0x0fd, 0x0c1, 0x0c9, 0x0cd, 0x0d3,
	0x0da, 0x0dd, 0x0e0, 0x0e8, 0x0ec, 0x0f2, 0x0f9, 0x0c0,
	0x0c8, 0x0cc, 0x0d2, 0x0d9, 0x0e2, 0x0ea, 0x0ee, 0x0f4,
	0x0fb, 0x0c2, 0x0ca, 0x0ce, 0x0d4, 0x0db, 0x0e5, 0x0c5,
	0x0f8, 0x0d8, 0x0e3, 0x0f1, 0x0f5, 0x0c3, 0x0d1, 0x0d5,
	0x0e6, 0x0c6, 0x0e7, 0x0c7, 0x0fe, 0x0f0, 0x0de, 0x0d0,
	0x0a3, 0x0bd, 0x0bc, 0x0a1, 0x0bf,
}

// unicode_to_zscii is a function that converts a Unicode character to ZSCII, returning 0 on failure.
func unicode_to_zscii(c uint8) uint8 {
	return c
}
