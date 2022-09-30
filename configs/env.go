package config

import (
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type environment int

const (
	dev environment = iota
	stg
	prod
)

var EnvConfig struct {
	Env   string
	Port  string
	Mongo struct {
		Protocol   string
		Username   string
		Pasword    string
		Host       string
		ReplicaSet string
		DbName     string
	}
}

// environment variable name
const (
	ENV  string = "ENV"
	PORT string = "PORT"

	// mongo
	MONGODB_PROTOCOL    string = "MONGODB_PROTOCOL"
	MONGODB_USERNAME    string = "MONGODB_USERNAME"
	MONGODB_PASSWORD    string = "MONGODB_PASSWORD"
	MONGODB_HOST        string = "MONGODB_HOST"
	MONGODB_REPLICA_SET string = "MONGODB_REPLICA_SET"
	MONGODB_NAME        string = "MONGODB_NAME"
)

func init() {
	EnvConfig.Env = os.Getenv(ENV)
	log.Printf("Environment is '%s'.\n", getEnvironment())
}

func (e environment) String() string {
	switch e {
	case dev:
		return "dev"
	case stg:
		return "stg"
	case prod:
		return "prod"
	default:
		return "unknown"
	}
}

func getEnvironment() environment {
	env := dev
	e := EnvConfig.Env
	if e == "stg" {
		env = stg
	}
	if e == "prod" {
		env = prod
	}

	return env
}

func IsProduction() bool {
	return getEnvironment() == prod
}

func IsDevelopment() bool {
	return getEnvironment() == dev
}

var apiUrlMap = map[environment]string{
	dev:  "https://dev-api.myblog.com",
	stg:  "https://stg-api.myblog.com",
	prod: "https://prod-api.myblog.com",
}

func GetApiUrl() string {
	return apiUrlMap[getEnvironment()]
}

// LoadEnvForApi load all environment variables for api
func LoadEnvForApi() error {
	EnvConfig.Env = os.Getenv(ENV)
	EnvConfig.Port = os.Getenv(PORT)
	EnvConfig.Mongo.Protocol = os.Getenv(MONGODB_PROTOCOL)
	EnvConfig.Mongo.Username = os.Getenv(MONGODB_USERNAME)
	EnvConfig.Mongo.Pasword = os.Getenv(MONGODB_PASSWORD)
	EnvConfig.Mongo.Host = os.Getenv(MONGODB_HOST)
	EnvConfig.Mongo.ReplicaSet = os.Getenv(MONGODB_REPLICA_SET)
	EnvConfig.Mongo.DbName = os.Getenv(MONGODB_NAME)

	envErrors := make([]string, 0)
	if EnvConfig.Env == "" {
		envErrors = append(envErrors, ENV)
	}
	if EnvConfig.Mongo.Protocol == "" {
		envErrors = append(envErrors, MONGODB_PROTOCOL)
	}
	if EnvConfig.Mongo.Username == "" {
		envErrors = append(envErrors, MONGODB_USERNAME)
	}
	if EnvConfig.Mongo.Pasword == "" {
		envErrors = append(envErrors, MONGODB_PASSWORD)
	}
	if EnvConfig.Mongo.Host == "" {
		envErrors = append(envErrors, MONGODB_HOST)
	}
	if EnvConfig.Mongo.DbName == "" {
		envErrors = append(envErrors, MONGODB_NAME)
	}

	if len(envErrors) == 0 {
		return nil
	}
	return errors.New("Can not get required environment variables: " + strings.Join(envErrors, ", "))
}
