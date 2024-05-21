package output

import (
	"encoding/json"
	"fmt"
	"io"
)

func JSONFormat(pretty bool) FormatWriter {
	return &jsonOutputWriter{
		Pretty: pretty,
	}
}

type jsonOutputWriter struct {
	Pretty bool
}

var _ FormatWriter = (*jsonOutputWriter)(nil)

func (j *jsonOutputWriter) Write(value any, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	if j.Pretty {
		encoder.SetIndent("", "  ")
	}
	err := encoder.Encode(value)
	if err != nil {
		return fmt.Errorf("failed to json-serialize output value %w", err)
	}

	return nil
}
