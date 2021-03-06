package apis

import (
	"encoding/json"
	"io/ioutil"
)

var APIMap map[string]API
var capacidades map[string]Capacidad

// GetApis devuelve un Map de apis
func GetApis() (map[string]API, error) {

	if len(APIMap) != 0 {
		return APIMap, nil
	}

	APIMap = make(map[string]API)
	var apis []API

	objs, err := ioutil.ReadFile("apis.json")
	err = json.Unmarshal(objs, &apis)
	if err != nil {
		return nil, err
	}

	// Convertimos el array de apis a Map
	for _, api := range apis {
		APIMap[api.ID] = api
	}
	return APIMap, err
}

// llenarCapacidades llena un Map de capacidades
func llenarCapacidades() error {

	if len(capacidades) != 0 {
		return nil
	}

	capacidades = make(map[string]Capacidad)
	var apis []API

	objs, err := ioutil.ReadFile("apis/apis.json")
	err = json.Unmarshal(objs, &apis)
	if err != nil {
		return err
	}

	// Convertimos el array de apis a Map
	for _, api := range apis {
		for _, capaidad := range api.Capacidades {
			if !api.Estado {
				capaidad.Estado = false
			}
			capacidades[capaidad.Metodo+"_"+api.Endpoint+capaidad.Path] = capaidad
		}
	}
	return err
}
