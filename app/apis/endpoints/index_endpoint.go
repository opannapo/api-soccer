package endpoints

import (
	"app/app/apis/endpoints/base"
	"github.com/gin-gonic/gin"
)

//NewIndexEndpoint new IndexEndpoint
func NewIndexEndpoint(g *gin.RouterGroup) {
	g.GET("/", home)
}

func home(c *gin.Context) {
	var res = struct {
		Name               string `json:"name"`
		Job                string `json:"job"`
		Phone              string `json:"phone"`
		Email              string `json:"email"`
		ProjectDescription string `json:"project_description"`
	}{}
	res.Name = "Taufan Alfazri"
	res.Job = "Software Development Engineer"
	res.Phone = "+62856-12345-32"
	res.Email = "taufan.alfazri@gmail.com"
	res.ProjectDescription = "KitaBisa Test"

	base.OutOk(c, res)
}
