package virtual

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
)

type SeekerFile interface {
	fs.File
	io.Seeker
}

type bufferFile struct {
	fs.File
	s []byte
	i int64
}

func SeekerOf(f fs.File) SeekerFile {
	if s, ok := f.(SeekerFile); ok {
		return s
	}
	return &bufferFile{File: f}
}

func (buf *bufferFile) Read(b []byte) (int, error) {
	n := copy(b, buf.s[buf.i:])
	if n == len(b) {
		return n, nil
	}
	x, err := buf.File.Read(b[n:])
	x += n
	if err == nil {
		buf.s = append(buf.s, b[n:x]...)
		buf.i += int64(x)
	}
	return x, err
}

func (buf *bufferFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		buf.i = offset
	case io.SeekCurrent:
		buf.i += offset
	case io.SeekEnd:
		b := bytes.NewBuffer(buf.s)
		_, err := io.Copy(b, buf.File)
		if err != nil {
			return 0, err
		}
		buf.i = int64(b.Len()) + offset
		buf.s = b.Bytes()
	default:
		return 0, errors.New("buffer.Seek: invalid whence")
	}
	if buf.i < 0 {
		return 0, errors.New("buffer.Seek: negative position")
	}
	n := int64(len(buf.s))
	if buf.i <= n {
		return buf.i, nil
	}
	buf.s = append(buf.s, make([]byte, buf.i-n)...)
	x, err := buf.File.Read(buf.s[n:])
	if err != nil {
		return 0, err
	}
	buf.i = n + int64(x)
	buf.s = buf.s[:buf.i]
	return buf.i, nil
}
