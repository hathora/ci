package workaround

import (
	"context"
	"flag"

	"github.com/urfave/cli/v3"
)

type IntFlag cli.IntFlag

func (i *IntFlag) unwrap() *cli.IntFlag {
	return (*cli.IntFlag)(i)
}

var _ cli.Flag = &IntFlag{}
var _ cli.DocGenerationFlag = &IntFlag{}
var _ cli.RequiredFlag = &IntFlag{}
var _ cli.VisibleFlag = &IntFlag{}
var _ cli.ActionableFlag = &IntFlag{}
var _ cli.DocGenerationMultiValueFlag = &IntFlag{}
var _ cli.CategorizableFlag = &IntFlag{}

// Apply implements cli.Flag.
func (i *IntFlag) Apply(fs *flag.FlagSet) error {
	return i.unwrap().Apply(fs)
}

// IsSet implements cli.Flag.
func (i *IntFlag) IsSet() bool {
	return i.unwrap().IsSet()
}

// Names implements cli.Flag.
func (i *IntFlag) Names() []string {
	return i.unwrap().Names()
}

// String implements cli.Flag.
func (i *IntFlag) String() string {
	return cli.FlagStringer(i)
}

// GetEnvVars implements cli.DocGenerationFlag.
func (i *IntFlag) GetEnvVars() []string {
	return i.unwrap().GetEnvVars()
}

// GetUsage implements cli.DocGenerationFlag.
func (i *IntFlag) GetUsage() string {
	return i.unwrap().GetUsage()
}

// GetValue implements cli.DocGenerationFlag.
func (i *IntFlag) GetValue() string {
	return i.unwrap().GetValue()
}

// TakesValue implements cli.DocGenerationFlag.
func (i *IntFlag) TakesValue() bool {
	return i.unwrap().TakesValue()
}

// GetDefaultText implements cli.DocGenerationFlag.
func (i *IntFlag) GetDefaultText() string {
	if i.DefaultText != "" {
		return i.DefaultText
	}
	return ""
}

// IsRequired implements cli.RequiredFlag.
func (i *IntFlag) IsRequired() bool {
	return i.unwrap().IsRequired()
}

// IsVisible implements cli.VisibleFlag.
func (i *IntFlag) IsVisible() bool {
	return i.unwrap().IsVisible()
}

// IsPersistent implements cli.DefaultVisibleFlag
func (i *IntFlag) IsDefaultVisible() bool {
	return i.unwrap().IsDefaultVisible()
}

// GetCategory implements cli.CategorizableFlag.
func (i *IntFlag) GetCategory() string {
	return i.unwrap().GetCategory()
}

// SetCategory implements cli.CategorizableFlag.
func (i *IntFlag) SetCategory(category string) {
	i.unwrap().SetCategory(category)
}

// IsMultiValueFlag implements cli.DocGenerationMultiValueFlag.
func (i *IntFlag) IsMultiValueFlag() bool {
	return i.unwrap().IsMultiValueFlag()
}

// RunAction implements cli.ActionableFlag.
func (i *IntFlag) RunAction(ctx context.Context, cmd *cli.Command) error {
	return i.unwrap().RunAction(ctx, cmd)
}
