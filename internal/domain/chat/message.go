package chat

// этот файл нужен для описания сущности сообщения

type Message struct {
	USER_ID    int
	MESSAGE_ID int
	MESSAGE    string
}

type MessageInterface interface {
	ReadMessage() *string
}
