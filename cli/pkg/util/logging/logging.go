package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewNamedLogger(name string) *zap.SugaredLogger {
	return zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.Lock(os.Stdout),
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return lvl >= zapcore.DebugLevel }),
		),
	).Named(fmt.Sprintf("[%s]", name)).Sugar()
}
