package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Operacion
const (
	ADD    = "ADD"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
)

// Tipo
const (
	REQUEST  = "REQUEST"
	RESPONSE = "RESPONSE"
	TRACE    = "TRACE"
)

// Nivel
const (
	INFO  = "INFO"
	ERROR = "ERROR"
)

// Headers
const (
	RequestID = "Request-Id"
	AppID     = "App-Id"
	Location  = "Location"
)

// Archivos Loggers
const (
	infoFile = "info.log"
)

// LocationObject objecto
type LocationObject struct {
	City          string `bson:"city" json:"city"`
	Country       string `bson:"country" json:"country"`
	CountryCode   string `bson:"countryCode" json:"country_code"`
	County        string `bson:"county" json:"county"`
	Neighbourhood string `bson:"neighbourhood" json:"neighbourhood"`
	Postcode      string `bson:"postcode" json:"postcode"`
	Road          string `bson:"road" json:"road"`
	State         string `bson:"state" json:"state"`
	StateDistrict string `bson:"stateDistrict" json:"state_district"`
	Suburb        string `bson:"suburb" json:"suburb"`
	Type          string `bson:"type" json:"type"`
}

// LoggerRequest loguea un request
type LoggerRequest struct {
	Fecha      time.Time   `bson:"fecha" json:"fecha"`
	Tipo       string      `bson:"tipo" json:"tipo"`
	Level      string      `bson:"level" json:"level"`
	RequestID  string      `bson:"requestId" json:"requestId"`
	AppID      string      `bson:"appId" json:"appId"`
	Location   interface{} `bson:"location" json:"location"`
	Username   string      `bson:"username" json:"username"`
	IPCliente  string      `bson:"ipCliente" json:"ipCliente"`
	IPServidor string      `bson:"ipServidor" json:"ipServidor"`
	API        string      `bson:"api" json:"api"`
	Metodo     string      `bson:"metodo" json:"metodo"`
	URI        string      `bson:"uri" json:"uri"`
	Body       interface{} `bson:"body" json:"body"`
}

// LoggerTrace loguea un trace
type LoggerTrace struct {
	Fecha      time.Time   `bson:"fecha" json:"fecha"`
	Tipo       string      `bson:"tipo" json:"tipo"`
	Level      string      `bson:"level" json:"level"`
	RequestID  string      `bson:"requestId" json:"requestId"`
	API        string      `bson:"api" json:"api"`
	Metodo     string      `bson:"metodo" json:"metodo"`
	Message    string      `bson:"message" json:"message"`
	Endpoint   interface{} `bson:"endpoint" json:"endpoint"`
	Body       interface{} `bson:"body" json:"body"`
	StatusCode interface{} `bson:"statusCode" json:"statusCode"`
	Error      string      `bson:"error" json:"error"`
}

// LoggerResponse loguea un response
type LoggerResponse struct {
	Fecha      time.Time   `bson:"fecha" json:"fecha"`
	Tipo       string      `bson:"tipo" json:"tipo"`
	Level      string      `bson:"level" json:"level"`
	RequestID  string      `bson:"requestId" json:"requestId"`
	AppID      string      `bson:"appId" json:"appId"`
	Username   string      `bson:"username" json:"username"`
	IPCliente  string      `bson:"ipCliente" json:"ipCliente"`
	IPServidor string      `bson:"ipServidor" json:"ipServidor"`
	API        string      `bson:"api" json:"api"`
	Metodo     string      `bson:"metodo" json:"metodo"`
	URI        string      `bson:"uri" json:"uri"`
	Body       interface{} `bson:"body" json:"body"`
	StatusCode int         `bson:"statusCode" json:"statusCode"`
	Message    string      `bson:"message" json:"message"`
}

// LoggerGeneral loguea un log general
type LoggerGeneral struct {
	Fecha      time.Time `bson:"fecha" json:"fecha"`
	Level      string    `bson:"level" json:"level"`
	RequestID  string    `bson:"requestId" json:"requestId"`
	AppID      string    `bson:"appId" json:"appId"`
	Username   string    `bson:"username" json:"username"`
	IPCliente  string    `bson:"ipCliente" json:"ipCliente"`
	IPServidor string    `bson:"ipServidor" json:"ipServidor"`
	API        string    `bson:"api" json:"api"`
	Method     string    `bson:"method" json:"method"`
	URI        string    `bson:"uri" json:"uri"`
	StatusCode int       `bson:"statusCode" json:"statusCode"`
	Message    string    `bson:"message" json:"message"`
}

// PrintLog escribe un log generico
func PrintLog(collectionName string, level string, ipCliente string, ipServidor string, method string, headers http.Header, username string, uri string, statusCode int, message string) {
	var loggerObj LoggerGeneral

	loggerObj.Level = level
	loggerObj.Fecha = time.Now()
	loggerObj.RequestID = headers.Get(RequestID)
	loggerObj.AppID = headers.Get(AppID)
	loggerObj.Username = username
	loggerObj.IPCliente = ipCliente
	loggerObj.IPServidor = ipServidor
	loggerObj.API = collectionName
	loggerObj.Method = method
	loggerObj.URI = uri
	loggerObj.StatusCode = statusCode
	loggerObj.Message = message

	Print(collectionName, loggerObj)
}

// PrintRequest escribe el request en un log
func PrintRequest(collectionName string, ipCliente, ipServidor string, headers http.Header, username string, api string, metodo string, uri string, body interface{}) {
	var loggerObj LoggerRequest

	loggerObj.Fecha = time.Now()
	loggerObj.Tipo = REQUEST
	loggerObj.Level = INFO
	loggerObj.RequestID = headers.Get(RequestID)
	loggerObj.AppID = headers.Get(AppID)

	// Mapeamos el string a la entidad
	if len(headers.Get(Location)) > 0 {
		var locationObject LocationObject
		if err := json.Unmarshal([]byte(headers.Get(Location)), &locationObject); err == nil {
			loggerObj.Location = locationObject
		}
	}

	loggerObj.Username = username
	loggerObj.IPCliente = ipCliente
	loggerObj.IPServidor = ipServidor
	loggerObj.API = api
	loggerObj.Metodo = metodo
	loggerObj.URI = uri
	loggerObj.Body = body

	// fmt.Println("PROBANDO---> ", collectionName, loggerObj)
	Print(collectionName, loggerObj)
}

// PrintTrace escribe la traza en un log
func PrintTrace(collectionName string, level string, headers http.Header, api string, metodo string, message string, endpoint interface{}, body interface{}, statusCode interface{}, err error) {
	var loggerObj LoggerTrace

	loggerObj.Fecha = time.Now()
	loggerObj.Tipo = TRACE
	loggerObj.Level = level
	loggerObj.RequestID = headers.Get(RequestID)
	loggerObj.API = api
	loggerObj.Metodo = metodo
	loggerObj.Endpoint = endpoint
	loggerObj.Body = body
	loggerObj.StatusCode = statusCode
	loggerObj.Message = message
	if err != nil {
		loggerObj.Error = err.Error()
	}

	// fmt.Println("PROBANDO---> ", collectionName, loggerObj)
	Print(collectionName, loggerObj)
}

// PrintResponse escribe el response en un log
func PrintResponse(collectionName string, ipCliente, ipServidor string, level string, headers http.Header, username string, api string, metodo string, uri string, body interface{}, statusCode int, message string) {
	var loggerObj LoggerResponse

	loggerObj.Fecha = time.Now()
	loggerObj.Tipo = RESPONSE
	loggerObj.Level = level
	loggerObj.RequestID = headers.Get(RequestID)
	loggerObj.AppID = headers.Get(AppID)
	loggerObj.Username = username
	loggerObj.IPCliente = ipCliente
	loggerObj.IPServidor = ipServidor
	loggerObj.API = api
	loggerObj.Metodo = metodo
	loggerObj.URI = uri
	loggerObj.Body = body
	loggerObj.StatusCode = statusCode
	loggerObj.Message = message

	// fmt.Println("PROBANDO---> ", collectionName, loggerObj)
	Print(collectionName, loggerObj)
}

// Print escribe un log
func Print(collectionName string, request interface{}) {

	// file, err := os.OpenFile(infoFile, os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer file.Close()

	// log.SetOutput(file)
	// log.Print("logger ->", request)

	var loggerRepository = LoggerRepository{}
	if err := loggerRepository.Create(collectionName, request); err != nil {
		fmt.Println("ERROR LOGGER ->", err)
	}
	// fmt.Println("logger ->", request)
}
