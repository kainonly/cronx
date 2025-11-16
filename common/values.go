package common

type Values struct {
	Address  string   `yaml:"address"`
	Node     string   `yaml:"node"`
	Key      string   `yaml:"key"`
	Origins  []string `yaml:"origins"`
	Database Database `yaml:"database"`
}

type Database struct {
	Path         string `yaml:"path"`
	Victorialogs string `yaml:"victorialogs"`
}
