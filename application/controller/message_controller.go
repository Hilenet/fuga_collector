package controller

import (
	"collector/domain/entity"
	"collector/domain/repository"
	"collector/usecase"
)

// DAOをいい感じにしてusecaseに投げる
type MessageController struct {
	Usecase usecase.IMessageUsecase
}

// initializer
func NewMessageController(repo repository.IMessageRepository) MessageController {
	usecase := usecase.NewMessageUsecase(repo)
	messageController := MessageController{
		Usecase: &usecase,
	}
	return messageController
}

// plain message

// MessageChanged
func (c *MessageController) HandleNewMessage(message entity.Message) error {
	_, err := c.Usecase.UpdateMessage(&message)
	if err != nil {
		return err
	}

	return nil
}

// MessageChanged
func (c *MessageController) HandleMessageChange(message entity.Message) error {
	_, err := c.Usecase.UpdateMessage(&message)
	if err != nil {
		return err
	}

	return nil
}

func (c *MessageController) HandleMessageDelete(message entity.Message) error {
	_, err := c.Usecase.DeleteMessage(&message)
	if err != nil {
		return err
	}

	return nil
}
