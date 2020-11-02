package typestruct

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" yaml:"model-path"`
	TableName string `mapstructure:"table-name" json:"tableName" yaml:"table-name"`
}
