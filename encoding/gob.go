
package vm_encoding

import (
	"io"
	"errors"
)

const MagicHeader = 0xbd57a496

var UnexceptMagicHeader = errors.New("Unexcept magic header")

type LabelMeta struct{
	Pkg string
	Path string
	Name string
	Program io.ReaderAt
	Offset int64
}

type Gob struct{
	parsed bool

	path string
	name string
	code io.ReaderAt
	offset int64

	labels map[string]LabelMeta
}

func NewGob(path string, code io.ReaderAt)(*Gob){
	return &Gob{
		path: path,
		code: code,
	}
}

func (g *Gob)Path()(string){
	return g.path
}

func (g *Gob)Name()(string){
	return g.name
}

func (g *Gob)ReaderAt()(io.ReaderAt){
	return g.code
}

func (g *Gob)Offset()(int64){
	return g.offset
}

func (g *Gob)ParseMetadata()(err error){
	if g.parsed {
		return
	}
	r := io.NewSectionReader(g.code, 0, -1)
	var head [8]byte
	_, err = r.Read(head[:4])
	if err != nil || BytesToUint32(head[:4]) != MagicHeader {
		return UnexceptMagicHeader
	}
	if _, err = r.Read(head[:2]); err != nil {
		return
	}
	buf := make([]byte, BytesToUint16(head[:2]))
	if _, err = r.Read(buf); err != nil {
		return
	}
	g.name = (string)(buf)
	if _, err = r.Read(head[:8]); err != nil {
		return
	}
	g.offset = (int64)(BytesToUint64(head[:8]))
	if _, err = r.Read(head[:4]); err != nil {
		return
	}
	if len(buf) < 256 {
		buf = make([]byte, 256)
	}
	for i := BytesToUint32(head[:4]); i > 0; i-- {
		var meta LabelMeta
		meta.Pkg = g.name
		if _, err = r.Read(head[:1]); err != nil {
			return
		}
		l := (int)(head[0])
		if _, err = r.Read(buf[:l]); err != nil {
			return
		}
		meta.Name = (string)(buf[:l])
		if _, err = r.Read(head[:2]); err != nil {
			return
		}
		if l = (int)(BytesToUint16(head[:2])); len(buf) < l {
			buf = make([]byte, l)
		}
		if _, err = r.Read(buf[:l]); err != nil {
			return
		}
		meta.Path = (string)(buf[:l])
		if _, err = r.Read(head[:8]); err != nil {
			return
		}
		meta.Offset = (int64)(BytesToUint64(head[:8]))
		g.labels[meta.Name] = meta
	}
	g.parsed = true
	return
}

func (g *Gob)Lookup(label string)(meta LabelMeta, ok bool){
	meta, ok = g.labels[label]
	return
}
