package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Jwt     JwtConfig      `yaml:"jwt"`
	Server  ServerConfig   `yaml:"server"`
	DB      PostgresConfig `yaml:"postgres"`
	Admin   AdminConfig    `yaml:"admin"`
	Mail    MailConfig     `yaml:"mail"`
	Captcha CaptchaConfig  `yaml:"captcha"`
	Redis   RedisConfig    `yaml:"redis"`
	Logger  LoggerConfig   `yaml:"logging"`
	ES      ESConfig       `yaml:"elasticsearch"`
	Oauth   OauthConfig    `yaml:"oauth"`
}
type JwtConfig struct {
	Secret       string   `yaml:"secret"`
	Expire       int64    `yaml:"expire"`
	SkippedPaths []string `yaml:"skippedPaths"`
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
	Workers  int    `yaml:"workers"`
}

type CaptchaConfig struct {
	Expire int    `yaml:"expire"`
	Resend int    `yaml:"resend"`
	Title  string `yaml:"title"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type LoggerConfig struct {
	Debug bool   `yaml:"debug"`
	Path  string `yaml:"path"`
}

type ESConfig struct {
	Host      string `yaml:"host"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	VerifyTls bool   `yaml:"verifyTls"`
}

type OauthConfig struct {
	ClientID     string `yaml:"clientID"`
	ClientSecret string `yaml:"clientSecret"`
	Scope        string `yaml:"scope"`
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
