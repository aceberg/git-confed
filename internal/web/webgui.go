package web

import (
	"log"
	"net/http"

	"github.com/aceberg/git-confed/internal/check"
	"github.com/aceberg/git-confed/internal/conf"
)

// Gui - start web server
func Gui(confPath, blocksPath, nodePath string) {

	AppConfig = conf.Get(confPath)
	AppConfig.ConfPath = confPath
	AppConfig.YamlPath = blocksPath
	AppConfig.NodePath = nodePath
	AppConfig.Icon = Icon
	log.Println("INFO: starting web gui with config", AppConfig.ConfPath)

	AppConfig.BlockMap = make(map[string]string)
	readBlocks()

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)                  // index.go
	http.HandleFunc("/blocks/", blocksHandler)          // blocks.go
	http.HandleFunc("/block_add/", blockAddHandler)     // blocks.go
	http.HandleFunc("/block_del/", blockDelHandler)     // blocks.go
	http.HandleFunc("/config/", configHandler)          // config.go
	http.HandleFunc("/config_save/", saveConfigHandler) // config.go
	http.HandleFunc("/edit/", editHandler)              // edit.go
	http.HandleFunc("/edit_block/", editBlockHandler)   // edit.go
	http.HandleFunc("/file_save/", saveFileHandler)     // edit.go
	http.HandleFunc("/filter/", filterHandler)          // filter.go
	http.HandleFunc("/folder_add/", addFolderHandler)   // folder.go
	http.HandleFunc("/folder_del/", delFolderHandler)   // folder.go
	http.HandleFunc("/sort/", sortHandler)              // sort.go
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
