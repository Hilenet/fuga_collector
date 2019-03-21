package main

import (
	"collector/infrastructure/handler"
	"collector/infrastructure/persistence/mariadb"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func init() {}

// production/developmentの管理とかベストプラクティスは？
func main() {
	// osの環境変数から接続情報抜いてきて接続
	// user:pass@tcp(host:port)/dbname
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("WRITE_USER"),
		os.Getenv("WRITE_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("PROD_MODE"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// TODO: DB周りの依存もっとinfraに寄せたいな
	// mariaDBでrepoをinit
	repos := mariadb.MariaDBInit(conn)

	// slackのRTMを購読
	slack, err := handler.NewSlackHandler(os.Getenv("SLACK_TOKEN"))
	if err != nil {
		panic(err)
	}

	// streamのhandle開始
	handler.HandleStream(&slack, &repos)
}
