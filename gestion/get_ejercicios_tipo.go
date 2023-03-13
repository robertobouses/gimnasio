package gestion

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEjerciciosPorTipo(c *gin.Context) {
	clase := c.Param("clase")
	var ejercicios []Ejercicio
	if err := db.Where("tipo_ejercicio.clase = ?", clase).Joins("JOIN tipo_ejercicio ON tipo_ejercicio.id = ejercicios.tipo_ejercicio_id").Find(&ejercicios).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, ejercicios)
	}
}
