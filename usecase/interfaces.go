/*
usecaseパッケージにはusecaseのinterfaceと具象クラスが含まれる．
usecaseはアプリケーションの単一機能(ユースケース)を表現する．

懸念

* 具象クラスが分離していないのは正直面倒になった為．injectionするならpackage切った方が良さそう
* 今回はUsecaseが単一entityに対応するものばかりなのでまとめたがこれまとめるべきじゃなくね？
* Interactorの分離してないので，Atomicじゃない事による問題があるかも．
*/
package usecase

import "collector/domain/entity"

// entity.Channelに係るApplication機能
type IChannelUsecase interface {
	// Slackから新規チャンネル
	AddChannel(channel *entity.Channel) (*entity.Channel, error)
	// Slackからチャンネル更新
	UpdateChannel(channel *entity.Channel) (*entity.Channel, error)
	// Slackからチャンネル削除
	DeleteChannel(channel *entity.Channel) (*entity.Channel, error)
}

// entity.Fileに係るApplication機能
type IFileUsecase interface {
	// Slackから新規ファイル
	AddFile(file *entity.File) (*entity.File, error)
	// Slackからファイル更新
	UpdateFile(file *entity.File) (*entity.File, error)
	// Slackからファイル削除
	DeleteFile(file *entity.File) (*entity.File, error)
}

// entity.Messageに係るApplication機能
type IMessageUsecase interface {
	// Slackから新規メッセージ
	AddMessage(message *entity.Message) (*entity.Message, error)
	// Slackからメッセージ更新
	UpdateMessage(message *entity.Message) (*entity.Message, error)
	// Slackからメッセージ削除
	DeleteMessage(message *entity.Message) (*entity.Message, error)
}

// repositoryハンドルしてるだけ
type IUserUsecase interface {
	// Slackから新規ユーザ
	AddUser(user *entity.User) (*entity.User, error)
	// Slackからユーザ更新
	UpdateUser(user *entity.User) (*entity.User, error)
	// Slackからユーザ削除
	DeleteUser(user *entity.User) (*entity.User, error)
}
