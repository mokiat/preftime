package prefix

import (
	"fmt"
	"io"
	"time"
)

type PrefixFunction func(io.Writer) (int, error)

func NewIndexPrefixFunction() PrefixFunction {
	index := 0
	return func(out io.Writer) (int, error) {
		index++
		return fmt.Fprintf(out, "[%d] ", index)
	}
}

func NewTimePrefixFunction() PrefixFunction {
	return func(out io.Writer) (int, error) {
		return fmt.Fprintf(out, "[%s] ", time.Now().Format("2006-01-02 15:04:05.000"))
	}
}
