package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func Init(filename string) *ZMachine {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d bytes\n", len(buffer))

	var header ZHeader
	header.read(buffer)

	if header.version != 3 && header.version != 5 {
		panic("Only Version 3 and 5 files supported. But found version " + strconv.Itoa(int(header.version)))
	}

	zm := NewZMachine(buffer, header)
	return zm
}

func Input() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func main() {
	filename := flag.String("file", "zork1.dat", "Z-Machine file to run")
	flag.Parse()

	zm := Init(*filename)
	//chat(zm)
	//return

	zm.input = Input
	for !zm.done {
		zm.InterpretInstruction()
		if zm.output.Len() > 0 {
			_, _ = os.Stdout.WriteString(zm.output.String())
			zm.output.Reset()
		}
	}

}
