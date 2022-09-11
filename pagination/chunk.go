package pagination

// Chunk represents a block of paginated data
type Chunk[T any] struct {
	Data []T `json:"data"`
	Pageable
}
