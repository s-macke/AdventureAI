package main

type ZStack struct {
	stack      []uint16
	top        int
	localFrame int
}

func NewStack() *ZStack {
	s := new(ZStack)
	s.stack = make([]uint16, MAX_STACK)
	s.top = MAX_STACK
	s.localFrame = s.top

	return s
}

func (s *ZStack) Push(value uint16) {
	if s.top == 0 {
		panic("Stack overflow")
	}
	s.top--
	s.stack[s.top] = value
}

func (s *ZStack) Pop() uint16 {
	if s.top == MAX_STACK {
		panic("Trying to pop from empty stack")
	}
	retValue := s.stack[s.top]

	s.top++
	return retValue
}

func (s *ZStack) Reset(newTop int) {
	if newTop > MAX_STACK || newTop < 0 {
		panic("Invalid stack top value")
	}
	s.top = newTop
}

func (s *ZStack) GetTopItem() uint16 {
	return s.stack[s.top]
}

func (s *ZStack) SaveFrame() {
	s.Push(uint16(s.localFrame))
	s.localFrame = s.top
}

// RestoreFrame returns caller address (where to return to)
func (s *ZStack) RestoreFrame() uint32 {

	// Discard local frame
	s.top = s.localFrame
	// Restore previous frame
	s.localFrame = int(s.Pop())

	retLo := s.Pop()
	retHi := s.Pop()

	return (uint32(retHi) << 16) | uint32(retLo)
}

func (s *ZStack) ValidateLocalVarIndex(localVarIndex int) {
	if localVarIndex > 0xF {
		panic("Local var index out of bounds")
	}
	if s.localFrame < localVarIndex {
		panic("Stack underflow")
	}
}
func (s *ZStack) GetLocalVar(localVarIndex int) uint16 {
	s.ValidateLocalVarIndex(localVarIndex)
	stackIndex := (s.localFrame - localVarIndex) - 1
	r := s.stack[stackIndex]
	return r
}

func (s *ZStack) SetLocalVar(localVarIndex int, value uint16) {
	s.ValidateLocalVarIndex(localVarIndex)
	stackIndex := (s.localFrame - localVarIndex) - 1
	s.stack[stackIndex] = value
}

func (s *ZStack) Dump() {
	DebugPrintf("Top = %d, local frame = %d\n", s.top, s.localFrame)

	for i := MAX_STACK - 1; i >= s.top; i-- {
		if i == s.localFrame {
			DebugPrintf("0x%X: 0x%X <------ local frame\n", i, s.stack[i])
		} else {
			DebugPrintf("0x%X: 0x%X\n", i, s.stack[i])
		}
	}
}
