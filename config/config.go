package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}
	/*once.Do(func() {
		jwtSecret = os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			log.Fatalf("JWT_SECRET environment variable not set")
		}
		log.Println("JWT_SECRET cargado correctamente")
	})*/
}
