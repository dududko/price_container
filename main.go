package main

import (
	"github.com/dududko/price_container/src"
)

func main() {
	s := src.NewServer()
	s.Start("3142")
}
