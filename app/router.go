package app

import (
	"app/app/apis/endpoints"
	"app/app/apis/endpoints/base"
	"app/app/injection/servicesinjection"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//SetupRoute Router Setup
func SetupRoute(services *servicesinjection.ServiceInjection) {
	gin.SetMode(viper.GetString("mode"))
	router := gin.Default()

	router.NoRoute(func(context *gin.Context) {
		err := base.ResError{
			Message: "Service Not Found",
			Code:    404,
		}
		res := base.ResDefault{
			Error:   err,
			Data:    nil,
			Success: false,
		}
		context.JSON(200, res)
	})

	v1 := router.Group("api/v1")
	{
		endpoints.NewIndexEndpoint(v1)
		endpoints.NewInstanceTeamEndpoint(v1, services)
	}

	_ = router.Run(viper.GetString("server.address"))
}
