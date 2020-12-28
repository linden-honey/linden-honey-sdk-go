package errors

import (
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
			rq.Equal(err.Error(), fmt.Sprintf("'%s' is required", tt.args.key))
		})
	}
}

func TestNewInvalidValueError(t *testing.T) {
	type args struct {
		key   string
		value interface{}
		rule  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				key:   "Title",
				value: "some title",
				rule:  "title should be in camelcase",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			err := NewInvalidValueError(tt.args.key, tt.args.value, tt.args.rule)
			rq.NotNil(err)
			rq.Equal(
				err.Error(),
				fmt.Sprintf(
					"'%s' has invalid value '%v' - %s",
					tt.args.key,
					tt.args.value,
					tt.args.rule,
				),
			)
		})
	}
}
