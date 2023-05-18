package main

import (
	"fmt"
	"strings"
)

// V3 only
// Returns offset pointing just after the string data
func (zm *ZMachine) DecodeZString(startOffset uint32) uint32 {

	done := false
	var zchars []uint8

	i := startOffset
	for !done {

		//--first byte-------   --second byte---
		//7    6 5 4 3 2  1 0   7 6 5  4 3 2 1 0
		//bit  --first--  --second---  --third--

		// Text in memory consists of a sequence of 2-byte words. Each word is divided into three 5-bit 'Z-characters', plus 1 bit left over, arranged as
		// --first byte-------   --second byte---
		// 7    6 5 4 3 2  1 0   7 6 5  4 3 2 1 0
		// bit  --first--  --second---  --third--
		w16 := zm.GetUint16(i)

		done = (w16 & 0x8000) != 0
		zchars = append(zchars, uint8((w16>>10)&0x1F), uint8((w16>>5)&0x1F), uint8(w16&0x1F))

		i += 2
	}

	alphabetType := 0

	for i := 0; i < len(zchars); i++ {
		zc := zchars[i]
		// z Characters 1,2,3 represent abbreviations, sometimes also called synonyms
		if zc > 0 && zc < 4 {
			//fmt.Println("Abbreviation", zc)
			// TODO, not sure if the if is correct
			var abbrevIndex uint8
			if i+1 >= len(zchars) {
				i++
				continue
			} else {
				abbrevIndex = zchars[i+1]
			}

			// "If z is the first Z-character (1, 2 or 3) and x the subsequent one,
			// then the interpreter must look up entry 32(z-1)+x in the abbreviations table"
			abbrevAddress := zm.GetUint16(zm.header.abbreviationTable + uint32(64*(zc-1)+abbrevIndex*2))
			zm.DecodeZString(zm.PackedAddress(uint32(abbrevAddress)))

			alphabetType = 0
			i++
			continue
		}
		if zc == 4 {
			alphabetType = 1
			continue
		} else if zc == 5 {
			alphabetType = 2
			continue
		}

		// Z-character 6 from A2 means that the two subsequent Z-characters specify a ten-bit ZSCII character code:
		// the next Z-character gives the top 5 bits and the one after the bottom 5.
		if alphabetType == 2 && zc == 6 {

			zc10 := (uint16(zchars[i+1]) << 5) | uint16(zchars[i+2])
			PrintZChar(&zm.output, zc10)

			i += 2

			alphabetType = 0
			continue
		}

		// z-character 0 is printed as a space
		if zc == 0 {
			_, _ = fmt.Fprintf(&zm.output, " ")
		} else {
			// If we're here zc >= 6. Alphabet tables are indexed starting at 6
			aindex := zc - 6
			_, _ = fmt.Fprintf(&zm.output, "%c", alphabets[alphabetType][aindex])
		}

		alphabetType = 0
	}

	return i
}

// NOTE: Doesn't support abbreviations.
func (zm *ZMachine) EncodeText(txt string) uint32 {

	encodedChars := make([]uint8, 12)
	encodedWords := make([]uint16, 2)
	padding := uint8(0x5)

	// Store 6 Z-chars. Clamp if longer, add padding if shorter
	i := 0
	j := 0
	for i < 6 {
		if j < len(txt) {
			c := txt[j]
			j++

			// See if we can find any alphabet
			ai := -1
			alphabetType := 0
			for a := 0; a < len(alphabets); a++ {
				index := strings.IndexByte(alphabets[a], c)
				if index >= 0 {
					ai = index
					alphabetType = a
					break
				}
			}
			if ai >= 0 {
				if alphabetType != 0 {
					// Alphabet change
					encodedChars[i] = uint8(alphabetType + 3)
					encodedChars[i+1] = uint8(ai + 6)
					i += 2
				} else {
					encodedChars[i] = uint8(ai + 6)
					i++
				}
			} else {
				// 10-bit ZC
				encodedChars[i] = 5
				encodedChars[i+1] = 6
				encodedChars[i+2] = (c >> 5)
				encodedChars[i+3] = (c & 0x1F)
				i += 4
			}
		} else {
			// Padding
			encodedChars[i] = padding
			i++
		}
	}

	for i := 0; i < 2; i++ {
		encodedWords[i] = (uint16(encodedChars[i*3+0]) << 10) | (uint16(encodedChars[i*3+1]) << 5) |
			uint16(encodedChars[i*3+2])
		if i == 1 {
			encodedWords[i] |= 0x8000
		}
	}

	return (uint32(encodedWords[0]) << 16) | uint32(encodedWords[1])
}
