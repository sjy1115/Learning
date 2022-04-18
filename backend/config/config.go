package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var Conf *Config

type Config struct {
	Web   Web   `yaml:"web" json:"web"`
	Mysql Mysql `yaml:"mysql"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

func (m Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		m.User, m.Password, m.Host, m.Port, m.DB)
}

type Web struct {
	Port     int    `yaml:"port" json:"port"`
	FilePath string `yaml:"file_path"`
}

func Init(confPath string) error {
	if Conf == nil {
		Conf = new(Config)
	}

	file, err := os.Open(confPath)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, Conf)
	return err
}
