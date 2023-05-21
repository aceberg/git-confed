package web

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"

	"github.com/aceberg/git-confed/internal/check"
	"github.com/aceberg/git-confed/internal/models"
)

func addNewRepo(folder, path string) {
	var oneRepo models.Repo

	sum := md5.Sum([]byte(folder))

	oneRepo.Folder = folder
	oneRepo.Color = fmt.Sprintf("%x", sum)[0:6]
	oneRepo.Path = path

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

	log.Println("ALL REPOS =", AllRepos)

	guiData.Repos = AllRepos

	execTemplate(w, "index", guiData)
}
