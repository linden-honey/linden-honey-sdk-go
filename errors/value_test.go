package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRequiredValueError(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				key: "Title",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			err := NewRequiredValueError(tt.args.key)
			rq.NotNil(err)
			rq.True(errors.Is(err, RequiredValueError))
			rq.EqualError(
				err,
				fmt.Errorf(
					"'%s' has invalid value: %w",
					tt.args.key,
					RequiredValueError,
				).Error(),
			)
		})
	}
}

func TestNewInvalidValueError(t *testing.T) {
	type args struct {
		key string
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				key: "Title",
				err: errors.New("title should be in camelcase"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			err := NewInvalidValueError(tt.args.key, tt.args.err)
			rq.NotNil(err)
			rq.EqualError(
				err,
				fmt.Errorf(
					"'%s' has invalid value: %w",
					tt.args.key,
					tt.args.err,
				).Error(),
			)
		})
	}
}
