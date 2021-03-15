package cmd

import (
	"log"

	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/adapters/database"
	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/adapters/http/api"
	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/usecases"
	"github.com/spf13/cobra"
)

var (
	addr string
	dsn  string
)

var HttpServerCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server",
	Long:  `Start HTTP server`,
	Run: func(cmd *cobra.Command, args []string) {
		linkStorage, err := database.NewPgEventStorage(dsn)
		if err != nil {
			log.Println(err)
			return
		}
		linkCases := &usecases.LinkService{
			LinkStorage: linkStorage,
		}
		server := &api.HttpServer{
			LinkUseCases: linkCases,
		}
		server.Serve(addr)
	},
}

func init() {
	rootCmd.AddCommand(HttpServerCmd)

	HttpServerCmd.Flags().StringVar(&addr, "addr", "localhost:8080", "host:port to listen")
	HttpServerCmd.Flags().StringVar(&dsn, "dsn", "host=127.0.0.1 user=user password=password dbname=postgres", "database connection string")
}
