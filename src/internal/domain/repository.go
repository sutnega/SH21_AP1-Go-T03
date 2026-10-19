package domain

type GameRepository interface {
	Save(game *Game) error
	Get(id string) (*Game, error)
}
