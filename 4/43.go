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
	
	for i := 0; i < n; i++ {
        char := b[i]
        if char >= 'a' && char <= 'z' {

        } else if char >= 'A' && char <= 'Z' {

        }
    }

    return n, err 

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
