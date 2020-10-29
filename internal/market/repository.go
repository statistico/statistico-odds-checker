package market

type Repository interface {
	Insert(m *Market) error
}
