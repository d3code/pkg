package log

import (
    "github.com/d3code/pkg/cfg"
    "go.uber.org/zap"
    "go.uber.org/zap/buffer"
    "go.uber.org/zap/zapcore"
    "os"
)

var Log *zap.Logger

func init() {

    if cfg.GetEnvironmentOrDefault("environment", "local") == "local" {
        config := encoderConfig
        config.EncodeTime = encodeTime
        config.EncodeLevel = encodeLevelColor()

        encoder := &consoleEncoder{
            Encoder: zapcore.NewConsoleEncoder(config),
            pool:    buffer.NewPool(),
        }
        Log = zap.New(
            zapcore.NewCore(
                encoder,
                os.Stdout,
                zapcore.DebugLevel,
            ),
            zap.ErrorOutput(os.Stderr),
            zap.AddStacktrace(zapcore.FatalLevel),
            zap.AddCaller(),
        )
    } else {
        encoder := &jsonEncoder{
            Encoder: zapcore.NewJSONEncoder(encoderConfig),
            pool:    buffer.NewPool(),
        }
        Log = zap.New(
            zapcore.NewCore(
                encoder,
                os.Stdout,
                zapcore.DebugLevel,
            ),
            zap.ErrorOutput(os.Stderr),
            zap.AddStacktrace(zapcore.FatalLevel),
            zap.AddCaller(),
        )
    }
}
