package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

var Language string = "fr"
var T Translation

type Translation struct {
	App struct {
		Title string `toml:"title"`
	} `toml:"app"`
	Home struct {
		Title string `toml:"title"`
	} `toml:"home"`
	AddNew struct {
		Title string `toml:"title"`
	} `toml:"add_new"`
	Label struct {
		Active   string `toml:"active"`
		Add      string `toml:"add"`
		AddedAt  string `toml:"added_at"`
		Cancel   string `toml:"cancel"`
		Clicks   string `toml:"clicks"`
		Link     string `toml:"link"`
		Next     string `toml:"next"`
		Previous string `toml:"previous"`
		URL      string `toml:"url"`
		Validate string `toml:"validate"`
	} `toml:"label"`
}

func GetTranslation() error {
	if os.Getenv("LOCALE") != "" {
		Language = os.Getenv("LOCALE")
	}
	pwd, _ := os.Getwd()
	file, err := os.ReadFile(pwd + "/i18n/translations/" + Language + ".toml")
	if err != nil {
		panic(err)
	}
	if err := toml.Unmarshal(file, &T); err != nil {
		return err
	}

	return nil
}
