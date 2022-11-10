package config

type Config interface {
	Setting() Settings
}

type DefaultConfig struct {
	settings Settings
}

func NewConfig(settings Settings) Config {
	return &DefaultConfig{settings: settings}
}

func (c *DefaultConfig) Setting() Settings {
	return c.settings
}
