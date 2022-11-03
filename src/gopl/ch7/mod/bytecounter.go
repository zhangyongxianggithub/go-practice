package mod

import (
	"bufio"
	"strings"
)

// 定义类型
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	count, err := retCount(p, bufio.ScanWords)
	if err == nil {
		*w += WordCounter(count)
	}
	return count, err
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	count, err := retCount(p, bufio.ScanLines)
	if err == nil {
		*l += LineCounter(count)
	}
	return count, err
}

func retCount(p []byte, fn bufio.SplitFunc) (int, error) {
	s := string(p)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(fn)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}
