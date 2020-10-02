package endpoints

import (
	"app/app/apis/endpoints/base"
	"app/app/injection/servicesinjection"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TeamEndpoint struct {
	services *servicesinjection.ServiceInjection
}

func NewInstanceTeamEndpoint(group *gin.RouterGroup, services *servicesinjection.ServiceInjection) {
	instance := &TeamEndpoint{
		services: services,
	}

	group.GET("/teams", instance.getAll)
	group.GET("/teams/:id", instance.getById)
}

func (instance *TeamEndpoint) getAll(c *gin.Context) {
	mysqlTeamService := instance.services.TeamService
	data, err := mysqlTeamService.GetAll()
	if err != nil {
		base.OutFailed(c, 0, err.Error())
	} else {
		base.OutOk(c, data)
	}
}

func (instance *TeamEndpoint) getById(c *gin.Context) {
	mysqlTeamService := instance.services.TeamService

	paramId := c.Param("id")
	teamId, _ := strconv.Atoi(paramId)

	data, err := mysqlTeamService.GetById(teamId)
	if err != nil {
		base.OutFailed(c, 0, err.Error())
	} else {
		base.OutOk(c, data)
	}
}
