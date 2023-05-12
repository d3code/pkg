package zlog

import (
    "github.com/d3code/pkg/shell"
    "go.uber.org/zap/buffer"
    "go.uber.org/zap/zapcore"
    "time"
)

type consoleEncoder struct {
    zapcore.Encoder
    pool buffer.Pool
}

func (e *consoleEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
    buf := e.pool.Get()

    if entry.Level == zapcore.DebugLevel {
        entry.Message = shell.ColorString(entry.Message, "grey")
    }

    entry.Time = entry.Time.Local()

    consoleBuffer, err := e.Encoder.EncodeEntry(entry, fields)
    if err != nil {
        return nil, err
    }

    _, err = buf.Write(consoleBuffer.Bytes())
    if err != nil {
        return nil, err
    }

    return buf, nil
}

func encodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    format := t.Format(time.RFC3339)
    greyTime := shell.ColorString(format, "grey")
    enc.AppendString(greyTime)
}

func encodeLevelColor() zapcore.LevelEncoder {
    debug := shell.ColorString("DEBUG", "grey")
    info := shell.ColorString("INFO", "blue")
    warning := shell.ColorString("WARNING", "yellow")
    errorLevel := shell.ColorString("ERROR", "red")
    critical := shell.ColorString("CRITICAL", "red")
    alert := shell.ColorString("ALERT", "red")
    emergency := shell.ColorString("EMERGENCY", "red")

    return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {

        switch l {
        case zapcore.DebugLevel:
            enc.AppendString(debug)
        case zapcore.InfoLevel:
            enc.AppendString(info)
        case zapcore.WarnLevel:
            enc.AppendString(warning)
        case zapcore.ErrorLevel:
            enc.AppendString(errorLevel)
        case zapcore.DPanicLevel:
            enc.AppendString(critical)
        case zapcore.PanicLevel:
            enc.AppendString(alert)
        case zapcore.FatalLevel:
            enc.AppendString(emergency)
        }
    }
}
