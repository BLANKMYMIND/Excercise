package main

import "golang.org/x/tour/reader"
import "fmt"
import "io"

type MyReader struct{}

// JOB: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (r MyReader) Read(b []byte) (res int, err error) {
	sum := 0
	for i := range b {
		b[i] = 65
		sum++
	}
	return sum, err
}

func main() {
	reader.Validate(MyReader{})

	var r MyReader
	b := make([]byte, 8)
	// self-test
	for i := 0; i < 4; i++ {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
