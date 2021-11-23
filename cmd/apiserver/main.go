package main

import (
	"log"

	apiserver "github.com/Kirship/apiserver/internal/app"
)

func main() {
	s := apiserver.New()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
