package commands

import (
	"fmt"
)

func missingRequiredFlag(flagName string) error {
	return fmt.Errorf("Required flag \"%s\" is not set", flagName)
}
