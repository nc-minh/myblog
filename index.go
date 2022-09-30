package main

import (
	configs "myblog/configs"
	logger "myblog/utils/logger"

	"myblog/routers/v1"

	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Api server start running...")

	err := godotenv.Load("dev.env")
	if err != nil {
		logger.Warn("Error while loading dev.env file")
	}

	err = configs.LoadEnvForApi()
	if err != nil {
		logger.Warn("Cannot load environment for api")
	}

	err = routers.InitializeRouters()
	if err != nil {
		logger.Warn("Cannot initialize routes")
	}
}
