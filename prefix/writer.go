package prefix

import (
	"bytes"
	"io"
)

func NewWriter(delegate io.Writer, prefixFunc Func) io.Writer {
	return &writer{
		delegate:   delegate,
		prefixFunc: prefixFunc,
		beginning:  true,
	}
}

type writer struct {
	delegate   io.Writer
	prefixFunc Func
	beginning  bool
}

func (w *writer) Write(data []byte) (int, error) {
	dataCount := len(data)

	index := bytes.IndexByte(data, '\n')
	for index >= 0 {
		if err := w.writeSegment(data[:index+1]); err != nil {
			return 0, err
		}
		data = data[index+1:]
		index = bytes.IndexByte(data, '\n')
	}

	if len(data) > 0 {
		if err := w.writeSegment(data); err != nil {
			return 0, err
		}
	}

	return dataCount, nil
}

func (w *writer) writeSegment(data []byte) error {
	if w.beginning {
		if _, err := w.prefixFunc(w.delegate); err != nil {
			return err
		}
		w.beginning = false
	}

	if _, err := w.delegate.Write(data); err != nil {
		return err
	}

	if data[len(data)-1] == '\n' {
		w.beginning = true
	}

	return nil
}
