package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Debug  bool
	Server struct {
		Address string
	}
}

var (
	// Config instance
	Config  config
	envPath string
)

func init() {

	flag.StringVar(&envPath, "env", "", "Environment file")
	flag.Parse()
}

// Read config file
func Read() {

	viper.SetConfigFile(envPath)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		log.Println("Service RUN on DEBUG mode")

		prettyJSON, err := json.MarshalIndent(Config, "", "    ")
		if err != nil {
			log.Fatal("Failed to generate json", err)
		}

		fmt.Printf("%s\n", string(prettyJSON))
	}
}
