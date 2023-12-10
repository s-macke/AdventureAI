package zmachine

import "strconv"

func ZOutputStream(zm *ZMachine, args []uint16, numArgs uint16) {
	switch int16(args[0]) {
	case -3: // memory_close
		zm.outputstream = 0
		break
	case 3: // memory_open // Begin output redirection to the memory of the Z-machine.
		zm.outputstream = 3
		// output should be empty
		zm.Output.Reset()
		table := uint32(args[1])
		zm.SetUint16(table, 0)
		zm.outputstreamtable = table
		//panic("Not implemented")
		break
	default:
		panic("ZFunctions_VAR 19 case " + strconv.Itoa(int(int16(args[0]))) + " not implemented")
	}

}
