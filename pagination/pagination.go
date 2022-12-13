package pagination

// Chunk is a block of paginated data.
type Chunk[T any] struct {
	Data []T `json:"data"`
	Pageable
}

// Pageable is a type for pagination support.
type Pageable struct {
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
	Sort   Sort `json:"sort"`
}

// Sort is a mapping [SortBy] -> [SortOrder].
type Sort map[SortBy]SortOrder

// SortBy is a type for defining a sort field.
type SortBy string

// SortOrder is a type for defining a sort order.
type SortOrder int

const (
	Ascending SortOrder = iota - 1
	Normal
	Descending
)
