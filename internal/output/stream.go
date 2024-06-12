package output

import (
	"bufio"
	"fmt"
	"io"
)

func StreamOutput(reader io.ReadCloser, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
