package data

import (
	"sync"
)
//класс хранилище текущих игр
type GameStorage struct {
	games sync.Map //потокобезопасная коллекция хранилищ игр
}

func NewGameStorage() *GameStorage {
	return &GameStorage{}
}

type GameDTO struct {
	ID    string
	Board [3][3]int  //модель игрового поля
}
