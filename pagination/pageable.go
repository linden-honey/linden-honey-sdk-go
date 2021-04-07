package pagination

// Pageable represents a type for pagination support
type Pageable struct {
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
	Sort   Sort `json:"sort"`
}

// PageableOption represents configuration option to construct Pageable value
type PageableOption func(p *Pageable)

// NewPageable returns a pointer to the new instance of Pageable or an error
func NewPageable(opts ...PageableOption) (*Pageable, error) {
	p := &Pageable{
		Limit:  0,
		Offset: 0,
		Sort:   make(map[SortBy]SortOrder),
	}

	for _, opt := range opts {
		opt(p)
	}

	if err := p.Validate(); err != nil {
		return nil, err
	}

	return p, nil
}

// WithLimit returns a PageableOption to set limit
func WithLimit(limit int) PageableOption {
	return func(p *Pageable) {
		p.Limit = limit
	}
}

// WithOffset returns a PageableOption to set offset
func WithOffset(offset int) PageableOption {
	return func(p *Pageable) {
		p.Offset = offset
	}
}

// WithSort returns a PageableOption to set sort
func WithSort(sort Sort) PageableOption {
	return func(p *Pageable) {
		for sb, so := range sort {
			p.Sort[sb] = so
		}
	}
}
