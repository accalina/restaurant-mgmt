package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type configImplement struct {
}

func (config *configImplement) Get(key string) string {
	return os.Getenv(key)
}

type Config interface {
	Get(key string) string
}

func New(filenames ...string) Config {
	if err := godotenv.Load(filenames...); err != nil {
		log.Fatal(err.Error())
	}
	return &configImplement{}
}
