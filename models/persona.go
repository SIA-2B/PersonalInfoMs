package models

type Persona struct {
	IDPersona             int64  `json:"idPersona"`
	NombrePersona         string `json:"nombrePersona"`
	ApellidoPersona       string `json:"apellidoPersona"`
	TipoDocumento         string `json:"tipoDocumento"`
	NUIPPersona           int64  `json:"NUIPPersona"`
	UsernamePersona       string `json:"usernamePersona"`
	LugarNacimiento       int64  `json:"lugarNacimiento"`
	LugarExpDocumento     int64  `json:"lugarExpDocumento"`
	EstadoCivil           int64  `json:"estadoCivil"`
	SexoBio               string `json:"sexoBio"`
	Etnia                 int64  `json:"etnia"`
	CorreoPersonal        string `json:"correoPersonal"`
	TelefonoMovil         int64  `json:"telefonoMovil"`
	TelefonoFijo          int64  `json:"telefonoFijo"`
	FechaNacimiento       string `json:"fechaNacimiento"`
	RolUsuario            string `json:"rolUsuario"`
	EPS                   int64  `json:"EPS"`
	GrupoSangre           string `json:"grupoSangre"`
	NivelAcademico        string `json:"nivelAcademico"`
	FactorRH              string `json:"factorRH"`
	DirResidencia         string `json:"dirResidencia"`
	LugarResidencia       int64  `json:"lugarResidencia"`
	EstratoSocioeconomico int64  `json:"estratoSocioeconomico"`
	LibretaMilitar        bool   `json:"libretaMilitar"`
	FechaRegistroSistema  string `json:"fechaRegistroSistema"`
	EstadoPersona         bool   `json:"estadoPersona"`
}
