package main

/* the Viper package appears to use 'mapstructure' internally
 *  GitHub: https://github.com/mitchellh/mapstructure
 *  Docs:   https://godoc.org/github.com/mitchellh/mapstructure
 */

import (
	"fmt"

	viper "github.com/spf13/viper"
)

type Simple struct {
	Key1 string `mapstructure:"key1"`
	Key2 string `mapstructure:"key2"`
	Key3 int    `mapstructure:"key3"`
}

type Complex struct {
	AltName string            `mapstructure:"name"`
	TS      uint64            `mapstructure:"ts"`
	Val     float64           `mapstructure:"val"`
	Tags    map[string]string `mapstructure:"tags"`
}

type VecElem struct {
	Name string            `mapstructure:"name"`
	Val  string            `mapstructure:"val"`
	Tags map[string]string `mapstructure:"tags"`
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

	var v []VecElem
	viper.UnmarshalKey("vec", &v)
	fmt.Printf("Vec:\n%v\n", v)
}

func (s Simple) String() string {
	retStr := fmt.Sprintf("Key1:  %v\n", s.Key1)
	retStr += fmt.Sprintf("Key2:  %v\n", s.Key2)
	retStr += fmt.Sprintf("Key3:  %v\n", s.Key3)
	return retStr
}

func (c Complex) String() string {
	retStr := fmt.Sprintf("AltName:  %v\n", c.AltName)
	retStr += fmt.Sprintf("TS:       %v\n", c.TS)
	retStr += fmt.Sprintf("Val:      %v\n", c.Val)
	retStr += fmt.Sprintf("Tags:     %v\n", c.Tags)
	return retStr
}
