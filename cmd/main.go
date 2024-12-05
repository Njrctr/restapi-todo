package main

import (
	"fmt"
	"log"
	"os"

	todo "github.com/Njrctr/restapi-todo"
	handler "github.com/Njrctr/restapi-todo/pkg/handlers"
	"github.com/Njrctr/restapi-todo/pkg/repository"
	"github.com/Njrctr/restapi-todo/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка инициализации конфига: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка получения переменных окружения: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("Ошибка инициализации Базы данных: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(todo.Server)
	log.Printf("Попытка запуска сервера на порту %s", viper.GetString("port"))
	if err := server.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
	fmt.Println(server)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
