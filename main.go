package main
// @title DevOps
// @version 0.1.0
// @description OWG DevOps Platform
// @termsOfService http://swagger.io/terms/

// @contact.name 胡泽润
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath 这里写base path
import (
	"fmt"
	"gin-DevOps/config"
	"gin-DevOps/initialize"
)

func main() {
	config.GdoVp = initialize.Viper() // 初始化Viper
	config.GdoLog = initialize.Zap()  // 初始化zap日志库
	config.GdoDb = initialize.GormMysql()
	r := initialize.SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
	defer config.GdoDb.Close()
}
