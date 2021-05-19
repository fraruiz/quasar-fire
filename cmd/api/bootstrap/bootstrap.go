package bootstrap

import (
	"database/sql"
	"fmt"

	server "github.com/franciscoruizar/quasar-fire/internal/infrastructure/server"
	"github.com/franciscoruizar/quasar-fire/internal/infrastructure/storage/mysql"
	"github.com/franciscoruizar/quasar-fire/internal/usecases"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "root"
	dbPass = "lechuga"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "quasar"
)

func Run() error {

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	sateliteRepository := mysql.NewSateliteRepository(db)

	topSecretCreator := usecases.NewTopSecretCreator(sateliteRepository)

	srv := server.New(host, port, topSecretCreator)
	return srv.Run()
}
