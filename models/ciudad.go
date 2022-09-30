package models

type Ciudad struct {
	IdCiudad     int    `json:"idCiudad"`
	NombreCiudad string `json:"nombreCiudad"`
	Pais_idPais  int    `json:"Pais_idPais"`
	EstadoCiudad bool   `json:"estadoCiudad"`
}
