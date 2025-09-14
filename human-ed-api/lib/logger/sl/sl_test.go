package sl

import (
	"errors"
	"log/slog"
	"reflect"
	"testing"
)

func TestErr_RegularErrors(t *testing.T) {
	tests := []struct {
		name     string
		inputErr error
		want     slog.Attr
	}{
		{
			name:     "valid regular error",
			inputErr: errors.New("valid error with regular text message"),
			want: slog.Attr{
				Key:   "error",
				Value: slog.StringValue("valid error with regular text message"),
			},
		},
	}

	for _, tt := range tests {
		res := Err(tt.inputErr)

		if res.Key != "error" || !reflect.DeepEqual(res.Value, slog.StringValue(tt.inputErr.Error())) {
			t.Errorf("invalid slog attr format! got: {Key: %s, Value: %s}, want: {Key: %s, Value: %s}", res.Key, res.Value, "error", tt.inputErr.Error())
		}
	}
}

func TestErr_NilError(t *testing.T) {
	test := struct {
		name     string
		inputErr error
		want     slog.Attr
	}{
		name:     "nil error",
		inputErr: nil,
		want: slog.Attr{
			Key:   "error",
			Value: slog.StringValue(""),
		},
	}

	res := Err(test.inputErr)

	if res.Key != "error" || !reflect.DeepEqual(res.Value, slog.StringValue("")) {
		t.Errorf("%s: invalid slog attr format! got: %v, want: %v", test.name, res, test.want)
	}
}
