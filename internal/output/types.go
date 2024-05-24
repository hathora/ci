package output

import "strings"

type Type int

const (
	UnknownType Type = iota
	JSON
	Text
	Value
)

func (o Type) String() string {
	switch o {
	case JSON:
		return "json"
	case Text:
		return "text"
	default:
		return "unknown"
	}
}

func ParseOutputType(s string) Type {
	lowercaseOutputType := strings.ToLower(s)
	if strings.HasPrefix(lowercaseOutputType, "value=") {
		return Value
	}

	switch lowercaseOutputType {
	case "json":
		return JSON
	case "text":
		return Text
	default:
		return UnknownType
	}
}
