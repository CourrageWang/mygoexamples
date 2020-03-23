package util

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	yaml2 "pkg/mod/gopkg.in/yaml.v2@v2.2.2"
)

type Config struct {
	DB struct {
		OLDTiDB_DSN string `yaml:"OLDTiDB_DSN"`
		TiDBDSN     string `yaml:"TiDB_DSN"`
		REDIS       struct {
			HOST string `yaml:"HOST"`
			PWD  string `yaml:"PWD"`
		} `yaml:"REDIS"`
	} `yaml:"DB"`
	MQ struct {
		MQDSN string `yaml:"MQ_DSN"`
	} `yaml:"MQ"`
	PIGEON struct {
		ADDR string `yaml:"ADDR"`
		USER string `yaml:"USER"`
		PWD  string `yaml:"PWD"`
	} `yaml:"PIGEON"`
	SERVER struct {
		// web绑定地址
		BindAddress string `yaml:"BIND_ADDRESS"`
		// dev, prod
		Mode string `yaml:"MODE"`

		LOG string `yaml:"LOG"`
	} `yaml:"SERVER"`
	DUMP struct {
		//未关联周期
		SPACE int `yaml:"SPACE"`
	} `yaml:DUMP`
}

var Conf *Config

func findConfigFile() string {
	var p = flag.String("config", "", "配置文件路径")
	flag.Parse()

	//优先加载参数里的配置文件
	if *p != "" {
		return *p
	}

	return "./conf/config.yaml"
}

func init() {
	confgPath := findConfigFile()
	if confgPath == "" {
		log.Printf("warnning: config file not found, use default")
		panic("config file not found")
	}
	if err := loadFile(confgPath); err != nil {
		log.Printf("load config file %s faild %v", confgPath, err)
		panic(fmt.Sprintf("load config file %s faild %v", confgPath, err))
	}
}

func loadFile(f string) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	return loadByBytes(data)
}

func loadByBytes(data []byte) error {
	var cfg Config
	if err := yaml2.Unmarshal(data, &cfg); err != nil {
		return err
	}
	if cfg.SERVER.BindAddress == "" {
		cfg.SERVER.BindAddress = ":80"
	}
	Conf = &cfg
	return nil
}
