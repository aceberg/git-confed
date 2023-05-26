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
	BlockMap map[string]string
	ListURL  []string
	Other    []string
}

// Repo - git repository
type Repo struct {
	Folder string
	Color  string
	Path   string
	Branch string
	User   string
	Remote []string
	URL    []string
	Other  []string
}

// GuiData - web gui data
type GuiData struct {
	Config Conf
	Repos  []Repo
	Themes []string
	Path   string
}
