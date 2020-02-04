package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	for i := range b {
		if b[i] >= 65 && b[i] <= 90 {
			b[i] += 13
			if b[i] > 90 {
				b[i] -= 26
			}
		} else if b[i] >= 97 && b[i] <= 122 {
			b[i] += 13
			if b[i] > 122 {
				b[i] -= 26
			}
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
