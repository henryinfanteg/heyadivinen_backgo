package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// Methods
const (
	Get    = "GET"
	Post   = "POST"
	Put    = "PUT"
	Delete = "DELETE"
	Path   = "PATH"
)

// Errores
const (
	errorResponseNull = "[RESPONSE - NULL]:"
	errorTimeOut      = "Server Timeout"
)

// GenerateHeaders genera un objecto http.header
func GenerateHeaders(requestID string, appID string, authorization string) http.Header {
	headers := make(http.Header)

	headers.Add(RequestID, requestID)
	headers.Add(AppID, appID)
	headers.Add(ContentType, "application/json")
	// fmt.Println("headers", headers)

	return headers
}

// GetResponse consume un servicio REST
func GetResponse(method string, endpoint string, timeout int, headers http.Header, request interface{}, response interface{}) (int, error) {
// func GetResponse(method string, endpoint string, timeout int, headers http.Header, body io.Reader, entity interface{}) (int, error) {

	var resp *http.Response
	var err error

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	requestByte, _ := json.Marshal(request)
	body := bytes.NewReader(requestByte)

	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return 500, err
	}

	// Mapeamos los headers
	req.Header = headers
	resp, err = client.Do(req)

	// Validamos TimeOut
	if err != nil {
		return 504, errors.New(errorTimeOut)
	}
	defer resp.Body.Close()

	// Validamos el statusCode es diferente de SUCCESS
	if !IsSuccessStatusCode(resp.StatusCode) {
		// fmt.Println("STATUS CODE IS ERROR ->", resp.StatusCode, resp.Status[4:len(resp.Status)], resp)
		return resp.StatusCode, errors.New(resp.Status)
	} else {
		// fmt.Println("STATUS CODE IS SUCCESS ->", resp.StatusCode, resp.Status[4:len(resp.Status)], resp)
	}

	// Convertimos la entidad
	if response != nil {
		err = ConvertBodyToEntity(resp.Body, &response)
	}

	return resp.StatusCode, err
}

// IsSuccessStatusCode valida si el statusCode es success
func IsSuccessStatusCode(statusCode int) bool {
	return (statusCode >= 200) && (statusCode <= 299)
}
