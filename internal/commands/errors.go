package commands

import (
	"fmt"
	"strconv"
)

func missingRequiredFlag(flagName string) error {
	return fmt.Errorf("Required flag \"%s\" is not set", flagName)
}

func invalidMemoryToCPURatio(memoryMB, cpu float64) error {
	return fmt.Errorf("invalid memory to CPU ratio of %s to %s; %s and %s must be in a %s:1 ratio",
		strconv.FormatFloat(memoryMB, 'f', -1, 64),
		strconv.FormatFloat(cpu, 'f', -1, 64),
		requestedMemoryFlag.Name,
		requestedCPUFlag.Name,
		strconv.FormatFloat(memoryMBPerCPU, 'f', -1, 64),
	)
}
