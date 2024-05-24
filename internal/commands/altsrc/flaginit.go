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

func InitializeValueSourcesFromFlags(ctx context.Context, cmd *cli.Command) error {
	var err error
	for _, flag := range cmd.Flags {
		sources := getSourcesFromFlag(flag)
		for _, source := range sources {
			if flagInitialized, ok := source.(FlagInitializedValueSource); ok {
				err = errors.Join(err, flagInitialized.Initialize(cmd))
				if err != nil {
					continue
				}
				for _, name := range flag.Names() {
					if value, ok := flagInitialized.Lookup(); ok {
						err = errors.Join(err, cmd.Set(name, value))
					}
				}
			}
		}
	}

	for _, sub := range cmd.Commands {
		oldBefore := sub.Before
		sub.Before = func(ctx context.Context, cmd *cli.Command) error {
			err := InitializeValueSourcesFromFlags(ctx, cmd)
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
