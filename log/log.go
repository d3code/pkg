package log

import (
    "github.com/d3code/pkg/common_util"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "log"
    "sync"
)

var Log = getLogger()

var (
    logger     *zap.SugaredLogger
    onceLogger sync.Once
)

func getLogger() *zap.SugaredLogger {
    onceLogger.Do(func() {
        InitLogger()
    })
    return logger
}

func InitLogger() {

    loggerConfig := zap.NewProductionConfig()
    loggerConfig.EncoderConfig = encoderConfig

    if common_util.GetEnvironmentOrDefault("environment", "local") == "local" {
        loggerConfig = zap.NewDevelopmentConfig()
    }

    loggerBuild, err := loggerConfig.Build(zap.AddStacktrace(zapcore.FatalLevel))
    if err != nil {
        log.Fatal(err)
    }

    logger = loggerBuild.Sugar()
}

var encoderConfig = zapcore.EncoderConfig{
    TimeKey:        "time",
    LevelKey:       "severity",
    NameKey:        "logger",
    CallerKey:      "caller",
    MessageKey:     "message",
    StacktraceKey:  "stacktrace",
    LineEnding:     zapcore.DefaultLineEnding,
    EncodeLevel:    encodeLevel(),
    EncodeTime:     zapcore.RFC3339TimeEncoder,
    EncodeDuration: zapcore.MillisDurationEncoder,
    EncodeCaller:   zapcore.ShortCallerEncoder,
}

func encodeLevel() zapcore.LevelEncoder {
    return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
        switch l {
        case zapcore.DebugLevel:
            enc.AppendString("DEBUG")
        case zapcore.InfoLevel:
            enc.AppendString("INFO")
        case zapcore.WarnLevel:
            enc.AppendString("WARNING")
        case zapcore.ErrorLevel:
            enc.AppendString("ERROR")
        case zapcore.DPanicLevel:
            enc.AppendString("CRITICAL")
        case zapcore.PanicLevel:
            enc.AppendString("ALERT")
        case zapcore.FatalLevel:
            enc.AppendString("EMERGENCY")
        }
    }
}
