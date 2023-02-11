package vm_cpu

import (
	"unsafe"
)

type Pointer = unsafe.Pointer

type StructField struct{
	Name string
	Size uint32
	Tags []string
}

type StructType struct {
	Fields []StructField
	Methods []string
}

type Value interface {
	Bool() bool
	Int8() int8
	Int16() int16
	Int32() int32
	Int64() int64
	Uint8() uint8
	Uint16() uint16
	Uint32() uint32
	Uint64() uint64
	Float32() float32
	Float64() float64
	Pointer() Pointer
	String() string

	Len() int64
	Cap() int64
	Index(n int64) Value
	SetIndex(n int64, val Value)
	Keys() []Value
	Values() []Value
	Items() [][2]Value
	Key(k Value) Value
	HasKey(k Value) bool
	SetKey(k Value, val Value)
}

type (
	nilItem     struct{}
	estructItem struct{}
	boolItem    struct{ v bool }
	int8Item    struct{ v int8 }
	int16Item   struct{ v int16 }
	int32Item   struct{ v int32 }
	int64Item   struct{ v int64 }
	uint8Item   struct{ v uint8 }
	uint16Item  struct{ v uint16 }
	uint32Item  struct{ v uint32 }
	uint64Item  struct{ v uint64 }
	float32Item struct{ v float32 }
	float64Item struct{ v float64 }
	pointerItem struct{ v Pointer }
	stringItem  struct{ v string }
	chanItem    struct{ v chan Value }
	echanItem   struct{ v chan struct{} }

	sliceItem   struct{ v []Value }
	esliceItem  struct{ n int }
	mapItem   struct{ v map[Value]Value }
)

var Nil Value = nilItem{}

func (v nilItem) Bool() bool       { panic("nil is not a bool") }
func (v nilItem) Int8() int8       { panic("nil is not a int8") }
func (v nilItem) Int16() int16     { panic("nil is not a int16") }
func (v nilItem) Int32() int32     { panic("nil is not a int32") }
func (v nilItem) Int64() int64     { panic("nil is not a int64") }
func (v nilItem) Uint8() uint8     { panic("nil is not a uint8") }
func (v nilItem) Uint16() uint16   { panic("nil is not a uint16") }
func (v nilItem) Uint32() uint32   { panic("nil is not a uint32") }
func (v nilItem) Uint64() uint64   { panic("nil is not a uint64") }
func (v nilItem) Float32() float32 { panic("nil is not a float32") }
func (v nilItem) Float64() float64 { panic("nil is not a float64") }
func (v nilItem) Pointer() Pointer { panic("nil is not a pointer") }
func (v nilItem) String() string   { panic("nil is not a string") }

func (v nilItem) Keys() []Value               { panic("Unsupported operation") }
func (v nilItem) Values() []Value             { panic("Unsupported operation") }
func (v nilItem) Items() [][2]Value           { panic("Unsupported operation") }
func (v nilItem) Len() int64                  { panic("Unsupported operation") }
func (v nilItem) Cap() int64                  { panic("Unsupported operation") }
func (v nilItem) Index(n int64) Value         { panic("Unsupported operation") }
func (v nilItem) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v nilItem) Key(k Value) Value           { panic("Unsupported operation") }
func (v nilItem) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v nilItem) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v estructItem) Bool() bool       { panic("struct is not a bool") }
func (v estructItem) Int8() int8       { panic("struct is not a int8") }
func (v estructItem) Int16() int16     { panic("struct is not a int16") }
func (v estructItem) Int32() int32     { panic("struct is not a int32") }
func (v estructItem) Int64() int64     { panic("struct is not a int64") }
func (v estructItem) Uint8() uint8     { panic("struct is not a uint8") }
func (v estructItem) Uint16() uint16   { panic("struct is not a uint16") }
func (v estructItem) Uint32() uint32   { panic("struct is not a uint32") }
func (v estructItem) Uint64() uint64   { panic("struct is not a uint64") }
func (v estructItem) Float32() float32 { panic("struct is not a float32") }
func (v estructItem) Float64() float64 { panic("struct is not a float64") }
func (v estructItem) Pointer() Pointer { panic("struct is not a pointer") }
func (v estructItem) String() string   { panic("struct is not a string") }

func (v estructItem) Keys() []Value               { panic("Unsupported operation") }
func (v estructItem) Values() []Value             { panic("Unsupported operation") }
func (v estructItem) Items() [][2]Value           { panic("Unsupported operation") }
func (v estructItem) Len() int64                  { panic("Unsupported operation") }
func (v estructItem) Cap() int64                  { panic("Unsupported operation") }
func (v estructItem) Index(n int64) Value         { panic("Unsupported operation") }
func (v estructItem) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v estructItem) Key(k Value) Value           { panic("Unsupported operation") }
func (v estructItem) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v estructItem) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v boolItem) Bool() bool { return v.v }
func (v boolItem) Int8() int8 {
	if v.v {
		return 1
	} else {
		return 0
	}
}
func (v boolItem) Int16() int16 {
	if v.v {
		return 1
	} else {
		return 0
	}
}
func (v boolItem) Int32() int32 {
	if v.v {
		return 1
	} else {
		return 0
	}
}
func (v boolItem) Int64() int64 {
	if v.v {
		return 1
	} else {
		return 0
	}
}
func (v boolItem) Uint8() uint8 {
	if v.v {
		return 1
	} else {
		return 0
	}
}
func (v boolItem) Uint16() uint16 {
	if v.v {
		return 1
	} else {
		return 0
	}
}
func (v boolItem) Uint32() uint32 {
	if v.v {
		return 1
	} else {
		return 0
	}
}
func (v boolItem) Uint64() uint64 {
	if v.v {
		return 1
	} else {
		return 0
	}
}
func (v boolItem) Float32() float32 { panic("bool is not a float32") }
func (v boolItem) Float64() float64 { panic("bool is not a float64") }
func (v boolItem) Pointer() Pointer { panic("not a pointer") }
func (v boolItem) String() string   { panic("not a string") }

func (v boolItem) Keys() []Value               { panic("Unsupported operation") }
func (v boolItem) Values() []Value             { panic("Unsupported operation") }
func (v boolItem) Items() [][2]Value           { panic("Unsupported operation") }
func (v boolItem) Len() int64                  { panic("Unsupported operation") }
func (v boolItem) Cap() int64                  { panic("Unsupported operation") }
func (v boolItem) Index(n int64) Value         { panic("Unsupported operation") }
func (v boolItem) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v boolItem) Key(k Value) Value           { panic("Unsupported operation") }
func (v boolItem) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v boolItem) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v int8Item) Bool() bool {
	if v.v == 0 {
		return false
	} else {
		return true
	}
}
func (v int8Item) Int8() int8       { return v.v }
func (v int8Item) Int16() int16     { return (int16)(v.v) }
func (v int8Item) Int32() int32     { return (int32)(v.v) }
func (v int8Item) Int64() int64     { return (int64)(v.v) }
func (v int8Item) Uint8() uint8     { return (uint8)(v.v) }
func (v int8Item) Uint16() uint16   { return (uint16)(v.v) }
func (v int8Item) Uint32() uint32   { return (uint32)(v.v) }
func (v int8Item) Uint64() uint64   { return (uint64)(v.v) }
func (v int8Item) Float32() float32 { return (float32)(v.v) }
func (v int8Item) Float64() float64 { return (float64)(v.v) }
func (v int8Item) Pointer() Pointer { panic("Cannot cast to a pointer") }
func (v int8Item) String() string   { panic("Cannot cast to a string") }

func (v int8Item) Keys() []Value               { panic("Unsupported operation") }
func (v int8Item) Values() []Value             { panic("Unsupported operation") }
func (v int8Item) Items() [][2]Value           { panic("Unsupported operation") }
func (v int8Item) Len() int64                  { panic("Unsupported operation") }
func (v int8Item) Cap() int64                  { panic("Unsupported operation") }
func (v int8Item) Index(n int64) Value         { panic("Unsupported operation") }
func (v int8Item) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v int8Item) Key(k Value) Value           { panic("Unsupported operation") }
func (v int8Item) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v int8Item) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v int16Item) Bool() bool {
	if v.v == 0 {
		return false
	} else {
		return true
	}
}
func (v int16Item) Int8() int8       { return (int8)(v.v) }
func (v int16Item) Int16() int16     { return v.v }
func (v int16Item) Int32() int32     { return (int32)(v.v) }
func (v int16Item) Int64() int64     { return (int64)(v.v) }
func (v int16Item) Uint8() uint8     { return (uint8)(v.v) }
func (v int16Item) Uint16() uint16   { return (uint16)(v.v) }
func (v int16Item) Uint32() uint32   { return (uint32)(v.v) }
func (v int16Item) Uint64() uint64   { return (uint64)(v.v) }
func (v int16Item) Float32() float32 { return (float32)(v.v) }
func (v int16Item) Float64() float64 { return (float64)(v.v) }
func (v int16Item) Pointer() Pointer { panic("not a pointer") }
func (v int16Item) String() string   { panic("not a string") }

func (v int16Item) Keys() []Value               { panic("Unsupported operation") }
func (v int16Item) Values() []Value             { panic("Unsupported operation") }
func (v int16Item) Items() [][2]Value           { panic("Unsupported operation") }
func (v int16Item) Len() int64                  { panic("Unsupported operation") }
func (v int16Item) Cap() int64                  { panic("Unsupported operation") }
func (v int16Item) Index(n int64) Value         { panic("Unsupported operation") }
func (v int16Item) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v int16Item) Key(k Value) Value           { panic("Unsupported operation") }
func (v int16Item) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v int16Item) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v int32Item) Bool() bool {
	if v.v == 0 {
		return false
	} else {
		return true
	}
}
func (v int32Item) Int8() int8       { return (int8)(v.v) }
func (v int32Item) Int16() int16     { return (int16)(v.v) }
func (v int32Item) Int32() int32     { return v.v }
func (v int32Item) Int64() int64     { return (int64)(v.v) }
func (v int32Item) Uint8() uint8     { return (uint8)(v.v) }
func (v int32Item) Uint16() uint16   { return (uint16)(v.v) }
func (v int32Item) Uint32() uint32   { return (uint32)(v.v) }
func (v int32Item) Uint64() uint64   { return (uint64)(v.v) }
func (v int32Item) Float32() float32 { return (float32)(v.v) }
func (v int32Item) Float64() float64 { return (float64)(v.v) }
func (v int32Item) Pointer() Pointer { panic("not a pointer") }
func (v int32Item) String() string   { panic("not a string") }

func (v int32Item) Keys() []Value               { panic("Unsupported operation") }
func (v int32Item) Values() []Value             { panic("Unsupported operation") }
func (v int32Item) Items() [][2]Value           { panic("Unsupported operation") }
func (v int32Item) Len() int64                  { panic("Unsupported operation") }
func (v int32Item) Cap() int64                  { panic("Unsupported operation") }
func (v int32Item) Index(n int64) Value         { panic("Unsupported operation") }
func (v int32Item) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v int32Item) Key(k Value) Value           { panic("Unsupported operation") }
func (v int32Item) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v int32Item) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v int64Item) Bool() bool {
	if v.v == 0 {
		return false
	} else {
		return true
	}
}
func (v int64Item) Int8() int8       { return (int8)(v.v) }
func (v int64Item) Int16() int16     { return (int16)(v.v) }
func (v int64Item) Int32() int32     { return (int32)(v.v) }
func (v int64Item) Int64() int64     { return v.v }
func (v int64Item) Uint8() uint8     { return (uint8)(v.v) }
func (v int64Item) Uint16() uint16   { return (uint16)(v.v) }
func (v int64Item) Uint32() uint32   { return (uint32)(v.v) }
func (v int64Item) Uint64() uint64   { return (uint64)(v.v) }
func (v int64Item) Float32() float32 { return (float32)(v.v) }
func (v int64Item) Float64() float64 { return (float64)(v.v) }
func (v int64Item) Pointer() Pointer { panic("not a pointer") }
func (v int64Item) String() string   { panic("not a string") }

func (v int64Item) Keys() []Value               { panic("Unsupported operation") }
func (v int64Item) Values() []Value             { panic("Unsupported operation") }
func (v int64Item) Items() [][2]Value           { panic("Unsupported operation") }
func (v int64Item) Len() int64                  { panic("Unsupported operation") }
func (v int64Item) Cap() int64                  { panic("Unsupported operation") }
func (v int64Item) Index(n int64) Value         { panic("Unsupported operation") }
func (v int64Item) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v int64Item) Key(k Value) Value           { panic("Unsupported operation") }
func (v int64Item) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v int64Item) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v uint8Item) Bool() bool {
	if v.v == 0 {
		return false
	} else {
		return true
	}
}
func (v uint8Item) Int8() int8       { return (int8)(v.v) }
func (v uint8Item) Int16() int16     { return (int16)(v.v) }
func (v uint8Item) Int32() int32     { return (int32)(v.v) }
func (v uint8Item) Int64() int64     { return (int64)(v.v) }
func (v uint8Item) Uint8() uint8     { return v.v }
func (v uint8Item) Uint16() uint16   { return (uint16)(v.v) }
func (v uint8Item) Uint32() uint32   { return (uint32)(v.v) }
func (v uint8Item) Uint64() uint64   { return (uint64)(v.v) }
func (v uint8Item) Float32() float32 { return (float32)(v.v) }
func (v uint8Item) Float64() float64 { return (float64)(v.v) }
func (v uint8Item) Pointer() Pointer { panic("not a pointer") }
func (v uint8Item) String() string   { panic("not a string") }

func (v uint8Item) Keys() []Value               { panic("Unsupported operation") }
func (v uint8Item) Values() []Value             { panic("Unsupported operation") }
func (v uint8Item) Items() [][2]Value           { panic("Unsupported operation") }
func (v uint8Item) Len() int64                  { panic("Unsupported operation") }
func (v uint8Item) Cap() int64                  { panic("Unsupported operation") }
func (v uint8Item) Index(n int64) Value         { panic("Unsupported operation") }
func (v uint8Item) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v uint8Item) Key(k Value) Value           { panic("Unsupported operation") }
func (v uint8Item) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v uint8Item) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v uint16Item) Bool() bool {
	if v.v == 0 {
		return false
	} else {
		return true
	}
}
func (v uint16Item) Int8() int8       { return (int8)(v.v) }
func (v uint16Item) Int16() int16     { return (int16)(v.v) }
func (v uint16Item) Int32() int32     { return (int32)(v.v) }
func (v uint16Item) Int64() int64     { return (int64)(v.v) }
func (v uint16Item) Uint8() uint8     { return (uint8)(v.v) }
func (v uint16Item) Uint16() uint16   { return v.v }
func (v uint16Item) Uint32() uint32   { return (uint32)(v.v) }
func (v uint16Item) Uint64() uint64   { return (uint64)(v.v) }
func (v uint16Item) Float32() float32 { return (float32)(v.v) }
func (v uint16Item) Float64() float64 { return (float64)(v.v) }
func (v uint16Item) Pointer() Pointer { panic("not a pointer") }
func (v uint16Item) String() string   { panic("not a string") }

func (v uint16Item) Keys() []Value               { panic("Unsupported operation") }
func (v uint16Item) Values() []Value             { panic("Unsupported operation") }
func (v uint16Item) Items() [][2]Value           { panic("Unsupported operation") }
func (v uint16Item) Len() int64                  { panic("Unsupported operation") }
func (v uint16Item) Cap() int64                  { panic("Unsupported operation") }
func (v uint16Item) Index(n int64) Value         { panic("Unsupported operation") }
func (v uint16Item) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v uint16Item) Key(k Value) Value           { panic("Unsupported operation") }
func (v uint16Item) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v uint16Item) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v uint32Item) Bool() bool {
	if v.v == 0 {
		return false
	} else {
		return true
	}
}
func (v uint32Item) Int8() int8       { return (int8)(v.v) }
func (v uint32Item) Int16() int16     { return (int16)(v.v) }
func (v uint32Item) Int32() int32     { return (int32)(v.v) }
func (v uint32Item) Int64() int64     { return (int64)(v.v) }
func (v uint32Item) Uint8() uint8     { return (uint8)(v.v) }
func (v uint32Item) Uint16() uint16   { return (uint16)(v.v) }
func (v uint32Item) Uint32() uint32   { return v.v }
func (v uint32Item) Uint64() uint64   { return (uint64)(v.v) }
func (v uint32Item) Float32() float32 { return (float32)(v.v) }
func (v uint32Item) Float64() float64 { return (float64)(v.v) }
func (v uint32Item) Pointer() Pointer { panic("not a pointer") }
func (v uint32Item) String() string   { panic("not a string") }

func (v uint32Item) Keys() []Value               { panic("Unsupported operation") }
func (v uint32Item) Values() []Value             { panic("Unsupported operation") }
func (v uint32Item) Items() [][2]Value           { panic("Unsupported operation") }
func (v uint32Item) Len() int64                  { panic("Unsupported operation") }
func (v uint32Item) Cap() int64                  { panic("Unsupported operation") }
func (v uint32Item) Index(n int64) Value         { panic("Unsupported operation") }
func (v uint32Item) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v uint32Item) Key(k Value) Value           { panic("Unsupported operation") }
func (v uint32Item) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v uint32Item) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v uint64Item) Bool() bool {
	if v.v == 0 {
		return false
	} else {
		return true
	}
}
func (v uint64Item) Int8() int8       { return (int8)(v.v) }
func (v uint64Item) Int16() int16     { return (int16)(v.v) }
func (v uint64Item) Int32() int32     { return (int32)(v.v) }
func (v uint64Item) Int64() int64     { return (int64)(v.v) }
func (v uint64Item) Uint8() uint8     { return (uint8)(v.v) }
func (v uint64Item) Uint16() uint16   { return (uint16)(v.v) }
func (v uint64Item) Uint32() uint32   { return (uint32)(v.v) }
func (v uint64Item) Uint64() uint64   { return v.v }
func (v uint64Item) Float32() float32 { return (float32)(v.v) }
func (v uint64Item) Float64() float64 { return (float64)(v.v) }
func (v uint64Item) Pointer() Pointer { panic("not a pointer") }
func (v uint64Item) String() string   { panic("not a string") }

func (v uint64Item) Keys() []Value               { panic("Unsupported operation") }
func (v uint64Item) Values() []Value             { panic("Unsupported operation") }
func (v uint64Item) Items() [][2]Value           { panic("Unsupported operation") }
func (v uint64Item) Len() int64                  { panic("Unsupported operation") }
func (v uint64Item) Cap() int64                  { panic("Unsupported operation") }
func (v uint64Item) Index(n int64) Value         { panic("Unsupported operation") }
func (v uint64Item) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v uint64Item) Key(k Value) Value           { panic("Unsupported operation") }
func (v uint64Item) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v uint64Item) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v float32Item) Bool() bool {
	if v.v == 0 {
		return false
	} else {
		return true
	}
}
func (v float32Item) Int8() int8       { return (int8)(v.v) }
func (v float32Item) Int16() int16     { return (int16)(v.v) }
func (v float32Item) Int32() int32     { return (int32)(v.v) }
func (v float32Item) Int64() int64     { return (int64)(v.v) }
func (v float32Item) Uint8() uint8     { return (uint8)(v.v) }
func (v float32Item) Uint16() uint16   { return (uint16)(v.v) }
func (v float32Item) Uint32() uint32   { return (uint32)(v.v) }
func (v float32Item) Uint64() uint64   { return (uint64)(v.v) }
func (v float32Item) Float32() float32 { return v.v }
func (v float32Item) Float64() float64 { return (float64)(v.v) }
func (v float32Item) Pointer() Pointer { panic("not a pointer") }
func (v float32Item) String() string   { panic("not a string") }

func (v float32Item) Keys() []Value               { panic("Unsupported operation") }
func (v float32Item) Values() []Value             { panic("Unsupported operation") }
func (v float32Item) Items() [][2]Value           { panic("Unsupported operation") }
func (v float32Item) Len() int64                  { panic("Unsupported operation") }
func (v float32Item) Cap() int64                  { panic("Unsupported operation") }
func (v float32Item) Index(n int64) Value         { panic("Unsupported operation") }
func (v float32Item) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v float32Item) Key(k Value) Value           { panic("Unsupported operation") }
func (v float32Item) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v float32Item) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v float64Item) Bool() bool {
	if v.v == 0 {
		return false
	} else {
		return true
	}
}
func (v float64Item) Int8() int8       { return (int8)(v.v) }
func (v float64Item) Int16() int16     { return (int16)(v.v) }
func (v float64Item) Int32() int32     { return (int32)(v.v) }
func (v float64Item) Int64() int64     { return (int64)(v.v) }
func (v float64Item) Uint8() uint8     { return (uint8)(v.v) }
func (v float64Item) Uint16() uint16   { return (uint16)(v.v) }
func (v float64Item) Uint32() uint32   { return (uint32)(v.v) }
func (v float64Item) Uint64() uint64   { return (uint64)(v.v) }
func (v float64Item) Float32() float32 { return (float32)(v.v) }
func (v float64Item) Float64() float64 { return v.v }
func (v float64Item) Pointer() Pointer { panic("not a pointer") }
func (v float64Item) String() string   { panic("not a string") }

func (v float64Item) Keys() []Value               { panic("Unsupported operation") }
func (v float64Item) Values() []Value             { panic("Unsupported operation") }
func (v float64Item) Items() [][2]Value           { panic("Unsupported operation") }
func (v float64Item) Len() int64                  { panic("Unsupported operation") }
func (v float64Item) Cap() int64                  { panic("Unsupported operation") }
func (v float64Item) Index(n int64) Value         { panic("Unsupported operation") }
func (v float64Item) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v float64Item) Key(k Value) Value           { panic("Unsupported operation") }
func (v float64Item) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v float64Item) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v pointerItem) Bool() bool       { return v.v == nil }
func (v pointerItem) Int8() int8       { panic("pointer is not a int8") }
func (v pointerItem) Int16() int16     { panic("pointer is not a int16") }
func (v pointerItem) Int32() int32     { panic("pointer is not a int32") }
func (v pointerItem) Int64() int64     { panic("pointer is not a int64") }
func (v pointerItem) Uint8() uint8     { panic("pointer is not a uint8") }
func (v pointerItem) Uint16() uint16   { panic("pointer is not a uint16") }
func (v pointerItem) Uint32() uint32   { panic("pointer is not a uint32") }
func (v pointerItem) Uint64() uint64   { panic("pointer is not a uint64") }
func (v pointerItem) Float32() float32 { panic("pointer is not a float32") }
func (v pointerItem) Float64() float64 { panic("pointer is not a float64") }
func (v pointerItem) Pointer() Pointer { return v.v }
func (v pointerItem) String() string   { panic("pointer is not a string") }

func (v pointerItem) Keys() []Value               { panic("Unsupported operation") }
func (v pointerItem) Values() []Value             { panic("Unsupported operation") }
func (v pointerItem) Items() [][2]Value           { panic("Unsupported operation") }
func (v pointerItem) Len() int64                  { panic("Unsupported operation") }
func (v pointerItem) Cap() int64                  { panic("Unsupported operation") }
func (v pointerItem) Index(n int64) Value         { panic("Unsupported operation") }
func (v pointerItem) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v pointerItem) Key(k Value) Value           { panic("Unsupported operation") }
func (v pointerItem) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v pointerItem) SetKey(k Value, val Value)   { panic("Unsupported operation") }

func (v stringItem) Bool() bool       { panic("string is not a bool") }
func (v stringItem) Int8() int8       { panic("string is not a int8") }
func (v stringItem) Int16() int16     { panic("string is not a int16") }
func (v stringItem) Int32() int32     { panic("string is not a int32") }
func (v stringItem) Int64() int64     { panic("string is not a int64") }
func (v stringItem) Uint8() uint8     { panic("string is not a uint8") }
func (v stringItem) Uint16() uint16   { panic("string is not a uint16") }
func (v stringItem) Uint32() uint32   { panic("string is not a uint32") }
func (v stringItem) Uint64() uint64   { panic("string is not a uint64") }
func (v stringItem) Float32() float32 { panic("string is not a float32") }
func (v stringItem) Float64() float64 { panic("string is not a float64") }
func (v stringItem) Pointer() Pointer { panic("string is not a pointer") }
func (v stringItem) String() string   { return v.v }

func (v stringItem) Keys() []Value               { panic("Unsupported operation") }
func (v stringItem) Values() []Value             { panic("Unsupported operation") }
func (v stringItem) Items() [][2]Value           { panic("Unsupported operation") }
func (v stringItem) Len() int64                  { panic("Unsupported operation") }
func (v stringItem) Cap() int64                  { panic("Unsupported operation") }
func (v stringItem) Index(n int64) Value         { panic("Unsupported operation") }
func (v stringItem) SetIndex(n int64, val Value) { panic("Unsupported operation") }
func (v stringItem) Key(k Value) Value           { panic("Unsupported operation") }
func (v stringItem) HasKey(k Value) bool         { panic("Unsupported operation") }
func (v stringItem) SetKey(k Value, val Value)   { panic("Unsupported operation") }

// func (v stringItem) Bool() bool       { panic("string is not a bool") }
// func (v stringItem) Int8() int8       { panic("string is not a int8") }
// func (v stringItem) Int16() int16     { panic("string is not a int16") }
// func (v stringItem) Int32() int32     { panic("string is not a int32") }
// func (v stringItem) Int64() int64     { panic("string is not a int64") }
// func (v stringItem) Uint8() uint8     { panic("string is not a uint8") }
// func (v stringItem) Uint16() uint16   { panic("string is not a uint16") }
// func (v stringItem) Uint32() uint32   { panic("string is not a uint32") }
// func (v stringItem) Uint64() uint64   { panic("string is not a uint64") }
// func (v stringItem) Float32() float32 { panic("string is not a float32") }
// func (v stringItem) Float64() float64 { panic("string is not a float64") }
// func (v stringItem) Pointer() Pointer { panic("string is not a pointer") }
// func (v stringItem) String() string   { return v.v }

// func (v stringItem) Keys() []Value               { panic("Unsupported operation") }
// func (v stringItem) Values() []Value             { panic("Unsupported operation") }
// func (v stringItem) Items() [][2]Value           { panic("Unsupported operation") }
// func (v stringItem) Len() int64                  { panic("Unsupported operation") }
// func (v stringItem) Cap() int64                  { panic("Unsupported operation") }
// func (v stringItem) Index(n int64) Value         { panic("Unsupported operation") }
// func (v stringItem) SetIndex(n int64, val Value) { panic("Unsupported operation") }
// func (v stringItem) Key(k Value) Value           { panic("Unsupported operation") }
// func (v stringItem) HasKey(k Value) bool         { panic("Unsupported operation") }
// func (v stringItem) SetKey(k Value, val Value)   { panic("Unsupported operation") }
