// AdamAuthor
package main

import (
	"os"
	"todoApp/configs"
	"todoApp/internal/handler"
	"todoApp/internal/repository"
	"todoApp/internal/service"
	"todoApp/pkg/logger"
	"todoApp/pkg/postgres"
	"todoApp/pkg/server"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	log := logger.GetLogger()

	if err := configs.InitConfig(); err != nil {
		log.Fatalf("error with reading configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error with loading env files: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(log, services)

	srv := new(server.Server)
	err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error with running server: %s", err.Error())
	}
}
