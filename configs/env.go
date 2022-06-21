package configs

import (
	"os"
)

func EnvMongoURI() string {
	return "mongodb://" +
		os.Getenv("MONGO_USERNAME") +
		":" + os.Getenv("MONGO_PASSWORD") +
		"@" + os.Getenv("MONGO_ENDPOINT")
}
