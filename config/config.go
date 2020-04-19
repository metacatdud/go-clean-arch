package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Debug bool
}

var (
	// Config instance
	Config config
	envArg string
)

// Read config file
func Read() {

	flag.StringVar(&envArg, "env", "", "Environment file")
	flag.Parse()

	viper.SetConfigFile(envArg)
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
