package gestion

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEjerciciosPorCardio(c *gin.Context) {
	cardio, _ := strconv.ParseBool(c.Param("cardio"))
	var ejercicios []Ejercicio
	if err := db.Where("tipo_ejercicio.cardio = ?", cardio).Joins("JOIN tipo_ejercicio ON tipo_ejercicio.id = ejercicios.tipo_ejercicio_id").Find(&ejercicios).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, ejercicios)
	}
}
