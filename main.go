package main

import (
	"bito_group/docs"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//docs "github.com/go-project-name/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func AddSinglePersonAndMatch(c *gin.Context) {

	name := c.PostForm("name")     // 姓名、
	height := c.PostForm("height") // 身高
	gender := c.PostForm("gender") // 性別

	log.Printf("name:%v, height:%v, gender:%v", name, height, gender)

	m := map[string]string{"status": "ok"}
	j, _ := json.Marshal(m)
	c.Data(http.StatusOK, "application/json", j)
}

func RemoveSinglePerson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"狀態": "ok",
	})
}

func QuerySinglePeople(c *gin.Context) {
	m := map[string]string{"status": "ok"}
	j, _ := json.Marshal(m)
	c.Data(http.StatusOK, "application/json", j)
}

func main() {

	router := gin.New()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{

		v1.POST("/AddSinglePersonAndMatch", AddSinglePersonAndMatch)
		v1.DELETE("/RemoveSinglePerson", RemoveSinglePerson)
		v1.POST("/QuerySinglePeople", QuerySinglePeople)

		v1.GET("/helloworld", Helloworld)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("0.0.0.0:8080")
}
