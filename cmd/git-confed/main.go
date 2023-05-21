package main

import (
	"flag"

	"github.com/aceberg/git-confed/internal/check"
	"github.com/aceberg/git-confed/internal/web"
)

const confPath = "config.yaml"

func main() {
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	flag.Parse()

	check.Path(*confPtr)

	web.Gui(*confPtr)
}
