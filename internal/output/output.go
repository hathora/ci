package output

import "io"

type FormatWriter interface {
	Write(value any, writer io.Writer) error
}
