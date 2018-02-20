package main

import (
	"fmt"

	viper "github.com/spf13/viper"
)

type Simple struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
	Key3 int    `json:"key3"`
}

type Complex struct {
	Name string            `json:"name"`
	TS   uint64            `json:"ts"`
	Val  float64           `json:"val"`
	Tags map[string]string `json:"tags"`
}

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var s Simple
	viper.UnmarshalKey("simple", &s)
	fmt.Printf("simple:\n%v\n", s)

	var c Complex
	viper.UnmarshalKey("complex", &c)
	fmt.Printf("Complex:\n%v\n", c)
}

func (s Simple) String() string {
	retStr := fmt.Sprintf("Key1:  %v\n", s.Key1)
	retStr += fmt.Sprintf("Key2:  %v\n", s.Key2)
	retStr += fmt.Sprintf("Key3:  %v\n", s.Key3)
	return retStr
}

func (c Complex) String() string {
	retStr := fmt.Sprintf("Name:  %v\n", c.Name)
	retStr += fmt.Sprintf("TS:    %v\n", c.TS)
	retStr += fmt.Sprintf("Val:   %v\n", c.Val)
	retStr += fmt.Sprintf("Tags:  %v\n", c.Tags)
	return retStr
}
