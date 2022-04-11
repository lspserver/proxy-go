package config

type Config struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	MetaData   MetaData `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type MetaData struct {
	Name string `yaml:"name"`
}

type Spec struct {
	Server []Server `yaml:"server"`
}

type Server struct {
	Name string   `yaml:"name"`
	Run  []string `yaml:"run"`
}

var (
	Build   string
	Version string
)

func New() *Config {
	return &Config{}
}
