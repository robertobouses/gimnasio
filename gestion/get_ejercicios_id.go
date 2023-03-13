package gestion

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEjercicioByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var ejercicio Ejercicio
	if err := db.Where("id = ?", id).First(&ejercicio).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, ejercicio)
	}
}
