package web

import (
	"net/http"

	"github.com/aceberg/git-confed/internal/check"
	"github.com/aceberg/git-confed/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig

	for _, folder := range AppConfig.Folders {
		if check.IsRepo(folder) {
			// Add to repos
		} else {
			// List subFolders
			// check.IsRepo
			// Add to repos
		}
	}

	execTemplate(w, "index", guiData)
}
