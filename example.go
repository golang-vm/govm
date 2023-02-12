
package main

import (
	"bytes"
	"fmt"
	"time"

	cpu "github.com/golang-vm/govm/cpu"
	. "github.com/golang-vm/govm/encoding"
	require "github.com/golang-vm/govm/require"
)

func main(){
	reg := require.NewRegistry()
	reg.Register("println", func(c *cpu.Context)(uint8){
		args0 := c.Args()
		args := make([]any, len(args0))
		for i, v := range args0 {
			args[i] = v.Int64()
		}
		fmt.Println(args...)
		return 0
	})
	reg.Register("sleep", func(c *cpu.Context)(uint8){
		args0 := c.Args()
		n := args0[0].Int64()
		fmt.Println("sleeping:", (time.Duration)(n) * time.Millisecond)
		time.Sleep((time.Duration)(n) * time.Millisecond)
		return 0
	})

	buf := bytes.NewBuffer(nil)
	buf.Write(Uint32ToBytes(MagicHeader, nil))
	name := "example"
	buf.Write(Uint16ToBytes((uint16)(len(name)), nil))
	buf.WriteString(name)
	strlst := "printlnsleepa@github.com/golang-vm/govm/examplemain@github.com/golang-vm/govm/example"
	buf.Write(Uint32ToBytes((uint32)(len(strlst)), nil))
	buf.WriteString(strlst)
	N := (uint32)(2)
	buf.Write(Uint32ToBytes(N, nil))
	S := (uint64)(buf.Len()) + (uint64)(16 * N)
	fmt.Printf("Offset: 0x%x\n", S)
	buf.Write(Uint32ToBytes(0x2f, nil))
	buf.Write(Uint32ToBytes(0x04, nil))
	buf.Write(Uint64ToBytes(S + 0x00, nil))
	buf.Write(Uint32ToBytes(0x0c, nil))
	buf.Write(Uint32ToBytes(0x01, nil))
	buf.Write(Uint64ToBytes(S + 0x2a, nil))
	buf.Write([]byte{ // 0x00
		// label: main@github.com/golang-vm/govm/example
		cpu.SETS, 0x0a,
			0x00, 0x00, 0x00, 0x0c,
			0x00, 0x00, 0x00, 0x23, // 0x0a
		cpu.SETS, 0x0b,
			0x00, 0x00, 0x00, 0x07,
			0x00, 0x00, 0x00, 0x05, // 0x14
		cpu.SETW, 0x1c, 0x00, 0x00,  // 0x18
		cpu.SETW, 0x1d, 0x03, 0x00,  // 0x1c
		cpu.PUSHU, 0x1c, // 0x1e
		cpu.CALL, 0x01, 0x0a, // 0x21
		cpu.PUSHU, 0x1d, // 0x23
		cpu.CALLN, 0x01, 0x0b, // 0x26
		cpu.INC, 0x1c, // 0x28
		cpu.JMP, 0xf1, // 0x2a jmp to 0x1c
		// label: a@github.com/golang-vm/govm/example
		cpu.ARG, 0x00, // 0x2c
		cpu.POPU, 0x00,
		cpu.SETS, 0x0c,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x07, // 0x71
		// cpu.SETI, 0x00, 0x00, 0x00, 0x00, 0x00, // 0x77
		cpu.SETB, 0x10, 0x03, // 0x7a
		cpu.SETB, 0x11, 0x02, // 0x7d
		cpu.SETB, 0x12, 0x70, // 0x80
		cpu.REM, 0x00, 0x10, 0x02, // 0x84
		cpu.JMP_IF, 0x02, 0x05, // 0x87 jmp to 0x8c
		cpu.PUSHU, 0x00, // 0x89
		cpu.CALLN, 0x01, 0x0c, // 0x8c
		cpu.ADD, 0x00, 0x11, 0x00, // 0x90
		cpu.CMP, 0x00, 0x12, 0x02, // 0x94
		cpu.EQ, 0x02, 0x01, 0x02, // 0x98
		cpu.JMP_IF, 0x02, 0x02, // 0x9b
		cpu.JMP, 0xe3, // 0x9d jmp to 0x80
		cpu.SETB, 0x01, 0x04,
		cpu.RET, 0x00,
	})
	/*
		func main(){
			for i := 0; ; i++ {
				a(i)
				sleep(0x300) // ms
			}
		}
		func a(n uint16){
			for {
				if n % 3 == 0 {
					println(n)
				}
				if n > 0x70 {
					break
				}
			}
		}
	*/
	gob := NewGob("github.com/golang-vm/govm/example", bytes.NewReader(buf.Bytes()))
	if err := gob.ParseMetadata(); err != nil {
		panic(err)
	}
	reg.AddPkg(gob)
	label, ok := gob.Lookup("main")
	if !ok {
		panic("Cannot lookup label 'main'")
	}
	c := cpu.NewCpu(cpu.NewContext(label, nil), reg)
	if err := c.Run(); err != nil {
		panic(err)
	}
}
