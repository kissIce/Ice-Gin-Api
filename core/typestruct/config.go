package typestruct

type Server struct {
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap       Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email     Email     `mapstructure:"email" json:"email" yaml:"email"`
	Casbin    Casbin    `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	Captcha   Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Md5Salt   string    `mapstructure:"md5salt" json:"md5salt" yaml:"md5salt"`
	InviteStr string    `mapstructure:"invitestr" json:"invitestr" yaml:"invitestr"`
	SnowFlake SnowFlake `mapstructure:"snowflake" json:"snowflake" yaml:"snowflake"`

	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu Qiniu `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
}
