package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ConfigImpl struct {
	MODE      string
	DBAddress []string
}

func NewConf(filenames ...string) (conf ConfigImpl) {
	LOG_FILE := "logs/log.txt"
	LogFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "no such") {
			err = os.Mkdir("./logs/", 0755)
			if err != nil {
				log.Println(err)
			}
			os.Create("./logs/log.txt")
		}
	}
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(LogFile)
	err = godotenv.Load(filenames...)
	if err != nil {
		log.Println(err)
	}
	conf.MODE = os.Getenv("DMS_MODE")
	conf.SetMode()
	return
}

func (conf *ConfigImpl) SetMode() {
	if conf.MODE == "debug" {
		conf.DBAddress = []string{"localhost:27017"}
		gin.SetMode(gin.DebugMode)
	} else if conf.MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
		conf.DBAddress = []string{os.Getenv("MONGO_DATABASE_HOST") + os.Getenv("MONGO_DATABASE_PORT")}
	}
}
func (config *ConfigImpl) Get(key string) string {
	return os.Getenv(key)
}

func GetDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}
