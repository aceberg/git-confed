package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/git-confed/internal/check"
	"github.com/aceberg/git-confed/internal/models"
)

// Get - read config from file or env
func Get(path string) models.Conf {
	var config models.Conf
	var folders []string

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8848")
	viper.SetDefault("THEME", "journal")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)

	err = viper.UnmarshalKey("folders", &folders)
	check.IfError(err)
	config.Folders = folders

	return config
}

// Write - write config to file
func Write(config models.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("host", config.Host)
	viper.Set("port", config.Port)
	viper.Set("theme", config.Theme)
	viper.Set("folders", config.Folders)

	err := viper.WriteConfig()
	check.IfError(err)
}
