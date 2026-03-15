package main

import (
	"go.uber.org/zap"
	"log/slog"
)

func main() {
	// correct
	slog.Info("starting server")
	slog.Error("failed to connect")
	zap.L().Info("all good")

	slog.Info("Starting server") // want "log message should start with a lowercase letter"
	slog.Error("Failed")         // want "log message should start with a lowercase letter"

	slog.Info("Запуск сервера") // want "log message should use only English letters" "log message should start with a lowercase letter"
	slog.Error("Ошибка")        // want "log message should use only English letters" "log message should start with a lowercase letter"

	slog.Info("server started!")  // want "log message should not contain special characters"
	slog.Warn("warning: message") // want "log message should not contain special characters"
	slog.Error("...failed")       // want "log message should not contain special characters"
	slog.Debug("emoji 😀")         // want "log message should not contain emoji"

	slog.Info("user password is 123") // want "log message may contain sensitive data: \"password\""
	slog.Debug("key is abc")          // want "log message may contain sensitive data: \"key\""
	zap.L().Info("token is xyz")      // want "log message may contain sensitive data: \"token\""
}
