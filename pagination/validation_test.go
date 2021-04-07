package pagination

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortBy_Validate(t *testing.T) {
	tests := []struct {
		name    string
		sb      SortBy
		wantErr bool
	}{
		{
			name: "ok",
			sb:   "name",
		},
		{
			name:    "err  empty",
			sb:      "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			err := tt.sb.Validate()
			if tt.wantErr {
				rq.Error(err)
			} else {
				rq.NoError(err)
			}
		})
	}
}

func TestSortOrder_Validate(t *testing.T) {
	tests := []struct {
		name    string
		so      SortOrder
		wantErr bool
	}{
		{
			name: "ok",
			so:   ASC,
		},
		{
			name:    "err  invalid value",
			so:      2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			err := tt.so.Validate()
			if tt.wantErr {
				rq.Error(err)
			} else {
				rq.NoError(err)
			}
		})
	}
}

func TestSort_Validate(t *testing.T) {
	tests := []struct {
		name    string
		sort    Sort
		wantErr bool
	}{
		{
			name: "ok",
			sort: Sort{
				"id":    ASC,
				"title": ASC,
			},
		},
		{
			name: "err  invalid sort by",
			sort: Sort{
				"":      ASC,
				"title": ASC,
			},
			wantErr: true,
		},
		{
			name: "err  invalid sort order",
			sort: Sort{
				"id":    2,
				"title": ASC,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			err := tt.sort.Validate()
			if tt.wantErr {
				rq.Error(err)
			} else {
				rq.NoError(err)
			}
		})
	}
}

func TestPageable_Validate(t *testing.T) {
	type fields struct {
		Limit  int
		Offset int
		Sort   Sort
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Limit:  0,
				Offset: 0,
				Sort:   make(Sort),
			},
		},
		{
			name: "err  invalid limit",
			fields: fields{
				Limit:  -1,
				Offset: 0,
				Sort:   make(Sort),
			},
			wantErr: true,
		},
		{
			name: "err  invalid offset",
			fields: fields{
				Limit:  0,
				Offset: -1,
				Sort:   make(Sort),
			},
			wantErr: true,
		},
		{
			name: "err  invalid sort",
			fields: fields{
				Limit:  0,
				Offset: 0,
				Sort: map[SortBy]SortOrder{
					"id": 2,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rq := require.New(t)

			p := Pageable{
				Limit:  tt.fields.Limit,
				Offset: tt.fields.Offset,
				Sort:   tt.fields.Sort,
			}

			err := p.Validate()
			if tt.wantErr {
				rq.Error(err)
			} else {
				rq.NoError(err)
			}
		})
	}
}
