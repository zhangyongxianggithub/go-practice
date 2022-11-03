package ch7

import (
	"bufio"
)

// 定义类型
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	advance, _, err := bufio.ScanWords(p, false)
	if err == nil {
		*w += WordCounter(advance)
	}
	return advance, err
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	advance, _, err := bufio.ScanLines(p, true)
	if err == nil {
		*l += LineCounter(advance)
	}
	return advance, err
}
