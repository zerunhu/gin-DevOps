package config

type Mysql struct {
	Host         string                              // 服务器地址
	Port         string                              // 端口
	Config       []string                            // 高级配置
	Dbname       string  `mapstructure:"db-name"`    // 数据库名
	Username     string                              // 数据库用户名
	Password     string                              // 数据库密码
}
