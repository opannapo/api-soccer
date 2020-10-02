package endpoints

import (
	"app/app/apis/endpoints/base"
	"app/app/injection/servicesinjection"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PlayerEndpoint struct {
	services *servicesinjection.ServiceInjection
}

func NewInstancePlayerEndpoint(group *gin.RouterGroup, services *servicesinjection.ServiceInjection) {
	instance := &PlayerEndpoint{
		services: services,
	}

	group.GET("/players", instance.getAll)
	group.GET("/players/:id", instance.getByTeam)
	group.GET("/player/:id", instance.getById)
}

func (instance *PlayerEndpoint) getAll(c *gin.Context) {
	mysqlService := instance.services.PlayerService
	data, err := mysqlService.GetAll()
	if err != nil {
		base.OutFailed(c, 0, err.Error())
	} else {
		base.OutOk(c, data)
	}
}

func (instance *PlayerEndpoint) getByTeam(c *gin.Context) {
	mysqlService := instance.services.PlayerService

	paramId := c.Param("id")
	teamId, _ := strconv.Atoi(paramId)

	data, err := mysqlService.GetByTeam(teamId)
	if err != nil {
		base.OutFailed(c, 0, err.Error())
	} else {
		base.OutOk(c, data)
	}
}

func (instance *PlayerEndpoint) getById(c *gin.Context) {
	mysqlService := instance.services.PlayerService

	paramId := c.Param("id")
	playerId, _ := strconv.Atoi(paramId)

	data, err := mysqlService.GetOne(playerId)
	if err != nil {
		base.OutFailed(c, 0, err.Error())
	} else {
		base.OutOk(c, data)
	}
}
