package fachada

import (
	capaaccesoadatos "almacenamiento/capaAccesoADatos"
	dtos "almacenamiento/capaFachadaServices/DTOs"
	"fmt"
)

type FachadaAlmacenamiento struct {
	repo *capaaccesoadatos.RepositorioCanciones
}

// Constructor de la fachada
func NuevaFachadaAlmacenamiento() *FachadaAlmacenamiento {
	fmt.Println("Inicializando fachada de almacenamiento...")

	repo := capaaccesoadatos.GetRepositorioCanciones()
	return &FachadaAlmacenamiento{
		repo: repo,
	}
}

// GuardarCancion invoca la operación asíncrona y la síncrona
func (thisF *FachadaAlmacenamiento) GuardarCancion(objCancion dtos.CancionAlmacenarDTOInput, data []byte) error {
	//Guarda archivo y registro en memoria
	//Delegar en el repositorio de canciones para guardar el archivo y el registro en memoria
	return thisF.repo.GuardarCancion(objCancion.Id, data)
}
