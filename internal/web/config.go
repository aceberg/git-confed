package web

import (
	"net/http"
	"strings"

	"github.com/aceberg/git-confed/internal/conf"
	"github.com/aceberg/git-confed/internal/models"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "emerald", "flatly", "grass", "grayscale", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "ocean", "pulse", "quartz", "sand", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "wood", "yeti", "zephyr"}

	execTemplate(w, "config", guiData)
}

func saveConfigHandler(w http.ResponseWriter, r *http.Request) {

	AppConfig.Host = r.FormValue("host")
	AppConfig.Port = r.FormValue("port")
	AppConfig.Theme = r.FormValue("theme")
	AppConfig.Color = r.FormValue("color")
	AppConfig.NodePath = r.FormValue("node")
	urls := r.FormValue("urls")
	other := r.FormValue("other")

	AppConfig.ListURL = strings.Split(urls, " ")
	list := []string{}
	for _, url := range AppConfig.ListURL {
		if url != "" {
			list = append(list, url)
		}
	}
	AppConfig.ListURL = list

	AppConfig.Other = strings.Split(other, " ")
	list = []string{}
	for _, url := range AppConfig.Other {
		if url != "" {
			list = append(list, url)
		}
	}
	AppConfig.Other = list

	conf.Write(AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
