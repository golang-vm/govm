
package main

import (
	"bytes"
	"fmt"

	cpu "github.com/golang-vm/govm/cpu"
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

	buf := bytes.NewReader([]byte{
		cpu.SETI, 0x00, 0x00, 0x00, 0x00, 0x00, // 0x06
		cpu.SETB, 0x10, 0x03, // 0x09
		cpu.SETB, 0x11, 0x01, // 0x0c
		cpu.SETB, 0x12, 0x70, // 0x0f
		cpu.REM, 0x00, 0x10, 0x02, // 0x13
		cpu.JMP_IF, 0x02, 0x0d, // 0x16
		cpu.PUSH, 0x00, // 0x18
		cpu.CALLN, 0x01,
			0x00, 0x07, 'p', 'r', 'i', 'n', 't', 'l', 'n', // 0x23
		cpu.ADD, 0x00, 0x11, 0x00, // 0x28
		cpu.CMP, 0x00, 0x12, 0x02, // 0x2c
		cpu.EQ, 0x02, 0x01, 0x02, // 0x30
		cpu.JMP_IF, 0x02, 0x02, // 0x33
		cpu.JMP, 0xdb, // 0x35
		cpu.PUSH, 0x00,
		cpu.CALLN, 0x01,
			0x00, 0x07, 'p', 'r', 'i', 'n', 't', 'l', 'n',
		cpu.SETB, 0x01, 0x04,
		cpu.EXIT, 0x1f,
	})
	c := cpu.NewCpu(cpu.NewContext(buf, 0), reg)
	err := c.Run()
	if err != nil {
		panic(err)
	}
}
