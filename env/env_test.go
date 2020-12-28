package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEnv(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name     string
		args     args
		envExist bool
		want     string
	}{
		{
			name: "ok",
			args: args{
				key:          "SOME_ENV",
				defaultValue: "OTHER_VALUE",
			},
			envExist: true,
			want:     "VALUE",
		},
		{
			name: "no env value",
			args: args{
				key:          "SOME_ENV",
				defaultValue: "OTHER_VALUE",
			},
			want:     "OTHER_VALUE",
			envExist: false,
		},
		{
			name: "missing key argument",
			args: args{
				key: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			if tt.envExist {
				err := os.Setenv(tt.args.key, tt.want)
				rq.NoError(err)

				t.Cleanup(func() {
					err := os.Unsetenv(tt.args.key)
					rq.NoError(err)
				})
			}

			got := GetEnv(tt.args.key, tt.args.defaultValue)
			rq.Equal(tt.want, got)
		})
	}
}

func TestGetRequiredEnv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		args      args
		envExist  bool
		want      string
		wantPanic bool
	}{
		{
			name: "ok",
			args: args{
				key: "SOME_ENV",
			},
			envExist: true,
			want:     "VALUE",
		},
		{
			name: "no env value",
			args: args{
				key: "SOME_ENV",
			},
			want:      "VALUE",
			envExist:  false,
			wantPanic: true,
		},
		{
			name: "missing key argument",
			args: args{
				key: "",
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			if tt.envExist {
				err := os.Setenv(tt.args.key, tt.want)
				rq.NoError(err)

				t.Cleanup(func() {
					err := os.Unsetenv(tt.args.key)
					rq.NoError(err)
				})
			}

			if tt.wantPanic {
				rq.Panics(func() {
					_ = GetRequiredEnv(tt.args.key)
				})
			} else {
				rq.NotPanics(func() {
					got := GetRequiredEnv(tt.args.key)
					rq.Equal(got, tt.want)
				})
			}
		})
	}
}
