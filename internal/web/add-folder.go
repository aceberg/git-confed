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
