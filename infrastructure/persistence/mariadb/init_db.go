package mariadb

import (
	"collector/domain/repository"
	"database/sql"
)

// MariaDBでRepositoryセットを作成
func MariaDBInit(conn *sql.DB) repository.RepositoryCollection {
	// channelRepo := NewChannelPersistence(&conn)
	// fileRepo := NewFilePersistence(&conn)
	// messageRepo := NewMessagePersistence(&conn)
	userRepo := NewUserPersistence(conn)

	repos := repository.RepositoryCollection{
		// ChannelRepository: channelRepo,
		// FileRepository:    fileRepo,
		// MessageRepository: messageRepo,
		UserRepository: userRepo,
	}

	return repos
}
