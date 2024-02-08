package errors

import (
	"errors"
	"testing"
)

func TestNewRequiredValueError(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				key: "Title",
			},
			want: "\"Title\" is required",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequiredValueError(tt.args.key); got != nil && got.Error() != tt.want {
				t.Errorf("NewRequiredValueError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInvalidValueError(t *testing.T) {
	type args struct {
		key string
		err error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				key: "Title",
				err: errors.New("title should be in camelcase"),
			},
			want: "\"Title\" is invalid: title should be in camelcase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInvalidValueError(tt.args.key, tt.args.err); got != nil && got.Error() != tt.want {
				t.Errorf("NewInvalidValueError() = %v, want %v", got, tt.want)
			}
		})
	}
}
