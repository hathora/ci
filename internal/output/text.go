package output

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"

	"text/tabwriter"
)

func TextFormat() FormatWriter {
	return &textOutputWriter{}
}

type textOutputWriter struct{}

var _ FormatWriter = (*textOutputWriter)(nil)

func (t *textOutputWriter) Write(value any, writer io.Writer) error {
	tw := tabwriter.NewWriter(writer, 0, 0, 2, ' ', 0)

	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// if we fail to marshal into a single map, try again as an array of maps
	fieldMap := map[string]any{}
	err = json.Unmarshal(jsonBytes, &fieldMap)
	if err != nil {
		var fieldMapArray []map[string]any
		err = json.Unmarshal(jsonBytes, &fieldMapArray)
		if err != nil {
			return err
		}
		keys := getKeys(fieldMapArray[0])

		err = printArray(fieldMapArray, keys, tw)
		if err != nil {
			return err
		}
	}

	keys := getKeys(fieldMap)

	err = printSingleton(fieldMap, keys, tw)
	if err != nil {
		return err
	}

	err = tw.Flush()
	return err
}

func printArray(arr []map[string]any, keys []string, writer io.Writer) error {
	err := printHeader(keys, writer)
	if err != nil {
		return err
	}

	for _, row := range arr {
		err := printRow(row, keys, writer)
		if err != nil {
			return err
		}
	}

	return nil
}

func printSingleton(m map[string]any, keys []string,  writer io.Writer) error {
	err := printHeader(keys, writer)
	if err != nil {
		return err
	}

	err = printRow(m, keys, writer)
	return err
}

func printHeader(keys []string, writer io.Writer) error {
	_, err := fmt.Fprintln(writer, strings.Join(keys[:], "\t"))

	return err
}

func printRow(m map[string]any, keys []string, writer io.Writer) error {
	values := make([]string, 0, len(m))

	for  _, key := range keys {
		values = append(values, fmt.Sprintf("%v", m[key]))
	}

	_, err := fmt.Fprintln(writer, strings.Join(values[:], "\t"))

	return err
}

func getKeys(m map[string]any) []string {
	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	return keys
}