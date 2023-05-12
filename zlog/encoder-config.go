package zlog

import "go.uber.org/zap/zapcore"

var encoderConfig = zapcore.EncoderConfig{
    TimeKey:        "time",
    LevelKey:       "severity",
    NameKey:        "logger",
    CallerKey:      "caller",
    MessageKey:     "message",
    StacktraceKey:  "stacktrace",
    LineEnding:     zapcore.DefaultLineEnding,
    EncodeTime:     zapcore.RFC3339TimeEncoder,
    EncodeDuration: zapcore.MillisDurationEncoder,
    EncodeCaller:   zapcore.ShortCallerEncoder,
    EncodeLevel: func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
        switch level {
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
    },
}
