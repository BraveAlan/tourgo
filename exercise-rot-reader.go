package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	if len(b) < 0 {
		return len(b), errors.New("length is empty")
	}
	n, err = r.r.Read(b)
	if err != nil {
		return len(b), err
	}
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return
}

func rot13(b byte) byte {
	var a byte
	switch {
	case b >= 'a' && b <= 'z':
		a = byte('a')
	case b >= 'A' && b <= 'Z':
		a = byte('A')
	default:
		return b
	}
	if b-a >= 13 {
		return b - 13
	} else {
		return b + 13
	}

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
