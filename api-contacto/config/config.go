package config

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Constantes
var (
	securityConfigFile   = "securityConfig.yaml"
	connectionConfigFile = "connectionConfig.yaml"
	constantConfigFile   = "constantConfig.yaml"
)

// SecurityConfig entity
type SecurityConfig struct {
	SecretKeyPassword string `yaml:"secretKeyPassword" json:"secretKeyPassword"`
}

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

// ConstantConfig entity
type ConstantConfig struct {
	URLBase     string `yaml:"urlBase" json:"urlBase"`
	URLImagenes string `yaml:"urlImagenes" json:"urlImagenes"`
}

var securityConfig SecurityConfig
var connectionConfig ConnectionConfig
var constantConfig ConstantConfig

// GetSecurityConfig obtiene la configuracion de seguridad
func GetSecurityConfig() *SecurityConfig {
	return &securityConfig
}

// GetConnectionConfig obtiene la configuracion de conexiones
func GetConnectionConfig() *ConnectionConfig {
	return &connectionConfig
}

// GetConstantConfig obtiene las constantes
func GetConstantConfig() *ConstantConfig {
	return &constantConfig
}

// LoadConfigFile carga los archivos de configuracion
func LoadConfigFile() {
	if err := readSecurityConfigFile(); err != nil {
		panic(err)
	}
	if err := readConnectionConfigFile(); err != nil {
		panic(err)
	}
	if err := readConstantConfigFile(); err != nil {
		panic(err)
	}
}

func readSecurityConfigFile() error {

	configPath := getPath()
	fileContent, err := ioutil.ReadFile(configPath + "/" + securityConfigFile)
	if err != nil {
		fmt.Printf("Error read config file: %v\n", err)
		return err
	}

	// expand environment variables
	fileContent = []byte(os.ExpandEnv(string(fileContent)))
	if err := yaml.Unmarshal(fileContent, &securityConfig); err != nil {
		fmt.Printf("Error Unmarshal: %v\n", err)
		return err
	}

	// fmt.Printf("Load security config: %v\n", securityConfig)
	return nil
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

// readConstantConfigFile read constants config file
func readConstantConfigFile() error {

	configPath := getPath()
	fileContent, err := ioutil.ReadFile(configPath + "/" + constantConfigFile)
	if err != nil {
		fmt.Println("Error read config file:", err)
		return err
	}

	// expand environment variables
	fileContent = []byte(os.ExpandEnv(string(fileContent)))
	if err := yaml.Unmarshal(fileContent, &constantConfig); err != nil {
		fmt.Println("Error Unmarshal:", err)
		return err
	}

	// fmt.Println("Load constant config:", constantConfig)
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
