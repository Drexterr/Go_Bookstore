package initializers

import (
	"github.com/lpernett/godotenv"
)

func LoadEnvVariables() {

	err := godotenv.Load()

	if err != nil {
		panic("Error loading env file")
	}
}
