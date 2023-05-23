package web

import (
	"net/http"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/aceberg/git-confed/internal/check"
	"github.com/aceberg/git-confed/internal/models"
)

func blocksHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	execTemplate(w, "blocks", guiData)
}

func readBlocks() {

	file, err := os.ReadFile(AppConfig.YamlPath)
	check.IfError(err)

	err = yaml.Unmarshal(file, &AppConfig.BlockMap)
	check.IfError(err)
}

func writeBlocks() {

	data, err := yaml.Marshal(&AppConfig.BlockMap)
	check.IfError(err)

	err = os.WriteFile(AppConfig.YamlPath, data, 0644)
	check.IfError(err)
}

func blockAddHandler(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	code := r.FormValue("code")

	if name != "" && code != "" {
		AppConfig.BlockMap[name] = code

		writeBlocks()
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func blockDelHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	delete(AppConfig.BlockMap, name)
	writeBlocks()

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
