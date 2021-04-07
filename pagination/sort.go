package pagination

// SortBy represents a type for defining a sort field
type SortBy string

// SortOrder represents a type for defining a sort order
type SortOrder int

const (
	ASC SortOrder = iota - 1
	NO
	DESC
)

// Sort represents a mapping SortBy -> SortOrder
type Sort map[SortBy]SortOrder
