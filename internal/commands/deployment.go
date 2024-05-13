package commands

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	allowedTransportTypes = []string{"tcp", "udp", "tls"}
	maxRoomsPerProcess    = 10000
	maxPort               = 65535
)

var Deployment = &cli.Command{
	Name:  "deployment",
	Usage: "options for deployments",
	Subcommands: []*cli.Command{
		{
			Name:  "get-deployment-info",
			Usage: "get a deployment by id",
			Flags: append([]cli.Flag{
				deploymentIDFlag,
			}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				zap.L().Info("Getting a deployment...")
				return nil
			},
		},
		{
			Name:  "get-latest-deployment",
			Usage: "get the latest deployment",
			Flags: append([]cli.Flag{}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				zap.L().Info("Getting the latest deployment...")
				return nil
			},
		},
		{
			Name:  "get-deployments",
			Usage: "get all deployments",
			Flags: append([]cli.Flag{		}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				zap.L().Info("Getting all deployments...")
				return nil
			},
		},
		{
			Name:  "create-deployment",
			Usage: "create a deployment",
			Flags: append([]cli.Flag{
				buildIDFlag,
				&cli.BoolFlag{
					Name:     "idle-timeout-enabled",
					EnvVars:  []string{"HATHORA_DEPLOYMENT_IDLE_TIMEOUT_ENABLED"},
					Usage:    "option to shut down processes that have had no new connections or rooms for five minutes",
					Required: true,
				},
				&cli.IntFlag{
					Name:     "rooms-per-process",
					EnvVars:  []string{"HATHORA_DEPLOYMENT_ROOMS_PER_PROCESS"},
					Usage:    "how many rooms can be scheduled in a process",
					Required: true,
					Action: func(ctx *cli.Context, v int) error {
						return requireIntInRange(v, 1, maxRoomsPerProcess, "rooms-per-process")
					},
				},
				&cli.StringFlag{
					Name:     "transport-type",
					EnvVars:  []string{"HATHORA_DEPLOYMENT_TRANSPORT_TYPE"},
					Usage:    "the underlying communication protocol to the exposed port",
					Required: true,
					Action: func(ctx *cli.Context, v string) error {
						return requireValidEnumValue(v, allowedTransportTypes, "transport-type")
					},
				},
				&cli.IntFlag{
					Name:     "container-port",
					EnvVars:  []string{"HATHORA_DEPLOYMENT_CONTAINER_PORT"},
					Usage:    "default server port",
					Required: true,
					Action: func(ctx *cli.Context, v int) error {
						return requireIntInRange(v, 1, maxPort, "container-port")
					},
				},
				&cli.Float64Flag{
					Name:     "requested-memory-mb",
					EnvVars:  []string{"HATHORA_DEPLOYMENT_REQUESTED_MEMORY_MB"},
					Usage:    "the amount of memory allocated to your process in MB",
					Required: true,
				},
				&cli.Float64Flag{
					Name:     "requested-cpu",
					EnvVars:  []string{"HATHORA_DEPLOYMENT_REQUESTED_CPU"},
					Usage:    "the number of cores allocated to your process",
					Required: true,
				},
				// TODO additional container ports
				// TODO env vars
			}, globalFlags...),
			Action: func(cCtx *cli.Context) error {
				zap.L().Info("Creating a deployment...")
				return nil
			},
		},
	},
}
