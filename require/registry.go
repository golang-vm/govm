
package vm_require

import (
	"strings"

	encoding "github.com/golang-vm/govm/encoding"
	cpu "github.com/golang-vm/govm/cpu"
)

type Registry struct{
	natives map[string]cpu.NativeCall
	labels map[string]encoding.LabelMeta
}

var _ cpu.Registry = (*Registry)(nil)

func NewRegistry()(r *Registry){
	return &Registry{
		natives: make(map[string]cpu.NativeCall),
		labels: make(map[string]encoding.LabelMeta),
	}
}

func (r *Registry)Lookup(label string)(meta encoding.LabelMeta, ok bool){
	// TODO: use package name
	i := strings.IndexByte(label, '@')
	name, pkg := label[:i], label[i + 1:]
	_, _ = name, pkg
	meta, ok = r.labels[label]
	return
}

func (r *Registry)Register(label string, fuc any){
	r.natives[label] = fuc.(cpu.NativeCall)
}

func (r *Registry)GetNative(label string)(cpu.NativeCall){
	return r.natives[label]
}

