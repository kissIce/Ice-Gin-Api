package typestruct

type System struct {
	Env         string `mapstructure:"env" json:"env" yaml:"env"`
	Addr        int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType      string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	Singlelogin bool `mapstructure:"single-login" json:"singleLogin" yaml:"single-login"`
	Cache       bool   `mapstructure:"cache" json:"cache" yaml:"cache"`
	Cachetype   string `mapstructure:"cache-type" json:"cacheType" yaml:"cache-type"`
	Language    string `mapstructure:"language" json:"language" yaml:"language"`
}
