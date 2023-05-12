package zlog

import (
    "github.com/d3code/pkg/cfg"
    "github.com/d3code/pkg/shell"
    "go.uber.org/zap"
    "go.uber.org/zap/buffer"
    "go.uber.org/zap/zapcore"
)

type jsonEncoder struct {
    zapcore.Encoder
    pool buffer.Pool
}

func (e *jsonEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
    buf := e.pool.Get()

    if entry.Level == zapcore.DebugLevel {
        entry.Message = shell.ColorString(entry.Message, "grey")
    }

    fields = append(fields, zap.String("environment", cfg.GetEnvironmentOrDefault("environment", "local")))
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
