package altsrc

import (
	"context"
	"errors"

	"github.com/urfave/cli/v3"
)

type FlagInitializedValueSource interface {
	cli.ValueSource
	Initialize(*cli.Command) error
}

func InitializeValueSourcesFromFlags(ctx context.Context, cmd *cli.Command, args []string) error {
	var err error
	for _, flag := range cmd.Flags {
		err = errors.Join(err, initializeFlagSources(cmd, flag, args))
	}

	for _, sub := range cmd.Commands {
		oldBefore := sub.Before
		sub.Before = func(ctx context.Context, cmd *cli.Command) error {
			err := InitializeValueSourcesFromFlags(ctx, cmd, args)
			if err != nil {
				return err
			}

			if oldBefore != nil {
				return oldBefore(ctx, cmd)
			}

			return nil
		}
	}

	return err
}

func initializeFlagSources(cmd *cli.Command, flag cli.Flag, args []string) error {
	var err error
	sources := getSourcesFromFlag(flag)

	// arg is explicitly set on the command line
	for _, flagName := range flag.Names() {
		shortName := "-" + flagName
		longName := "--" + flagName
		for _, arg := range args {
			if arg == shortName || arg == longName {
				return nil
			}
		}
	}

	firstSourceValuePrecident := len(sources) + 1
	var firstPrecidentSourceValue string
	firstFlagInitializedValuePrecident := len(sources) + 1
	var firstPrecidentFlagInitializedValue string

	for i, source := range sources {
		if _, alreadyResolved := source.Lookup(); alreadyResolved {
			if firstPrecidentSourceValue == "" {
				firstSourceValuePrecident = i
			}
			break
		}
		// don't keep looking for flag initialized values if a higher precident value is already set
		if firstPrecidentFlagInitializedValue != "" {
			continue
		}
		if flagInitialized, ok := source.(FlagInitializedValueSource); ok {
			initErr := flagInitialized.Initialize(cmd)
			if initErr != nil {
				err = errors.Join(err, initErr)
				continue
			}
			if value, ok := flagInitialized.Lookup(); ok {
				firstPrecidentFlagInitializedValue = value
				firstFlagInitializedValuePrecident = i
			}
		}
	}

	// if no flag initialized values are set, then there is nothing to do
	if firstPrecidentFlagInitializedValue == "" {
		return nil
	}

	// if the highest precident source value is before the highest precident flag initialized value
	//   then the source value should win
	if firstSourceValuePrecident < firstFlagInitializedValuePrecident {
		return nil
	}

	// if the highest precident flag initialized value is before the highest precident source value
	//   then the flag initialized value should win
	for _, name := range flag.Names() {
		err = errors.Join(err, cmd.Set(name, firstPrecidentFlagInitializedValue))
	}

	return err
}

func getSourcesFromFlag(flag cli.Flag) []cli.ValueSource {
	var sources []cli.ValueSource

	if flag, ok := flag.(*cli.StringFlag); ok {
		if flag.Sources.Chain != nil {
			sources = append(sources, flag.Sources.Chain...)
		}
	}

	if flag, ok := flag.(*cli.BoolFlag); ok {
		if flag.Sources.Chain != nil {
			sources = append(sources, flag.Sources.Chain...)
		}
	}

	if flag, ok := flag.(*cli.BoolWithInverseFlag); ok {
		if flag.Sources.Chain != nil {
			sources = append(sources, flag.Sources.Chain...)
		}
	}

	if flag, ok := flag.(*cli.IntFlag); ok {
		if flag.Sources.Chain != nil {
			sources = append(sources, flag.Sources.Chain...)
		}
	}

	if flag, ok := flag.(*cli.UintFlag); ok {
		if flag.Sources.Chain != nil {
			sources = append(sources, flag.Sources.Chain...)
		}
	}

	if flag, ok := flag.(*cli.FloatFlag); ok {
		if flag.Sources.Chain != nil {
			sources = append(sources, flag.Sources.Chain...)
		}
	}

	if flag, ok := flag.(*cli.DurationFlag); ok {
		if flag.Sources.Chain != nil {
			sources = append(sources, flag.Sources.Chain...)
		}
	}

	if flag, ok := flag.(*cli.TimestampFlag); ok {
		if flag.Sources.Chain != nil {
			sources = append(sources, flag.Sources.Chain...)
		}
	}

	if flag, ok := flag.(*cli.StringSliceFlag); ok {
		if flag.Sources.Chain != nil {
			sources = append(sources, flag.Sources.Chain...)
		}
	}

	if flag, ok := flag.(*cli.IntSliceFlag); ok {
		if flag.Sources.Chain != nil {
			sources = append(sources, flag.Sources.Chain...)
		}
	}

	if flag, ok := flag.(*cli.UintSliceFlag); ok {
		if flag.Sources.Chain != nil {
			sources = append(sources, flag.Sources.Chain...)
		}
	}

	if flag, ok := flag.(*cli.FloatSliceFlag); ok {
		if flag.Sources.Chain != nil {
			sources = append(sources, flag.Sources.Chain...)
		}
	}

	if flag, ok := flag.(*cli.StringMapFlag); ok {
		if flag.Sources.Chain != nil {
			sources = append(sources, flag.Sources.Chain...)
		}
	}

	return sources
}
