package main

import (
	"github.com/cmczk/human-ed/internal/config"
)

func main() {
	cfg := config.MustLoad()

	_ = cfg
}
