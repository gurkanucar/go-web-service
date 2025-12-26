package logger

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ctxKey string

const TraceIDKey ctxKey = "trace_id"

type TraceHandler struct {
	slog.Handler
}

func (h TraceHandler) Handle(ctx context.Context, r slog.Record) error {
	if traceID, ok := ctx.Value(TraceIDKey).(string); ok {
		r.AddAttrs(slog.String("trace_id", traceID))
	}
	return h.Handler.Handle(ctx, r)
}

func (h TraceHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return TraceHandler{h.Handler.WithAttrs(attrs)}
}

func (h TraceHandler) WithGroup(name string) slog.Handler {
	return TraceHandler{h.Handler.WithGroup(name)}
}

func New(isProd bool) *slog.Logger {
	level := slog.LevelDebug
	if isProd {
		level = slog.LevelInfo
	}

	cwd, _ := os.Getwd()
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				if src, ok := a.Value.Any().(*slog.Source); ok {
					rel := strings.TrimPrefix(src.File, cwd+string(filepath.Separator))
					a.Value = slog.StringValue(rel + ":" + strconv.Itoa(src.Line))
				}
			}
			return a
		},
	}

	return slog.New(TraceHandler{slog.NewJSONHandler(os.Stdout, opts)})
}
