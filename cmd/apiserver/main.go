package main

import (
	"log"

	"github.com/kirship/apiserver/"
)

func main() {
	s := apiserver.New()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
