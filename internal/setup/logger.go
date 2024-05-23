package setup

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logger(verbosity int) (*zap.Logger, func()) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(LevelFromVerbosity(verbosity)),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}

	logger := zap.Must(config.Build())
	zap.ReplaceGlobals(logger)

	return logger, func() {
		_ = logger.Sync()
	}
}

func LevelFromVerbosity(verbosity int) zapcore.Level {
	switch verbosity {
	case 0:
		return zap.WarnLevel
	case 1:
		return zap.InfoLevel
	case 2:
		return zap.DebugLevel
	case 3:
		return zap.DebugLevel
	default:
		if verbosity < 0 {
			return LevelFromVerbosity(0)
		}
		return LevelFromVerbosity(3)
	}
}
