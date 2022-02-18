package validation

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDelegate(t *testing.T) {
	tests := []struct {
		name    string
		want    *Delegate
		wantErr bool
	}{
		{
			name:    "ok",
			want:    &Delegate{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			got, err := NewDelegate()
			if tt.wantErr {
				rq.Error(err)
			} else {
				rq.NoError(err)
				rq.Equal(tt.want, got)
			}
		})
	}
}

func TestDelegate_Validate(t *testing.T) {
	type args struct {
		in Validator
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				in: ValidatorStub{
					Err: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "validation failed",
			args: args{
				in: ValidatorStub{
					Err: errors.New("some error"),
				},
			},
			wantErr: true,
		},
		{
			name: "nil input",
			args: args{
				in: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			d := &Delegate{}

			err := d.Validate(tt.args.in)

			if tt.wantErr {
				rq.Error(err)
			} else {
				rq.NoError(err)
			}
		})
	}
}
