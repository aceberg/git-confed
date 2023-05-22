package models

// Conf - web gui config
type Conf struct {
	Host     string
	Port     string
	Theme    string
	Icon     string
	ConfPath string
	YamlPath string
	Folders  []string
}

// Repo - git repository
type Repo struct {
	Folder string
	Color  string
	Name   string
	Path   string
	Branch string
	User   string
	Remote []string
}

// GuiData - web gui data
type GuiData struct {
	Config Conf
	Repos  []Repo
	Themes []string
}
