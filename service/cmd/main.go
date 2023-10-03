package main

import (
	"github.com/spf13/viper"
	"log"
	"service/pkg/database"
	"service/pkg/hanlder"
	"service/pkg/server"
	"service/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	dbConfig := database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("port"),
		User:     viper.GetString("db.username"),
		DB:       viper.GetString("db.dbname"),
		Password: viper.GetString("db.password"),
	}

	connect, err := database.NewConnectToDatabase(dbConfig)
	if err != nil {
		return
	}

	db := database.NewDatabase(connect)
	services := service.NewService(db)
	handlers := handler.NewHandler(services)
	router := handlers.InitRouter()

	serv := new(server.Server)
	err = serv.InitServer("8080", router)
	if err != nil {
		log.Fatalf("Server can't be opened: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
