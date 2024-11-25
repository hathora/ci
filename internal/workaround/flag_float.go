package workaround

import (
	"context"
	"flag"

	"github.com/urfave/cli/v3"
)

type FloatFlag cli.FloatFlag

func (f *FloatFlag) unwrap() *cli.FloatFlag {
	return (*cli.FloatFlag)(f)
}

var _ cli.Flag = &FloatFlag{}
var _ cli.DocGenerationFlag = &FloatFlag{}
var _ cli.RequiredFlag = &FloatFlag{}
var _ cli.VisibleFlag = &FloatFlag{}
var _ cli.ActionableFlag = &FloatFlag{}
var _ cli.DocGenerationMultiValueFlag = &FloatFlag{}
var _ cli.CategorizableFlag = &FloatFlag{}

// Apply implements cli.Flag.
func (f *FloatFlag) Apply(fs *flag.FlagSet) error {
	return f.unwrap().Apply(fs)
}

// IsSet implements cli.Flag.
func (f *FloatFlag) IsSet() bool {
	return f.unwrap().IsSet()
}

// Names implements cli.Flag.
func (f *FloatFlag) Names() []string {
	return f.unwrap().Names()
}

// String implements cli.Flag.
func (f *FloatFlag) String() string {
	return cli.FlagStringer(f)
}

// GetEnvVars implements cli.DocGenerationFlag.
func (f *FloatFlag) GetEnvVars() []string {
	return f.unwrap().GetEnvVars()
}

// GetUsage implements cli.DocGenerationFlag.
func (f *FloatFlag) GetUsage() string {
	return f.unwrap().GetUsage()
}

// GetValue implements cli.DocGenerationFlag.
func (f *FloatFlag) GetValue() string {
	return f.unwrap().GetValue()
}

// TakesValue implements cli.DocGenerationFlag.
func (f *FloatFlag) TakesValue() bool {
	return f.unwrap().TakesValue()
}

// GetDefaultText implements cli.DocGenerationFlag.
func (i *FloatFlag) GetDefaultText() string {
	if i.DefaultText != "" {
		return i.DefaultText
	}
	return ""
}

// IsRequired implements cli.RequiredFlag.
func (f *FloatFlag) IsRequired() bool {
	return f.unwrap().IsRequired()
}

// IsVisible implements cli.VisibleFlag.
func (f *FloatFlag) IsVisible() bool {
	return f.unwrap().IsVisible()
}

// IsDefaultVisible implements cli.DefaultVisibleFlag.
func (f *FloatFlag) IsDefaultVisible() bool {
	return f.unwrap().IsDefaultVisible()
}

// RunAction implements cli.ActionableFlag.
func (f *FloatFlag) RunAction(ctx context.Context, cmd *cli.Command) error {
	return f.unwrap().RunAction(ctx, cmd)
}

// IsMultiValueFlag implements cli.DocGenerationMultiValueFlag.
func (f *FloatFlag) IsMultiValueFlag() bool {
	return f.unwrap().IsMultiValueFlag()
}

// GetCategory implements cli.CategorizableFlag.
func (f *FloatFlag) GetCategory() string {
	return f.unwrap().GetCategory()
}

// SetCategory implements cli.CategorizableFlag.
func (f *FloatFlag) SetCategory(category string) {
	f.unwrap().SetCategory(category)
}
