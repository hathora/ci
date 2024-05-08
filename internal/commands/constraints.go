package commands

import (
	"fmt"
	"slices"
	"strings"
)

func requireValidEnumValue(value string, allowedValues []string, enumName string) error {
	if !slices.Contains(allowedValues, value) {
		return fmt.Errorf("flag %s must be one of [%s]", enumName, strings.Join(allowedTransportTypes, ", "))
	}
	
	return nil
}

func requireIntInRange(value int, lower int, upper int, flagName string) error {
	if value < lower || value > upper {
		return fmt.Errorf("flag %s must be between %d and %d", flagName, lower, upper)
	}
	
	return nil
}