package repository

type RepositoryCollection struct {
	ChannelRepository IChannelRepository
	FileRepository    IFileRepository
	MessageRepository IMessageRepository
	UserRepository    IUserRepository
}
