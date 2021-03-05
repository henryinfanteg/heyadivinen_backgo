package apis

// API entidad
type API struct {
	ID          string      `json:"id"`
	Nombre      string      `json:"nombre"`
	Descripcion string      `json:"descripcion"`
	Endpoint    string      `json:"endpoint"`
	Estado      bool        `json:"estado"`
	Capacidades []Capacidad `json:"capacidades"`
}

// Capacidad entidad
type Capacidad struct {
	Nombre    string `json:"nombre"`
	Metodo    string `json:"metodo"`
	Path      string `json:"path"`
	RolMinimo int    `json:"rolMinimo"`
	Permiso   string `json:"permiso"`
	Estado    bool   `json:"estado"`
}
