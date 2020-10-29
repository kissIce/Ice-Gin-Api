package typestruct

type Captcha struct {
	Length int `mapstructure:"length" json:"length" yaml:"long"`
	Width  int `mapstructure:"width" json:"width" yaml:"width"`
	Height int `mapstructure:"height" json:"height" yaml:"height"`
}
