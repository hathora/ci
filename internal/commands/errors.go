package commands

import (
	"fmt"
)

func missingRequiredFlag(flagName string) error {
	return fmt.Errorf("required flag \"%s\" is not set", flagName)
}
