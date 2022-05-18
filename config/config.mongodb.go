package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

func (conf *ConfigImpl) NewMongoDatabase() *mgo.Database {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    conf.DBAddress,
		Timeout:  10 * time.Second,
		Database: os.Getenv("MONGO_DATABASE"),
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Println(err)
	}
	log.Println("Connected to MongoDB!")
	session.SetMode(mgo.Monotonic, true)
	return session.DB(os.Getenv("MONGO_DATABASE"))
}
