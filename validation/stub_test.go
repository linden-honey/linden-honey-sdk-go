package validation

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidatorStub_Validate(t *testing.T) {
	type fields struct {
		Err error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Err: errors.New("some error"),
			},
			wantErr: true,
		},
		{
			name: "validation failed",
			fields: fields{
				Err: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			vs := ValidatorStub{
				Err: tt.fields.Err,
			}
			err := vs.Validate()

			if tt.wantErr {
				rq.Error(err)
				rq.Equal(tt.fields.Err, err)
			} else {
				rq.NoError(err)
			}
		})
	}
}
