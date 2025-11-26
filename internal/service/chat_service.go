package service

import (
	"github.com/MIrrox27/REST-API/internal/domain/chat"
)

type ChatReposytory interface {
	SaveMessage(msg chat.Message) error
}
