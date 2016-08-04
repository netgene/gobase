package main

import (
    "fmt"
    "os"
    "github.com/jinzhu/configor"
    "github.com/op/go-logging"
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
    log.Info(s)
}

var log = logging.MustGetLogger("example")
var format = logging.MustStringFormatter(
    `[%{time:15:04:05 000000}] [%{level:.4s}] %{message}`,
)
var config Config

func main() {
    backend1 := logging.NewLogBackend(os.Stderr, "", 0)
    backend2 := logging.NewLogBackend(os.Stderr, "", 0)
    backend2Formatter := logging.NewBackendFormatter(backend2, format)
    backend1Leveled := logging.AddModuleLevel(backend1)
    backend1Leveled.SetLevel(logging.ERROR, "")
    logging.SetBackend(backend1Leveled, backend2Formatter)

    configor.Load(&config, "/media/sf_share/gobase/conf/config.yml")
    config.Println()
}
