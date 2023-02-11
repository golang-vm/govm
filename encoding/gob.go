
package encoding

import (
	"io"
	"errors"
)

const MagicHeader = 0xbd57a496

var UnexceptMagicHeader = errors.New("Unexcept magic header")

// A structure that used to save metadata of labels(methods)
type LabelMeta struct{
	Pkg *Gob
	Name string
	Offset int64
}

// The type of a program
type Gob struct{
	parsed bool

	path string
	name string
	code io.ReaderAt

	strlst string
	labels map[string]LabelMeta
}

func NewGob(path string, code io.ReaderAt) *Gob {
	return &Gob{
		path: path,
		code: code,
		labels: make(map[string]LabelMeta),
	}
}

func (g *Gob) Path() string {
	return g.path
}

func (g *Gob) Name() string {
	return g.name
}

func (g *Gob) Program() io.ReaderAt {
	return g.code
}

func (g *Gob) ParseMetadata() (err error) {
	if g.parsed {
		return
	}
	r := io.NewSectionReader(g.code, 0, -1)
	var head [8]byte
	if _, err = r.Read(head[:4]); err != nil || BytesToUint32(head[:4]) != MagicHeader {
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
	if _, err = r.Read(head[:4]); err != nil {
		return
	}
	l := BytesToUint32(head[:4])
	if (uint32)(len(buf)) < l {
		buf = make([]byte, l)
	}
	if _, err = r.Read(buf[:l]); err != nil {
		return
	}
	g.strlst = (string)(buf)
	if _, err = r.Read(head[:4]); err != nil {
		return
	}
	if len(buf) < 256 {
		buf = make([]byte, 256)
	}
	for i := BytesToUint32(head[:4]); i > 0; i-- {
		var meta LabelMeta
		meta.Pkg = g
		if _, err = r.Read(head[:8]); err != nil {
			return
		}
		p, l := BytesToUint32(head[0:4]), BytesToUint32(head[4:8])
		meta.Name = g.strlst[p:p + l]
		if len(meta.Name) == 0 {
			panic("Label name cannot be empty")
		}
		if _, err = r.Read(head[:8]); err != nil {
			return
		}
		meta.Offset = (int64)(BytesToUint64(head[:8]))
		g.labels[meta.Name] = meta
	}
	g.parsed = true
	return
}

func (g *Gob) Strlst() string {
	return g.strlst
}

func (g *Gob) Labels() (labels []LabelMeta) {
	labels = make([]LabelMeta, 0, len(g.labels))
	for _, l := range g.labels {
		labels = append(labels, l)
	}
	return
}

func (g *Gob) Lookup(label string) (meta LabelMeta, ok bool) {
	meta, ok = g.labels[label]
	return
}
