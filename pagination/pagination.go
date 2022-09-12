package pagination

// SortBy represents a type for defining a sort field
type SortBy string

// SortOrder represents a type for defining a sort order
type SortOrder int

const (
	Ascending SortOrder = iota - 1
	Normal
	Descending
)

// Sort represents a mapping SortBy -> SortOrder
type Sort map[SortBy]SortOrder

// Pageable represents a type for pagination support
type Pageable struct {
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
	Sort   Sort `json:"sort"`
}

// Chunk represents a block of paginated data
type Chunk[T any] struct {
	Data []T `json:"data"`
	Pageable
}
