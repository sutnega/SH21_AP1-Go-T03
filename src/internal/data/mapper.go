package data
import "tic-tac-toe/internal/domain"

//Распаковывает игру из хранилища DTO в Domain
func ToDomain(dto *GameDTO) *domain.Game {
	return &domain.Game{
		ID:    dto.ID,
		Board: dto.Board,
	}
}
//упаковывает игру из Domain в DTO для переноса между слоями
func ToDTO(game *domain.Game) *GameDTO {
	return &GameDTO{
		ID:    game.ID,
		Board: game.Board,
	}
}
