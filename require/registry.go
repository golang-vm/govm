
package require

import (
	"strings"

	encoding "github.com/golang-vm/govm/encoding"
	cpu "github.com/golang-vm/govm/cpu"
)

type Registry struct{
	natives map[string]cpu.NativeCall
	labels map[string]encoding.LabelMeta
	pkgs map[string]*encoding.Gob
}

var _ cpu.Registry = (*Registry)(nil)

func NewRegistry()(r *Registry){
	return &Registry{
		natives: make(map[string]cpu.NativeCall),
		labels: make(map[string]encoding.LabelMeta),
		pkgs: make(map[string]*encoding.Gob),
	}
}

func (r *Registry)Lookup(label string)(meta encoding.LabelMeta, ok bool){
	i := strings.IndexByte(label, '@')
	if i >= 0 {
		lb, pkg := label[:i], label[i + 1:]
		var p *encoding.Gob
		if p, ok = r.pkgs[pkg]; !ok {
			return
		}
		return p.Lookup(lb)
	}
	meta, ok = r.labels[label]
	return
}

func (r *Registry)AddPkg(p *encoding.Gob){
	r.pkgs[p.Path()] = p
}

func (r *Registry)Register(label string, fuc any){
	r.natives[label] = fuc.(cpu.NativeCall)
}

func (r *Registry)GetNative(label string)(cpu.NativeCall){
	return r.natives[label]
}

