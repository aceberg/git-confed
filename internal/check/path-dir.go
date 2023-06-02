package check

import (
	"os"
	"path/filepath"
)

// Path - create path if not exist
func Path(path string) {

	_, err := os.Stat(path)

	if path != "" && err != nil {

		dir := filepath.Dir(path)

		err = os.MkdirAll(dir, os.ModePerm)
		IfError(err)

		_, err = os.Create(path)
		IfError(err)
	}
}

// IsDir - check if path is a directory
func IsDir(path string) bool {

	info, err := os.Stat(path)
	IfError(err)

	return info.IsDir()
}

// ListDir - get list of directories
func ListDir(path string) []string {
	var list []string

	f, err := os.Open(path)
	IfError(err)

	files, err := f.Readdir(0)
	IfError(err)

	for _, v := range files {
		if v.IsDir() {
			list = append(list, path+"/"+v.Name())
		}
	}

	return list
}

// Name - return dir name from path
func Name(path string) string {

	return filepath.Base(path)
}
