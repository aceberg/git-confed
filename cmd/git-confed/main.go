package main

import (
	"flag"

	"github.com/aceberg/git-confed/internal/check"
	"github.com/aceberg/git-confed/internal/web"
)

const confPath = "/data/git-confed/config.yaml"
const blocksPath = "/data/git-confed/blocks.yaml"
const nodePath = ""

func main() {
	blocksPtr := flag.String("b", blocksPath, "Path to blocks yaml file")
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	nodePtr := flag.String("n", nodePath, "Path to node modules")
	flag.Parse()

	check.Path(*blocksPtr)
	check.Path(*confPtr)

	web.Gui(*confPtr, *blocksPtr, *nodePtr) // webgui.go
}
