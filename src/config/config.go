package config

import (
    "fmt"
    "github.com/jinzhu/configor"
)

type Config struct {
    APPName string `yaml:"APPName"`

    DB struct {
        Name     string `yaml:"Name"`
        User     string `yaml:"User"`
        Password string `yaml:"Password"`
        Port     uint `yaml:"Port"`
    } `yaml:"DB"`

    Contacts []struct {
        Name  string `yaml:"Name"`
        Email string `yaml:"Email"`
    } `yaml:"Contacts"`
}

func (c *Config) Println() {
    s := fmt.Sprintf("config: %+v", c)
    fmt.Println(s)
}

var config Config

func test() {
    //fmt.Printf("config: %+v", config)
    configor.Load(&config, "/media/sf_share/gobase/conf/config.yml")
    config.Println()
}
