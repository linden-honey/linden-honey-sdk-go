package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

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
			want: "'Title' has invalid value: title should be in camelcase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			got := NewInvalidValueError(tt.args.key, tt.args.err)
			rq.NotNil(got)
			rq.Equal(
				tt.want,
				got.Error(),
			)
		})
	}
}

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
			want: "'Title' has invalid value: required value is missing",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			got := NewRequiredValueError(tt.args.key)
			rq.NotNil(got)
			rq.Equal(
				tt.want,
				got.Error(),
			)
		})
	}
}
