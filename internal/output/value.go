package output

import (
	"fmt"
	"io"
	"strings"

	"github.com/hathora/ci/internal/sdk/models/shared"
)

func ValueFormat(field string) FormatWriter {
	return &valueOutputWriter{
		Field: field,
	}
}

type valueOutputWriter struct {
	Field string
}

var _ FormatWriter = (*valueOutputWriter)(nil)

func (j *valueOutputWriter) Write(value any, writer io.Writer) error {
	buildPtr, test := value.(*shared.Build)
	if test {
		found, err := getBuildValuePtr(buildPtr, j.Field)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(writer, "%v\n", found)
		return err
	}

	build, test := value.(shared.Build)
	if test {
		found, err := getBuildValue(build, j.Field)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(writer, "%v\n", found)
		return err
	}

	builds, test := value.([]shared.Build)
	if test {
		for _, b := range builds {
			found, err := getBuildValue(b, j.Field)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintf(writer, "%v\n", found)
			if err != nil {
				return err
			}
		}

		return nil
	}

	// default to just writing the text
	_, err := fmt.Fprintf(writer, "%v\n", value)
	return err
}

func getBuildValuePtr(build *shared.Build, key string) (any, error) {
	return getBuildValue(*build, key)
}

func getBuildValue(build shared.Build, key string) (any, error) {
	switch strings.ToLower(key) {
	case "buildid":
		return build.BuildID, nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("key %s not supported", key))
	}
}
