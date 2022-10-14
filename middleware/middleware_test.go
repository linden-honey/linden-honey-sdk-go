package middleware

import (
	"reflect"
	"testing"
)

type testService func(calls *[]string) []string

func testMiddleware1() Middleware[testService] {
	return func(next testService) testService {
		return func(calls *[]string) []string {
			*calls = append(*calls, "testMiddleware1")

			return next(calls)
		}
	}
}

func testMiddleware2() Middleware[testService] {
	return func(next testService) testService {
		return func(calls *[]string) []string {
			*calls = append(*calls, "testMiddleware2")

			return next(calls)
		}
	}
}

func TestCompose(t *testing.T) {
	type args struct {
		mw []Middleware[testService]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ok",
			args: args{
				mw: []Middleware[testService]{
					testMiddleware1(),
					testMiddleware2(),
				},
			},
			want: []string{
				"testMiddleware1",
				"testMiddleware2",
				"testService",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := make([]string, 0)
			svc := func(calls *[]string) []string {
				*calls = append(*calls, "testService")

				return *calls
			}
			if got := Compose(tt.args.mw...)(svc)(&calls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compose(tt.args.mw...)(svc)(&calls) = %v, want %v", got, tt.want)
			}
		})
	}
}
