package vm_cpu

import (
	"io"

	encoding "github.com/golang-vm/govm/encoding"
)

type Context struct {
	prog io.ReaderAt
	pc int64

	stack *StackFrame
	args []Value
	variables map[string]Value

	parent *Context
}

var _ io.Reader = (*Context)(nil)

func NewContext(prog io.ReaderAt, off int64) *Context {
	return &Context{
		prog: prog,
		pc: off,
		stack: GetStackFrame(),
		variables: make(map[string]Value),
	}
}

func (c *Context) Read(buf []byte) (n int, err error) {
	if c.prog == nil {
		panic("Cannot read program when calling native method")
	}
	n, err = c.prog.ReadAt(buf, c.pc)
	if err != nil {
		return
	}
	c.pc += (int64)(n)
	return
}

func (c *Context) IsNative() bool {
	return c.prog == nil
}

func (c *Context) GetCmd() (cmd Cmd) {
	var buf [1]byte
	_, err := c.Read(buf[:])
	if err != nil {
		panic(err)
	}
	cmd = (Cmd)(buf[0])
	return
}

func (c *Context) Stack() *StackFrame {
	return c.stack
}

func (c *Context) Args() []Value {
	return c.args
}

func (c *Context) Parent() (*Context) {
	return c.parent
}

func (c *Context) NewCall(label encoding.LabelMeta) (nc *Context) {
	nc = &Context{
		prog: label.Program,
		pc: label.Offset,
		stack: GetStackFrame(),
		variables: make(map[string]Value),
		parent: c,
	}
	return
}

func (c *Context) NewNativeCall(label string) (nc *Context) {
	nc = &Context{
		stack: GetStackFrame(),
		variables: make(map[string]Value),
		parent: c,
	}
	return
}
