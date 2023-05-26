package web

import (
	"log"
	"net/http"

	"github.com/aceberg/git-confed/internal/models"
)

func inSlice(slice []string, word string) bool {

	for _, str := range slice {
		if str == word {
			return true
		}
	}
	return false
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig

	tag := r.URL.Query().Get("tag")
	name := r.URL.Query().Get("name")

	newRepos := []models.Repo{}
	for _, repo := range AllRepos {
		switch tag {
		case "folder":
			if name == repo.Folder {
				newRepos = append(newRepos, repo)
			}
		case "branch":
			if name == repo.Branch {
				newRepos = append(newRepos, repo)
			}
		case "user":
			if name == repo.User {
				newRepos = append(newRepos, repo)
			}
		case "remote":
			if inSlice(repo.Remote, name) {
				newRepos = append(newRepos, repo)
			}
		case "url":
			if inSlice(repo.URL, name) {
				newRepos = append(newRepos, repo)
			}
		case "other":
			if inSlice(repo.Other, name) {
				newRepos = append(newRepos, repo)
			}
		default:
			log.Println("ERROR: wrong filter tag")
		}
	}
	AllRepos = newRepos
	guiData.Repos = newRepos

	execTemplate(w, "index", guiData)
}
