package config

type Casbin struct {
	Path  string  `mapstructure:"model-path" json:"modelPath" yaml:"model-path"`  //配置文件地址
}
