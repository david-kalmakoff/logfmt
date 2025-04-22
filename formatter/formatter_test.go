package formatter_test

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/david-kalmakoff/logfmt/formatter"
	"github.com/stretchr/testify/require"
)

func TestFormatter(t *testing.T) {
	var inB []byte
	inBuf := bytes.NewBuffer(inB)

	var outB []byte
	outBuf := bytes.NewBuffer(outB)

	go formatter.New(inBuf, outBuf)

	h := slog.NewJSONHandler(inBuf, &slog.HandlerOptions{})
	l := slog.New(h)
	slog.SetDefault(l)

	slog.Log(context.Background(), slog.LevelInfo, "message")

	time.Sleep(time.Second)

	msg1 := fmt.Sprintf("%sINFO%s => ", formatter.Green, formatter.Reset)
	msg2 := " \"message\"\n"
	require.Contains(t, outBuf.String(), msg1)
	require.Contains(t, outBuf.String(), msg2)
}
