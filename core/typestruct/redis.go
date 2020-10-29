package typestruct

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Uri     string `mapstructure:"uri" json:"uri" yaml:"uri"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}