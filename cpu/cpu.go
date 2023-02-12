
package cpu

import (
	"fmt"
	"io"
	"math"
	"os"

	. "github.com/golang-vm/govm/encoding"
)

type Callable = func([]Value) []Value
type NativeCall = func(ctx *Context) (n uint8)

type Registry interface {
	Lookup(label string) (meta LabelMeta, ok bool)
	GetNative(label string) NativeCall
}

type ExitErr struct {
	Code int
}

func (e *ExitErr) Error() string {
	return fmt.Sprintf("Exit status %d", e.Code)
}

type LookupErr struct {
	Label string
}

func (e *LookupErr) Error() string {
	return fmt.Sprintf("Label %s not fonud", e.Label)
}

type Cpu struct {
	context *Context

	R [32]uint64
	P [32]Pointer
	S [16]string

	reg Registry

	retn     uint8
	panicv   Value
	panicing bool
}

func NewCpu(context *Context, reg Registry) *Cpu {
	return &Cpu{
		context: context,
		reg:     reg,
	}
}

func (cpu *Cpu) Context() *Context {
	return cpu.context
}

func (cpu *Cpu) Registry() Registry {
	return cpu.reg
}

func (cpu *Cpu) Run() (err error) {
	for {
		if err = cpu.Tick(); err != nil {
			if e0, ok := err.(*ExitErr); ok {
				os.Exit(e0.Code)
			}
			return
		}
	}
}

func (cpu *Cpu) Tick() (err error) {
	var r io.Reader = cpu.context
	// vals := cpu.context.Vals()
	// fmt.Printf("DEBUG: Reading 0x%x\n", cpu.context.pc)
	cmd := cpu.context.GetCmd()
	// fmt.Printf("DEBUG: Exe command 0x%x\n", cmd)
	switch cmd {
	case SKIP:
		return
	case EXIT:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		return &ExitErr{(int)(cpu.R[i])}
	case MOV:
		var a, b byte
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = cpu.R[a]
	case SETB:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		var v uint8
		if v, err = ReadUint8(r); err != nil {
			return
		}
		cpu.R[i] = (uint64)(v)
	case LOADB:
	case STOREB:
	case SETW:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		var v uint16
		if v, err = ReadUint16(r); err != nil {
			return
		}
		cpu.R[i] = (uint64)(v)
	case LOADW:
	case STOREW:
	case SETI:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		var v uint32
		if v, err = ReadUint32(r); err != nil {
			return
		}
		cpu.R[i] = (uint64)(v)
	case LOADI:
	case STOREI:
	case SETQ:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		var v uint64
		if v, err = ReadUint64(r); err != nil {
			return
		}
		cpu.R[i] = v
	case LOADQ:
	case STOREQ:
	case NEW:
		typid := MustReadUint32(r)
		_ = typid
		panic("TODO")
	case LOADP:
	case STOREP:
	case SETS:
		var (
			i uint8
			p uint32
			l uint32
		)
		if i, err = ReadUint8(r); err != nil {
			return
		}
		if p, err = ReadUint32(r); err != nil {
			return
		}
		if l, err = ReadUint32(r); err != nil {
			return
		}
		if l == 0 {
			cpu.S[i] = ""
		} else {
			cpu.S[i] = cpu.context.Strlst()[p:p + l]
		}
	case LOADS:
	case STORES:
	case PUSHU:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		cpu.context.Push(uint64Item{cpu.R[i]})
	case PUSHB:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		cpu.context.Push(int8Item{(int8)((uint8)(cpu.R[i]))})
	case PUSHW:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		cpu.context.Push(int16Item{(int16)((uint16)(cpu.R[i]))})
	case PUSHI:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		cpu.context.Push(int32Item{(int32)((uint32)(cpu.R[i]))})
	case PUSHQ:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		cpu.context.Push(int64Item{(int64)((uint64)(cpu.R[i]))})
	case PUSHP:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		cpu.context.Push(pointerItem{cpu.P[i]})
	case PUSHS:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		cpu.context.Push(stringItem{cpu.S[i]})
	case POP:
		var n uint8
		if n, err = ReadUint8(r); err != nil {
			return
		}
		cpu.context.stack = cpu.context.stack[:len(cpu.context.stack) - (int)(n)]
	case POPU:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		cpu.R[i] = cpu.context.Pop().Uint64()
	case POPP:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		cpu.P[i] = cpu.context.Pop().Pointer()
	case POPS:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		cpu.S[i] = cpu.context.Pop().String()
	// Number cast
	case B2Q:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)((int8)(cpu.R[a]))
	case W2Q:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)((int16)(cpu.R[a]))
	case I2Q:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)((int32)(cpu.R[a]))
	case I2F:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)(math.Float32bits(
			(float32)((int32)(cpu.R[a]))))
	case I2D:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = math.Float64bits(
			(float64)((int32)(cpu.R[a])))
	case UI2F:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)(math.Float32bits(
			(float32)((uint32)(cpu.R[a]))))
	case UI2D:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = math.Float64bits(
			(float64)((uint32)(cpu.R[a])))
	case Q2F:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)(math.Float32bits(
			(float32)((int64)(cpu.R[a]))))
	case Q2D:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = math.Float64bits(
			(float64)((int64)(cpu.R[a])))
	case UQ2F:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)(math.Float32bits(
			(float32)(cpu.R[a])))
	case UQ2D:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = math.Float64bits(
			(float64)(cpu.R[a]))
	case F2I:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)((int32)(math.Float32frombits((uint32)(cpu.R[a]))))
	case F2UI:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)((uint32)(math.Float32frombits((uint32)(cpu.R[a]))))
	case F2Q:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)((int64)(math.Float32frombits((uint32)(cpu.R[a]))))
	case F2UQ:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)(math.Float32frombits((uint32)(cpu.R[a])))
	case F2D:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = math.Float64bits(
			(float64)(math.Float32frombits((uint32)(cpu.R[a]))))
	case D2I:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)((int32)(math.Float64frombits(cpu.R[a])))
	case D2UI:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)((uint32)(math.Float64frombits(cpu.R[a])))
	case D2Q:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)((int64)(math.Float64frombits(cpu.R[a])))
	case D2UQ:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)(math.Float64frombits(cpu.R[a]))
	case D2F:
		var a, b uint8
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)(math.Float32bits(
			(float32)(math.Float64frombits(cpu.R[a]))))

	// Number opers
	case ADD:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = cpu.R[a] + cpu.R[b]
	case SUB:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = cpu.R[a] - cpu.R[b]
	case MUL:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = cpu.R[a] * cpu.R[b]
	case SMUL:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = (uint64)((int64)(cpu.R[a]) * (int64)(cpu.R[b]))
	case QUO:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = cpu.R[a] / cpu.R[b]
	case SQUO:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = (uint64)((int64)(cpu.R[a]) / (int64)(cpu.R[b]))
	case REM:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = cpu.R[a] % cpu.R[b]
	case SREM:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = (uint64)((int64)(cpu.R[a]) % (int64)(cpu.R[b]))
	//
	case NOT:
		var a, b byte
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = ^cpu.R[a]
	case AND:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = cpu.R[a] & cpu.R[b]
	case OR:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = cpu.R[a] | cpu.R[b]
	case XOR:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = cpu.R[a] ^ cpu.R[b]
	case SHL:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = cpu.R[a] << cpu.R[b]
	case SHR:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = cpu.R[a] >> cpu.R[b]
	case SSHR:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = (uint64)((int8)(cpu.R[a]) >> cpu.R[b])
	case SSHRW:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = (uint64)((int16)(cpu.R[a]) >> cpu.R[b])
	case SSHRI:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = (uint64)((int32)(cpu.R[a]) >> cpu.R[b])
	case SSHRQ:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = (uint64)((int64)(cpu.R[a]) >> cpu.R[b])
	case INC:
		var i byte
		if i, err = ReadUint8(r); err != nil {
			return
		}
		cpu.R[i]++
	case DEC:
		var i byte
		if i, err = ReadUint8(r); err != nil {
			return
		}
		cpu.R[i]--
	// Float opers
	case FNEG:
		var a, b byte
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = (uint64)(math.Float32bits(
			-math.Float32frombits((uint32)(cpu.R[a]))))
	case FADD:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = (uint64)(math.Float32bits(
			math.Float32frombits((uint32)(cpu.R[a])) + math.Float32frombits((uint32)(cpu.R[b]))))
	case FSUB:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = (uint64)(math.Float32bits(
			math.Float32frombits((uint32)(cpu.R[a])) - math.Float32frombits((uint32)(cpu.R[b]))))
	case FMUL:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = (uint64)(math.Float32bits(
			math.Float32frombits((uint32)(cpu.R[a])) * math.Float32frombits((uint32)(cpu.R[b]))))
	case FQUO:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = (uint64)(math.Float32bits(
			math.Float32frombits((uint32)(cpu.R[a])) / math.Float32frombits((uint32)(cpu.R[b]))))
	case DNEG:
		var a, b byte
		if a, b, err = ReadByte2(r); err != nil {
			return
		}
		cpu.R[b] = math.Float64bits(
			-math.Float64frombits(cpu.R[a]))
	case DADD:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = math.Float64bits(
			math.Float64frombits(cpu.R[a]) + math.Float64frombits(cpu.R[b]))
	case DSUB:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = math.Float64bits(
			math.Float64frombits(cpu.R[a]) - math.Float64frombits(cpu.R[b]))
	case DMUL:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = math.Float64bits(
			math.Float64frombits(cpu.R[a]) * math.Float64frombits(cpu.R[b]))
	case DQUO:
		var a, b, c byte
		if a, b, c, err = ReadByte3(r); err != nil {
			return
		}
		cpu.R[c] = math.Float64bits(
			math.Float64frombits(cpu.R[a]) / math.Float64frombits(cpu.R[b]))
	// Bool opers
	case LNOT:
		var ai, bi byte
		if ai, bi, err = ReadByte2(r); err != nil {
			return
		}
		if cpu.R[ai] == 0 {
			cpu.R[bi] = 1
		} else {
			cpu.R[bi] = 0
		}
	case LAND:
		var ai, bi, ci byte
		if ai, bi, ci, err = ReadByte3(r); err != nil {
			return
		}
		if cpu.R[ai] & cpu.R[bi] == 0 {
			cpu.R[ci] = 0
		} else {
			cpu.R[ci] = 1
		}
	case LOR:
		var ai, bi, ci byte
		if ai, bi, ci, err = ReadByte3(r); err != nil {
			return
		}
		if cpu.R[ai] | cpu.R[bi] == 0 {
			cpu.R[ci] = 0
		} else {
			cpu.R[ci] = 1
		}
	case EQ:
		var ai, n, bi byte
		if ai, n, bi, err = ReadByte3(r); err != nil {
			return
		}
		if (uint8)(cpu.R[ai]) == n {
			cpu.R[bi] = 1
		} else {
			cpu.R[bi] = 0
		}
	case NE:
		var ai, n, bi byte
		if ai, n, bi, err = ReadByte3(r); err != nil {
			return
		}
		if (uint8)(cpu.R[ai]) == n {
			cpu.R[bi] = 0
		} else {
			cpu.R[bi] = 1
		}
	case GT:
		var ai, n, bi byte
		if ai, n, bi, err = ReadByte3(r); err != nil {
			return
		}
		if (int8)((uint8)(cpu.R[ai])) > (int8)(n) {
			cpu.R[bi] = 1
		} else {
			cpu.R[bi] = 0
		}
	case LT:
		var ai, n, bi byte
		if ai, n, bi, err = ReadByte3(r); err != nil {
			return
		}
		if (int8)((uint8)(cpu.R[ai])) < (int8)(n) {
			cpu.R[bi] = 1
		} else {
			cpu.R[bi] = 0
		}
	case CMP:
		var ai, bi, ci byte
		if ai, bi, ci, err = ReadByte3(r); err != nil {
			return
		}
		a, b := cpu.R[ai], cpu.R[bi]
		if a == b {
			cpu.R[ci] = 0
		} else if a > b {
			cpu.R[ci] = 1
		} else {
			cpu.R[ci] = NEG1
		}
	case ICMP:
		var ai, bi, ci byte
		if ai, bi, ci, err = ReadByte3(r); err != nil {
			return
		}
		a, b := (int64)(cpu.R[ai]), (int64)(cpu.R[bi])
		if a == b {
			cpu.R[ci] = 0
		} else if a > b {
			cpu.R[ci] = 1
		} else {
			cpu.R[ci] = NEG1
		}
	case FCMP:
		var ai, bi, ci byte
		if ai, bi, ci, err = ReadByte3(r); err != nil {
			return
		}
		a, b := math.Float32frombits((uint32)(cpu.R[ai])), math.Float32frombits((uint32)(cpu.R[bi]))
		if a == b {
			cpu.R[ci] = 0
		} else if a > b {
			cpu.R[ci] = 1
		} else {
			cpu.R[ci] = NEG1
		}
	case DCMP:
		var ai, bi, ci byte
		if ai, bi, ci, err = ReadByte3(r); err != nil {
			return
		}
		a, b := math.Float64frombits(cpu.R[ai]), math.Float64frombits(cpu.R[bi])
		if a == b {
			cpu.R[ci] = 0
		} else if a > b {
			cpu.R[ci] = 1
		} else {
			cpu.R[ci] = NEG1
		}
	case CAT:
		var ai, bi, ci byte
		if ai, bi, ci, err = ReadByte3(r); err != nil {
			return
		}
		cpu.S[ci] = cpu.S[ai] + cpu.S[bi]
	case CUT:
		var ai, bi, ci, di byte
		if ai, bi, ci, di, err = ReadByte4(r); err != nil {
			return
		}
		cpu.S[di] = cpu.S[ai][cpu.R[bi]:cpu.R[ci]]
	case GOTO:
		var step0 uint32
		if step0, err = ReadUint24(r); err != nil {
			return
		}
		cpu.context.pc += (int64)(Uint24ToInt24(step0))
		// fmt.Printf("DEBUG: Jumped to 0x%x\n", cpu.context.pc)
	case JMP:
		var step0 uint8
		if step0, err = ReadUint8(r); err != nil {
			return
		}
		cpu.context.pc += (int64)((int8)(step0))
		// fmt.Printf("DEBUG: Jumped to 0x%x\n", cpu.context.pc)
	case JMP_IF:
		var i, step0 byte
		if i, step0, err = ReadByte2(r); err != nil {
			return
		}
		if cpu.R[i] != 0 {
			cpu.context.pc += (int64)((int8)(step0))
			// fmt.Printf("DEBUG: IF Jumped to 0x%x\n", cpu.context.pc)
		}
	case JMP_IFN:
		var i, step0 byte
		if i, step0, err = ReadByte2(r); err != nil {
			return
		}
		if cpu.R[i] == 0 {
			cpu.context.pc += (int64)((int8)(step0))
			// fmt.Printf("DEBUG: IFN Jumped to 0x%x\n", cpu.context.pc)
		}
	case CALL:
		var n, s byte
		if n, s, err = ReadByte2(r); err != nil {
			return
		}
		label1 := cpu.S[s]
		var label LabelMeta
		var ok bool
		if label, ok = cpu.reg.Lookup(label1); !ok {
			return &LookupErr{label1}
		}
		cpu.context = NewContext(label, cpu.context, cpu.context.PopN((uint32)(n))...)
	case CALLN:
		var n, s byte
		if n, s, err = ReadByte2(r); err != nil {
			return
		}
		label := cpu.S[s]
		cpu.context = cpu.context.NewNativeCall(label, cpu.context.PopN((uint32)(n))...)
		call := cpu.reg.GetNative(label)
		if call == nil {
			return fmt.Errorf("Native label '%s' not found", label)
		}
		cpu.retn = call(cpu.context)
		parent := cpu.context.Parent()
		parent.Push(cpu.context.PopN((uint32)(cpu.retn))...)
		cpu.context = cpu.context.Parent()
	case CALLG:
		var n, s byte
		if n, s, err = ReadByte2(r); err != nil {
			return
		}
		label1 := cpu.S[s]
		var label LabelMeta
		var ok bool
		if label, ok = cpu.reg.Lookup(label1); !ok {
			return &LookupErr{label1}
		}
		ctx := NewContext(label, cpu.context)
		ctx.args = ctx.Parent().PopN((uint32)(n))
		cpu2 := NewCpu(ctx, cpu.reg)
		go func(){
			cpu2.Run()
		}()
	// case CALLD:
	// 	pkg := (string)(readFull(r, make([]byte, MustReadUint16(r))))
	// 	label := (string)(readFull(r, make([]byte, MustReadUint8(r))))
	// 	cpu.context.AddDefer(pkg, label)
	case ARG:
		var i uint8
		if i, err = ReadUint8(r); err != nil {
			return
		}
		cpu.context.Push(cpu.context.Arg(i))
	case RET:
		var n uint8
		if n, err = ReadUint8(r); err != nil {
			return
		}
		cpu.retn = n
		parent := cpu.context.Parent()
		parent.Push(cpu.context.PopN((uint32)(n))...)
		cpu.context = parent
	case PAN:
		panic("TODO")
		cpu.panicv = cpu.context.Pop()
		cpu.panicing = true
	case REC:
		if cpu.panicing {
			cpu.context.Push(cpu.panicv)
		} else {
			cpu.context.Push(nil)
		}
	default:
		panic(fmt.Sprintf("Unknown command 0x%x", cmd))
	}
	return
}
