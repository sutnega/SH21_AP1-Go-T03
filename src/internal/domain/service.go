package domain
//Интерфейс работающий с порядком игры
type GameService interface {
	GetNextMove(board Board) (row, col int, err error)
	ValidateBoard(current, updated Board) error
	IsGameOver(board Board) bool
}
