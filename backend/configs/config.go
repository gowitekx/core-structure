package configs

import (
	"context"
	"log"

	"github.com/Sirupsen/logrus"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type Param struct {
	Port             string
	SiteURL          string
	Environment      int8
	DatabaseServer   string
	DatabaseName     string
	Username         string
	Password         string
	Logfile          string
	JWTSecretKey     string
}

var Config Param
var Log = logrus.New()

// Read and parse the configuration file
func (c *Param) Read(environment string) {
	var filepath = "./configs/"
	if environment == "testing" {
		filepath = "github.com/infinity-framework/backend/configs/" //required to change path
	}
	viper.SetConfigName("config") // no need to include file extension
	viper.AddConfigPath(filepath) // os.Getenv("configpath") optionally look for config in the working directory
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		Ld.Logger(context.Background(), ERROR, "Read Config file ", err)
		return
	}
	if err := viper.Unmarshal(&c); err != nil {
		Ld.Logger(context.Background(), ERROR, "Unmarshal Config ", err)
	}
	log.Println("Logfile:", Config.Logfile)

}
