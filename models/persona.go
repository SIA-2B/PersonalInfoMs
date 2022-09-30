package models

type Persona struct {
	IDPersona             int    `json:"idPersona"`
	NombrePersona         string `json:"nombrePersona"`
	ApellidoPersona       string `json:"apellidoPersona"`
	TipoDocumento         string `json:"tipoDocumento"`
	NUIPPersona           int    `json:"NUIPPersona"`
	LugarNacimiento       int    `json:"lugarNacimiento"`
	LugarExpDocumento     int    `json:"lugarExpDocumento"`
	EstadoCivil           int    `json:"estadoCivil"`
	SexoBio               string `json:"sexoBio"`
	Etnia                 int    `json:"etnia"`
	CorreoPersonal        string `json:"correoPersonal"`
	TelefonoMovil         int    `json:"telefonoMovil"`
	TelefonoFijo          int    `json:"telefonoFijo"`
	FechaNacimiento       string `json:"fechaNacimiento"`
	EPS                   int    `json:"EPS"`
	GrupoSangre           string `json:"grupoSangre"`
	NivelAcademico        string `json:"nivelAcademico"`
	FactorRH              string `json:"factorRH"`
	DirResidencia         string `json:"dirResidencia"`
	LugarResidencia       int    `json:"lugarResidencia"`
	EstratoSocioeconomico int    `json:"estratoSocioeconomico"`
	LibretaMilitar        bool   `json:"libretaMilitar"`
	FechaRegistroSistema  string `json:"fechaRegistroSistema"`
	EstadoPersona         bool   `json:"estadoPersona"`
}
