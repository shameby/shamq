package config

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config 配置
type Config = struct {
	Common CommonConfig `yaml:"common"`
}

// CommonConfig
type CommonConfig struct {
	ProENV   *bool
	GoPATH   string
	LogPATH  string `yaml:"logPath"`
	HTTPPort string `yaml:"httpPort"`
}

// Conf config 实例
var Conf *Config

func init() {
	Conf = new(Config)
	// 获取命令行参数
	// 如果是开发环境，直接启动或make run即可,log文件和配置所需的config.yml文件均在源码文件夹中
	// 如果是生产环境，即二进制程序单独部署时，需加上 -P,log文件和配置所需的config.yml文件在二进制程序同一级 或加上 -C 指定配置路径
	Conf.Common.ProENV = flag.Bool("P", false, "project running on production env or not")
	confPath := flag.String("C", "./config.yml", "config.yml path")
	flag.Parse()

	// 生产环境， 配置所需的 config.yml 文件在二进制程序同一级
	yamlFile, err := ioutil.ReadFile(*confPath)
	if *Conf.Common.ProENV == false {
		// 开发环境， 配置所需的 config.yml 文件在源码文件中
		Conf.Common.GoPATH = os.Getenv("GOPATH")
		filePath := filepath.Join(Conf.Common.GoPATH + "/src/shamq/config.yml")
		yamlFile, err = ioutil.ReadFile(filePath)
	}

	if err != nil {
		panic("yamlFile get err:")
	}

	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		panic("config yml Unmarshal err")
	}
}
