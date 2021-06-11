package initialize

import (
	"flag"
	"fmt"
	config2 "gin-DevOps/config"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func Viper() *viper.Viper {
	// 获取config配置文件路径， 优先级: 命令行(-c xxx) > 环境变量 > 默认值
	var config string
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config == "" {
		if configEnv := os.Getenv(config2.DefaultConfigEnv); configEnv == "" {
			config = config2.DefaultConfigFile
			fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config2.DefaultConfigFile)
		} else {
			config = configEnv
			fmt.Printf("您正在使用GdoConfig环境变量,config的路径为%v\n", config)
		}
	} else {
		fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
	}

	// 读取配置文件数据
	vp := viper.New()
	vp.SetConfigFile(config)
	vp.SetConfigType("yaml")
	//vp.SetConfigName("config")
	//vp.AddConfigPath("./")
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file, file: %s, err: %s \n", config, err))
	}

    // 监控实时变化
	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := vp.Unmarshal(&config2.GdoConfig); err != nil {
			fmt.Println(err)
		}
	})

	if err := vp.Unmarshal(&config2.GdoConfig); err != nil {
		fmt.Println(err)
	}
	return vp
}
