package model

type AppConfigKV struct {
	Key   string `gorm:"primaryKey" json:"key"`
	Value string `json:"value"`
}

func (AppConfigKV) TableName() string {
	return "app_config"
}

const GravatarConfigKey = "gravatar"

type GravatarConfig struct {
	BaseURL string `json:"base_url"`
}

var DefaultGravatarConfig = GravatarConfig{
	BaseURL: "https://www.gravatar.com/avatar/",
}

type AppConfig struct {
	Gravatar GravatarConfig `json:"gravatar"`
}
