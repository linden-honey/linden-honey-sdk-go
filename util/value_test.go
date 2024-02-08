package util

import (
	"reflect"
	"testing"
)

func TestPtr(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "ok",
			args: args{
				v: "test",
			},
			want: Ptr("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ptr(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVal(t *testing.T) {
	type args struct {
		p *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				p: Ptr("test"),
			},
			want: "test",
		},
		{
			name: "nil pointer",
			args: args{
				p: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Val(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Val() = %v, want %v", got, tt.want)
			}
		})
	}
}
