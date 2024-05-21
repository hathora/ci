package output

import (
	"fmt"
	"io"
)

func TextFormat() FormatWriter {
	return &textOutputWriter{}
}

type textOutputWriter struct{}

var _ FormatWriter = (*textOutputWriter)(nil)

func (t *textOutputWriter) Write(value any, writer io.Writer) error {
	// TODO: Use go templates to define text output for different output types
	_, err := fmt.Fprintf(writer, "%v\n", value)
	return err
}
