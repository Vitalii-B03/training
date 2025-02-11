package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func getYAML(config []Config) (string, error) {
	yamlData, err := yaml.Marshal(config[0])
	if err != nil {
		log.Fatal(err)
	}
	return string(yamlData), nil
}
func main() {
	config := []Config{
		{
			Server: Server{
				Port: "8080",
			},
			Db: Db{
				Host:     "localhost",
				Port:     "5432",
				User:     "postgres",
				Password: "postgres",
			},
		},
	}
	yamlData,err := getYAML(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(yamlData)
}
