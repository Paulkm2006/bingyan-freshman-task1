package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Jwt    JwtConfig      `yaml:"jwt"`
	Server ServerConfig   `yaml:"server"`
	DB     PostgresConfig `yaml:"postgres"`
	Admin  AdminConfig    `yaml:"admin"`
	Mail   MailConfig     `yaml:"mail"`
	Redis  RedisConfig    `yaml:"redis"`
}
type JwtConfig struct {
	Secret string `yaml:"secret"`
	Expire int64  `yaml:"expire"`
}
type ServerConfig struct {
	Port string `yaml:"port"`
	Ver  string `yaml:"ver"`
}
type PostgresConfig struct {
	Dsn string `yaml:"dsn"`
}
type AdminConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type MailConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Nickname string `yaml:"nickname"`
	Expire   int    `yaml:"expire"`
	Resend   int    `yaml:"resend"`
}
type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

var Config ConfigStruct

func InitConfig() {
	var configFile []byte
	var err error
	configFile, err = os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		panic(err)
	}
}
