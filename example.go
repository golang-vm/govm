
package main

import (
	"bytes"
	"fmt"

	cpu "github.com/golang-vm/govm/cpu"
	. "github.com/golang-vm/govm/encoding"
	require "github.com/golang-vm/govm/require"
)

func main(){
	reg := require.NewRegistry()
	reg.Register("println", func(c *cpu.Context)(err error){
		args0 := c.Args()
		args := make([]any, len(args0))
		for i, v := range args0 {
			args[i] = v.Int64()
		}
		fmt.Println(args...)
		return
	})

	buf := bytes.NewBuffer(nil)
	buf.Write(Uint32ToBytes(MagicHeader, nil))
	name := "example"
	buf.Write(Uint16ToBytes((uint16)(len(name)), nil))
	buf.WriteString(name)
	strlst := "mainprintln"
	buf.Write(Uint32ToBytes((uint32)(len(strlst)), nil))
	buf.WriteString(strlst)
	buf.Write(Uint32ToBytes(1, nil))
	buf.Write(Uint32ToBytes(0, nil))
	buf.Write(Uint32ToBytes(4, nil))
	buf.Write(Uint64ToBytes((uint64)(buf.Len()) + 8, nil))
	buf.Write([]byte{
		cpu.SETS, 0x00,
			0x00, 0x00, 0x00, 0x04,
			0x00, 0x00, 0x00, 0x07, // 0x0a
		cpu.SETI, 0x00, 0x00, 0x00, 0x00, 0x00, // 0x10
		cpu.SETB, 0x10, 0x03, // 0x13
		cpu.SETB, 0x11, 0x01, // 0x16
		cpu.SETB, 0x12, 0x70, // 0x19
		cpu.REM, 0x00, 0x10, 0x02, // 0x1d
		cpu.JMP_IF, 0x02, 0x05, // 0x20 jmp to 0x25
		cpu.PUSH, 0x00, // 0x22
		cpu.CALLN, 0x01, 0x00, // 0x25
		cpu.ADD, 0x00, 0x11, 0x00, // 0x29
		cpu.CMP, 0x00, 0x12, 0x02, // 0x2d
		cpu.EQ, 0x02, 0x01, 0x02, // 0x31
		cpu.JMP_IF, 0x02, 0x02, // 0x34
		cpu.JMP, 0xe3, // 0x36 jmp to 0x19
		cpu.PUSH, 0x00,
		cpu.CALLN, 0x01, 0x00,
		cpu.SETB, 0x01, 0x04,
		cpu.EXIT, 0x1f,
	})
	gob := NewGob("github.com/golang-vm/govm/example", bytes.NewReader(buf.Bytes()))
	if err := gob.ParseMetadata(); err != nil {
		panic(err)
	}
	label, ok := gob.Lookup("main")
	if !ok {
		panic("Cannot lookup label 'main'")
	}
	c := cpu.NewCpu(cpu.NewContext(label, nil), reg)
	if err := c.Run(); err != nil {
		panic(err)
	}
}
