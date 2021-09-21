package ports

type GamesRepository interface {
	Get(id string) (domain.Game, error)
	Save(domain.Game) error
}

type GamesService interface {
	Get(id string) (domain.Game, error)
	Create(name string, size uint, bombs uint) (domain.Game, error)
	Reveal(id string, row uint, col uint) (domain.Game, error)
}
