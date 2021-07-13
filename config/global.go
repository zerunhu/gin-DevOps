package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GdoDb     *gorm.DB
	GdoVp     *viper.Viper
	GdoLog    *zap.Logger
	GdoConfig Server
)

type Server struct {
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql    Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT      JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Casbin   Casbin   `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
}