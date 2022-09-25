package models

type Pais struct {
	IdPais     int    `json:"idPais"`
	NombrePais string `json:"nombrePais"`
	EstadoPais bool   `json:"estadoPais"`
}
