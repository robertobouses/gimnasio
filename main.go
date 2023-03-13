package main

import (
	"log"

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
	r.GET("/ejercicios/:id", gestion.GetEjercicioByID)
	r.GET("/ejercicios/tipo/:clase", gestion.GetEjerciciosPorTipo)
	r.GET("/ejercicios/cardio/:cardio", gestion.GetEjerciciosPorCardio)
	r.POST("/ejercicios", gestion.CrearEjercicio)
	r.Run()
}
