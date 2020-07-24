package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type Cfg struct {
	AppName       string `yaml:"app_name"`
	AppKey        string `yaml:"app_key"`
	AppDebug      bool   `yaml:"app_debug"`
	AppUrl        string `yaml:"app_url"`
	DbConnection  string `yaml:"db_connection"`
	DbHost        string `yaml:"db_host"`
	DbPort        uint16 `yaml:"db_port"`
	DbDatabase    string `yaml:"db_database"`
	DbUsername    string `yaml:"db_username"`
	DbPassword    string `yaml:"db_password"`
	RedisHost     string `yaml:"redis_host"`
	RedisPassword string `yaml:"redis_password"`
	RedisPort     string `yaml:"redis_port"`
	Port          string `yaml:"server_port"`
	SessMaxAge    int    `yaml:"sess_max_age"`
}

var cfg *Cfg

func LoadCfg(path string) error {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config Cfg
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	cfg = &config
	return err
}

func GetCfg() *Cfg {
	return cfg
}
