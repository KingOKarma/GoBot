package internal

import (
	"fmt"
	util "github.com/Floor-Gang/utilpkg/config"
	"log"
)

type Config struct {
	Token  string `yaml:"bot_token"`
	Prefix string `yaml:"bot_prefix"`
}

const ConfigPath = "config.yml"

func GetConfig(location string) (config Config) {
	fmt.Println("Inside GetConfig")
	config = Config{
		Prefix: "!",
	}
	err := util.GetConfig(location, &config)

	if err != nil {
		log.Fatalln(err)
		fmt.Println("util Getconfig error")
	}
	return config
}
