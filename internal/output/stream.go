package output

import (
	"fmt"
	"io"
)

func StreamOutput(reader io.ReadCloser, writer io.Writer) error {
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}

		if n > 0 {
			_, err = fmt.Fprintf(writer, "%v\n", string(buffer[:n]))
			if err != nil {
				return err
			}
		}

		if err == io.EOF {
			break
		}
	}
	
	return nil
}
