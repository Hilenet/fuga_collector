package usecase

import (
	"collector/domain/entity"
	"collector/domain/repository"
)

type MessageUsecase struct {
	repo repository.IMessageRepository
}

// initializer
func NewMessageUsecase(repo repository.IMessageRepository) MessageUsecase {
	MessageUsecase := MessageUsecase{
		repo: repo,
	}
	return MessageUsecase
}

// Slackから新規メッセージ
func (u *MessageUsecase) AddMessage(message *entity.Message) (*entity.Message, error) {
	res, err := u.repo.Create(message)
	if nil != err {
		return nil, err
	}

	return res, nil
}

// Slackからメッセージ情報更新
func (u *MessageUsecase) UpdateMessage(message *entity.Message) (*entity.Message, error) {
	message.Edited = true
	res, err := u.repo.Update(message)
	if nil != err {
		return nil, err
	}

	return res, nil
}

// Slackからメッセージ削除
func (u *MessageUsecase) DeleteMessage(message *entity.Message) (*entity.Message, error) {
	message.Deleted = true
	res, err := u.repo.Update(message)
	if nil != err {
		return nil, err
	}

	return res, nil
}
