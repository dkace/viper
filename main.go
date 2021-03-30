package main

import (
	"github.com/gin-gonic/gin"
	"viper/common"
	"viper/router"
)

func main() {
	common.InitDB()

	r := gin.Default()
	r = router.CollectRoute(r)

	panic(r.Run(":8080"))

}
