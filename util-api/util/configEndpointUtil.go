package util

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Constantes
var (
	endpointConfigFile = "endpointConfig.yaml"
)

// EndpointConfig entity
type EndpointConfig struct {
	// Endpoints Api Palabras
	EndpointAPIPalabras string `yaml:"endpointApiPalabras" json:"endpointApiPalabras"`
	TimeoutAPIPalabras  int    `yaml:"timeoutApiPalabras" json:"timeoutApiPalabras"`

}

var endpointConfig EndpointConfig

// GetEndpointConfig obtiene la configuracion de los endpoints
func GetEndpointConfig() *EndpointConfig {
	// Validamos si la config esta vacia
	if (EndpointConfig{}) == endpointConfig {
		if err := readEndpointConfigFile(); err != nil {
			panic(err)
		}
	}
	return &endpointConfig
}

// readEndpointConfigFile read endpoint config file
func readEndpointConfigFile() error {

	configPath := getPath()
	fileContent, err := ioutil.ReadFile(configPath + "/" + endpointConfigFile)
	if err != nil {
		fmt.Printf("Error read config file: %v\n", err)
		return err
	}

	// expand environment variables
	fileContent = []byte(os.ExpandEnv(string(fileContent)))
	if err := yaml.Unmarshal(fileContent, &endpointConfig); err != nil {
		fmt.Printf("Error Unmarshal: %v\n", err)
		return err
	}

	// fmt.Printf("Load endpoint config: %v\n", endpointConfig)
	return nil
}

func getPath() string {
	// Obtenemos la ruta de la env sino devolvemos la ruta por defecto
	envConfigPath := os.Getenv("MAXADIVINA_CONFIG_PATH")
	if len(envConfigPath) > 0 {
		return envConfigPath
	}
	return "config"
}
