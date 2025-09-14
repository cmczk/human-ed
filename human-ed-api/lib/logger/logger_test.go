package logger

import (
	"log/slog"
	"os"
	"reflect"
	"testing"
)

func TestSetup(t *testing.T) {
	os.Setenv("CONFIG_PATH", "./config/local.yaml")
	expectedLogType := reflect.TypeOf(&slog.Logger{})

	t.Run("local env config", func(t *testing.T) {
		log := Setup(envLocal)
		_ = log.Handler()

		if gotType := reflect.TypeOf(log); gotType != reflect.TypeOf(&slog.Logger{}) {
			t.Errorf("%s: unexpected type of logger: %v, want: %v", envLocal, gotType, expectedLogType)
		}
	})

	t.Run("dev env config", func(t *testing.T) {
		log := Setup(envLocal)
		_ = log.Handler()

		if gotType := reflect.TypeOf(log); gotType != reflect.TypeOf(&slog.Logger{}) {
			t.Errorf("%s: unexpected type of logger: %v, want: %v", envLocal, gotType, expectedLogType)
		}
	})

	t.Run("prod env config", func(t *testing.T) {
		log := Setup(envLocal)
		_ = log.Handler()

		if gotType := reflect.TypeOf(log); gotType != reflect.TypeOf(&slog.Logger{}) {
			t.Errorf("%s: unexpected type of logger: %v, want: %v", envLocal, gotType, expectedLogType)
		}
	})
}
