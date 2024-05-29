package commands

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func requireValidEnumValue(value string, allowedValues []string, enumName string) error {
	if !slices.Contains(allowedValues, value) {
		return fmt.Errorf("flag %s must be one of [%s]", enumName, strings.Join(allowedTransportTypes, ", "))
	}

	return nil
}

func requireIntInRange(value int64, lower int64, upper int64, flagName string) error {
	if value < lower || value > upper {
		return fmt.Errorf("flag %s must be between %d and %d", flagName, lower, upper)
	}

	return nil
}

func requireFloatInRange(value float64, lower float64, upper float64, flagName string) error {
	if value < lower || value > upper {
		return fmt.Errorf("flag %s must be between %s and %s",
			flagName,
			strconv.FormatFloat(lower, 'f', -1, 64),
			strconv.FormatFloat(upper, 'f', -1, 64))
	}

	return nil
}

func requireMaxDecimals(value float64, max int, flagName string) error {
	var decimals int
	stringValue := strconv.FormatFloat(value, 'f', -1, 64)
	index := strings.IndexByte(stringValue, '.')
	if index > -1 {
		decimals = len(stringValue) - index - 1
	}

	if decimals > max {
		return fmt.Errorf("flag %s must have a maximum of %d decimal point(s)", flagName, max)
	}

	return nil
}