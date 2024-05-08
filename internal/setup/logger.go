package setup

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logger() (*zap.Logger, func()) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
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
