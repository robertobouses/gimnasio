package gestion

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type TipoEjercicio struct {
	Clase  string `json:"clase"`
	Cardio bool   `json:"cardio"`
}

type Ejercicio struct {
	ID            int           `json:"id"`
	Nombre        string        `json:"nombre"`
	TipoEjercicio TipoEjercicio `json:"tipoEjercicio"`
}

func GetEjercicios(c *gin.Context) {

	db, _ := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=mysecretpassword sslmode=disable")

	var ejercicios []Ejercicio
	if err := db.Find(&ejercicios).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, ejercicios)
	}
}
