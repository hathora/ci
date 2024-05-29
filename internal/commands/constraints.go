package commands

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type StringLike interface {
	~string
}

func requireValidEnumValue[S StringLike](value S, allowedValues []string, enumName string) error {
	if !slices.Contains(allowedValues, string(value)) {
		return fmt.Errorf("flag %s must be one of [%s]", enumName, strings.Join(allowedTransportTypes, ", "))
	}

	return nil
}

type IntLike interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func requireIntInRange[N IntLike](
	value N,
	lower N,
	upper N,
	flagName string,
) error {
	if value < lower || value > upper {
		return fmt.Errorf("flag %s must be between %d and %d", flagName, lower, upper)
	}

	return nil
}

type FloatLike interface {
	~float32 | ~float64
}

func requireFloatInRange[N FloatLike](value N, lower N, upper N, flagName string) error {
	if value < lower || value > upper {
		return fmt.Errorf("flag %s must be between %s and %s",
			flagName,
			strconv.FormatFloat(float64(lower), 'f', -1, 64),
			strconv.FormatFloat(float64(upper), 'f', -1, 64))
	}

	return nil
}

func requireMaxDecimals[N FloatLike](value N, max int, flagName string) error {
	var decimals int
	stringValue := strconv.FormatFloat(float64(value), 'f', -1, 64)
	index := strings.IndexByte(stringValue, '.')
	if index > -1 {
		decimals = len(stringValue) - index - 1
	}

	if decimals > max {
		return fmt.Errorf("flag %s must have a maximum of %d decimal point(s)", flagName, max)
	}

	return nil
}
