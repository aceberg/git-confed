package main

import (
	"flag"

	"github.com/aceberg/git-confed/internal/check"
	"github.com/aceberg/git-confed/internal/models"
	"github.com/aceberg/git-confed/internal/web"
)

const confPath = "/data/git-confed/config.yaml"

func main() {
	var conf models.Conf

	confPtr := flag.String("c", confPath, "Path to config yaml file")
	flag.Parse()

	conf.ConfPath = *confPtr
	check.Path(conf.ConfPath)

	web.Gui(conf)
}
