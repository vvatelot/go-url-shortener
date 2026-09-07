package config

import (
	"embed"

	"github.com/pelletier/go-toml/v2"
)

var Language string = "fr"
var T Translation
var EmbedDirI18n embed.FS

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
	Edit struct {
		Title string `toml:"title"`
	} `toml:"edit"`
	Detail struct {
		Title      string `toml:"title"`
		StatsTitle string `toml:"stats_title"`
	} `toml:"detail"`
	Label struct {
		Active      string `toml:"active"`
		Add         string `toml:"add"`
		AddedAt     string `toml:"added_at"`
		Back        string `toml:"back"`
		Cancel      string `toml:"cancel"`
		Clicks      string `toml:"clicks"`
		Detail      string `toml:"detail"`
		Edit        string `toml:"edit"`
		Last7Days   string `toml:"last_7_days"`
		Last30Days  string `toml:"last_30_days"`
		Link        string `toml:"link"`
		Next        string `toml:"next"`
		No          string `toml:"no"`
		PeriodTotal string `toml:"period_total"`
		Previous    string `toml:"previous"`
		Title       string `toml:"title"`
		URL         string `toml:"url"`
		Validate    string `toml:"validate"`
		Yes         string `toml:"yes"`
	} `toml:"label"`
}

func GetTranslation() error {
	file, err := EmbedDirI18n.ReadFile("i18n/translations/" + Language + ".toml")
	if err != nil {
		panic(err)
	}
	if err := toml.Unmarshal(file, &T); err != nil {
		return err
	}

	return nil
}
