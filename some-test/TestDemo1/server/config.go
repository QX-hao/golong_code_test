package server

import (
	"os"
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig `yaml:"server" mapstructure:"server"`
}

type ServerConfig struct {
	Port int `yaml:"port" mapstructure:"port"`
}


// 使用yaml来解析
func InitConfig_v1(contentPath string) (*Config, error){
	// if data, err := os.ReadFile(contentPath); err != nil {
	// 	return nil, fmt.Errorf("读取文件失败%s:%s",data,err)
	// }
	data,err := os.ReadFile(contentPath)
	if err != nil {
		return nil, fmt.Errorf("读取%s失败:%w", contentPath, err)
	}

	c := &Config{}
	if err := yaml.Unmarshal(data,c); err != nil {
		return nil, fmt.Errorf("解析文件失败:%w", err)
	}

	return c,nil
}


// 使用viper来解析
func IntiConfig_v2(contentPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(contentPath)

	ext := filepath.Ext(contentPath)
	switch ext {
	case ".yml", ".yaml":
		v.SetConfigType("yaml")
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败:%w", err)
	}

	c := &Config{}
	if err := v.Unmarshal(c); err != nil {
		return nil, fmt.Errorf("解析配置文件失败:%w", err)
	}

	return c, nil
}
