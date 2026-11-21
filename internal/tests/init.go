package tests

import "log/slog"

func init() {
	slog.SetDefault(slog.New(slog.DiscardHandler))
}
