package cla

import (
	"bufio"
	"fmt"
	"os"
)

type Console struct {
	reader *bufio.Reader
}

func NewConsole() *Console {
	return &Console{reader: bufio.NewReader(os.Stdin)}
}

func (c *Console) Print(s string) {
	fmt.Print(s)
}

func (c *Console) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (c *Console) ReadLine() (string, error) {
	text, _ := c.reader.ReadString('\n')
	return text[:len(text)-1], nil
}
