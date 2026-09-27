package main

import (
	"cliente/capaFachadaServices"
)

func main() {
	fachada := capaFachadaServices.NuevaFachadaCliente()
	fachada.EjecutarMenuPrincipal()
}
