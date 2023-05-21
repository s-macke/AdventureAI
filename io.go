package main

import (
	"bytes"
	"strings"
)

func ZRead(zm *ZMachine, args []uint16, numArgs uint16) {
	textAddress := args[0]
	maxChars := uint16(zm.buf[textAddress])
	if maxChars == 0 {
		panic("Invalid max chars")
	}
	maxChars--

	//reader := bufio.NewReader(os.Stdin)
	//input, _ := reader.ReadString('\n')
	input := zm.input()
	DebugPrintf("Input: %s\n", input)

	input = strings.ToLower(input)
	input = strings.Trim(input, "\r\n")
	input = strings.Trim(input, "\n")
	input = strings.TrimSpace(input)
	//fmt.Println(args, maxChars)
	//fmt.Println(input)

	// copy text and terminate with 0
	copy(zm.buf[textAddress+1:textAddress+maxChars], input)
	zm.buf[textAddress+uint16(len(input))+1] = 0

	var words []string
	var wordStarts []uint16
	var stringBuffer bytes.Buffer
	prevWordStart := uint16(1)
	for i := uint16(1); zm.buf[textAddress+i] != 0; i++ {
		ch := zm.buf[textAddress+i]
		if ch == ' ' {
			if prevWordStart < 0xFFFF {
				words = append(words, stringBuffer.String())
				wordStarts = append(wordStarts, prevWordStart)
				stringBuffer.Truncate(0)
			}
			prevWordStart = 0xFFFF
		} else {
			stringBuffer.WriteByte(ch)
			if prevWordStart == 0xFFFF {
				prevWordStart = i
			}
		}
	}
	// Last word
	if prevWordStart < 0xFFFF {
		words = append(words, stringBuffer.String())
		wordStarts = append(wordStarts, prevWordStart)
	}
	// TODO: include other separators, not only spaces

	parseAddress := uint32(args[1])
	maxTokens := zm.buf[parseAddress]
	//DebugPrintf("Max tokens: %d\n", maxTokens)
	parseAddress++
	numTokens := uint8(len(words))
	if numTokens > maxTokens {
		numTokens = maxTokens
	}
	zm.buf[parseAddress] = numTokens
	parseAddress++

	// "Each block consists of the byte address of the word in the dictionary, if it is in the dictionary, or 0 if it isn't;
	// followed by a byte giving the number of letters in the word; and finally a byte giving the position in the text-buffer
	// of the first letter of the word.
	for i, w := range words {

		if uint8(i) >= maxTokens {
			break
		}

		DebugPrintf("w = %s, %d\n", w, wordStarts[i])
		dictionaryAddress := zm.FindInDictionary(w)
		DebugPrintf("Dictionary address: 0x%X\n", dictionaryAddress)

		zm.SetUint16(parseAddress, dictionaryAddress)
		zm.buf[parseAddress+2] = uint8(len(w))
		zm.buf[parseAddress+3] = uint8(wordStarts[i])
		parseAddress += 4
	}

	/* Store key */
	if zm.header.version >= 5 {
		c := unicode_to_zscii(13)
		if c == 0 {
			c = '?'
		}
		zm.StoreResult(113)
	}
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

	zm.output.Reset()
	for i := uint16(0); i < numEntries; i++ {
		foundAddress := entriesAddress + uint32(i)*uint32(entryLength)
		zm.DecodeZString(foundAddress)
		if zm.output.String() == str {
			zm.output.Reset()
			return uint16(foundAddress)
		}
		zm.output.Reset()
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
 * unicode_to_zscii
 *
 * Convert a Unicode character to ZSCII, returning 0 on failure.
 *
 */
func unicode_to_zscii(c uint8) uint8 {
	return c
	/*
		  var i int

		if c >= ZC_LATIN1_MIN {
		// game has its own Unicode table
		if (z_header.x_unicode_table != 0) {
		zbyte N
		int i

		LOW_BYTE(z_header.x_unicode_table, N)
		for (i = 0x9b; i < 0x9b + N; i++) {
		zword addr =
		z_header.x_unicode_table + 1 + 2 * (i - 0x9b);
		zword unicode;

		LOW_WORD(addr, unicode)
		if (c == unicode)
		return (zbyte) i;
		}
		return 0;
		} else {	// game uses standard set
		for (i = 0x9b; i <= 0xdf; i++) {
		if (c == zscii_to_latin1[i - 0x9b])
		return (zbyte) i;
		}
		return 0;
		}
		}

		return (zbyte) c;
	*/
}
