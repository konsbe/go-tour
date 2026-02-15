package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(p []byte) (int, error) {
	n, err := rr.r.Read(p)

	for i := range n {
		c := p[i]

		switch {
		case c >= 'a' && c <= 'z':
			p[i] = 'a' + (c-'a'+13)%26
		case c >= 'A' && c <= 'Z':
			p[i] = 'A' + (c-'A'+13)%26
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
