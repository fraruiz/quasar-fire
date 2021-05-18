package bootstrap

import (
	server "github.com/franciscoruizar/quasar-fire/internal/platform/server"
	"github.com/franciscoruizar/quasar-fire/internal/platform/storage/memory"
	"github.com/franciscoruizar/quasar-fire/internal/usecases"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	sateliteRepository := memory.NewInMemorySateliteRepository()

	topSecretCreator := usecases.NewTopSecretCreator(sateliteRepository)
	topSecretSplitCreator := usecases.NewTopSecretSplitCreator()
	sateliteFinder := usecases.NewSateliteFinder(sateliteRepository)

	srv := server.New(host, port, topSecretCreator, sateliteFinder, topSecretSplitCreator)
	return srv.Run()
}
