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
	var jsonBytes []byte
	var err error

	if j.Pretty {
		jsonBytes, err = json.MarshalIndent(value, "", "  ")
	} else {
		jsonBytes, err = json.Marshal(value)
	}
	if err != nil {
		return fmt.Errorf("failed to json-serialize output value %w", err)
	}

	_, err = writer.Write(jsonBytes)
	return err
}
