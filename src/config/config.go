package main

import (
    "fmt"
    "flag"
    "os"
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
var confFile *string

func getOpt() (exit bool) {
    confFile = flag.String("c", "", "configure file")
    bUsage := flag.Bool("h", false, "Usage")

    flag.Parse()
    if flag.NFlag() == 0 {
        usage()
        return false
    }
    if *bUsage {
        usage()
        return false
    }
    return true
}

func usage() {
    fmt.Printf("Usage :%s  [c:h]\n", os.Args[0])
    fmt.Println("        -c configure file")
    fmt.Println("        -h Show help")
}

func main() {
    
    if !getOpt() {
        return
    }

    configor.Load(&config, *confFile)
    config.Println()
}
