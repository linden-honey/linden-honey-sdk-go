package pagination

import (
	"reflect"
	"testing"
)

func TestNewPageable(t *testing.T) {
	type args struct {
		opts []PageableOption
	}
	tests := []struct {
		name    string
		args    args
		want    *Pageable
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				opts: []PageableOption{
					WithLimit(0),
					WithOffset(0),
					WithSort(make(Sort)),
				},
			},
			want: &Pageable{
				Limit:  0,
				Offset: 0,
				Sort:   make(Sort),
			},
		},
		{
			name: "err  invalid limit",
			args: args{
				opts: []PageableOption{
					WithLimit(-1),
					WithOffset(0),
					WithSort(make(Sort)),
				},
			},
			wantErr: true,
		},
		{
			name: "err  invalid offset",
			args: args{
				opts: []PageableOption{
					WithLimit(0),
					WithOffset(-1),
					WithSort(make(Sort)),
				},
			},
			wantErr: true,
		},
		{
			name: "err  invalid sort",
			args: args{
				opts: []PageableOption{
					WithLimit(0),
					WithOffset(-1),
					WithSort(map[SortBy]SortOrder{
						"id": 2,
					}),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPageable(tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPageable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPageable() = %v, want %v", got, tt.want)
			}
		})
	}
}
