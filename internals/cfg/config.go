package cfg

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Cfg struct {
	Port   string
	DbName string
	DbUser string
	DbPass string
	DbHost string
	DbPort string
}

func LoadAndStoreConfig() Cfg {
	v := viper.New()
	v.SetEnvPrefix("SERV")
	v.SetDefault("PORT", "8080") // порт сервера
	v.SetDefault("DBNAME", "parking_database") // название базыданных
	v.SetDefault("DBUSER", "postgres") // пользователь бд
	v.SetDefault("DBPASS", "mount7890") // пароль
	v.SetDefault("DBHOST", "localhost") // хост
	v.SetDefault("DBPORT", "5432") // порт бд по дефолту 5432 в psql
	v.AutomaticEnv()

	var cfg Cfg

	err := v.Unmarshal(&cfg)

	if err != nil {
		log.Panic(err)
	}
	return cfg
}

func (cfg *Cfg) GetDBString() string {
	
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost,cfg.DbPort,cfg.DbName)
}
