package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	apiserver "github.com/Kirship/apiserver/internal/app"
)

var (
	configpath string
)

func init() {
	flag.StringVar(&configpath, "config-path", "configs/apiserver.toml", "path to config-file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configpath, config)
	if err != nil {
		fmt.Println(err)
	}
	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
