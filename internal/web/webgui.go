package web

import (
	"log"
	"net/http"

	"github.com/aceberg/git-confed/internal/check"
	"github.com/aceberg/git-confed/internal/conf"
)

// Gui - start web server
func Gui(confPath string) {

	AppConfig = conf.Get(confPath)
	AppConfig.ConfPath = confPath
	AppConfig.Icon = Icon
	log.Println("INFO: starting web gui with config", AppConfig.ConfPath)

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/file_save/", saveFileHandler)
	http.HandleFunc("/folder_add/", addFolderHandler)
	http.HandleFunc("/folder_del/", delFolderHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/config_save/", saveConfigHandler)
	http.HandleFunc("/edit/", editHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
