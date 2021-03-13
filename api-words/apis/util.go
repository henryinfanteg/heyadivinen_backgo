package apis

import (
	"encoding/json"
	"io/ioutil"
)

var APIMap map[string]API
var capabilities map[string]Capability

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
func fillCapabilities() error {

	if len(capabilities) != 0 {
		return nil
	}

	capabilities = make(map[string]Capability)
	var apis []API

	objs, err := ioutil.ReadFile("apis/apis.json")
	err = json.Unmarshal(objs, &apis)
	if err != nil {
		return err
	}

	// Convertimos el array de apis a Map
	for _, api := range apis {
		for _, capability := range api.Capabilities {
			if !api.Status {
				capability.Status = false
			}
			capabilities[capability.Method+"_"+api.Endpoint+capability.Path] = capability
		}
	}
	return err
}

