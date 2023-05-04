package cmd

import (
	"github.com/amirhosseinmoayedi/Project-template/internall/interface/http"
	"github.com/amirhosseinmoayedi/Project-template/internall/log"

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

	server := http.NewServer()
	server.Serve()
	server.WaitForShutDownSignal()

}
