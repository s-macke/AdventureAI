package zmachine

import "fmt"

var (
	OBJECT_PARENT_INDEX            uint32 = 4
	OBJECT_SIBLING_INDEX           uint32 = 5
	OBJECT_CHILD_INDEX             uint32 = 6
	NULL_OBJECT_INDEX              uint16 = 0
	OBJECT_ENTRY_SIZE              uint16 = 9
	OBJECT_PROPERTY_ADDRESS_OFFSET uint32 = 7
	OBJECT_PROPERTY_DEFAULTS_WORDS uint32 = 31
	MAX_ATTRIBUTES                 uint16 = 32
)

func (zm *ZMachine) InitObjectsConstants() {
	if zm.header.Version <= 3 {
		OBJECT_PARENT_INDEX = 4
		OBJECT_SIBLING_INDEX = 5
		OBJECT_CHILD_INDEX = 6
		NULL_OBJECT_INDEX = 0
		OBJECT_ENTRY_SIZE = 9
		OBJECT_PROPERTY_ADDRESS_OFFSET = 7
		OBJECT_PROPERTY_DEFAULTS_WORDS = 31
		MAX_ATTRIBUTES = 32
	} else {
		OBJECT_PARENT_INDEX = 6
		OBJECT_SIBLING_INDEX = 8
		OBJECT_CHILD_INDEX = 10
		NULL_OBJECT_INDEX = 0
		OBJECT_ENTRY_SIZE = 14
		OBJECT_PROPERTY_ADDRESS_OFFSET = 12
		OBJECT_PROPERTY_DEFAULTS_WORDS = 63
		MAX_ATTRIBUTES = 48
	}
}

func (zm *ZMachine) GetObjectEntryAddress(objectIndex uint16) uint32 {
	if objectIndex > MAX_OBJECT || objectIndex == 0 {
		fmt.Printf("Index: %d\n", objectIndex)
		panic("Invalid object index")
	}

	// Convert from 1-based (0 = NULL = no object) to 0-based

	objectIndex--
	// Skip default props
	var objectEntryAddress uint32
	objectEntryAddress = zm.header.objTableAddress + (OBJECT_PROPERTY_DEFAULTS_WORDS * 2) + uint32(objectIndex*OBJECT_ENTRY_SIZE)
	return objectEntryAddress
}

func (zm *ZMachine) SetObjectProperty(objectIndex uint16, propertyId uint16, value uint16) {
	propData, numBytes := zm.GetObjectPropertyInfo(objectIndex, propertyId)
	if propData == 0 {
		panic("Property not found")
	}

	if numBytes == 1 {
		zm.buf[propData] = uint8(value & 0xFF)
	} else if numBytes == 2 {
		zm.SetUint16(uint32(propData), value)
	} else {
		panic("SetObjectProperty only supports 1/2 byte properties")
	}
}

func (zm *ZMachine) GetFirstPropertyAddress(objectIndex uint16) uint16 {
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	propertiesAddress := zm.GetUint16(objectEntryAddress + OBJECT_PROPERTY_ADDRESS_OFFSET)

	//size := uint16(zm.buf[propertiesAddress] & 0xff)
	//return propertiesAddress + 1 + 2*size

	nameLength := uint16(zm.buf[propertiesAddress]) * 2 // in 2-byte words
	propData := propertiesAddress + 1 + nameLength
	return propData
}

// Returns prop data address, number of property bytes
// (0 if not found)
func (zm *ZMachine) GetObjectPropertyInfo(objectIndex uint16, propertyId uint16) (uint16, uint16) {
	propData := zm.GetFirstPropertyAddress(objectIndex)

	var propNo uint16
	var propSize uint16

	for {
		value := uint16(zm.buf[propData])
		// a property list is terminated by a size byte of 0
		if value == 0 {
			break
		}
		propData++
		addOne := uint16(0)

		if zm.header.Version <= 3 {
			propNo = value & 0x1F
			propSize = value >> 5
		} else {
			propNo = value & 0x3F
			if (value & 0x80) == 0 {
				propSize = value >> 6
			} else {
				propSize = uint16(zm.buf[propData])
				propSize &= 0x3f
				if propSize == 0 {
					propSize = 64 /* demanded by Spec 1.0 */
				}
				addOne = 1
			}
		}

		if propNo == propertyId {
			return propData + addOne, propSize + 1
		}
		propData += propSize + 1
	}

	return uint16(0), uint16(0)
}

func (zm *ZMachine) GetObjectPropertyAddress(objectIndex uint16, propertyId uint16) uint16 {
	address, _ := zm.GetObjectPropertyInfo(objectIndex, propertyId)
	return address
}

func (zm *ZMachine) GetNextObjectProperty(objectIndex uint16, propertyId uint16) uint16 {
	//DebugPrintf("GetNextObjectProperty(%d, %d)\n", objectIndex, propertyId)
	nextPropSize := uint8(0)

	// " if called with zero, it gives the first property number present."
	if propertyId == 0 {
		propData := zm.GetFirstPropertyAddress(objectIndex)
		nextPropSize = zm.buf[propData]
	} else {
		propData, numBytes := zm.GetObjectPropertyInfo(objectIndex, propertyId)
		if propData == 0 {
			panic("GetNextObjectProperty - non existent property")
		}
		nextPropSize = zm.buf[propData+numBytes]
	}
	// "zero, indicating the end of the property list"
	if nextPropSize == 0 {
		return 0
	} else {
		return uint16(nextPropSize & 0x1F)
	}
}

func (zm *ZMachine) GetObjectProperty(objectIndex uint16, propertyId uint16) uint16 {
	propData, numBytes := zm.GetObjectPropertyInfo(objectIndex, propertyId)
	result := uint16(0)

	if propData == 0 {
		// Get a default one
		result = zm.GetPropertyDefault(propertyId)
		//DebugPrintf("Default prop %d = 0x%X\n", propertyId, result)
	} else {
		if numBytes == 1 {
			result = uint16(zm.buf[propData])
		} else if numBytes == 2 {
			result = zm.GetUint16(uint32(propData))
		} else {
			panic("GetObjectProperty only supports 1/2 byte properties")
		}
	}

	return result
}

// True if set
func (zm *ZMachine) TestObjectAttr(objectIndex uint16, attribute uint16) bool {
	if attribute > MAX_ATTRIBUTES {
		panic("Attribute out of bounds")
	}
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)

	byteIndex := uint32(attribute >> 3)
	shift := 7 - (attribute & 0x7)

	return (zm.buf[objectEntryAddress+byteIndex] & (1 << shift)) != 0
}

func (zm *ZMachine) SetObjectAttr(objectIndex uint16, attribute uint16) {
	if attribute > MAX_ATTRIBUTES {
		panic("Attribute out of bounds")
	}
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)

	byteIndex := uint32(attribute >> 3)
	shift := 7 - (attribute & 0x7)

	zm.buf[objectEntryAddress+byteIndex] |= 1 << shift
}

func (zm *ZMachine) ClearObjectAttr(objectIndex uint16, attribute uint16) {
	if attribute > MAX_ATTRIBUTES {
		panic("Attribute out of bounds")
	}

	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	byteIndex := uint32(attribute >> 3)
	shift := 7 - (attribute & 0x7)

	zm.buf[objectEntryAddress+byteIndex] &= ^(1 << shift)
}

func (zm *ZMachine) IsDirectParent(childIndex uint16, parentIndex uint16) bool {
	return zm.GetParentObjectIndex(childIndex) == parentIndex
}

// Unlink object from its parent
func (zm *ZMachine) UnlinkObject(objectIndex uint16) {
	currentParentIndex := zm.GetParentObjectIndex(objectIndex)

	// Unlink from current parent first
	if currentParentIndex == NULL_OBJECT_INDEX {
		return
	}

	// If we're the first child -> move to sibling
	if zm.GetChildIndex(currentParentIndex) == objectIndex {
		zm.SetChildObjectIndex(currentParentIndex, zm.GetSiblingIndex(objectIndex))
	} else {
		childIter := zm.GetChildIndex(currentParentIndex)
		prevChild := NULL_OBJECT_INDEX
		for childIter != objectIndex && childIter != NULL_OBJECT_INDEX {
			prevChild = childIter
			childIter = zm.GetSiblingIndex(childIter)
		}
		// Sanity checks
		if childIter == NULL_OBJECT_INDEX {
			panic("Object not found on parent children list")
		}
		if prevChild == NULL_OBJECT_INDEX {
			panic("Corrupted data")
		}
		sibling := zm.GetSiblingIndex(objectIndex)
		zm.SetSiblingObjectIndex(prevChild, sibling)
	}
	zm.SetParentObjectIndex(objectIndex, NULL_OBJECT_INDEX)
}

func (zm *ZMachine) ReparentObject(objectIndex uint16, newParentIndex uint16) {
	currentParentIndex := zm.GetParentObjectIndex(objectIndex)

	if currentParentIndex == newParentIndex {
		return
	}

	zm.UnlinkObject(objectIndex)

	// Make the first child of our new parent
	zm.SetSiblingObjectIndex(objectIndex, zm.GetChildIndex(newParentIndex))
	zm.SetChildObjectIndex(newParentIndex, objectIndex)
	zm.SetParentObjectIndex(objectIndex, newParentIndex)
}

func (zm *ZMachine) GetParentObjectIndex(objectIndex uint16) uint16 {
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	if zm.header.Version <= 3 {
		return uint16(zm.buf[objectEntryAddress+OBJECT_PARENT_INDEX])
	} else {
		return zm.GetUint16(objectEntryAddress + OBJECT_PARENT_INDEX)
	}
}

func (zm *ZMachine) SetParentObjectIndex(objectIndex uint16, parentIndex uint16) {
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	if zm.header.Version <= 3 {
		zm.buf[objectEntryAddress+OBJECT_PARENT_INDEX] = uint8(parentIndex)
	} else {
		zm.SetUint16(objectEntryAddress+OBJECT_PARENT_INDEX, parentIndex)
	}
}

func (zm *ZMachine) GetChildIndex(objectIndex uint16) uint16 {
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	if zm.header.Version <= 3 {
		return uint16(zm.buf[objectEntryAddress+OBJECT_CHILD_INDEX])
	} else {
		return zm.GetUint16(objectEntryAddress + OBJECT_CHILD_INDEX)
	}
}

func (zm *ZMachine) SetChildObjectIndex(objectIndex uint16, childIndex uint16) {
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	if zm.header.Version <= 3 {
		zm.buf[objectEntryAddress+OBJECT_CHILD_INDEX] = uint8(childIndex)
	} else {
		zm.SetUint16(objectEntryAddress+OBJECT_CHILD_INDEX, childIndex)
	}
}

func (zm *ZMachine) GetSiblingIndex(objectIndex uint16) uint16 {
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	if zm.header.Version <= 3 {
		return uint16(zm.buf[objectEntryAddress+OBJECT_SIBLING_INDEX])
	} else {
		return zm.GetUint16(objectEntryAddress + OBJECT_SIBLING_INDEX)
	}
}

func (zm *ZMachine) SetSiblingObjectIndex(objectIndex uint16, siblingIndex uint16) {
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	if zm.header.Version <= 3 {
		zm.buf[objectEntryAddress+OBJECT_SIBLING_INDEX] = uint8(siblingIndex)
	} else {
		zm.SetUint16(objectEntryAddress+OBJECT_SIBLING_INDEX, siblingIndex)
	}
}

func (zm *ZMachine) PrintObjectName(objectIndex uint16) {
	objectEntryAddress := zm.GetObjectEntryAddress(objectIndex)
	propertiesAddress := uint32(zm.GetUint16(objectEntryAddress + OBJECT_PROPERTY_ADDRESS_OFFSET))
	zm.DecodeZString(propertiesAddress + 1)
}
