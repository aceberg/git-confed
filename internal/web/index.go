package web

import (
	// "log"
	"net/http"

	"github.com/aceberg/git-confed/internal/check"
	"github.com/aceberg/git-confed/internal/models"
)

func addNewRepo(folder, path string) {
	var oneRepo models.Repo

	oneRepo.Folder = folder
	oneRepo.Color = check.Color(folder)
	oneRepo.Path = path
	oneRepo.Branch = check.Branch(path)
	oneRepo.User, oneRepo.Remote = check.ParseConfig(path)
	oneRepo.URL = check.URL(path, AppConfig.ListURL)
	oneRepo.Other = check.URL(path, AppConfig.Other)

	AllRepos = append(AllRepos, oneRepo)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig

	AllRepos = []models.Repo{}

	for _, folder := range AppConfig.Folders {
		if check.IsRepo(folder) {
			addNewRepo(folder, folder)
		} else {
			for _, subFolder := range check.ListDir(folder) {
				if check.IsRepo(subFolder) {
					addNewRepo(folder, subFolder)
				}
			}
		}
	}

	// log.Println("ALL REPOS =", AllRepos)

	guiData.Repos = AllRepos

	execTemplate(w, "index", guiData)
}
