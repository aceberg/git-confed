package web

import (
	"embed"

	"github.com/aceberg/git-confed/internal/models"
)

var (
	// AppConfig - config for Web Gui
	AppConfig models.Conf
	// AllRepos - all repositories
	AllRepos []models.Repo
)

// TemplHTML - html templates
//
//go:embed templates/*
var TemplHTML embed.FS

// TemplPath - path to html templates
const TemplPath = "templates/"

// // TemplPath - path to html templates
// const TemplPath = "../../internal/web/templates/"
