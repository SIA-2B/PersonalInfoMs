package models

type Acudiente struct {
	IdAcudiente       int    `json:"idAcudiente"`
	NombreAcudiente   string `json:"nombreAcudiente"`
	ApellidoAcudiente string `json:"apellidoAcudiente"`
	Relacion_Persona  string `json:"Relacion_Con_Persona"`
	TipoDocumento     string `json:"tipoDocumento"`
	NUIPAcudiente     int    `json:"NUIPAcudiente"`
	Persona_idPersona int    `json:"Persona_idPersona"`
	EstadoAcudiente   bool   `json:"estadoAcudiente"`
}
