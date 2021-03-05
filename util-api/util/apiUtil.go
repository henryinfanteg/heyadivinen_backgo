package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

// Headers
const (
	RequestID     = "Request-Id"
	AppID         = "App-Id"
	ContentType   = "content-type"
)

// Nivel
const (
	INFO  = "INFO"
	ERROR = "ERROR"
)

// Errores
const (
	ErrorGenerico        = "Ocurrio un error al procesar la petición"
	ErrorServiceNotFound = "El servicio no existe o esta inactivo"
	ErrorHeaderNotFound  = "Headers obligatorios no encontrados"
	ErrorInvalidToken    = "Token invalido"
)

// ConvertBodyToString convierte un body en un string
func ConvertBodyToString(body io.Reader) string {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(body)
	return buffer.String()
}

// ConvertBodyToEntity convierte un body en una entidad
func ConvertBodyToEntity(body io.Reader, entity interface{}) error {

	// Obtenemos el body en bytes
	bodyArrayBytes, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("Failed reading the request body ->", err)
		return err
	}

	// Mapeamos los bytes a la entidad
	err = json.Unmarshal(bodyArrayBytes, entity)
	return err
}

// ConvertEntityToString convierte una entidad en un string
func ConvertEntityToString(entity interface{}) string {
	return fmt.Sprintf("%v", entity)
}

// ConvertEntityToJSONString convierte una entidad en un json string
func ConvertEntityToJSONString(entity interface{}) string {
	arrayBytes, _ := json.Marshal(entity)
	return string(arrayBytes)
}

// ConvertStringToEntity convierte un string en una entidad
func ConvertStringToEntity(objString string, entity interface{}) error {

	// Mapeamos el string a la entidad
	err := json.Unmarshal([]byte(objString), entity)
	return err
}

// ValidarHeaders valida los headers obligatorios
func ValidarHeaders(headers http.Header) bool {
	valid := false

	if len(headers.Get(RequestID)) != 0 && len(headers.Get(AppID)) != 0 {
		valid = true
	}

	return valid
}

// GetAPIToPath obtiene el nombre del api de un path
func GetAPIToPath(palabraSiguiente string, path string) string {
	vec := strings.Split(path, "/")

	siguiente := false
	if len(vec) > 2 {
		for _, palabra := range vec {
			if siguiente {
				return palabra
			}

			if palabraSiguiente == palabra {
				siguiente = true
			}
		}
	}
	return ""
}

// GetIPServer obtiene la ip de la maquina
func GetIPServer() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
