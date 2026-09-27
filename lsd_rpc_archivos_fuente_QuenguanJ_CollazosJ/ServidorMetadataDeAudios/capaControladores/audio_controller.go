package controller

import (
	"net/http"
	"strconv"

	service "microservicio/capaFachadaServices"

	"github.com/gin-gonic/gin"
)

type AudioController struct {
	service *service.AudioService
}

func NewAudioController(service *service.AudioService) *AudioController {
	return &AudioController{service: service}
}

// ListarPorTipo - GET /tipos/:idTipo/audios
func (this *AudioController) ListarPorTipo(ctx *gin.Context) {
	idTipo, err := strconv.Atoi(ctx.Param("idTipo"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"codigo": http.StatusBadRequest, "mensaje": "El id del tipo es inválido"})
		return
	}

	resumen := this.service.ListarPorTipo(idTipo)
	ctx.JSON(http.StatusOK, gin.H{"codigo": http.StatusOK, "mensaje": "Audios del tipo consultado", "audios": resumen})
}

// ConsultarDetalle - GET /tipos/:idTipo/audios/:id
func (this *AudioController) ConsultarDetalle(ctx *gin.Context) {
	idTipo, err1 := strconv.Atoi(ctx.Param("idTipo"))
	id, err2 := strconv.Atoi(ctx.Param("id"))
	if err1 != nil || err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"codigo": http.StatusBadRequest, "mensaje": "Los ids son inválidos"})
		return
	}

	detalle, encontrado := this.service.ConsultarDetalle(idTipo, id)
	if !encontrado {
		ctx.JSON(http.StatusNotFound, gin.H{"codigo": http.StatusNotFound, "mensaje": "El audio no se encontró"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"codigo": http.StatusOK, "mensaje": "Detalle del audio", "objAudio": detalle})
}
