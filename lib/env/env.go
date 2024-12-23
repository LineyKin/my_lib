package env

import (
	"os"

	env "github.com/joho/godotenv"
)

const port_key string = "PORT"
const db_key string = "DBFILE"
const kafka_port_key = "KAFKA_PORT"

func getByKey(key string) string {
	err := env.Load(".ENV")

	// костыль для тестов
	if err != nil {
		err = env.Load("../.ENV")
	}

	if err != nil {
		panic("Невозможно загрузить .ENV")
	}

	return os.Getenv(key)
}

func GetPort() string {
	return getByKey(port_key)
}

func GetDbName() string {
	return getByKey(db_key)
}

func GetKafkaPort() string {
	return getByKey(kafka_port_key)
}
