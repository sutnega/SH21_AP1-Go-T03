package di

import (
	"tic-tac-toe/internal/application"
	"tic-tac-toe/internal/domain"
	"tic-tac-toe/internal/data"
	"tic-tac-toe/internal/web"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		data.NewGameStorage,    // singleton класс хранилище
		data.NewGameRepository, // репозиторий для работы с хранилищем
		application.NewGameService,       // сервис  для работы с репозиторием, правила игры и МИНИМАКС
		web.NewHandler,                   // хендлер
	),
	fx.Invoke(func(domain.GameRepository) {}), // гарантируем инициализацию
)
