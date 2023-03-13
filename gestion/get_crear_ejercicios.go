package gestion

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrearEjercicio(c *gin.Context) {
	var ejercicio Ejercicio
	if err := c.BindJSON(&ejercicio); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	if err := db.Create(&ejercicio).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, ejercicio)
}
