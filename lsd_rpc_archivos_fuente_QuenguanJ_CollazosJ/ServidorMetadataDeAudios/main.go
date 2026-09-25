package main

import (
	"microservicio/capaAccesoADatos/repository"
	controller "microservicio/capaControladores"
	service "microservicio/capaFachadaServices"

	"github.com/gin-gonic/gin"
)

func main() {

	audioRepository := repository.NewMetadataAudioRepository()
	audioService := service.NewMetadataAudioService(audioRepository)
	audioController := controller.NewMetadataAudioController(audioService)

	router := gin.Default()

	router.POST("/audios", audioController.RegistrarAudio)
	router.GET("/audios/:titulo", audioController.ConsultarAudio)

	router.Run(":8080")
}
