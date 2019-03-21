/*
controllerパッケージはそのままcontroller．
handlerから来た外界のイベントをアプリケーション内部イベントに変換．

具体的にはデータを加工してusecaseに引き渡す．
usecaseはInputDataをサボって省いた為，entityに変換する．
上下レイヤがentityに対応してるのでこれもentity毎．
*/
package controller

import "collector/domain/repository"

type ControllerCollection struct {
	ChannelController ChannelController
	FileController    FileController
	UserController    UserController
	MessageController MessageController
}

func NewControllerCollection(repos *repository.RepositoryCollection) ControllerCollection {
	collection := ControllerCollection{
		ChannelController: NewChannelController(repos.ChannelRepository),
		FileController:    NewFileController(repos.FileRepository),
		UserController:    NewUserController(repos.UserRepository),
		MessageController: NewMessageController(repos.MessageRepository),
	}
	return collection
}
