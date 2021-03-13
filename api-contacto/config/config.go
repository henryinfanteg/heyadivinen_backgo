package config

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Constantes
var (
	connectionConfigFile = "connectionConfig.yaml"
)

// Database entity
type Database struct {
	Username string   `yaml:"username" json:"username"`
	Password string   `yaml:"password" json:"password"`
	Database string   `yaml:"database" json:"database"`
	Host     []string `yaml:"host" json:"host"`
}

// Cache entity
type Cache struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Database string `yaml:"database" json:"database"`
	Host     string `yaml:"host" json:"host"`
}

// ConnectionConfig entity
type ConnectionConfig struct {
	Database Database `yaml:"database" json:"database"`
	Cache    Cache    `yaml:"cache" json:"cache"`
}

var connectionConfig ConnectionConfig

// GetConnectionConfig obtiene la configuracion de conexiones
func GetConnectionConfig() *ConnectionConfig {
	return &connectionConfig
}

// LoadConfigFile carga los archivos de configuracion
func LoadConfigFile() {
	if err := readConnectionConfigFile(); err != nil {
		panic(err)
	}
}

// readConnectionConfigFile read connection config file
func readConnectionConfigFile() error {

	configPath := getPath()
	fileContent, err := ioutil.ReadFile(configPath + "/" + connectionConfigFile)
	if err != nil {
		fmt.Println("Error read config file:", err)
		return err
	}

	// expand environment variables
	fileContent = []byte(os.ExpandEnv(string(fileContent)))
	if err := yaml.Unmarshal(fileContent, &connectionConfig); err != nil {
		fmt.Println("Error Unmarshal:", err)
		return err
	}

	// fmt.Println("Load connection config:", connectionConfig)
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
