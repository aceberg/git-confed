package web

import (
	"net/http"
	"os"

	"github.com/aceberg/git-confed/internal/check"
	"github.com/aceberg/git-confed/internal/models"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig

	path := r.URL.Query().Get("path")
	guiData.Path = path

	file, err := os.ReadFile(path + "/.git/config")
	check.IfError(err)
	guiData.Themes = append(guiData.Themes, string(file))

	execTemplate(w, "edit", guiData)
}

func saveFileHandler(w http.ResponseWriter, r *http.Request) {

	path := r.FormValue("path")
	text := r.FormValue("text")

	err := os.WriteFile(path+"/.git/config", []byte(text), 0644)
	check.IfError(err)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func editBlockHandler(w http.ResponseWriter, r *http.Request) {

	path := r.FormValue("path")
	blockName := r.FormValue("blockname")

	block := check.ReplaceReponame(AppConfig.BlockMap[blockName], check.Name(path))

	file, err := os.ReadFile(path + "/.git/config")
	check.IfError(err)

	text := string(file) + "\n" + block

	err = os.WriteFile(path+"/.git/config", []byte(text), 0644)
	check.IfError(err)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
