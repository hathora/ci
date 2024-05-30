package output

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"slices"
	"sort"

	"text/tabwriter"
)

func TextFormat(opts ...TextFormatterOption) FormatWriter {
	registry := &textFormatterRegistry{
		FieldOrders: make(map[string][]string),
		OmitFields:  make(map[string][]string),
		Formatters:  make(map[string]formatter),
	}

	for _, opt := range opts {
		opt(registry)
	}

	return &textOutputWriter{
		registry: registry,
	}
}

type formatter func(any) string

type textFormatterRegistry struct {
	FieldOrders map[string][]string
	OmitFields  map[string][]string
	Formatters  map[string]formatter
}

type TextFormatterOption func(*textFormatterRegistry)

func WithFieldOrder[T any](empty T, fields ...string) TextFormatterOption {
	return func(r *textFormatterRegistry) {
		typeKey := reflect.TypeOf(empty).String()
		r.FieldOrders[typeKey] = fields
	}
}

func WithoutFields[T any](empty T, fields ...string) TextFormatterOption {
	return func(r *textFormatterRegistry) {
		typeKey := reflect.TypeOf(empty).String()
		r.OmitFields[typeKey] = fields
	}
}

func WithFormatter[T any](empty T, f func(T) string) TextFormatterOption {
	return func(r *textFormatterRegistry) {
		typeKey := reflect.TypeOf(empty).String()
		r.Formatters[typeKey] = func(value any) string {
			return f(value.(T))
		}
	}
}

type textOutputWriter struct {
	registry *textFormatterRegistry
}

var _ FormatWriter = (*textOutputWriter)(nil)

func (t *textOutputWriter) Write(value any, writer io.Writer) error {
	tw := tabwriter.NewWriter(writer, 0, 0, 2, ' ', 0)

	v := reflect.ValueOf(value)
	valueType := v.Type()

	switch v.Kind() {
	case reflect.Slice:
		if v.Len() == 0 {
			return nil
		}
		elemType := valueType.Elem()
		err := t.printFieldNames(v.Index(0), elemType, tw)
		if err != nil {
			return err
		}
		for i := 0; i < v.Len(); i++ {
			err = t.printFieldValues(v.Index(i), tw)
			if err != nil {
				return err
			}
		}
	case reflect.Struct:
		err := t.printFieldNames(v, valueType, tw)
		if err != nil {
			return err
		}
		err = t.printFieldValues(v, tw)
		if err != nil {
			return err
		}
	case reflect.Ptr:
		if v.IsNil() {
			return nil
		}
		return t.Write(v.Elem().Interface(), writer)
	default:
		return fmt.Errorf("unsupported type: %s", v.Kind())
	}
	fmt.Fprintln(writer)

	return tw.Flush()
}

func (t *textOutputWriter) printFieldNames(v reflect.Value, valueType reflect.Type, writer io.Writer) error {
	var fieldNames []string

	for i := 0; i < v.NumField(); i++ {
		fieldName := valueType.Field(i).Name
		omissions, hasOmissions := t.registry.OmitFields[valueType.String()]
		if hasOmissions && slices.Contains(omissions, fieldName) {
			continue
		}

		fieldNames = append(fieldNames, valueType.Field(i).Name)
	}

	sort.Slice(fieldNames, func(i, j int) bool {
		registryFieldOrder, hasFieldOrder := t.registry.FieldOrders[valueType.String()]
		if !hasFieldOrder {
			return fieldNames[i] < fieldNames[j]
		}

		iIndex := slices.Index(registryFieldOrder, fieldNames[i])
		jIndex := slices.Index(registryFieldOrder, fieldNames[j])

		if iIndex == -1 && jIndex == -1 {
			return fieldNames[i] < fieldNames[j]
		}

		if iIndex == -1 {
			return false
		}

		if jIndex == -1 {
			return true
		}

		return iIndex < jIndex
	})

	for _, fieldName := range fieldNames {
		fmt.Fprintf(writer, "%s\t", fieldName)
	}
	fmt.Fprintln(writer)

	return nil
}

func (t *textOutputWriter) printFieldValues(v reflect.Value, writer io.Writer) error {
	var fieldNames []string

	valueType := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldName := valueType.Field(i).Name
		omissions, hasOmissions := t.registry.OmitFields[valueType.String()]
		if hasOmissions && slices.Contains(omissions, fieldName) {
			continue
		}

		fieldNames = append(fieldNames, fieldName)
	}

	sort.Slice(fieldNames, func(i, j int) bool {
		registryFieldOrder, hasFieldOrder := t.registry.FieldOrders[valueType.String()]
		if !hasFieldOrder {
			return fieldNames[i] < fieldNames[j]
		}

		iIndex := slices.Index(registryFieldOrder, fieldNames[i])
		jIndex := slices.Index(registryFieldOrder, fieldNames[j])

		if iIndex == -1 && jIndex == -1 {
			return fieldNames[i] < fieldNames[j]
		}

		if iIndex == -1 {
			return false
		}

		if jIndex == -1 {
			return true
		}

		return iIndex < jIndex
	})

	var err error
	for _, fieldName := range fieldNames {
		field := v.FieldByName(fieldName)
		if !field.IsValid() {
			continue
		}

		err = errors.Join(err, t.printFieldValue(field, writer))
		fmt.Fprintf(writer, "\t")
	}
	fmt.Fprintln(writer)

	return nil
}

func (t *textOutputWriter) printFieldValue(v reflect.Value, writer io.Writer) error {
	typeKey := v.Type().String()
	formatter, hasFormatter := t.registry.Formatters[typeKey]
	if hasFormatter {
		fmt.Fprintf(writer, "%s", formatter(v.Interface()))
		return nil
	}

	switch v.Kind() {
	case reflect.String:
		fmt.Fprintf(writer, "%s", v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(writer, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(writer, "%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(writer, "%f", v.Float())
	case reflect.Bool:
		fmt.Fprintf(writer, "%t", v.Bool())
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				fmt.Fprintf(writer, ",")
			}
			err := t.printFieldValue(v.Index(i), writer)
			if err != nil {
				return err
			}
		}
	case reflect.Struct:
		fmt.Fprintf(writer, "{")
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				fmt.Fprintf(writer, ",")
			}
			name := v.Type().Field(i).Name
			fmt.Fprintf(writer, "%s:", name)
			field := v.Field(i)
			err := t.printFieldValue(field, writer)
			if err != nil {
				return err
			}
		}
		fmt.Fprintf(writer, "}")
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Fprintf(writer, "null")
		} else {
			return t.printFieldValue(v.Elem(), writer)
		}
	default:
		if stringer, ok := v.Interface().(fmt.Stringer); ok {
			fmt.Fprintf(writer, "%s", stringer.String())
		} else {
			return fmt.Errorf("unsupported type: %s", v.Kind())
		}
	}
	return nil
}
