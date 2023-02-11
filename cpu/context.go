
package cpu

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
	strlst string

	parent *Context
}

var _ io.Reader = (*Context)(nil)

func NewContext(label encoding.LabelMeta, p *Context) *Context {
	return &Context{
		prog: label.Pkg.Program(),
		pc: label.Offset,
		stack: GetStackFrame(),
		variables: make(map[string]Value),
		strlst: label.Pkg.Strlst(),
		parent: p,
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

func (c *Context) Strlst() string {
	return c.strlst
}

func (c *Context) Parent() *Context {
	return c.parent
}

func (c *Context) NewNativeCall(label string) (nc *Context) {
	nc = &Context{
		stack: GetStackFrame(),
		variables: make(map[string]Value),
		parent: c,
	}
	return
}
