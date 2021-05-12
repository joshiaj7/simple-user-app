package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/joshiaj7/simple-user-app/internal/model"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	Hostname string `envconfig:"HOSTNAME"`
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
	DbName   string `envconfig:"NAME"`
	Port     string `envconfig:"PORT"`
}

// DB for global use
var DB *gorm.DB

// SetupDB to setup db
func SetupDB() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	var c config
	err = envconfig.Process("DB", &c)
	if err != nil {
		fmt.Println(err)
	}

	format := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"

	dsn := fmt.Sprintf(format, c.Hostname, c.User, c.Password, c.DbName, c.Port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	DB.AutoMigrate(&model.User{}, &model.UserLogIn{})

	return err
}
