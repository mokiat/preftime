package prefix

import (
	"bytes"
	"io"
)

func NewPrefixWriter(dest io.Writer, prefixFunc PrefixFunction) io.Writer {
	return &prefixWriter{
		beginning:  true,
		prefixFunc: prefixFunc,
		dest:       dest,
	}
}

type prefixWriter struct {
	beginning  bool
	prefixFunc PrefixFunction
	dest       io.Writer
}

func (w *prefixWriter) Write(data []byte) (int, error) {
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

func (w *prefixWriter) writeSegment(data []byte) error {
	if w.beginning {
		if _, err := w.prefixFunc(w.dest); err != nil {
			return err
		}
		w.beginning = false
	}

	if _, err := w.dest.Write(data); err != nil {
		return err
	}

	if data[len(data)-1] == '\n' {
		w.beginning = true
	}

	return nil
}
