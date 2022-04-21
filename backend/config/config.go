package config

import (
	"fmt"
	"io/ioutil"
	"learning/db/cache"
	"learning/pkg/log"
	"os"

	"gopkg.in/yaml.v2"
)

type Loader interface {
	Load()
}

var Conf *Config

type Config struct {
	Web       Web          `yaml:"web" json:"web"`
	Mysql     Mysql        `yaml:"mysql" json:"mysql"`
	Jwt       Jwt          `yaml:"jwt" json:"jwt"`
	Redis     cache.Config `yaml:"redis" json:"redis"`
	Log       log.Log      `yaml:"log" json:"log"`
	IgnoreUrl []string     `yaml:"ignoreUrl" json:"ignoreUrl"`
}

type Jwt struct {
	Secret string `yaml:"secret" json:"secret"`
}

func (j *Jwt) Load() {
	j.Secret = "fjkalsfashfhasf"
}

type Mysql struct {
	Host       string `json:"host" yaml:"host"`
	DockerHost string `json:"docker_host" yaml:"docker_host"`
	Port       int    `json:"port" yaml:"port"`
	DockerPort int    `json:"docker_port" yaml:"docker_port"`
	User       string `json:"user" yaml:"user"`
	Password   string `json:"password" yaml:"password"`
	DB         string `json:"db" yaml:"db"`
}

func (m *Mysql) Load() {
	m.Host = "121.199.167.227"
	m.DockerHost = "mysql"
	m.Port = 3307
	m.DockerPort = 3306
	m.User = "root"
	m.Password = "admin"
	m.DB = "learning"
}

func (m Mysql) DSN(docker bool) string {
	if docker {
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			m.User, m.Password, m.DockerHost, m.DockerPort, m.DB)
	} else {
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			m.User, m.Password, m.Host, m.Port, m.DB)
	}
}

type Web struct {
	Port   int  `yaml:"port" json:"port"`
	Docker bool `yaml:"docker" json:"docker"`
}

func (w *Web) Load() {
	w.Port = 8081
	w.Docker = true
}

func (c *Config) Load() {
	c.Mysql.Load()
	c.Web.Load()
	c.Jwt.Load()
	c.Redis.Load()
	c.Log.Load()
	c.IgnoreUrl = []string{
		"/api/static/*",
		"/api/user/login",
		"/api/user/logout",
		"/api/user/register",
		"/api/course/chat",
		"/api/user/verifycode",
	}
}

func Init(confPath string) error {
	if Conf == nil {
		Conf = new(Config)
	}

	if confPath == "" {
		Conf.Load()
		return nil
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
