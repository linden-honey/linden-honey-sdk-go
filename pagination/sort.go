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
