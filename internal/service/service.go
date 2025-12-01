package service

import (
	"github.com/MIrrox27/REST-API/internal/adapter/postgress"
	"github.com/MIrrox27/REST-API/internal/domain/chat"
)

type ChatReposytory interface {
	SaveMessage(msg chat.Message) error
}
type ChatServiceImpl struct { // структура для работы с бд, получает зависимость из
	repo postgress.PostgresChat
}
