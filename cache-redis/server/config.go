package server

// ConectionCache entidad
type ConectionCache struct {
	Host     string `json:"host"`
	Database int   `json:"database"`
	Password string   `json:"password"`
}
