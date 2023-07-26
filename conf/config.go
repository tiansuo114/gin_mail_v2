package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var Config *Conf

type Conf struct {
	EncryptSecret *EncryptSecret `yaml:"encryptSecret"`
	PhotoPath     *PhotoPath     `yaml:"photoPath"`
	System        *System        `yaml:"system"`
	Mysql         *Mysql         `yaml:"mysql"`
	Redis         *Redis         `yaml:"redis"`
	Cache         *Cache         `yaml:"cache"`
	Email         *Email         `yaml:"email"`
}

type Mysql struct {
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Dialect  string `yaml:"dialect"`
	DbHost   string `yaml:"dbHost"`
	DbPort   string `yaml:"dbPort"`
	DbName   string `yaml:"dbName"`
}

type Redis struct {
	RedisHost     string `yaml:"redisHost"`
	RedisPort     int    `yaml:"redisPort"`
	RedisPassword int    `yaml:"redisPassword"`
	RedisNetwork  string `yaml:"redisNetwork"`
	RedisDbName   int    `yaml:"redisDbName"`
}

type Cache struct {
	CacheEmpires int         `yaml:"cacheEmpires"`
	CacheWarmUp  interface{} `yaml:"cacheWarmUp"`
	CacheServer  interface{} `yaml:"cacheServer"`
	CacheType    string      `yaml:"cacheType"`
}

type Email struct {
	SmtpHost  interface{} `yaml:"smtpHost"`
	SmtpEmail interface{} `yaml:"smtpEmail"`
	SmtpPass  interface{} `yaml:"smtpPass"`
	Address   string      `yaml:"address"`
}

type EncryptSecret struct {
	JwtSecret   string `yaml:"jwtSecret"`
	EmailSecret string `yaml:"emailSecret"`
	PhoneSecret string `yaml:"phoneSecret"`
}

type PhotoPath struct {
	PhotoHost   string `yaml:"photoHost"`
	ProductPath string `yaml:"ProductPath"`
	AvatarPath  string `yaml:"AvatarPath"`
}

type System struct {
	Host        string  `yaml:"Host"`
	UploadModel string  `yaml:"UploadModel"`
	Domain      string  `yaml:"domain"`
	Version     float64 `yaml:"version"`
	Env         string  `yaml:"env"`
	HttpPort    string  `yaml:"HttpPort"`
}

func LoadYamlConfig() {
	absPath, err := filepath.Abs("conf/config.yaml")
	fmt.Println("absPath:", absPath)
	dataBytes, err := os.ReadFile(absPath)
	if err != nil {
		panic("读取配置文件失败" + err.Error())
	}
	conf_to_load := Conf{}
	err = yaml.Unmarshal(dataBytes, &conf_to_load)
	if err != nil {
		panic("解析配置文件失败" + err.Error())
	}

	Config = &conf_to_load
}
