package output

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"slices"
	"sort"
	"strconv"

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
	FieldOrders        map[string][]string
	OmitFields         map[string][]string
	Formatters         map[string]formatter
	PropertyFormatters map[string]map[string]formatter
	RenameFields       map[string]map[string]string
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

func RenameField[T any](empty T, renameField string, renameTo string) TextFormatterOption {
	return func(r *textFormatterRegistry) {
		typeKey := reflect.TypeOf(empty).String()
		if r.RenameFields == nil {
			r.RenameFields = make(map[string]map[string]string)
		}
		_, hasTypeRenameFields := r.RenameFields[typeKey]
		if !hasTypeRenameFields {
			r.RenameFields[typeKey] = make(map[string]string)
		}
		r.RenameFields[typeKey][renameField] = renameTo
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

func WithPropertyFormatter[P any, T any](emptyParent P, propertyName string, f func(T) string) TextFormatterOption {
	return func(r *textFormatterRegistry) {
		typeKey := reflect.TypeOf(emptyParent).String()
		if r.PropertyFormatters == nil {
			r.PropertyFormatters = make(map[string]map[string]formatter)
		}
		if r.PropertyFormatters[typeKey] == nil {
			r.PropertyFormatters[typeKey] = make(map[string]formatter)
		}
		r.PropertyFormatters[typeKey][propertyName] = func(value any) string {
			return f(value.(T))
		}
	}
}

type textOutputWriter struct {
	registry *textFormatterRegistry
}

var _ FormatWriter = (*textOutputWriter)(nil)

func (t *textOutputWriter) Write(value any, writer io.Writer) error {
	tw := tabwriter.NewWriter(writer, 0, 1, 2, ' ', 0)

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
		renameFields, hasRenameFields := t.registry.RenameFields[valueType.String()]
		if hasRenameFields {
			renameField, hasRenameField := renameFields[fieldName]
			if hasRenameField {
				fieldName = renameField
			}
		}

		fmt.Fprintf(writer, "%s\t", fieldName)
	}
	fmt.Fprintln(writer)

	return nil
}

func (t *textOutputWriter) printFieldValues(v reflect.Value, writer io.Writer) error {
	var propertyNames []string

	valueType := v.Type()
	valueTypeName := valueType.String()
	for i := 0; i < v.NumField(); i++ {
		fieldName := valueType.Field(i).Name
		omissions, hasOmissions := t.registry.OmitFields[valueType.String()]
		if hasOmissions && slices.Contains(omissions, fieldName) {
			continue
		}

		propertyNames = append(propertyNames, fieldName)
	}

	sort.Slice(propertyNames, func(i, j int) bool {
		registryFieldOrder, hasFieldOrder := t.registry.FieldOrders[valueType.String()]
		if !hasFieldOrder {
			return propertyNames[i] < propertyNames[j]
		}

		iIndex := slices.Index(registryFieldOrder, propertyNames[i])
		jIndex := slices.Index(registryFieldOrder, propertyNames[j])

		if iIndex == -1 && jIndex == -1 {
			return propertyNames[i] < propertyNames[j]
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
	for _, propertyName := range propertyNames {
		field := v.FieldByName(propertyName)
		if !field.IsValid() {
			continue
		}

		err = errors.Join(err, t.printFieldValue(valueTypeName, propertyName, field, writer))
		fmt.Fprintf(writer, "\t")
	}
	fmt.Fprintln(writer)

	return err
}

func (t *textOutputWriter) printFieldValue(parentType, propertyName string, v reflect.Value, writer io.Writer) error {
	typeKey := v.Type().String()
	parentFormatter, hasParentFormatter := t.registry.PropertyFormatters[parentType]
	if hasParentFormatter {
		propertyFormatter, hasPropertyFormatter := parentFormatter[propertyName]
		if hasPropertyFormatter {
			fmt.Fprintf(writer, "%s", propertyFormatter(v.Interface()))
			return nil
		}
	}

	typeFormatter, hasTypeFormatter := t.registry.Formatters[typeKey]
	if hasTypeFormatter {
		fmt.Fprintf(writer, "%s", typeFormatter(v.Interface()))
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
		fmt.Fprintf(writer, "%s", strconv.FormatFloat(v.Float(), 'f', -1, 64))
	case reflect.Bool:
		fmt.Fprintf(writer, "%t", v.Bool())
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				fmt.Fprintf(writer, ",")
			}
			elementFieldName := fmt.Sprintf("%s[%d]", propertyName, i)
			err := t.printFieldValue(parentType, elementFieldName, v.Index(i), writer)
			if err != nil {
				return err
			}
		}
	case reflect.Struct:
		fmt.Fprintf(writer, "{")
		typeName := v.Type().String()
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				fmt.Fprintf(writer, ",")
			}
			name := v.Type().Field(i).Name
			fmt.Fprintf(writer, "%s:", name)
			field := v.Field(i)
			err := t.printFieldValue(typeName, propertyName, field, writer)
			if err != nil {
				return err
			}
		}
		fmt.Fprintf(writer, "}")
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Fprintf(writer, "null")
		} else {
			return t.printFieldValue(parentType, propertyName, v.Elem(), writer)
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
