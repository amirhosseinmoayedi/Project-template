package cmd

import (
	"github.com/amirhosseinmoayedi/Project-template/internall/config"
	"github.com/amirhosseinmoayedi/Project-template/internall/interface/http"
	"github.com/amirhosseinmoayedi/Project-template/internall/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the webserver",
	Long:  `create the containers and run the http server`,
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)

}

func serve(_ *cobra.Command, _ []string) {
	log.Logger.Info("serving the application")

	log.Logger.Info("connecting to database")
	_, err := gorm.Open(postgres.New(postgres.Config{
		DSN: config.Configs.DataBase.DSN(),
	}), &gorm.Config{Logger: log.GetGormLogger()})

	if err != nil {
		log.Logger.WithError(err).Fatal("can't connect to db")
	}

	server := http.NewServer()
	server.Serve()
	server.WaitForShutDownSignal()

}
