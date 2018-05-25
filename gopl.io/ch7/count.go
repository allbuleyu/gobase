package ch7

import (
	"io"
	"fmt"
)

type myReader struct {
	s string
	i int64
	prevRune int
}

func (r *myReader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	fmt.Println(n)
	r.i += int64(n)
	return
}

func WordLenCount() {

}

//
func NewReader(s string) io.Reader {

	//return strings.NewReader(s)
	return &myReader{s:s, i:0, prevRune:-1}
}

type LimitedReader struct {
	R io.Reader
	N int64
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{R:r, N:n}
}