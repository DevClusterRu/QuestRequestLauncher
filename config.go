package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"path"
)

type Link struct {
	Uri      string `yaml:"uri"`
	Interval int    `yaml:"interval"`
}

type Config struct {
	LINKS_COLLECTION []Link `yaml:"LINKS_COLLECTION"`
}

func newConfig(fileName string) (c *Config, err error) {
	log.Printf("reading config from '%s'", fileName)
	if ext := path.Ext(fileName); ext != ".yaml" && ext != ".yml" {
		err = fmt.Errorf("invalid  file '%s' extenstion, expected 'yaml' or 'yml'", ext)
		return
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = fmt.Errorf("can't read file '%s'", fileName)
		return
	}

	conf := new(Config)
	if err = yaml.Unmarshal(file, conf); err != nil {
		err = fmt.Errorf("file %s yaml unmarshal error: %v", fileName, err)
	}

	return conf, err
}
