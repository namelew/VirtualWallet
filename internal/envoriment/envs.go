package envoriment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadFile(file string) {
	err := godotenv.Load(file)

	if err != nil && err == os.ErrNotExist {
		log.Println("unable to load env variables from file. ", err.Error())
	}
}

func GetVar(key string) string {
	return os.Getenv(key)
}
