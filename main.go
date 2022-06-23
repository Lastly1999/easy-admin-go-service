package main

import (
	"easy-admin-go-service/pkg/viper"
	"easy-admin-go-service/routers"
	"log"
	"strconv"
)

// @title easy-admin信息化
// @version 1.0
// @description 接口
// @termsOfService http://swagger.io/terms/

// @contact.name lastly
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:5000
// @BasePath /v1
func main() {
	port := strconv.Itoa(viper.SrvConf.Port)
	log.Printf("swagger-ui start at http://127.0.0.1:5000/swagger/index.html services")
	err := routers.InitRouters().Run(":" + port)
	if err != nil {
		panic(err)
	}
}
