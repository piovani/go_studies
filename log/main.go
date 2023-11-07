package main

import (
	"log/slog"
	"os"
)

func main() {
	slog.Info("Hello world!")
	slog.Error("Error")
	slog.Warn("Warn")

	// l := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// l.Info("Hello world!")
	// l.Error("Error")
	// l.Warn("Warn")

	// slog.SetDefault(l)

	// slog.Info("Hello world!")
	// slog.Error("Error")
	// slog.Warn("Warn")

	level := new(slog.LevelVar)
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))

	l.Info("Hello world!")
	l.Error("Error")
	l.Warn("Warn")
}
