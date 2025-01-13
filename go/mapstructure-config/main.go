package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/go-viper/mapstructure/v2"
	"gopkg.in/yaml.v3"
)

type ComponentOne struct {
	Attr    string    `mapstructure:"attr" yaml:"attr"`
	Measure int       `mapstructure:"measure" yaml:"measure"`
	Factor  float64   `mapstructure:"factor" yaml:"factor"`
	TS      time.Time `mapstructure:"timestamp" yaml:"timestamp"` // implements TextUnmarshaller
	IP6Addr net.IP    `mapstructure:"ip6-addr" yaml:"ip6-addr"`   // implements TextUnmarshaller
	IP4Addr net.IP    `mapstructure:"ip4-addr" yaml:"ip4-addr"`   // implements TextUnmarshaller
}

func main() {
	rawCfg, err := loadCfgFile("demo1.yaml")
	if err != nil {
		panic(err)
	}

	c1, err := decodeComponentOne(rawCfg, "comp1")
	if err != nil {
		panic(err)
	}

	fmt.Println("-- comp1 -----")
	fmt.Printf("%#v\n", c1)
	fmt.Println("---")

	// -> YAML and output
	bytes, err := yaml.Marshal(c1)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(bytes))
	fmt.Println("...")

	fmt.Println("done.")
}

// decodeComponentOne pulls the "object" at `key` and attempts to decode it into
// a `ComponentOne` struct, returning the results.
func decodeComponentOne(data map[string]any, key string) (*ComponentOne, error) {
	var c1 ComponentOne

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:   nil,
		DecodeHook: mapstructure.TextUnmarshallerHookFunc(), // <-- any existing TextUnmarshaller is applied
		Result:     &c1,
	})
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(data[key]); err != nil {
		return nil, err
	}
	return &c1, nil
}

// loadCfgFile is responsible for loading the (config) file into the generic
// map[string]any structure that mapstructure will later be used to process.
func loadCfgFile(filename string) (map[string]any, error) {

	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data := make(map[string]any)
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
