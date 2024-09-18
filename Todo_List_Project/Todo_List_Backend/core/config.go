package core

import (
	"todo-list/model"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

var appConfig *model.Config

func LoadConfig() *model.Config {
	if (appConfig) != nil {
		return appConfig
	}

	fmt.Println("Loading configuration file...")

	file, err := os.Open("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Unable to open config file.")
		panic(err)
	}

	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Unable to read config file.")
		panic(err)
	}

	err = yaml.Unmarshal(buffer, &appConfig)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Unable to parse config file.")
		panic(err)
	}

	fmt.Println("Configuration loaded.")

	return appConfig
}
