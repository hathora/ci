package output

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

func ValueFormat(empty any, field string) (FormatWriter, error) {
	tType := reflect.TypeOf(empty)
	for i := 0; i < tType.NumField(); i++ {
		structField := tType.Field(i)
		jsonTag := structField.Tag.Get("json")
		tagName := strings.Split(jsonTag, ",")[0]

		if tagName == "-" {
			continue
		}

		if tagName == field {
			return &valueOutputWriter{
				outputType: tType,
				field:      structField.Name,
			}, nil
		}
	}

	return nil, fmt.Errorf("field \"%s\" not supported for output type: %s", field, tType.Name())
}

type valueOutputWriter struct {
	outputType reflect.Type
	field      string
}

var _ FormatWriter = (*valueOutputWriter)(nil)

func (j *valueOutputWriter) Write(value any, writer io.Writer) error {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	fieldValue := v.FieldByName(j.field)
	if !fieldValue.IsValid() {
		return fmt.Errorf("field %s not found", j.field)
	}

	fmt.Fprintln(writer, fieldValue.Interface())

	return nil
}
