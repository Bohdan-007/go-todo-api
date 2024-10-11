package main

import (
	"log"
	"os"

	"github.com/Bohdan-007/go-todo-api"
	handler "github.com/Bohdan-007/go-todo-api/pkg/handlers"
	repository "github.com/Bohdan-007/go-todo-api/pkg/repositories"
	service "github.com/Bohdan-007/go-todo-api/pkg/services"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	resvices := service.NewService(repos)
	handlers := handler.NewHandler(resvices)

	srv := new(todo.Server)

	// if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occurred while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
