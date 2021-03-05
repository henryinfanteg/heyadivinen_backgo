package util

import (
	"fmt"
	"net/http"
)

// Headers
const (
	RequestID = "Request-Id"
	AppID     = "App-Id"
)

// Acciones
const (
	ADD    = "ADD"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
)

// Auditar guarda la auditoria en BD
func Auditar(collectionName string, usuario string, headers http.Header, idRegistro interface{}, data interface{}, accion string) {
	var auditoriaRepository = AuditoriaRepository{}
	if err := auditoriaRepository.Create(collectionName, usuario, headers.Get(RequestID), headers.Get(AppID), idRegistro, data, accion, ""); err != nil {
		fmt.Println("ERROR AUDITORIA ->", err)
	}
}

// AuditarWithComentario guarda la auditoria en BD y un comentario
func AuditarWithComentario(collectionName string, usuario string, headers http.Header, idRegistro interface{}, data interface{}, accion string, comentario string) {
	var auditoriaRepository = AuditoriaRepository{}
	if err := auditoriaRepository.Create(collectionName, usuario, headers.Get(RequestID), headers.Get(AppID), idRegistro, data, accion, comentario); err != nil {
		fmt.Println("ERROR AUDITORIA ->", err)
	}
}
