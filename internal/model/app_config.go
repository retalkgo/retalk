package model

type AppConfigKV struct {
	Key   string `gorm:"primaryKey" json:"key"`
	Value string `json:"value"`
}

func (AppConfigKV) TableName() string {
	return "app_config"
}

type GravatarConfig struct {
	BaseURL string `json:"base_url"`
}

type AppConfig struct {
	Gravatar GravatarConfig `json:"gravatar"`
}
