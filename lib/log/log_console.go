package log

import (
	"os"

	"go.uber.org/zap/zapcore"
)

// NewConsoleAdapter create a console adapter.
func NewConsoleAdapter() Instance {
	return func() zapcore.Core {
		return zapcore.NewCore(
			zapcore.NewJSONEncoder(DefaultEncoderConfig),
			zapcore.AddSync(os.Stdout),
			DefaultLevelEnablerFunc(),
		)
	}
}
