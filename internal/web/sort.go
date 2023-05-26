package web

import (
	"net/http"
	"sort"

	"github.com/aceberg/git-confed/internal/models"
)

func sortHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig

	tag := r.URL.Query().Get("tag")

	if tag == "path-up" {
		sort.Slice(AllRepos, func(i, j int) bool {
			return AllRepos[i].Path < AllRepos[j].Path
		})
	} else if tag == "path-down" {
		sort.Slice(AllRepos, func(i, j int) bool {
			return AllRepos[i].Path > AllRepos[j].Path
		})
	}

	guiData.Repos = AllRepos

	execTemplate(w, "index", guiData)
}
