package envloader

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Load loads the environment variables from the .env file
func MustLoad() error {
	ex, _ := os.Getwd()

	err := godotenv.Load(fmt.Sprintf("%s/.env", ex))

	if err != nil {
		return err
	}

	return nil
}
