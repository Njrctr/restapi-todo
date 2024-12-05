package main

import (
	"fmt"
	"log"

	todo "github.com/Njrctr/restapi-todo"
	handler "github.com/Njrctr/restapi-todo/pkg/handlers"
	"github.com/Njrctr/restapi-todo/pkg/repository"
	"github.com/Njrctr/restapi-todo/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка инициализации конфига: %s", err.Error())
	}
	repos := repository.NewRepository()
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
