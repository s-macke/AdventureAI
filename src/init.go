package mainsrc

import (
	"github.com/s-macke/AdventureAI/src/zmachine"
	"os"
	"path/filepath"
	"strconv"
)

func Init(filename string) *zmachine.ZMachine {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("Read %d bytes\n", len(buffer))

	var header zmachine.ZHeader
	header.Read(buffer)

	if header.Version != 3 && header.Version != 4 && header.Version != 5 && header.Version != 8 {
		panic("Only Version 3, 4, 5 or 8 files supported. But found version " + strconv.Itoa(int(header.Version)))
	}

	zm := zmachine.NewZMachine(filepath.Base(filename), buffer, header)
	return zm
}
