package web

import (
	"log"
	"net/http"

	"github.com/aceberg/git-confed/internal/check"
	"github.com/aceberg/git-confed/internal/conf"
)

func addFolderHandler(w http.ResponseWriter, r *http.Request) {

	folderPath := r.FormValue("path")

	if folderPath != "" && check.IsDir(folderPath) {

		AppConfig.Folders = append(AppConfig.Folders, folderPath)
		conf.Write(AppConfig)

		log.Println("INFO: Folder added =", folderPath)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func delFolderHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Query().Get("path")

	log.Println("INFO: Remove folder", path)
	folders := []string{}

	for _, oneFolder := range AppConfig.Folders {
		if oneFolder != path {
			folders = append(folders, oneFolder)
		}
	}
	AppConfig.Folders = folders
	conf.Write(AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
