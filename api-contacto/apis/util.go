package apis

import (
	"encoding/json"
	"io/ioutil"

	authUtil "github.com/henryinfanteg/heyadivinen_backgo/util-auth/util"
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

// ValidarPermisos valida los permisos de un servicio
func ValidarPermisos(method string, path string, token string) (acceso bool, err error) {
	// Llenamos las capacidades
	err = llenarCapacidades()
	if err != nil {
		return false, err
	}

	// fmt.Println(method + "_" + path)

	// Buscamos la capacidad
	capacidad := capacidades[method+"_"+path]
	if &capacidad == nil {
		return false, err
	}

	return authUtil.ValidarAccesoByToken(token, capacidad.RolMinimo, capacidad.Permiso)
}
