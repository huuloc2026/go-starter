package main

import (
	"fmt"
)

func main() {
	var w Writer = &ConsoleWriter{}
	w.Writer([]byte("Hello struct"))
}

type Writer interface {
	Writer([]byte) (int, error)
}
type ConsoleWriter struct {
}

func (cw *ConsoleWriter) Writer(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
