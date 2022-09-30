package models

type Etnia struct {
	IdEtnia         int    `json:"idEtnia"`
	CategoriaEtnica string `json:"categoriaEtnica"`
	EstadoEtnia     bool   `json:"estadoEtnia"`
}
