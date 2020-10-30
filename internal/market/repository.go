package market

type RepositoryQuery struct {
	EventID    *uint64
	MarketName *string
	Side       *string
	SortBy     *string
}

type Repository interface {
	Insert(m *Market) error
	Get(q *RepositoryQuery) ([]*Market, error)
}
