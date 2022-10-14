package pagination

import (
	"testing"
)

func TestChunk_Validate(t *testing.T) {
	type fields struct {
		Data     []string
		Pageable Pageable
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Data: []string{
					"one",
					"two",
				},
				Pageable: Pageable{},
			},
		},
		{
			name: "err  invalid pageable",
			fields: fields{
				Data: []string{
					"one",
					"two",
				},
				Pageable: Pageable{
					Limit:  -1,
					Offset: 0,
					Sort:   make(Sort),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Chunk[string]{
				Data:     tt.fields.Data,
				Pageable: tt.fields.Pageable,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Chunk[T].Validate() error = %v, wantErr %v", err, tt.wantErr)
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
			p := Pageable{
				Limit:  tt.fields.Limit,
				Offset: tt.fields.Offset,
				Sort:   tt.fields.Sort,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Pageable.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSort_Validate(t *testing.T) {
	tests := []struct {
		name    string
		s       Sort
		wantErr bool
	}{
		{
			name: "ok",
			s: Sort{
				"id":    Ascending,
				"title": Ascending,
			},
		},
		{
			name: "err  invalid sort by",
			s: Sort{
				"":      Ascending,
				"title": Ascending,
			},
			wantErr: true,
		},
		{
			name: "err  invalid sort order",
			s: Sort{
				"id":    2,
				"title": Ascending,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Sort.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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
			if err := tt.sb.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("SortBy.Validate() error = %v, wantErr %v", err, tt.wantErr)
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
			so:   Ascending,
		},
		{
			name:    "err  invalid value",
			so:      2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.so.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("SortOrder.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
