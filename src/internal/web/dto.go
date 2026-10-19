package web
//отправка запроса от клиента в domain
type GameRequest struct {
	Board [3][3]int `json:"board"`
}
// получение ответа из domain клиенту
type GameResponse struct {
	ID      string    `json:"id"`
	Board   [3][3]int `json:"board"`
	Over    bool      `json:"over"`
	Message string    `json:"message,omitempty"`
}
