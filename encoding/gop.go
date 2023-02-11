
package encoding

import (
	"archive/zip"
	"bytes"
	"io"
	// "path"
)

type Gop struct{
	path string
	file *zip.ReadCloser

	cache map[string]io.ReaderAt
}

func OpenGop(path string)(p *Gop, err error){
	p = &Gop{
		path: path,
	}
	err = p.Init()
	if err != nil {
		return nil, err
	}
	return
}

func (p *Gop)Init()(err error){
	if p.file != nil {
		panic("Gop already inited")
	}
	p.file, err = zip.OpenReader(p.path)
	if err != nil { return }
	return
}

func (p *Gop)Close()(err error){
	if p.file == nil {
		return nil
	}
	err = p.file.Close()
	p.file = nil
	return
}

func (p *Gop)Open(name string)(bf io.ReaderAt, err error){
	bf, ok := p.cache[name]
	if ok {
		return
	}
	if p.file == nil {
		if err = p.Init(); err != nil {
			return
		}
	}
	r, err := p.file.Open(name)
	if err != nil { return }
	buf, err := io.ReadAll(r)
	if err != nil { return }
	bf = bytes.NewReader(buf)
	p.cache[name] = bf
	return
}
