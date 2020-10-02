package endpoints

import (
	"app/app/apis/endpoints/base"
	"app/app/apis/param"
	"app/app/injection/servicesinjection"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type TeamEndpoint struct {
	services *servicesinjection.ServiceInjection
}

func NewInstanceTeamEndpoint(group *gin.RouterGroup, services *servicesinjection.ServiceInjection) {
	instance := &TeamEndpoint{
		services: services,
	}

	group.GET("/teams", instance.getAll)
	group.GET("/team/:id", instance.getById)
	group.POST("/team/create", instance.create)
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

func (instance *TeamEndpoint) create(c *gin.Context) {
	mysqlTeamService := instance.services.TeamService

	var p param.TeamCreateParameter
	err := c.ShouldBindJSON(&p)
	if err != nil {
		base.OutFailed(c, http.StatusBadRequest, err.Error())
		panic(err)
		return
	}

	if p.Players == nil {
		base.OutFailed(c, http.StatusBadRequest, "PARAM PLAYER REQUIRED")
		c.AbortWithStatus(200)
		panic(err)
		return
	}

	err = mysqlTeamService.Create(&p)
	if err != nil {
		if strings.ContainsAny("Error 1062", err.Error()) {
			errMsg := "GAGAL MENGINUT DATA, DUPLICATE ID TEAM / PLAYER"
			base.OutFailed(c, 500, errMsg)
		} else {
			base.OutFailed(c, http.StatusBadRequest, err.Error())
		}

		panic(err)
		return
	}

	base.OutOk(c, p)
}
