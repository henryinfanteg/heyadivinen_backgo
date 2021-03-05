package server

// ConectionMQ entidad
type ConectionMQ struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	QueueName string `json:"queueName"`
}
