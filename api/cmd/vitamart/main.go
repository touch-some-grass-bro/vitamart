package main

import (
	"fmt"

	"github.com/touch-some-grass-bro/vitamart/server"
)

func main() {
  fmt.Println("Starting...")
	s := server.New()
	s.RunServer()
}

