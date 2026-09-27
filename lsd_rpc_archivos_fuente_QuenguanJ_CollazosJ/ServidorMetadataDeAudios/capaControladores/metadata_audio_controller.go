package controller

import (
	"net/http"
	"strconv"

	service "microservicio/capaFachadaServices"
	"microservicio/capaFachadaServices/dto"

	"github.com/gin-gonic/gin"
)

// MetadataAudioController expone los servicios REST del catálogo de tipos de audio.
type MetadataAudioController struct {
	service *service.MetadataAudioService
}

func NewMetadataAudioController(service *service.MetadataAudioService) *MetadataAudioController {
	return &MetadataAudioController{service: service}
}

// RegistrarTipo - POST /tipos
func (this *MetadataAudioController) RegistrarTipo(ctx *gin.Context) {
	var tipoDTO dto.TipoAudioDTO

	if err := ctx.ShouldBindJSON(&tipoDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"codigo":  http.StatusBadRequest,
			"mensaje": "Los datos del tipo de audio son inválidos: " + err.Error(),
		})
		return
	}

	this.service.RegistrarTipo(tipoDTO)

	ctx.JSON(http.StatusCreated, gin.H{
		"codigo":  http.StatusCreated,
		"mensaje": "Tipo de audio registrado correctamente",
		"objTipo": tipoDTO,
	})
}

// ConsultarTipo - GET /tipos/:idTipo
func (this *MetadataAudioController) ConsultarTipo(ctx *gin.Context) {
	idParam := ctx.Param("idTipo")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"codigo":  http.StatusBadRequest,
			"mensaje": "El id del tipo de audio es inválido",
		})
		return
	}

	respuesta := this.service.ConsultarTipo(id)
	ctx.JSON(respuesta.Codigo, respuesta)
}

// ListarTipos - GET /tipos
func (this *MetadataAudioController) ListarTipos(ctx *gin.Context) {
	tipos := this.service.ListarTipos()
	ctx.JSON(http.StatusOK, gin.H{
		"codigo":  http.StatusOK,
		"mensaje": "Tipos de audio disponibles",
		"tipos":   tipos,
	})
}
