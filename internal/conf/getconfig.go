package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/git-confed/internal/check"
	"github.com/aceberg/git-confed/internal/models"
)

// Get - read config from file or env
func Get(path string) models.Conf {
	var config models.Conf
	var folders, urls, other []string

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8848")
	viper.SetDefault("THEME", "darkly")
	viper.SetDefault("COLOR", "light")
	viper.SetDefault("URLS", []string{"bitbucket", "github", "gitlab"})

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)
	config.Color, _ = viper.Get("COLOR").(string)

	err = viper.UnmarshalKey("folders", &folders)
	check.IfError(err)
	config.Folders = folders

	err = viper.UnmarshalKey("urls", &urls)
	check.IfError(err)
	config.ListURL = urls

	err = viper.UnmarshalKey("other", &other)
	check.IfError(err)
	config.Other = other

	return config
}

// Write - write config to file
func Write(config models.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("host", config.Host)
	viper.Set("port", config.Port)
	viper.Set("theme", config.Theme)
	viper.Set("color", config.Color)
	viper.Set("folders", config.Folders)
	viper.Set("urls", config.ListURL)
	viper.Set("other", config.Other)

	err := viper.WriteConfig()
	check.IfError(err)
}
