
package cpu

import (
	"io"

	encoding "github.com/golang-vm/govm/encoding"
)

type Context struct {
	prog io.ReaderAt
	pc int64

	vals  []Value
	stack []Value
	args  []Value
	variables map[string]Value
	strlst string

	parent *Context
}

var _ io.Reader = (*Context)(nil)

func NewContext(label encoding.LabelMeta, p *Context, args ...Value) *Context {
	return &Context{
		prog: label.Pkg.Program(),
		pc: label.Offset,
		args: args,
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

func (c *Context) Vals() []Value {
	return c.vals
}

func (c *Context) Val(i int32) *Value {
	return &c.vals[i]
}

func (c *Context) SetVal(i int32, v Value) {
	c.vals[i] = v
}

func (c *Context) GrowVals(n uint32) {
	if l := (uint32)(len(c.vals)); l < n {
		c.vals = append(c.vals, make([]Value, l - n)...)
	}
}

func (c *Context) Push(v ...Value) {
	c.stack = append(c.stack, v...)
}

func (c *Context) Pop() (v Value) {
	l := len(c.stack) - 1
	v, c.stack = c.stack[l], c.stack[:l]
	return
}

func (c *Context) PopN(n uint32) (v []Value) {
	return c.PopArr(make([]Value, n))
}

func (c *Context) PopArr(v []Value) []Value {
	n := len(c.stack) - len(v)
	if n < 0 {
		panic("Stack index out of bounds")
	}
	copy(v, c.stack[n:])
	c.stack = c.stack[:n]
	return v
}

func (c *Context) Args() []Value {
	return c.args
}

func (c *Context) Arg(i uint8) Value {
	return c.args[i]
}

func (c *Context) Strlst() string {
	return c.strlst
}

func (c *Context) Parent() *Context {
	return c.parent
}

func (c *Context) NewNativeCall(label string, args ...Value) (nc *Context) {
	nc = &Context{
		args: args,
		variables: make(map[string]Value),
		parent: c,
	}
	return
}
