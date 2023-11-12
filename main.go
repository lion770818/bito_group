package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	router.POST("/AddSinglePersonAndMatch", AddSinglePersonAndMatch)
	router.DELETE("/RemoveSinglePerson", RemoveSinglePerson)
	router.POST("/QuerySinglePeople", QuerySinglePeople)

	router.Run("0.0.0.0:8080")
}
