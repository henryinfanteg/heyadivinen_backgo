package server

// ConectionDB entidad
type ConectionDB struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Database string   `json:"database"`
	Host     []string `json:"host"`
}
