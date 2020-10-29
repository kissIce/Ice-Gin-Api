package typestruct

type JWT struct {
	Key    string `mapstructure:"key" json:"key" yaml:"key"`
	Expire int64  `mapstructure:"expire" json:"expire" yaml:"expire"`
	Domain string `mapstructure:"domain" json:"domain" yaml:"domain"`
}

