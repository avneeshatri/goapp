package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var wc WriterCloser = NewBufferWriteCloser()
	wc.Write([]byte("hi there how are you"))
	wc.Close()

	bwc := wc.(*BufferedWriterClose)
	fmt.Println(bwc)

	var myObj interface{} = NewBufferWriteCloser()
	if wc, ok := myObj.(BufferedWriterClose); ok {
		wc.Write([]byte("Hi there how are you"))
		wc.Close()
	}
	r, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

}

type Writer interface {
	Write(data []byte) (int, error)
}
type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterClose struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterClose) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}

	return nil
}

func (bwc *BufferedWriterClose) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)

	if err != nil {
		return 0, nil
	}

	v := make([]byte, 8)

	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, nil
		}

		_, err = fmt.Println(string(v))

		if err != nil {
			return 0, nil
		}
	}
	return n, nil
}

func NewBufferWriteCloser() *BufferedWriterClose {
	return &BufferedWriterClose{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
