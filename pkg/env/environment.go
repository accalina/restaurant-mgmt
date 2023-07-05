package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Environment                            string
	ServerPort                             string

	// database
	IsRunMigration string
	DBHost, DBPort, DBUser, DBPass, DBName string

	// chace
	RedisHost, RedisPort, RedisPass, RedisPoolMaxSize, RedisPoolMinIdleSize string
}

var env Env

// BaseEnv get base global environment
func BaseEnv() Env {
	return env
}

func Load(filenames ...string) {
	if err := godotenv.Load(filenames...); err != nil {
		log.Fatal(err.Error())
	}

	env.Environment = os.Getenv("ENVIRONMENT")
	env.ServerPort = os.Getenv("SERVER_PORT")

	parseDatabaseEnv()
	parseChaceEnv()

}

func parseDatabaseEnv() {
	env.DBHost = os.Getenv("DB_HOST")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBUser = os.Getenv("DB_USER")
	env.DBPass = os.Getenv("DB_PASS")
	env.DBName = os.Getenv("DB_NAME")
	env.IsRunMigration = os.Getenv("IS_RUN_MIGRATION")
}

func parseChaceEnv() {
	env.RedisHost = os.Getenv(("REDIS_HOST"))
	env.RedisPort = os.Getenv(("REDIS_PORT"))
	env.RedisPass = os.Getenv(("REDIS_PASS"))
	env.RedisPoolMaxSize = os.Getenv(("REDIS_POOL_MAX_SIZE"))
	env.RedisPoolMinIdleSize = os.Getenv(("REDIS_POOL_MIN_IDLE_SIZE"))
}
