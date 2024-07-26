package main

import (
	"log"
	"phone-contact/config"
	"phone-contact/infrastructure/database"
	userRepository "phone-contact/src/entity/user/repository"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error on reading env file: %s", err)
	}

	var config config.Config

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error on unmarshalling config: %s", err)
	}

	connection, err := database.CreateDBConnection(config)
	if err != nil {
		log.Fatalf("Error on connecting to database: %s", err)
	}

	userRepo := userRepository.CreateUserRepository(connection)
	log.Println("Database connection successful", userRepo)
}
