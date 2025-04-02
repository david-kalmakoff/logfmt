package main

import (
	"context"
	"log/slog"
	"os"
	"time"
)

func main() {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	l := slog.New(h)
	slog.SetDefault(l)

	levels := []slog.Level{slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError}
	i := 0

	for {
		slog.Log(context.Background(), levels[i], "message", "i", i)

		i = (i + 1) % len(levels)

		time.Sleep(time.Second * 1)
	}
}
