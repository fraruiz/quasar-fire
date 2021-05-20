package bootstrap

import (
	"database/sql"
	"fmt"

	server "github.com/franciscoruizar/quasar-fire/internal/infrastructure/server"
	"github.com/franciscoruizar/quasar-fire/internal/infrastructure/storage/mysql"
	"github.com/franciscoruizar/quasar-fire/internal/usecases"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

func Run() error {

	var cfg config
	err := envconfig.Process("APP", &cfg)
	if err != nil {
		return err
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	sateliteRepository := mysql.NewSateliteRepository(db)

	topSecretCreator := usecases.NewTopSecretCreator(sateliteRepository)

	srv := server.New(cfg.Host, cfg.Port, topSecretCreator)
	return srv.Run()
}

type config struct {
	// Server configuration
	Host string `default:"localhost"`
	Port uint   `default:"8080"`
	// Database configuration
	DbUser string `default:"root"`
	DbPass string `default:"lechuga"`
	DbHost string `default:"localhost"`
	DbPort uint   `default:"3306"`
	DbName string `default:"quasar"`
}
