package bootstrap

import (
	server "github.com/franciscoruizar/quasar-fire/internal/infrastructure/server"
	"github.com/franciscoruizar/quasar-fire/internal/infrastructure/storage/memory"

	"github.com/franciscoruizar/quasar-fire/internal/usecases"
	"github.com/kelseyhightower/envconfig"
)

func Run() error {

	var cfg config
	err := envconfig.Process("APP", &cfg)
	if err != nil {
		return err
	}

	sateliteRepository := memory.NewSateliteRepository()
	topSecretCreator := usecases.NewTopSecretCreator(sateliteRepository)
	topSecretSplitCreator := usecases.NewTopSecretSplitCreator(sateliteRepository)

	srv := server.New(cfg.Host, cfg.Port, topSecretCreator, topSecretSplitCreator)
	return srv.Run()
}

type config struct {
	// Server configuration
	Host string `default:"localhost"`
	Port uint   `default:"8080"`
}
