package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	name := viper.GetString("name")
	fmt.Printf("name = '%s'\n", name)
}
