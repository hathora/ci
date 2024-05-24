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
	if strings.HasPrefix(s, "value=") {
		return Value
	}

	switch strings.ToLower(s) {
	case "json":
		return JSON
	case "text":
		return Text
	default:
		return UnknownType
	}
}
