package config

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Warningf("%v", err)
	}
	log.Info("Using config file: ", viper.ConfigFileUsed())
}

func Env() string {
	return viper.GetString("env")
}

func Port() string {
	if !viper.IsSet("ports") {
		return "8080"
	}
	return viper.GetString("ports")
}

func DBHost() string {
	return viper.GetString("database.host")
}

func DBDatabase() string {
	return viper.GetString("database.database")
}

func DBUser() string {
	return viper.GetString("database.username")
}

func DBPassword() string {
	return viper.GetString("database.password")
}

func MaxIdleConns() int {
	if !viper.IsSet("database.maxIdleConns") {
		return 3
	}
	return viper.GetInt("database.maxIdleConns")
}

func MaxOpenConns() int {
	if !viper.IsSet("database.maxOpenConns") {
		return 15
	}
	return viper.GetInt("database.maxOpenConns")
}

func ConnMaxLifeTime() time.Duration {
	time := viper.GetString("database.connMaxLifeTime")
	return parseTimeDuration(time, DefaultConnMaxLifeTime)
}

func ConnMaxIdleTime() time.Duration {
	time := viper.GetString("database.connMaxIdleTime")
	return parseTimeDuration(time, DefaultConnMaxIdleTime)
}

func DBDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s", DBUser(), DBPassword(), DBHost(), DBDatabase())
}

func RedisHost() string {
	return viper.GetString("redis.host")
}

func RedisDB() int {
	if !viper.IsSet("redis.db") {
		return 0
	}
	return viper.GetInt("redis.db")
}

func RedisPoolSize() int {
	if !viper.IsSet("redis.poolSize") {
		return 10
	}
	return viper.GetInt("redis.poolSize")
}

func RedisMaxIdleConns() int {
	if !viper.IsSet("redis.maxIdleConns") {
		return 10
	}
	return viper.GetInt("redis.maxIdleConns")
}
