package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/robertobouses/gimnasio/gestion"
)

type Ejercicio struct {
	ID            int           `json:"id"`
	Nombre        string        `json:"nombre"`
	TipoEjercicio TipoEjercicio `json:"tipoEjercicio"`
}

type TipoEjercicio struct {
	Clase  string `json:"clase"`
	Cardio bool   `json:"cardio"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=mysecretpassword sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&Ejercicio{})

	r := gin.Default()
	r.GET("/ejercicios", gestion.GetEjercicios)
	r.GET("/ejercicios/:id", GetEjercicioByID)
	r.GET("/ejercicios/tipo/:clase", GetEjerciciosPorTipo)
	r.GET("/ejercicios/cardio/:cardio", GetEjerciciosPorCardio)
	r.POST("/ejercicios", CrearEjercicio)
	r.Run()
}

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
